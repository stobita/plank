package repository

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/pkg/errors"
	"github.com/stobita/plank/internal/model"
	"github.com/stobita/plank/internal/rdb"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (r *repository) GetCard(id uint) (*model.Card, error) {
	ctx := context.Background()
	row, err := rdb.Cards(
		rdb.CardWhere.ID.EQ(id),
		qm.Load(rdb.CardRels.Section),
		qm.Load(
			qm.Rels(
				rdb.CardRels.CardsLabels,
				rdb.CardsLabelRels.Label,
			),
		),
		qm.Load(
			qm.Rels(
				rdb.CardRels.Section,
				rdb.SectionRels.Board,
			),
		),
	).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return &model.Card{
		ID:          row.ID,
		Name:        row.Name,
		Description: row.Description,
		Section: &model.Section{
			ID: row.R.Section.ID,
			Board: &model.Board{
				ID: row.R.Section.BoardID,
			},
		},
	}, nil
}

func (r *repository) SaveCard(m *model.Card) error {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}
	row, err := rdb.Cards(
		rdb.CardWhere.ID.EQ(m.ID),
		qm.Load(
			qm.Rels(
				rdb.CardRels.CardsLabels,
				rdb.CardsLabelRels.Label,
			),
		),
	).One(ctx, tx)
	if err != nil {
		return err
	}

	row.Name = m.Name
	row.Description = m.Description
	row.LimitTime = null.TimeFrom(*m.LimitTime)

	if _, err := row.Update(ctx, tx, boil.Whitelist(
		rdb.CardColumns.Name,
		rdb.CardColumns.Description,
		rdb.CardColumns.LimitTime,
	)); err != nil {
		return err
	}

	if err := updateCardLabels(row, row.R.CardsLabels, m.Labels, tx); err != nil {
		return errors.Wrap(err, "update card labels error")
	}

	tx.Commit()

	if err := r.saveCardDocument(m); err != nil {
		return errors.Wrap(err, "save card document error")
	}

	return nil
}

func (r *repository) saveCardDocument(card *model.Card) error {
	ctx := context.Background()
	doc := cardDocument{
		Name:        card.Name,
		Description: card.Description,
	}
	_, err := r.esClient.Update().
		Index(cardIndex).
		Id(strconv.Itoa(int(card.ID))).
		Doc(doc).
		Do(ctx)
	if err != nil {
		return errors.Wrap(err, "update document error")
	}
	return nil
}

