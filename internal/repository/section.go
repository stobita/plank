package repository

import (
	"context"
	"database/sql"
	"log"
	"sort"
	"strconv"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/stobita/plank/internal/model"
	"github.com/stobita/plank/internal/rdb"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

const cardIndex = "card"

type cardDocument struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *repository) GetSection(id uint) (*model.Section, error) {
	ctx := context.Background()
	row, err := rdb.Sections(
		rdb.SectionWhere.ID.EQ(id),
		qm.Load(rdb.SectionRels.Board),
	).One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.Section{
		ID:   row.ID,
		Name: row.Name,
		Board: &model.Board{
			ID: row.R.Board.ID,
		},
	}, nil
}

func (r *repository) convertSections(rows rdb.SectionSlice) ([]*model.Section, error) {
	var sections []*model.Section
	for _, row := range rows {
		var cards []*model.Card
		for _, card := range row.R.Cards {
			labels := make([]*model.Label, len(card.R.CardsLabels))
			for i, v := range card.R.CardsLabels {
				labels[i] = &model.Label{
					ID:   v.R.Label.ID,
					Name: v.R.Label.Name,
				}
			}
			images := make([]string, len(card.R.CardsImages))
			log.Printf("image: %v", card.R.CardsImages)
			for i, v := range card.R.CardsImages {
				images[i] = v.URL
			}
			cards = append(cards, &model.Card{
				ID:          card.ID,
				Name:        card.Name,
				Description: card.Description,
				Position:    card.R.SectionsCardsPosition.Position,
				Labels:      labels,
				LimitTime:   &card.LimitTime.Time,
				Images:      images,
			})
		}
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Position < cards[j].Position
		})
		sections = append(sections, &model.Section{
			ID:    row.ID,
			Name:  row.Name,
			Cards: cards,
			Board: &model.Board{
				ID: row.R.Board.ID,
			},
			Position: row.R.BoardsSectionsPosition.Position,
		})
	}
	sort.Slice(sections, func(i, j int) bool {
		return sections[i].Position < sections[j].Position
	})
	return sections, nil
}

func (r *repository) GetBoardSectionsWithCards(board *model.Board) ([]*model.Section, error) {
	ctx := context.Background()
	rows, err := rdb.Sections(
		rdb.SectionWhere.BoardID.EQ(board.ID),
		qm.Load(rdb.SectionRels.Cards),
		qm.Load(rdb.SectionRels.BoardsSectionsPosition),
		qm.Load(
			qm.Rels(
				rdb.SectionRels.Cards,
				rdb.CardRels.SectionsCardsPosition,
			),
		),
		qm.Load(
			qm.Rels(
				rdb.SectionRels.Cards,
				rdb.CardRels.CardsLabels,
				rdb.CardsLabelRels.Label,
			),
		),
		qm.Load(
			qm.Rels(
				rdb.SectionRels.Cards,
				rdb.CardRels.CardsImages,
			),
		),
		qm.Load(rdb.SectionRels.Board),
	).All(ctx, r.db)
	if err != nil {
		return nil, errors.Wrap(err, "repository: get All error")
	}
	return r.convertSections(rows)
}
func (r *repository) SearchBoardSectionsWithCards(board *model.Board, word string) ([]*model.Section, error) {
	ctx := context.Background()
	query := elastic.NewMultiMatchQuery(word, "name", "description")
	result, err := r.esClient.Search().
		Index(cardIndex).
		Query(query).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	var ids []interface{}
	if result.Hits.TotalHits.Value > 0 {
		for _, hit := range result.Hits.Hits {
			id, err := strconv.Atoi(hit.Id)
			if err != nil {
				return nil, err
			}
			ids = append(ids, id)
		}
	}
	if len(ids) < 1 {
		return nil, errors.New("not match word")
	}
	rows, err := rdb.Sections(
		rdb.SectionWhere.BoardID.EQ(board.ID),
		qm.Load(
			rdb.SectionRels.Cards,
			qm.WhereIn("cards.id in ?", ids...),
		),
		qm.Load(rdb.SectionRels.BoardsSectionsPosition),
		qm.Load(
			qm.Rels(
				rdb.SectionRels.Cards,
				rdb.CardRels.SectionsCardsPosition,
			),
		),
		qm.Load(
			qm.Rels(
				rdb.SectionRels.Cards,
				rdb.CardRels.CardsLabels,
				rdb.CardsLabelRels.Label,
			),
		),
		qm.Load(
			qm.Rels(
				rdb.SectionRels.Cards,
				rdb.CardRels.CardsImages,
			),
		),
		qm.Load(rdb.SectionRels.Board),
	).All(ctx, r.db)
	if err != nil {
		return nil, errors.Wrap(err, "repository: get All error")
	}
	return r.convertSections(rows)
}

