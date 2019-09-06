package repository

import (
	"context"
	"database/sql"

	"github.com/stobita/plank/internal/model"
	"github.com/stobita/plank/internal/rdb"
	"github.com/volatiletech/sqlboiler/boil"
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
			ID:    v.ID,
			Title: v.Title,
		})
	}
	return m, nil
}

func (r *repository) SaveNewLink(m *model.Board) error {
	ctx := context.Background()
	row := rdb.Board{
		Title: m.Title,
	}
	if err := row.Insert(ctx, r.db, boil.Whitelist("title")); err != nil {
		return err
	}
	m.ID = row.ID
	return nil
}
