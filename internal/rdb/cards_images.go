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

// CardsImage is an object representing the database table.
type CardsImage struct {
	ID        uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CardID    uint      `boil:"card_id" json:"card_id" toml:"card_id" yaml:"card_id"`
	URL       string    `boil:"url" json:"url" toml:"url" yaml:"url"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *cardsImageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cardsImageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CardsImageColumns = struct {
	ID        string
	CardID    string
	URL       string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	CardID:    "card_id",
	URL:       "url",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

var CardsImageWhere = struct {
	ID        whereHelperuint
	CardID    whereHelperuint
	URL       whereHelperstring
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperuint{field: "`cards_images`.`id`"},
	CardID:    whereHelperuint{field: "`cards_images`.`card_id`"},
	URL:       whereHelperstring{field: "`cards_images`.`url`"},
	CreatedAt: whereHelpernull_Time{field: "`cards_images`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`cards_images`.`updated_at`"},
}

// CardsImageRels is where relationship names are stored.
var CardsImageRels = struct {
	Card string
}{
	Card: "Card",
}

// cardsImageR is where relationships are stored.
type cardsImageR struct {
	Card *Card
}

// NewStruct creates a new relationship struct
func (*cardsImageR) NewStruct() *cardsImageR {
	return &cardsImageR{}
}

// cardsImageL is where Load methods for each relationship are stored.
type cardsImageL struct{}

var (
	cardsImageAllColumns            = []string{"id", "card_id", "url", "created_at", "updated_at"}
	cardsImageColumnsWithoutDefault = []string{"card_id", "url"}
	cardsImageColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	cardsImagePrimaryKeyColumns     = []string{"id"}
)

type (
	// CardsImageSlice is an alias for a slice of pointers to CardsImage.
	// This should generally be used opposed to []CardsImage.
	CardsImageSlice []*CardsImage
	// CardsImageHook is the signature for custom CardsImage hook methods
	CardsImageHook func(context.Context, boil.ContextExecutor, *CardsImage) error

	cardsImageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cardsImageType                 = reflect.TypeOf(&CardsImage{})
	cardsImageMapping              = queries.MakeStructMapping(cardsImageType)
	cardsImagePrimaryKeyMapping, _ = queries.BindMapping(cardsImageType, cardsImageMapping, cardsImagePrimaryKeyColumns)
	cardsImageInsertCacheMut       sync.RWMutex
	cardsImageInsertCache          = make(map[string]insertCache)
	cardsImageUpdateCacheMut       sync.RWMutex
	cardsImageUpdateCache          = make(map[string]updateCache)
	cardsImageUpsertCacheMut       sync.RWMutex
	cardsImageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var cardsImageBeforeInsertHooks []CardsImageHook
var cardsImageBeforeUpdateHooks []CardsImageHook
var cardsImageBeforeDeleteHooks []CardsImageHook
var cardsImageBeforeUpsertHooks []CardsImageHook

var cardsImageAfterInsertHooks []CardsImageHook
var cardsImageAfterSelectHooks []CardsImageHook
var cardsImageAfterUpdateHooks []CardsImageHook
var cardsImageAfterDeleteHooks []CardsImageHook
var cardsImageAfterUpsertHooks []CardsImageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CardsImage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsImageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CardsImage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsImageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CardsImage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsImageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CardsImage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsImageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CardsImage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsImageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CardsImage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsImageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CardsImage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsImageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CardsImage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsImageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CardsImage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range cardsImageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCardsImageHook registers your hook function for all future operations.
func AddCardsImageHook(hookPoint boil.HookPoint, cardsImageHook CardsImageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cardsImageBeforeInsertHooks = append(cardsImageBeforeInsertHooks, cardsImageHook)
	case boil.BeforeUpdateHook:
		cardsImageBeforeUpdateHooks = append(cardsImageBeforeUpdateHooks, cardsImageHook)
	case boil.BeforeDeleteHook:
		cardsImageBeforeDeleteHooks = append(cardsImageBeforeDeleteHooks, cardsImageHook)
	case boil.BeforeUpsertHook:
		cardsImageBeforeUpsertHooks = append(cardsImageBeforeUpsertHooks, cardsImageHook)
	case boil.AfterInsertHook:
		cardsImageAfterInsertHooks = append(cardsImageAfterInsertHooks, cardsImageHook)
	case boil.AfterSelectHook:
		cardsImageAfterSelectHooks = append(cardsImageAfterSelectHooks, cardsImageHook)
	case boil.AfterUpdateHook:
		cardsImageAfterUpdateHooks = append(cardsImageAfterUpdateHooks, cardsImageHook)
	case boil.AfterDeleteHook:
		cardsImageAfterDeleteHooks = append(cardsImageAfterDeleteHooks, cardsImageHook)
	case boil.AfterUpsertHook:
		cardsImageAfterUpsertHooks = append(cardsImageAfterUpsertHooks, cardsImageHook)
	}
}

// One returns a single cardsImage record from the query.
func (q cardsImageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CardsImage, error) {
	o := &CardsImage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rdb: failed to execute a one query for cards_images")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CardsImage records from the query.
func (q cardsImageQuery) All(ctx context.Context, exec boil.ContextExecutor) (CardsImageSlice, error) {
	var o []*CardsImage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "rdb: failed to assign all query results to CardsImage slice")
	}

	if len(cardsImageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CardsImage records in the query.
func (q cardsImageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to count cards_images rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cardsImageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "rdb: failed to check if cards_images exists")
	}

	return count > 0, nil
}

// Card pointed to by the foreign key.
func (o *CardsImage) Card(mods ...qm.QueryMod) cardQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.CardID),
	}

	queryMods = append(queryMods, mods...)

	query := Cards(queryMods...)
	queries.SetFrom(query.Query, "`cards`")

	return query
}

// LoadCard allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (cardsImageL) LoadCard(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCardsImage interface{}, mods queries.Applicator) error {
	var slice []*CardsImage
	var object *CardsImage

	if singular {
		object = maybeCardsImage.(*CardsImage)
	} else {
		slice = *maybeCardsImage.(*[]*CardsImage)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &cardsImageR{}
		}
		args = append(args, object.CardID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &cardsImageR{}
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

	if len(cardsImageAfterSelectHooks) != 0 {
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
		foreign.R.CardsImages = append(foreign.R.CardsImages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CardID == foreign.ID {
				local.R.Card = foreign
				if foreign.R == nil {
					foreign.R = &cardR{}
				}
				foreign.R.CardsImages = append(foreign.R.CardsImages, local)
				break
			}
		}
	}

	return nil
}

// SetCard of the cardsImage to the related item.
// Sets o.R.Card to related.
// Adds o to related.R.CardsImages.
func (o *CardsImage) SetCard(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Card) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `cards_images` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"card_id"}),
		strmangle.WhereClause("`", "`", 0, cardsImagePrimaryKeyColumns),
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
		o.R = &cardsImageR{
			Card: related,
		}
	} else {
		o.R.Card = related
	}

	if related.R == nil {
		related.R = &cardR{
			CardsImages: CardsImageSlice{o},
		}
	} else {
		related.R.CardsImages = append(related.R.CardsImages, o)
	}

	return nil
}

// CardsImages retrieves all the records using an executor.
func CardsImages(mods ...qm.QueryMod) cardsImageQuery {
	mods = append(mods, qm.From("`cards_images`"))
	return cardsImageQuery{NewQuery(mods...)}
}

// FindCardsImage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCardsImage(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*CardsImage, error) {
	cardsImageObj := &CardsImage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cards_images` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, cardsImageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rdb: unable to select from cards_images")
	}

	return cardsImageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CardsImage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("rdb: no cards_images provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cardsImageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cardsImageInsertCacheMut.RLock()
	cache, cached := cardsImageInsertCache[key]
	cardsImageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cardsImageAllColumns,
			cardsImageColumnsWithDefault,
			cardsImageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cardsImageType, cardsImageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cardsImageType, cardsImageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cards_images` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cards_images` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cards_images` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, cardsImagePrimaryKeyColumns))
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
		return errors.Wrap(err, "rdb: unable to insert into cards_images")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cardsImageMapping["ID"] {
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
		return errors.Wrap(err, "rdb: unable to populate default values for cards_images")
	}

CacheNoHooks:
	if !cached {
		cardsImageInsertCacheMut.Lock()
		cardsImageInsertCache[key] = cache
		cardsImageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CardsImage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CardsImage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	cardsImageUpdateCacheMut.RLock()
	cache, cached := cardsImageUpdateCache[key]
	cardsImageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cardsImageAllColumns,
			cardsImagePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("rdb: unable to update cards_images, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cards_images` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, cardsImagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cardsImageType, cardsImageMapping, append(wl, cardsImagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "rdb: unable to update cards_images row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to get rows affected by update for cards_images")
	}

	if !cached {
		cardsImageUpdateCacheMut.Lock()
		cardsImageUpdateCache[key] = cache
		cardsImageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q cardsImageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to update all for cards_images")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to retrieve rows affected for cards_images")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CardsImageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cardsImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cards_images` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cardsImagePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to update all in cardsImage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to retrieve rows affected all in update all cardsImage")
	}
	return rowsAff, nil
}

var mySQLCardsImageUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CardsImage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("rdb: no cards_images provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cardsImageColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCardsImageUniqueColumns, o)

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

	cardsImageUpsertCacheMut.RLock()
	cache, cached := cardsImageUpsertCache[key]
	cardsImageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cardsImageAllColumns,
			cardsImageColumnsWithDefault,
			cardsImageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cardsImageAllColumns,
			cardsImagePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("rdb: unable to upsert cards_images, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "cards_images", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cards_images` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(cardsImageType, cardsImageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cardsImageType, cardsImageMapping, ret)
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
		return errors.Wrap(err, "rdb: unable to upsert for cards_images")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == cardsImageMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(cardsImageType, cardsImageMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "rdb: unable to retrieve unique values for cards_images")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "rdb: unable to populate default values for cards_images")
	}

CacheNoHooks:
	if !cached {
		cardsImageUpsertCacheMut.Lock()
		cardsImageUpsertCache[key] = cache
		cardsImageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CardsImage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CardsImage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("rdb: no CardsImage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cardsImagePrimaryKeyMapping)
	sql := "DELETE FROM `cards_images` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to delete from cards_images")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to get rows affected by delete for cards_images")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cardsImageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("rdb: no cardsImageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to delete all from cards_images")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to get rows affected by deleteall for cards_images")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CardsImageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(cardsImageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cardsImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cards_images` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cardsImagePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rdb: unable to delete all from cardsImage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rdb: failed to get rows affected by deleteall for cards_images")
	}

	if len(cardsImageAfterDeleteHooks) != 0 {
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
func (o *CardsImage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCardsImage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CardsImageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CardsImageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cardsImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cards_images`.* FROM `cards_images` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, cardsImagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "rdb: unable to reload all in CardsImageSlice")
	}

	*o = slice

	return nil
}

// CardsImageExists checks if the CardsImage row exists.
func CardsImageExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cards_images` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "rdb: unable to check if cards_images exists")
	}

	return exists, nil
}
