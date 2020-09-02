package repository

import (
	"context"

	"github.com/stobita/plank/internal/model"
	"github.com/stobita/plank/internal/rdb"
	"github.com/volatiletech/sqlboiler/boil"
)

func (r *repository) SaveNewLabel(m *model.Label) error {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}
	row := rdb.Label{
		Name:    m.Name,
		BoardID: m.BoardID,
	}
	if err := row.Insert(ctx, tx, boil.Whitelist(
		rdb.LabelColumns.Name,
		rdb.LabelColumns.BoardID,
	)); err != nil {
		tx.Rollback()
		return err
	}
	m.ID = row.ID
	tx.Commit()
	return nil
}

func (r *repository) GetLabelByName(name string) (*model.Label, error) {
	ctx := context.Background()
	label, err := rdb.Labels(
		rdb.LabelWhere.Name.EQ(name),
	).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return &model.Label{
		ID:   label.ID,
		Name: label.Name,
	}, nil
}

func (r *repository) GetLabel(id uint) (*model.Label, error) {
	ctx := context.Background()
	label, err := rdb.Labels(
		rdb.LabelWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return &model.Label{
		ID:   label.ID,
		Name: label.Name,
	}, nil
}
