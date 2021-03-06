// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package rdb

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// CardsLabel is an object representing the database table.
type CardsLabel struct {
	ID        uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CardID    uint      `boil:"card_id" json:"card_id" toml:"card_id" yaml:"card_id"`
	LabelID   uint      `boil:"label_id" json:"label_id" toml:"label_id" yaml:"label_id"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *cardsLabelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cardsLabelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CardsLabelColumns = struct {
	ID        string
	CardID    string
	LabelID   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	CardID:    "card_id",
	LabelID:   "label_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

var CardsLabelWhere = struct {
	ID        whereHelperuint
	CardID    whereHelperuint
	LabelID   whereHelperuint
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperuint{field: "`cards_labels`.`id`"},
	CardID:    whereHelperuint{field: "`cards_labels`.`card_id`"},
	LabelID:   whereHelperuint{field: "`cards_labels`.`label_id`"},
	CreatedAt: whereHelpernull_Time{field: "`cards_labels`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`cards_labels`.`updated_at`"},
}

// CardsLabelRels is where relationship names are stored.
var CardsLabelRels = struct {
	Card  string
	Label string
}{
	Card:  "Card",
	Label: "Label",
}

// cardsLabelR is where relationships are stored.
type cardsLabelR struct {
	Card  *Card
	Label *Label
}

// NewStruct creates a new relationship struct
func (*cardsLabelR) NewStruct() *cardsLabelR {
	return &cardsLabelR{}
}

// cardsLabelL is where Load methods for each relationship are stored.
type cardsLabelL struct{}

var (
	cardsLabelAllColumns            = []string{"id", "card_id", "label_id", "created_at", "updated_at"}
	cardsLabelColumnsWithoutDefault = []string{"card_id", "label_id"}
	cardsLabelColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	cardsLabelPrimaryKeyColumns     = []string{"id"}
)

type (
	// CardsLabelSlice is an alias for a slice of pointers to CardsLabel.
	// This should generally be used opposed to []CardsLabel.
	CardsLabelSlice []*CardsLabel
	// CardsLabelHook is the signature for custom CardsLabel hook methods
	CardsLabelHook func(context.Context, boil.ContextExecutor, *CardsLabel) error

	cardsLabelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cardsLabelType                 = reflect.TypeOf(&CardsLabel{})
	cardsLabelMapping              = queries.MakeStructMapping(cardsLabelType)
	cardsLabelPrimaryKeyMapping, _ = queries.BindMapping(cardsLabelType, cardsLabelMapping, cardsLabelPrimaryKeyColumns)
	cardsLabelInsertCacheMut       sync.RWMutex
	cardsLabelInsertCache          = make(map[string]insertCache)
	cardsLabelUpdateCacheMut       sync.RWMutex
	cardsLabelUpdateCache          = make(map[string]updateCache)
	cardsLabelUpsertCacheMut       sync.RWMutex
	cardsLabelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cardsLabelBeforeInsertHooks []CardsLabelHook
var cardsLabelBeforeUpdateHooks []CardsLabelHook
var cardsLabelBeforeDeleteHooks []CardsLabelHook
var cardsLabelBeforeUpsertHooks []CardsLabelHook

var cardsLabelAfterInsertHooks []CardsLabelHook
var cardsLabelAfterSelectHooks []CardsLabelHook
var cardsLabelAfterUpdateHooks []CardsLabelHook
var cardsLabelAfterDeleteHooks []CardsLabelHook
var cardsLabelAfterUpsertHooks []CardsLabelHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CardsLabel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsLabelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CardsLabel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsLabelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CardsLabel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsLabelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CardsLabel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsLabelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CardsLabel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsLabelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CardsLabel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsLabelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CardsLabel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsLabelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CardsLabel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsLabelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CardsLabel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsLabelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCardsLabelHook registers your hook function for all future operations.
func AddCardsLabelHook(hookPoint boil.HookPoint, cardsLabelHook CardsLabelHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cardsLabelBeforeInsertHooks = append(cardsLabelBeforeInsertHooks, cardsLabelHook)
	case boil.BeforeUpdateHook:
		cardsLabelBeforeUpdateHooks = append(cardsLabelBeforeUpdateHooks, cardsLabelHook)
	case boil.BeforeDeleteHook:
		cardsLabelBeforeDeleteHooks = append(cardsLabelBeforeDeleteHooks, cardsLabelHook)
	case boil.BeforeUpsertHook:
		cardsLabelBeforeUpsertHooks = append(cardsLabelBeforeUpsertHooks, cardsLabelHook)
	case boil.AfterInsertHook:
		cardsLabelAfterInsertHooks = append(cardsLabelAfterInsertHooks, cardsLabelHook)
	case boil.AfterSelectHook:
		cardsLabelAfterSelectHooks = append(cardsLabelAfterSelectHooks, cardsLabelHook)
	case boil.AfterUpdateHook:
		cardsLabelAfterUpdateHooks = append(cardsLabelAfterUpdateHooks, cardsLabelHook)
	case boil.AfterDeleteHook:
		cardsLabelAfterDeleteHooks = append(cardsLabelAfterDeleteHooks, cardsLabelHook)
	case boil.AfterUpsertHook:
		cardsLabelAfterUpsertHooks = append(cardsLabelAfterUpsertHooks, cardsLabelHook)
	}
}

// One returns a single cardsLabel record from the query.
func (q cardsLabelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CardsLabel, error) {
	o := &CardsLabel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rdb: failed to execute a one query for cards_labels")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CardsLabel records from the query.
func (q cardsLabelQuery) All(ctx context.Context, exec boil.ContextExecutor) (CardsLabelSlice, error) {
	var o []*CardsLabel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "rdb: failed to assign all query results to CardsLabel slice")
	}

	if len(cardsLabelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CardsLabel records in the query.
func (q cardsLabelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to count cards_labels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cardsLabelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "rdb: failed to check if cards_labels exists")
	}

	return count > 0, nil
}

// Card pointed to by the foreign key.
func (o *CardsLabel) Card(mods ...qm.QueryMod) cardQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.CardID),
	}

	queryMods = append(queryMods, mods...)

	query := Cards(queryMods...)
	queries.SetFrom(query.Query, "`cards`")

	return query
}

// Label pointed to by the foreign key.
func (o *CardsLabel) Label(mods ...qm.QueryMod) labelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.LabelID),
	}

	queryMods = append(queryMods, mods...)

	query := Labels(queryMods...)
	queries.SetFrom(query.Query, "`labels`")

	return query
}

// LoadCard allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (cardsLabelL) LoadCard(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCardsLabel interface{}, mods queries.Applicator) error {
	var slice []*CardsLabel
	var object *CardsLabel

	if singular {
		object = maybeCardsLabel.(*CardsLabel)
	} else {
		slice = *maybeCardsLabel.(*[]*CardsLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &cardsLabelR{}
		}
		args = append(args, object.CardID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &cardsLabelR{}
			}

			for _, a := range args {
				if a == obj.CardID {
					continue Outer
				}
			}

			args = append(args, obj.CardID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`cards`), qm.WhereIn(`cards.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Card")
	}

	var resultSlice []*Card
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Card")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for cards")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for cards")
	}

	if len(cardsLabelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Card = foreign
		if foreign.R == nil {
			foreign.R = &cardR{}
		}
		foreign.R.CardsLabels = append(foreign.R.CardsLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CardID == foreign.ID {
				local.R.Card = foreign
				if foreign.R == nil {
					foreign.R = &cardR{}
				}
				foreign.R.CardsLabels = append(foreign.R.CardsLabels, local)
				break
			}
		}
	}

	return nil
}

