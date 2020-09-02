package repository

import (
	"context"
	"database/sql"

	"github.com/stobita/plank/internal/model"
	"github.com/stobita/plank/internal/rdb"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

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

func (r *repository) GetBoardLabels(boardID uint) ([]*model.Label, error) {
	ctx := context.Background()
	rows, err := rdb.Labels(
		rdb.LabelWhere.BoardID.EQ(boardID),
	).All(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	m := make([]*model.Label, len(rows))
	for i, v := range rows {
		m[i] = &model.Label{
			ID:   v.ID,
			Name: v.Name,
		}
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

func (r *repository) SaveBoard(m *model.Board) error {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	row, err := rdb.Boards(
		rdb.BoardWhere.ID.EQ(m.ID),
	).One(ctx, tx)
	if err != nil {
		return err
	}
	row.Name = m.Name
	if _, err := row.Update(ctx, tx, boil.Whitelist(
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
