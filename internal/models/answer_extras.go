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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AnswerExtras is an object representing the database table.
type AnswerExtras struct {
	AnswerID  string      `boil:"answer_id" json:"answer_id" toml:"answer_id" yaml:"answer_id"`
	Port      null.Int64  `boil:"port" json:"port,omitempty" toml:"port" yaml:"port,omitempty"`
	Priority  null.Int64  `boil:"priority" json:"priority,omitempty" toml:"priority" yaml:"priority,omitempty"`
	Protocol  null.String `boil:"protocol" json:"protocol,omitempty" toml:"protocol" yaml:"protocol,omitempty"`
	Weight    null.String `boil:"weight" json:"weight,omitempty" toml:"weight" yaml:"weight,omitempty"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *answerExtrasR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L answerExtrasL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AnswerExtrasColumns = struct {
	AnswerID  string
	Port      string
	Priority  string
	Protocol  string
	Weight    string
	CreatedAt string
	UpdatedAt string
}{
	AnswerID:  "answer_id",
	Port:      "port",
	Priority:  "priority",
	Protocol:  "protocol",
	Weight:    "weight",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var AnswerExtrasTableColumns = struct {
	AnswerID  string
	Port      string
	Priority  string
	Protocol  string
	Weight    string
	CreatedAt string
	UpdatedAt string
}{
	AnswerID:  "answer_extras.answer_id",
	Port:      "answer_extras.port",
	Priority:  "answer_extras.priority",
	Protocol:  "answer_extras.protocol",
	Weight:    "answer_extras.weight",
	CreatedAt: "answer_extras.created_at",
	UpdatedAt: "answer_extras.updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AnswerExtrasWhere = struct {
	AnswerID  whereHelperstring
	Port      whereHelpernull_Int64
	Priority  whereHelpernull_Int64
	Protocol  whereHelpernull_String
	Weight    whereHelpernull_String
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	AnswerID:  whereHelperstring{field: "\"answer_extras\".\"answer_id\""},
	Port:      whereHelpernull_Int64{field: "\"answer_extras\".\"port\""},
	Priority:  whereHelpernull_Int64{field: "\"answer_extras\".\"priority\""},
	Protocol:  whereHelpernull_String{field: "\"answer_extras\".\"protocol\""},
	Weight:    whereHelpernull_String{field: "\"answer_extras\".\"weight\""},
	CreatedAt: whereHelpertime_Time{field: "\"answer_extras\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"answer_extras\".\"updated_at\""},
}

// AnswerExtrasRels is where relationship names are stored.
var AnswerExtrasRels = struct {
	Answer string
}{
	Answer: "Answer",
}

// answerExtrasR is where relationships are stored.
type answerExtrasR struct {
	Answer *Answer `boil:"Answer" json:"Answer" toml:"Answer" yaml:"Answer"`
}

// NewStruct creates a new relationship struct
func (*answerExtrasR) NewStruct() *answerExtrasR {
	return &answerExtrasR{}
}

func (r *answerExtrasR) GetAnswer() *Answer {
	if r == nil {
		return nil
	}
	return r.Answer
}

// answerExtrasL is where Load methods for each relationship are stored.
type answerExtrasL struct{}

var (
	answerExtrasAllColumns            = []string{"answer_id", "port", "priority", "protocol", "weight", "created_at", "updated_at"}
	answerExtrasColumnsWithoutDefault = []string{"answer_id"}
	answerExtrasColumnsWithDefault    = []string{"port", "priority", "protocol", "weight", "created_at", "updated_at"}
	answerExtrasPrimaryKeyColumns     = []string{"answer_id"}
	answerExtrasGeneratedColumns      = []string{}
)

type (
	// AnswerExtrasSlice is an alias for a slice of pointers to AnswerExtras.
	// This should almost always be used instead of []AnswerExtras.
	AnswerExtrasSlice []*AnswerExtras
	// AnswerExtrasHook is the signature for custom AnswerExtras hook methods
	AnswerExtrasHook func(context.Context, boil.ContextExecutor, *AnswerExtras) error

	answerExtrasQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	answerExtrasType                 = reflect.TypeOf(&AnswerExtras{})
	answerExtrasMapping              = queries.MakeStructMapping(answerExtrasType)
	answerExtrasPrimaryKeyMapping, _ = queries.BindMapping(answerExtrasType, answerExtrasMapping, answerExtrasPrimaryKeyColumns)
	answerExtrasInsertCacheMut       sync.RWMutex
	answerExtrasInsertCache          = make(map[string]insertCache)
	answerExtrasUpdateCacheMut       sync.RWMutex
	answerExtrasUpdateCache          = make(map[string]updateCache)
	answerExtrasUpsertCacheMut       sync.RWMutex
	answerExtrasUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var answerExtrasAfterSelectHooks []AnswerExtrasHook

var answerExtrasBeforeInsertHooks []AnswerExtrasHook
var answerExtrasAfterInsertHooks []AnswerExtrasHook

var answerExtrasBeforeUpdateHooks []AnswerExtrasHook
var answerExtrasAfterUpdateHooks []AnswerExtrasHook

var answerExtrasBeforeDeleteHooks []AnswerExtrasHook
var answerExtrasAfterDeleteHooks []AnswerExtrasHook

var answerExtrasBeforeUpsertHooks []AnswerExtrasHook
var answerExtrasAfterUpsertHooks []AnswerExtrasHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AnswerExtras) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerExtrasAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AnswerExtras) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerExtrasBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AnswerExtras) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerExtrasAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AnswerExtras) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerExtrasBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AnswerExtras) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerExtrasAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AnswerExtras) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerExtrasBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AnswerExtras) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerExtrasAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AnswerExtras) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerExtrasBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AnswerExtras) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range answerExtrasAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnswerExtrasHook registers your hook function for all future operations.
func AddAnswerExtrasHook(hookPoint boil.HookPoint, answerExtrasHook AnswerExtrasHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		answerExtrasAfterSelectHooks = append(answerExtrasAfterSelectHooks, answerExtrasHook)
	case boil.BeforeInsertHook:
		answerExtrasBeforeInsertHooks = append(answerExtrasBeforeInsertHooks, answerExtrasHook)
	case boil.AfterInsertHook:
		answerExtrasAfterInsertHooks = append(answerExtrasAfterInsertHooks, answerExtrasHook)
	case boil.BeforeUpdateHook:
		answerExtrasBeforeUpdateHooks = append(answerExtrasBeforeUpdateHooks, answerExtrasHook)
	case boil.AfterUpdateHook:
		answerExtrasAfterUpdateHooks = append(answerExtrasAfterUpdateHooks, answerExtrasHook)
	case boil.BeforeDeleteHook:
		answerExtrasBeforeDeleteHooks = append(answerExtrasBeforeDeleteHooks, answerExtrasHook)
	case boil.AfterDeleteHook:
		answerExtrasAfterDeleteHooks = append(answerExtrasAfterDeleteHooks, answerExtrasHook)
	case boil.BeforeUpsertHook:
		answerExtrasBeforeUpsertHooks = append(answerExtrasBeforeUpsertHooks, answerExtrasHook)
	case boil.AfterUpsertHook:
		answerExtrasAfterUpsertHooks = append(answerExtrasAfterUpsertHooks, answerExtrasHook)
	}
}

// One returns a single answerExtras record from the query.
func (q answerExtrasQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AnswerExtras, error) {
	o := &AnswerExtras{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for answer_extras")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AnswerExtras records from the query.
func (q answerExtrasQuery) All(ctx context.Context, exec boil.ContextExecutor) (AnswerExtrasSlice, error) {
	var o []*AnswerExtras

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AnswerExtras slice")
	}

	if len(answerExtrasAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AnswerExtras records in the query.
func (q answerExtrasQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count answer_extras rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q answerExtrasQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if answer_extras exists")
	}

	return count > 0, nil
}

// Answer pointed to by the foreign key.
func (o *AnswerExtras) Answer(mods ...qm.QueryMod) answerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"answer_id\" = ?", o.AnswerID),
	}

	queryMods = append(queryMods, mods...)

	return Answers(queryMods...)
}

// LoadAnswer allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (answerExtrasL) LoadAnswer(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAnswerExtras interface{}, mods queries.Applicator) error {
	var slice []*AnswerExtras
	var object *AnswerExtras

	if singular {
		var ok bool
		object, ok = maybeAnswerExtras.(*AnswerExtras)
		if !ok {
			object = new(AnswerExtras)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAnswerExtras)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAnswerExtras))
			}
		}
	} else {
		s, ok := maybeAnswerExtras.(*[]*AnswerExtras)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAnswerExtras)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAnswerExtras))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &answerExtrasR{}
		}
		args = append(args, object.AnswerID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &answerExtrasR{}
			}

			for _, a := range args {
				if a == obj.AnswerID {
					continue Outer
				}
			}

			args = append(args, obj.AnswerID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`answers`),
		qm.WhereIn(`answers.answer_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Answer")
	}

	var resultSlice []*Answer
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Answer")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for answers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for answers")
	}

	if len(answerExtrasAfterSelectHooks) != 0 {
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
		object.R.Answer = foreign
		if foreign.R == nil {
			foreign.R = &answerR{}
		}
		foreign.R.AnswerExtras = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AnswerID == foreign.AnswerID {
				local.R.Answer = foreign
				if foreign.R == nil {
					foreign.R = &answerR{}
				}
				foreign.R.AnswerExtras = local
				break
			}
		}
	}

	return nil
}

