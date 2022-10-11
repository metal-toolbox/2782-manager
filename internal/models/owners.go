// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Owner is an object representing the database table.
type Owner struct {
	OwnerID      string    `boil:"owner_id" json:"owner_id" toml:"owner_id" yaml:"owner_id"`
	OwnerName    string    `boil:"owner_name" json:"owner_name" toml:"owner_name" yaml:"owner_name"`
	OwnerOrigin  string    `boil:"owner_origin" json:"owner_origin" toml:"owner_origin" yaml:"owner_origin"`
	OwnerService string    `boil:"owner_service" json:"owner_service" toml:"owner_service" yaml:"owner_service"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *ownerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ownerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OwnerColumns = struct {
	OwnerID      string
	OwnerName    string
	OwnerOrigin  string
	OwnerService string
	CreatedAt    string
	UpdatedAt    string
}{
	OwnerID:      "owner_id",
	OwnerName:    "owner_name",
	OwnerOrigin:  "owner_origin",
	OwnerService: "owner_service",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

var OwnerTableColumns = struct {
	OwnerID      string
	OwnerName    string
	OwnerOrigin  string
	OwnerService string
	CreatedAt    string
	UpdatedAt    string
}{
	OwnerID:      "owners.owner_id",
	OwnerName:    "owners.owner_name",
	OwnerOrigin:  "owners.owner_origin",
	OwnerService: "owners.owner_service",
	CreatedAt:    "owners.created_at",
	UpdatedAt:    "owners.updated_at",
}

// Generated where

var OwnerWhere = struct {
	OwnerID      whereHelperstring
	OwnerName    whereHelperstring
	OwnerOrigin  whereHelperstring
	OwnerService whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpertime_Time
}{
	OwnerID:      whereHelperstring{field: "\"owners\".\"owner_id\""},
	OwnerName:    whereHelperstring{field: "\"owners\".\"owner_name\""},
	OwnerOrigin:  whereHelperstring{field: "\"owners\".\"owner_origin\""},
	OwnerService: whereHelperstring{field: "\"owners\".\"owner_service\""},
	CreatedAt:    whereHelpertime_Time{field: "\"owners\".\"created_at\""},
	UpdatedAt:    whereHelpertime_Time{field: "\"owners\".\"updated_at\""},
}

// OwnerRels is where relationship names are stored.
var OwnerRels = struct {
	Answers string
}{
	Answers: "Answers",
}

// ownerR is where relationships are stored.
type ownerR struct {
	Answers AnswerSlice `boil:"Answers" json:"Answers" toml:"Answers" yaml:"Answers"`
}

// NewStruct creates a new relationship struct
func (*ownerR) NewStruct() *ownerR {
	return &ownerR{}
}

func (r *ownerR) GetAnswers() AnswerSlice {
	if r == nil {
		return nil
	}
	return r.Answers
}

// ownerL is where Load methods for each relationship are stored.
type ownerL struct{}

var (
	ownerAllColumns            = []string{"owner_id", "owner_name", "owner_origin", "owner_service", "created_at", "updated_at"}
	ownerColumnsWithoutDefault = []string{"owner_name", "owner_origin", "owner_service"}
	ownerColumnsWithDefault    = []string{"owner_id", "created_at", "updated_at"}
	ownerPrimaryKeyColumns     = []string{"owner_name", "owner_origin", "owner_service"}
	ownerGeneratedColumns      = []string{}
)

type (
	// OwnerSlice is an alias for a slice of pointers to Owner.
	// This should almost always be used instead of []Owner.
	OwnerSlice []*Owner
	// OwnerHook is the signature for custom Owner hook methods
	OwnerHook func(context.Context, boil.ContextExecutor, *Owner) error

	ownerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ownerType                 = reflect.TypeOf(&Owner{})
	ownerMapping              = queries.MakeStructMapping(ownerType)
	ownerPrimaryKeyMapping, _ = queries.BindMapping(ownerType, ownerMapping, ownerPrimaryKeyColumns)
	ownerInsertCacheMut       sync.RWMutex
	ownerInsertCache          = make(map[string]insertCache)
	ownerUpdateCacheMut       sync.RWMutex
	ownerUpdateCache          = make(map[string]updateCache)
	ownerUpsertCacheMut       sync.RWMutex
	ownerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var ownerAfterSelectHooks []OwnerHook

var ownerBeforeInsertHooks []OwnerHook
var ownerAfterInsertHooks []OwnerHook

var ownerBeforeUpdateHooks []OwnerHook
var ownerAfterUpdateHooks []OwnerHook

var ownerBeforeDeleteHooks []OwnerHook
var ownerAfterDeleteHooks []OwnerHook

var ownerBeforeUpsertHooks []OwnerHook
var ownerAfterUpsertHooks []OwnerHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Owner) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ownerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Owner) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ownerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Owner) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ownerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Owner) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ownerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Owner) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ownerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Owner) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ownerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Owner) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ownerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Owner) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ownerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Owner) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ownerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOwnerHook registers your hook function for all future operations.
func AddOwnerHook(hookPoint boil.HookPoint, ownerHook OwnerHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		ownerAfterSelectHooks = append(ownerAfterSelectHooks, ownerHook)
	case boil.BeforeInsertHook:
		ownerBeforeInsertHooks = append(ownerBeforeInsertHooks, ownerHook)
	case boil.AfterInsertHook:
		ownerAfterInsertHooks = append(ownerAfterInsertHooks, ownerHook)
	case boil.BeforeUpdateHook:
		ownerBeforeUpdateHooks = append(ownerBeforeUpdateHooks, ownerHook)
	case boil.AfterUpdateHook:
		ownerAfterUpdateHooks = append(ownerAfterUpdateHooks, ownerHook)
	case boil.BeforeDeleteHook:
		ownerBeforeDeleteHooks = append(ownerBeforeDeleteHooks, ownerHook)
	case boil.AfterDeleteHook:
		ownerAfterDeleteHooks = append(ownerAfterDeleteHooks, ownerHook)
	case boil.BeforeUpsertHook:
		ownerBeforeUpsertHooks = append(ownerBeforeUpsertHooks, ownerHook)
	case boil.AfterUpsertHook:
		ownerAfterUpsertHooks = append(ownerAfterUpsertHooks, ownerHook)
	}
}

// One returns a single owner record from the query.
func (q ownerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Owner, error) {
	o := &Owner{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for owners")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Owner records from the query.
func (q ownerQuery) All(ctx context.Context, exec boil.ContextExecutor) (OwnerSlice, error) {
	var o []*Owner

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Owner slice")
	}

	if len(ownerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Owner records in the query.
func (q ownerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count owners rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ownerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if owners exists")
	}

	return count > 0, nil
}

// Answers retrieves all the answer's Answers with an executor.
func (o *Owner) Answers(mods ...qm.QueryMod) answerQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"answers\".\"owner_id\"=?", o.OwnerID),
	)

	return Answers(queryMods...)
}

// LoadAnswers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (ownerL) LoadAnswers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOwner interface{}, mods queries.Applicator) error {
	var slice []*Owner
	var object *Owner

	if singular {
		var ok bool
		object, ok = maybeOwner.(*Owner)
		if !ok {
			object = new(Owner)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeOwner)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeOwner))
			}
		}
	} else {
		s, ok := maybeOwner.(*[]*Owner)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeOwner)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeOwner))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ownerR{}
		}
		args = append(args, object.OwnerID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ownerR{}
			}

			for _, a := range args {
				if a == obj.OwnerID {
					continue Outer
				}
			}

			args = append(args, obj.OwnerID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`answers`),
		qm.WhereIn(`answers.owner_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load answers")
	}

	var resultSlice []*Answer
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice answers")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on answers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for answers")
	}

	if len(answerAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Answers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &answerR{}
			}
			foreign.R.Owner = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.OwnerID == foreign.OwnerID {
				local.R.Answers = append(local.R.Answers, foreign)
				if foreign.R == nil {
					foreign.R = &answerR{}
				}
				foreign.R.Owner = local
				break
			}
		}
	}

	return nil
}