// LoadLabel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (cardsLabelL) LoadLabel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCardsLabel interface{}, mods queries.Applicator) error {
	var slice []*CardsLabel
	var object *CardsLabel

	if singular {
		object = maybeCardsLabel.(*CardsLabel)
	} else {
		slice = *maybeCardsLabel.(*[]*CardsLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &cardsLabelR{}
		}
		args = append(args, object.LabelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &cardsLabelR{}
			}

			for _, a := range args {
				if a == obj.LabelID {
					continue Outer
				}
			}

			args = append(args, obj.LabelID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`labels`), qm.WhereIn(`labels.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Label")
	}

	var resultSlice []*Label
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Label")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for labels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for labels")
	}

	if len(cardsLabelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Label = foreign
		if foreign.R == nil {
			foreign.R = &labelR{}
		}
		foreign.R.CardsLabels = append(foreign.R.CardsLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LabelID == foreign.ID {
				local.R.Label = foreign
				if foreign.R == nil {
					foreign.R = &labelR{}
				}
				foreign.R.CardsLabels = append(foreign.R.CardsLabels, local)
				break
			}
		}
	}

	return nil
}

// SetCard of the cardsLabel to the related item.
// Sets o.R.Card to related.
// Adds o to related.R.CardsLabels.
func (o *CardsLabel) SetCard(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Card) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `cards_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"card_id"}),
		strmangle.WhereClause("`", "`", 0, cardsLabelPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CardID = related.ID
	if o.R == nil {
		o.R = &cardsLabelR{
			Card: related,
		}
	} else {
		o.R.Card = related
	}

	if related.R == nil {
		related.R = &cardR{
			CardsLabels: CardsLabelSlice{o},
		}
	} else {
		related.R.CardsLabels = append(related.R.CardsLabels, o)
	}

	return nil
}