// SetAnswer of the answerExtras to the related item.
// Sets o.R.Answer to related.
// Adds o to related.R.AnswerExtras.
func (o *AnswerExtras) SetAnswer(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Answer) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"answer_extras\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"answer_id"}),
		strmangle.WhereClause("\"", "\"", 2, answerExtrasPrimaryKeyColumns),
	)
	values := []interface{}{related.AnswerID, o.AnswerID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AnswerID = related.AnswerID
	if o.R == nil {
		o.R = &answerExtrasR{
			Answer: related,
		}
	} else {
		o.R.Answer = related
	}

	if related.R == nil {
		related.R = &answerR{
			AnswerExtras: o,
		}
	} else {
		related.R.AnswerExtras = o
	}

	return nil
}

// AnswerExtras retrieves all the records using an executor.
func AnswerExtras(mods ...qm.QueryMod) answerExtrasQuery {
	mods = append(mods, qm.From("\"answer_extras\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"answer_extras\".*"})
	}

	return answerExtrasQuery{q}
}

// FindAnswerExtras retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAnswerExtras(ctx context.Context, exec boil.ContextExecutor, answerID string, selectCols ...string) (*AnswerExtras, error) {
	answerExtrasObj := &AnswerExtras{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"answer_extras\" where \"answer_id\"=$1", sel,
	)

	q := queries.Raw(query, answerID)

	err := q.Bind(ctx, exec, answerExtrasObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from answer_extras")
	}

	if err = answerExtrasObj.doAfterSelectHooks(ctx, exec); err != nil {
		return answerExtrasObj, err
	}

	return answerExtrasObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AnswerExtras) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no answer_extras provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(answerExtrasColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	answerExtrasInsertCacheMut.RLock()
	cache, cached := answerExtrasInsertCache[key]
	answerExtrasInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			answerExtrasAllColumns,
			answerExtrasColumnsWithDefault,
			answerExtrasColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(answerExtrasType, answerExtrasMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(answerExtrasType, answerExtrasMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"answer_extras\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"answer_extras\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into answer_extras")
	}

	if !cached {
		answerExtrasInsertCacheMut.Lock()
		answerExtrasInsertCache[key] = cache
		answerExtrasInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AnswerExtras.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AnswerExtras) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	answerExtrasUpdateCacheMut.RLock()
	cache, cached := answerExtrasUpdateCache[key]
	answerExtrasUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			answerExtrasAllColumns,
			answerExtrasPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update answer_extras, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"answer_extras\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, answerExtrasPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(answerExtrasType, answerExtrasMapping, append(wl, answerExtrasPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update answer_extras row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for answer_extras")
	}

	if !cached {
		answerExtrasUpdateCacheMut.Lock()
		answerExtrasUpdateCache[key] = cache
		answerExtrasUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q answerExtrasQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for answer_extras")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for answer_extras")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AnswerExtrasSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerExtrasPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"answer_extras\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, answerExtrasPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in answerExtras slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all answerExtras")
	}
	return rowsAff, nil
}

