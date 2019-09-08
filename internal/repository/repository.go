package repository

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/stobita/plank/internal/model"
	"github.com/stobita/plank/internal/rdb"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
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

func (r *repository) GetBoardByID(id uint) (*model.Board, error) {
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

func (r *repository) GetSectionByID(id uint) (*model.Section, error) {
	ctx := context.Background()
	row, err := rdb.Sections(rdb.SectionWhere.ID.EQ(id)).One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.Section{
		ID:   row.ID,
		Name: row.Name,
	}, nil
}

func (r *repository) GetBoardSectionsWithCards(board *model.Board) ([]*model.Section, error) {
	ctx := context.Background()
	rows, err := rdb.Sections(
		rdb.SectionWhere.BoardID.EQ(board.ID),
		qm.Load(rdb.SectionRels.Cards),
	).All(ctx, r.db)
	if err != nil {
		return nil, errors.Wrap(err, "repository: get All error")
	}
	var sections []*model.Section
	for _, row := range rows {
		var cards []*model.Card
		for _, card := range row.R.Cards {
			cards = append(cards, &model.Card{
				Name:        card.Name,
				Description: card.Description,
			})
		}
		sections = append(sections, &model.Section{
			ID:    row.ID,
			Name:  row.Name,
			Cards: cards,
		})
	}
	return sections, nil
}

func (r *repository) SaveNewSection(m *model.Section) error {
	ctx := context.Background()
	row := rdb.Section{
		Name:    m.Name,
		BoardID: m.Board.ID,
	}
	if err := row.Insert(ctx, r.db, boil.Whitelist(
		rdb.SectionColumns.Name,
		rdb.SectionColumns.BoardID,
	)); err != nil {
		return err
	}
	m.ID = row.ID
	return nil
}

func (r *repository) SaveNewCard(m *model.Card) error {
	ctx := context.Background()
	row := rdb.Card{
		Name:        m.Name,
		Description: m.Description,
		SectionID:   m.Section.ID,
	}
	if err := row.Insert(ctx, r.db, boil.Whitelist(
		rdb.CardColumns.Name,
		rdb.CardColumns.Description,
		rdb.CardColumns.SectionID,
	)); err != nil {
		return err
	}
	m.ID = row.ID
	return nil
}