// SetLabel of the cardsLabel to the related item.
// Sets o.R.Label to related.
// Adds o to related.R.CardsLabels.
func (o *CardsLabel) SetLabel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Label) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `cards_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"label_id"}),
		strmangle.WhereClause("`", "`", 0, cardsLabelPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.LabelID = related.ID
	if o.R == nil {
		o.R = &cardsLabelR{
			Label: related,
		}
	} else {
		o.R.Label = related
	}

	if related.R == nil {
		related.R = &labelR{
			CardsLabels: CardsLabelSlice{o},
		}
	} else {
		related.R.CardsLabels = append(related.R.CardsLabels, o)
	}

	return nil
}

// CardsLabels retrieves all the records using an executor.
func CardsLabels(mods ...qm.QueryMod) cardsLabelQuery {
	mods = append(mods, qm.From("`cards_labels`"))
	return cardsLabelQuery{NewQuery(mods...)}
}

// FindCardsLabel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCardsLabel(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*CardsLabel, error) {
	cardsLabelObj := &CardsLabel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cards_labels` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cardsLabelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rdb: unable to select from cards_labels")
	}

	return cardsLabelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CardsLabel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("rdb: no cards_labels provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cardsLabelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cardsLabelInsertCacheMut.RLock()
	cache, cached := cardsLabelInsertCache[key]
	cardsLabelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cardsLabelAllColumns,
			cardsLabelColumnsWithDefault,
			cardsLabelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cardsLabelType, cardsLabelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cardsLabelType, cardsLabelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cards_labels` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cards_labels` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cards_labels` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cardsLabelPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "rdb: unable to insert into cards_labels")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cardsLabelMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "rdb: unable to populate default values for cards_labels")
	}

CacheNoHooks:
	if !cached {
		cardsLabelInsertCacheMut.Lock()
		cardsLabelInsertCache[key] = cache
		cardsLabelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CardsLabel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CardsLabel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cardsLabelUpdateCacheMut.RLock()
	cache, cached := cardsLabelUpdateCache[key]
	cardsLabelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cardsLabelAllColumns,
			cardsLabelPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("rdb: unable to update cards_labels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cards_labels` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cardsLabelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cardsLabelType, cardsLabelMapping, append(wl, cardsLabelPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to update cards_labels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to get rows affected by update for cards_labels")
	}

	if !cached {
		cardsLabelUpdateCacheMut.Lock()
		cardsLabelUpdateCache[key] = cache
		cardsLabelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cardsLabelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to update all for cards_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to retrieve rows affected for cards_labels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CardsLabelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("rdb: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cardsLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cards_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cardsLabelPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to update all in cardsLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to retrieve rows affected all in update all cardsLabel")
	}
	return rowsAff, nil
}

var mySQLCardsLabelUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CardsLabel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("rdb: no cards_labels provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cardsLabelColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCardsLabelUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	cardsLabelUpsertCacheMut.RLock()
	cache, cached := cardsLabelUpsertCache[key]
	cardsLabelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cardsLabelAllColumns,
			cardsLabelColumnsWithDefault,
			cardsLabelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cardsLabelAllColumns,
			cardsLabelPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("rdb: unable to upsert cards_labels, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "cards_labels", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cards_labels` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cardsLabelType, cardsLabelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cardsLabelType, cardsLabelMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "rdb: unable to upsert for cards_labels")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cardsLabelMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cardsLabelType, cardsLabelMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "rdb: unable to retrieve unique values for cards_labels")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "rdb: unable to populate default values for cards_labels")
	}

CacheNoHooks:
	if !cached {
		cardsLabelUpsertCacheMut.Lock()
		cardsLabelUpsertCache[key] = cache
		cardsLabelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CardsLabel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CardsLabel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("rdb: no CardsLabel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cardsLabelPrimaryKeyMapping)
	sql := "DELETE FROM `cards_labels` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to delete from cards_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to get rows affected by delete for cards_labels")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cardsLabelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("rdb: no cardsLabelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to delete all from cards_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to get rows affected by deleteall for cards_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CardsLabelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cardsLabelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cardsLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cards_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cardsLabelPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to delete all from cardsLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to get rows affected by deleteall for cards_labels")
	}

	if len(cardsLabelAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CardsLabel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCardsLabel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CardsLabelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CardsLabelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cardsLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cards_labels`.* FROM `cards_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cardsLabelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "rdb: unable to reload all in CardsLabelSlice")
	}

	*o = slice

	return nil
}

// CardsLabelExists checks if the CardsLabel row exists.
func CardsLabelExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cards_labels` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "rdb: unable to check if cards_labels exists")
	}

	return exists, nil
}