// AddAnswers adds the given related objects to the existing relationships
// of the owner, optionally inserting them as new records.
// Appends related to o.R.Answers.
// Sets related.R.Owner appropriately.
func (o *Owner) AddAnswers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Answer) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.OwnerID = o.OwnerID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"answers\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"owner_id"}),
				strmangle.WhereClause("\"", "\"", 2, answerPrimaryKeyColumns),
			)
			values := []interface{}{o.OwnerID, rel.RecordID, rel.OwnerID, rel.AnswerID, rel.AnswerTarget, rel.AnswerType}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.OwnerID = o.OwnerID
		}
	}

	if o.R == nil {
		o.R = &ownerR{
			Answers: related,
		}
	} else {
		o.R.Answers = append(o.R.Answers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &answerR{
				Owner: o,
			}
		} else {
			rel.R.Owner = o
		}
	}
	return nil
}

// Owners retrieves all the records using an executor.
func Owners(mods ...qm.QueryMod) ownerQuery {
	mods = append(mods, qm.From("\"owners\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"owners\".*"})
	}

	return ownerQuery{q}
}

// FindOwner retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOwner(ctx context.Context, exec boil.ContextExecutor, ownerName string, ownerOrigin string, ownerService string, selectCols ...string) (*Owner, error) {
	ownerObj := &Owner{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"owners\" where \"owner_name\"=$1 AND \"owner_origin\"=$2 AND \"owner_service\"=$3", sel,
	)

	q := queries.Raw(query, ownerName, ownerOrigin, ownerService)

	err := q.Bind(ctx, exec, ownerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from owners")
	}

	if err = ownerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return ownerObj, err
	}

	return ownerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Owner) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no owners provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ownerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ownerInsertCacheMut.RLock()
	cache, cached := ownerInsertCache[key]
	ownerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ownerAllColumns,
			ownerColumnsWithDefault,
			ownerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ownerType, ownerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ownerType, ownerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"owners\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"owners\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into owners")
	}

	if !cached {
		ownerInsertCacheMut.Lock()
		ownerInsertCache[key] = cache
		ownerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Owner.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Owner) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ownerUpdateCacheMut.RLock()
	cache, cached := ownerUpdateCache[key]
	ownerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ownerAllColumns,
			ownerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update owners, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"owners\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ownerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ownerType, ownerMapping, append(wl, ownerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update owners row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for owners")
	}

	if !cached {
		ownerUpdateCacheMut.Lock()
		ownerUpdateCache[key] = cache
		ownerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q ownerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for owners")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for owners")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OwnerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ownerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"owners\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, ownerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in owner slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all owner")
	}
	return rowsAff, nil
}