func (r *repository) SaveNewSection(m *model.Section) error {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	row := rdb.Section{
		Name:    m.Name,
		BoardID: m.Board.ID,
	}
	if err := row.Insert(ctx, tx, boil.Whitelist(
		rdb.SectionColumns.Name,
		rdb.SectionColumns.BoardID,
	)); err != nil {
		return err
	}
	m.ID = row.ID

	if _, err := tx.ExecContext(
		ctx,
		"UPDATE boards_sections_positions set position = position+1 WHERE board_id = ? AND position >= ? ORDER BY position DESC;",
		row.BoardID,
		m.Position,
	); err != nil {
		tx.Rollback()
		return err
	}

	if err := row.SetBoardsSectionsPosition(ctx, tx, true, &rdb.BoardsSectionsPosition{BoardID: row.BoardID, Position: m.Position}); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *repository) SaveSection(m *model.Section) error {
	ctx := context.Background()
	row, err := rdb.Sections(
		rdb.SectionWhere.ID.EQ(m.ID),
	).One(ctx, r.db)
	if err != nil {
		return err
	}
	row.Name = m.Name
	if _, err := row.Update(ctx, r.db, boil.Whitelist(
		rdb.SectionColumns.Name,
	)); err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteSection(m *model.Section) error {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	row, err := rdb.Sections(
		rdb.SectionWhere.ID.EQ(m.ID),
		qm.Load(rdb.SectionRels.Cards),
		qm.Load(rdb.SectionRels.SectionsCardsPositions),
	).One(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err := row.R.SectionsCardsPositions.DeleteAll(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	if _, err := row.R.Cards.DeleteAll(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	if _, err := row.Delete(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	_, err = r.esClient.Delete().
		Index(cardIndex).
		Id(strconv.Itoa(int(m.ID))).
		Do(ctx)
	if err != nil {
		return errors.Wrap(err, "update document error")
	}
	return nil
}
func (r *repository) ReorderSectionPosition(id uint, position uint) error {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed transaction begin")
	}
	section, err := rdb.Sections(
		rdb.SectionWhere.ID.EQ(id),
	).One(ctx, tx)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed get section")
	}

	oldPosition, err := rdb.BoardsSectionsPositions(
		rdb.BoardsSectionsPositionWhere.SectionID.EQ(section.ID),
		rdb.BoardsSectionsPositionWhere.BoardID.EQ(section.BoardID),
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
		"UPDATE boards_sections_positions set position = position-1 WHERE board_id = ? AND position > ? ORDER BY position;",
		section.BoardID,
		oldPosition.Position,
	); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.ExecContext(
		ctx,
		"UPDATE boards_sections_positions set position = position+1 WHERE board_id = ? AND position >= ? ORDER BY position DESC;",
		section.BoardID,
		position,
	); err != nil {
		tx.Rollback()
		return err
	}

	newPosition := rdb.BoardsSectionsPosition{
		SectionID: section.ID,
		BoardID:   section.BoardID,
		Position:  position,
	}
	if err := newPosition.Insert(ctx, tx, boil.Infer()); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