// Delete deletes a single AnswerExtras record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AnswerExtras) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AnswerExtras provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), answerExtrasPrimaryKeyMapping)
	sql := "DELETE FROM \"answer_extras\" WHERE \"answer_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from answer_extras")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for answer_extras")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q answerExtrasQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no answerExtrasQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from answer_extras")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for answer_extras")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AnswerExtrasSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(answerExtrasBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerExtrasPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"answer_extras\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, answerExtrasPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from answerExtras slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for answer_extras")
	}

	if len(answerExtrasAfterDeleteHooks) != 0 {
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
func (o *AnswerExtras) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAnswerExtras(ctx, exec, o.AnswerID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnswerExtrasSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AnswerExtrasSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), answerExtrasPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"answer_extras\".* FROM \"answer_extras\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, answerExtrasPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AnswerExtrasSlice")
	}

	*o = slice

	return nil
}

// AnswerExtrasExists checks if the AnswerExtras row exists.
func AnswerExtrasExists(ctx context.Context, exec boil.ContextExecutor, answerID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"answer_extras\" where \"answer_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, answerID)
	}
	row := exec.QueryRowContext(ctx, sql, answerID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if answer_extras exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AnswerExtras) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no answer_extras provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(answerExtrasColumnsWithDefault, o)

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

	answerExtrasUpsertCacheMut.RLock()
	cache, cached := answerExtrasUpsertCache[key]
	answerExtrasUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			answerExtrasAllColumns,
			answerExtrasColumnsWithDefault,
			answerExtrasColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			answerExtrasAllColumns,
			answerExtrasPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert answer_extras, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(answerExtrasPrimaryKeyColumns))
			copy(conflict, answerExtrasPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"answer_extras\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(answerExtrasType, answerExtrasMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(answerExtrasType, answerExtrasMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert answer_extras")
	}

	if !cached {
		answerExtrasUpsertCacheMut.Lock()
		answerExtrasUpsertCache[key] = cache
		answerExtrasUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
