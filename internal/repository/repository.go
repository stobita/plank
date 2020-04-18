package repository

import (
	"context"
	"database/sql"
	"log"
	"os"
	"sort"

	"github.com/pkg/errors"
	"github.com/stobita/plank/internal/model"
	"github.com/stobita/plank/internal/rdb"
	"github.com/stobita/plank/internal/usecase"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type repository struct {
	db *sql.DB
}

// New ...
func New(db *sql.DB) usecase.Repository {
	if os.Getenv("PRODUCTION") != "true" {
		boil.DebugMode = true
	}
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllBoards() ([]*model.Board, error) {
	ctx := context.Background()
	boardRows, err := rdb.Boards().All(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var m []*model.Board
	for _, v := range boardRows {
		m = append(m, &model.Board{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return m, nil
}

func (r *repository) SaveNewBoard(m *model.Board) error {
	ctx := context.Background()
	row := rdb.Board{
		Name: m.Name,
	}
	if err := row.Insert(ctx, r.db, boil.Whitelist(rdb.BoardColumns.Name)); err != nil {
		return err
	}
	m.ID = row.ID
	return nil
}

func (r *repository) GetBoard(id uint) (*model.Board, error) {
	ctx := context.Background()
	row, err := rdb.Boards(rdb.BoardWhere.ID.EQ(id)).One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.Board{
		ID:   row.ID,
		Name: row.Name,
	}, nil
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
		qm.Load(rdb.SectionRels.Board),
	).All(ctx, r.db)
	if err != nil {
		return nil, errors.Wrap(err, "repository: get All error")
	}
	var sections []*model.Section
	for _, row := range rows {
		var cards []*model.Card
		for _, card := range row.R.Cards {
			cards = append(cards, &model.Card{
				ID:          card.ID,
				Name:        card.Name,
				Description: card.Description,
				Position:    card.R.SectionsCardsPosition.Position,
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
	}
	if err := row.Insert(ctx, tx, boil.Whitelist(
		rdb.CardColumns.Name,
		rdb.CardColumns.Description,
		rdb.CardColumns.SectionID,
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

	tx.Commit()

	return nil
}

func (r *repository) GetCard(id uint) (*model.Card, error) {
	ctx := context.Background()
	row, err := rdb.Cards(
		rdb.CardWhere.ID.EQ(id),
		qm.Load(rdb.CardRels.Section),
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
		},
	}, nil
}

func (r *repository) SaveCard(m *model.Card) error {
	ctx := context.Background()
	row, err := rdb.Cards(
		rdb.CardWhere.ID.EQ(m.ID),
	).One(ctx, r.db)
	if err != nil {
		return err
	}

	row.Name = m.Name
	row.Description = m.Description

	if _, err := row.Update(ctx, r.db, boil.Whitelist(
		rdb.CardColumns.Name,
		rdb.CardColumns.Description,
	)); err != nil {
		return err
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
	row, err := rdb.Cards(
		rdb.CardWhere.ID.EQ(m.ID),
		qm.Load(rdb.CardRels.SectionsCardsPosition),
	).One(ctx, r.db)
	if err != nil {
		return err
	}
	if _, err := row.R.SectionsCardsPosition.Delete(ctx, r.db); err != nil {
		return err
	}
	if _, err := row.Delete(ctx, r.db); err != nil {
		return err
	}
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
	return nil
}

func (r *repository) SaveBoard(m *model.Board) error {
	ctx := context.Background()
	row, err := rdb.Boards(
		rdb.BoardWhere.ID.EQ(m.ID),
	).One(ctx, r.db)
	if err != nil {
		return err
	}
	row.Name = m.Name
	if _, err := row.Update(ctx, r.db, boil.Whitelist(
		rdb.BoardColumns.Name,
	)); err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteBoard(m *model.Board) error {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	row, err := rdb.Boards(
		rdb.BoardWhere.ID.EQ(m.ID),
		qm.Load(
			qm.Rels(
				rdb.BoardRels.Sections,
				rdb.SectionRels.Cards,
			),
		),
		qm.Load(
			qm.Rels(
				rdb.BoardRels.Sections,
				rdb.SectionRels.SectionsCardsPositions,
			),
		),
	).One(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, v := range row.R.Sections {
		if _, err := v.R.SectionsCardsPositions.DeleteAll(ctx, tx); err != nil {
			tx.Rollback()
			return err
		}
		if _, err := v.R.Cards.DeleteAll(ctx, tx); err != nil {
			tx.Rollback()
			return err
		}
	}
	if _, err := row.R.Sections.DeleteAll(ctx, tx); err != nil {
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