// Delete deletes a single Owner record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Owner) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Owner provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ownerPrimaryKeyMapping)
	sql := "DELETE FROM \"owners\" WHERE \"owner_name\"=$1 AND \"owner_origin\"=$2 AND \"owner_service\"=$3"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from owners")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for owners")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ownerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ownerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from owners")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for owners")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OwnerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(ownerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ownerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"owners\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ownerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from owner slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for owners")
	}

	if len(ownerAfterDeleteHooks) != 0 {
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
func (o *Owner) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOwner(ctx, exec, o.OwnerName, o.OwnerOrigin, o.OwnerService)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OwnerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OwnerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ownerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"owners\".* FROM \"owners\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ownerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OwnerSlice")
	}

	*o = slice

	return nil
}

// OwnerExists checks if the Owner row exists.
func OwnerExists(ctx context.Context, exec boil.ContextExecutor, ownerName string, ownerOrigin string, ownerService string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"owners\" where \"owner_name\"=$1 AND \"owner_origin\"=$2 AND \"owner_service\"=$3 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, ownerName, ownerOrigin, ownerService)
	}
	row := exec.QueryRowContext(ctx, sql, ownerName, ownerOrigin, ownerService)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if owners exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Owner) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no owners provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ownerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	ownerUpsertCacheMut.RLock()
	cache, cached := ownerUpsertCache[key]
	ownerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ownerAllColumns,
			ownerColumnsWithDefault,
			ownerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			ownerAllColumns,
			ownerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert owners, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ownerPrimaryKeyColumns))
			copy(conflict, ownerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"owners\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(ownerType, ownerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ownerType, ownerMapping, ret)
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
		_, _ = fmt.Fprintln(boil.DebugWriter, cache.query)
		_, _ = fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // CockcorachDB doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert owners")
	}

	if !cached {
		ownerUpsertCacheMut.Lock()
		ownerUpsertCache[key] = cache
		ownerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}