func updateCardLabels(card *rdb.Card, savedLabels rdb.CardsLabelSlice, inputLabels []*model.Label, tx *sql.Tx) error {
	beforeLabels := make([]string, len(savedLabels))
	for i, v := range savedLabels {
		beforeLabels[i] = v.R.Label.Name
	}
	afterLabels := make([]string, len(inputLabels))
	for i, v := range inputLabels {
		afterLabels[i] = v.Name
	}

	labelDiff := stringArrayDiff(beforeLabels, afterLabels)

	addLabelIDs := []interface{}{}
	for _, v := range labelDiff.Inc {
		for _, vv := range inputLabels {
			if v == vv.Name {
				addLabelIDs = append(addLabelIDs, vv.ID)
			}
		}
	}

	removeLabelIDs := []interface{}{}
	for _, v := range labelDiff.Dec {
		for _, vv := range savedLabels {
			if v == vv.R.Label.Name {
				removeLabelIDs = append(removeLabelIDs, vv.R.Label.ID)
			}
		}
	}

	ctx := context.Background()

	if len(addLabelIDs) > 0 {
		rels := []*rdb.CardsLabel{}
		for _, v := range addLabelIDs {
			i := &rdb.CardsLabel{
				LabelID: v.(uint),
			}
			rels = append(rels, i)
		}
		if err := card.AddCardsLabels(ctx, tx, true, rels...); err != nil {
			return err
		}
	}

	if len(removeLabelIDs) > 0 {
		if _, err := card.CardsLabels(
			qm.WhereIn("label_id in ?", removeLabelIDs...),
		).DeleteAll(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) MoveCardPosition(id uint, position uint, newSectionID uint) error {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	card, err := rdb.Cards(
		rdb.CardWhere.ID.EQ(id),
	).One(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := sortCards(position, card, newSectionID, tx); err != nil {
		return errors.Wrap(err, "sort card error")
	}

	card.SectionID = newSectionID
	if _, err := card.Update(ctx, tx, boil.Whitelist(rdb.CardColumns.SectionID)); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil

}
func (r *repository) SaveCardImage(cardID uint, url string) error {
	ctx := context.Background()
	row, err := rdb.Cards(
		rdb.CardWhere.ID.EQ(cardID),
	).One(ctx, r.db)
	if err != nil {
		return err
	}
	if err := row.AddCardsImages(ctx, r.db, true, &rdb.CardsImage{
		CardID: cardID,
		URL:    url,
	}); err != nil {
		return err
	}
	return nil
}
func (r *repository) ReorderCardPosition(id uint, position uint) error {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	card, err := rdb.Cards(
		rdb.CardWhere.ID.EQ(id),
	).One(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := sortCards(position, card, card.SectionID, tx); err != nil {
		return errors.Wrap(err, "sort cards error")
	}

	tx.Commit()
	return nil
}

func sortCards(position uint, card *rdb.Card, targetSectionID uint, tx *sql.Tx) error {
	ctx := context.Background()
	// delete
	oldPosition, err := rdb.SectionsCardsPositions(
		rdb.SectionsCardsPositionWhere.CardID.EQ(card.ID),
		rdb.SectionsCardsPositionWhere.SectionID.EQ(card.SectionID),
	).One(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := oldPosition.Delete(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := updateLowerCardPositions(oldPosition.Position, card.SectionID, tx); err != nil {
		return errors.Wrap(err, "update lower card position error")
	}

	if err := updateHigherCardPositions(position, targetSectionID, tx); err != nil {
		return errors.Wrap(err, "update higher card position error")
	}

	newPosition := rdb.SectionsCardsPosition{
		CardID:    card.ID,
		SectionID: targetSectionID,
		Position:  position,
	}

	if err := newPosition.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func updateHigherCardPositions(position uint, sectionID uint, tx *sql.Tx) error {
	ctx := context.Background()
	if _, err := tx.ExecContext(
		ctx,
		"UPDATE sections_cards_positions set position = position+1 WHERE section_id = ? AND position >= ? ORDER BY position DESC;",
		sectionID,
		position,
	); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func updateLowerCardPositions(oldPosition uint, sectionID uint, tx *sql.Tx) error {
	ctx := context.Background()
	if _, err := tx.ExecContext(
		ctx,
		"UPDATE sections_cards_positions set position = position-1 WHERE section_id = ? AND position > ? ORDER BY position;",
		sectionID,
		oldPosition,
	); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (r *repository) DeleteCard(m *model.Card) error {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}
	row, err := rdb.Cards(
		rdb.CardWhere.ID.EQ(m.ID),
		qm.Load(rdb.CardRels.SectionsCardsPosition),
		qm.Load(rdb.CardRels.CardsLabels),
	).One(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err := row.R.SectionsCardsPosition.Delete(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	if _, err := row.R.CardsLabels.DeleteAll(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	if _, err := row.Delete(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *repository) SaveNewCard(m *model.Card) error {
	ctx := context.Background()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	row := rdb.Card{
		Name:        m.Name,
		Description: m.Description,
		SectionID:   m.Section.ID,
		LimitTime:   null.TimeFrom(*m.LimitTime),
	}
	if err := row.Insert(ctx, tx, boil.Whitelist(
		rdb.CardColumns.Name,
		rdb.CardColumns.Description,
		rdb.CardColumns.SectionID,
		rdb.CardColumns.LimitTime,
	)); err != nil {
		tx.Rollback()
		return err
	}

	m.ID = row.ID

	if err := sortForNewCard(row, m.Position, tx); err != nil {
		return errors.Wrap(err, "sort for new card error")
	}

	if err := addCardsLabels(row, m.Labels, tx); err != nil {
		return errors.Wrap(err, "add cards labels error")
	}

	tx.Commit()

	if err := r.addCardDocument(m); err != nil {
		return errors.Wrap(err, "add card document error")
	}

	return nil
}

func sortForNewCard(row rdb.Card, position uint, tx *sql.Tx) error {
	ctx := context.Background()
	if _, err := tx.ExecContext(
		ctx,
		"UPDATE sections_cards_positions set position = position+1 WHERE section_id = ? AND position >= ? ORDER BY position DESC;",
		row.SectionID,
		position,
	); err != nil {
		tx.Rollback()
		return err
	}
	if err := row.SetSectionsCardsPosition(ctx, tx, true, &rdb.SectionsCardsPosition{SectionID: row.SectionID, Position: position}); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func addCardsLabels(card rdb.Card, labels []*model.Label, tx *sql.Tx) error {
	ctx := context.Background()
	rels := make([]*rdb.CardsLabel, len(labels))
	for i, v := range labels {
		value := &rdb.CardsLabel{
			LabelID: v.ID,
		}
		rels[i] = value
	}
	if err := card.AddCardsLabels(ctx, tx, true, rels...); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "link.AddLinksTags error")
	}
	return nil
}

func (r *repository) addCardDocument(m *model.Card) error {
	ctx := context.Background()
	doc := cardDocument{
		Name:        m.Name,
		Description: m.Description,
	}
	_, err := r.esClient.Index().Index(cardIndex).
		Id(strconv.Itoa(int(m.ID))).BodyJson(doc).Do(ctx)
	if err != nil {
		return errors.Wrap(err, "create document error")
	}
	return nil
}
