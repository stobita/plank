package repository

import (
	"context"
	"log"
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

	beforeLabels := make([]string, len(row.R.CardsLabels))
	for i, v := range row.R.CardsLabels {
		beforeLabels[i] = v.R.Label.Name
	}
	log.Println("beforeLabels: ", beforeLabels)
	afterLabels := make([]string, len(m.Labels))
	for i, v := range m.Labels {
		afterLabels[i] = v.Name
	}
	log.Println("afterLabels: ", afterLabels)

	labelDiff := stringArrayDiff(beforeLabels, afterLabels)

	addLabelIDs := []interface{}{}
	for _, v := range labelDiff.Inc {
		for _, vv := range m.Labels {
			if v == vv.Name {
				addLabelIDs = append(addLabelIDs, vv.ID)
			}
		}
	}
	log.Println("addLabelIDs: ", addLabelIDs)
	removeLabelIDs := []interface{}{}
	for _, v := range labelDiff.Dec {
		for _, vv := range row.R.CardsLabels {
			if v == vv.R.Label.Name {
				removeLabelIDs = append(removeLabelIDs, vv.R.Label.ID)
			}
		}
	}
	log.Println("removeLabelIDs: ", removeLabelIDs)

	if len(addLabelIDs) > 0 {
		rels := []*rdb.CardsLabel{}
		for _, v := range addLabelIDs {
			i := &rdb.CardsLabel{
				LabelID: v.(uint),
			}
			rels = append(rels, i)
		}
		if err := row.AddCardsLabels(ctx, tx, true, rels...); err != nil {
			return err
		}
	}

	if len(removeLabelIDs) > 0 {
		if _, err := row.CardsLabels(
			qm.WhereIn("label_id in ?", removeLabelIDs...),
		).DeleteAll(ctx, tx); err != nil {
			return err
		}
	}

	tx.Commit()

	doc := cardDocument{
		Name:        m.Name,
		Description: m.Description,
	}
	_, err = r.esClient.Update().
		Index(cardIndex).
		Id(strconv.Itoa(int(m.ID))).
		Doc(doc).
		Do(ctx)
	if err != nil {
		return errors.Wrap(err, "update document error")
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

	oldPosition, err := rdb.SectionsCardsPositions(
		rdb.SectionsCardsPositionWhere.CardID.EQ(id),
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

	if _, err := tx.ExecContext(
		ctx,
		"UPDATE sections_cards_positions set position = position-1 WHERE section_id = ? AND position > ? ORDER BY position;",
		card.SectionID,
		oldPosition.Position,
	); err != nil {
		tx.Rollback()
		return err
	}

	// higher
	if _, err := tx.ExecContext(
		ctx,
		"UPDATE sections_cards_positions set position = position+1 WHERE section_id = ? AND position >= ? ORDER BY position DESC;",
		newSectionID,
		position,
	); err != nil {
		tx.Rollback()
		return err
	}

	newPosition := rdb.SectionsCardsPosition{
		CardID:    card.ID,
		SectionID: newSectionID,
		Position:  position,
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := newPosition.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return err
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

	// delete
	oldPosition, err := rdb.SectionsCardsPositions(
		rdb.SectionsCardsPositionWhere.CardID.EQ(card.ID),
		rdb.SectionsCardsPositionWhere.SectionID.EQ(card.SectionID),
	).One(ctx, tx)

	if _, err := oldPosition.Delete(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}

	log.Printf("oldPosition.Position: %v", oldPosition.Position)

	if _, err := tx.ExecContext(
		ctx,
		"UPDATE sections_cards_positions set position = position-1 WHERE section_id = ? AND position > ? ORDER BY position;",
		card.SectionID,
		oldPosition.Position,
	); err != nil {
		tx.Rollback()
		return err
	}

	log.Println("add position")

	// higher
	if _, err := tx.ExecContext(
		ctx,
		"UPDATE sections_cards_positions set position = position+1 WHERE section_id = ? AND position >= ? ORDER BY position DESC;",
		card.SectionID,
		position,
	); err != nil {
		tx.Rollback()
		return err
	}

	newPosition := rdb.SectionsCardsPosition{
		CardID:    card.ID,
		SectionID: card.SectionID,
		Position:  position,
	}
	if err := newPosition.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
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

	if _, err := tx.ExecContext(
		ctx,
		"UPDATE sections_cards_positions set position = position+1 WHERE section_id = ? AND position >= ? ORDER BY position DESC;",
		row.SectionID,
		m.Position,
	); err != nil {
		tx.Rollback()
		return err
	}

	if err := row.SetSectionsCardsPosition(ctx, tx, true, &rdb.SectionsCardsPosition{SectionID: row.SectionID, Position: m.Position}); err != nil {
		tx.Rollback()
		return err
	}

	rels := make([]*rdb.CardsLabel, len(m.Labels))
	for i, v := range m.Labels {
		value := &rdb.CardsLabel{
			LabelID: v.ID,
		}
		rels[i] = value
	}
	if err := row.AddCardsLabels(ctx, tx, true, rels...); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "link.AddLinksTags error")
	}

	tx.Commit()

	doc := cardDocument{
		Name:        m.Name,
		Description: m.Description,
	}
	_, err = r.esClient.Index().Index(cardIndex).
		Id(strconv.Itoa(int(m.ID))).BodyJson(doc).Do(ctx)
	if err != nil {
		return errors.Wrap(err, "create document error")
	}

	return nil
}
