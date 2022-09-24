// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/vogtp/go-win-session/ent/activity"
	"github.com/vogtp/go-win-session/ent/day"
	"github.com/vogtp/go-win-session/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeActivity = "Activity"
	TypeDay      = "Day"
)

// ActivityMutation represents an operation that mutates the Activity nodes in the graph.
type ActivityMutation struct {
	config
	op            Op
	typ           string
	id            *int
	username      *string
	activity      *int64
	addactivity   *int64
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Activity, error)
	predicates    []predicate.Activity
}

var _ ent.Mutation = (*ActivityMutation)(nil)

// activityOption allows management of the mutation configuration using functional options.
type activityOption func(*ActivityMutation)

// newActivityMutation creates new mutation for the Activity entity.
func newActivityMutation(c config, op Op, opts ...activityOption) *ActivityMutation {
	m := &ActivityMutation{
		config:        c,
		op:            op,
		typ:           TypeActivity,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withActivityID sets the ID field of the mutation.
func withActivityID(id int) activityOption {
	return func(m *ActivityMutation) {
		var (
			err   error
			once  sync.Once
			value *Activity
		)
		m.oldValue = func(ctx context.Context) (*Activity, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Activity.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withActivity sets the old Activity of the mutation.
func withActivity(node *Activity) activityOption {
	return func(m *ActivityMutation) {
		m.oldValue = func(context.Context) (*Activity, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ActivityMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ActivityMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ActivityMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ActivityMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Activity.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUsername sets the "username" field.
func (m *ActivityMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *ActivityMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the Activity entity.
// If the Activity object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ActivityMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *ActivityMutation) ResetUsername() {
	m.username = nil
}

// SetActivity sets the "activity" field.
func (m *ActivityMutation) SetActivity(i int64) {
	m.activity = &i
	m.addactivity = nil
}

// Activity returns the value of the "activity" field in the mutation.
func (m *ActivityMutation) Activity() (r int64, exists bool) {
	v := m.activity
	if v == nil {
		return
	}
	return *v, true
}

// OldActivity returns the old "activity" field's value of the Activity entity.
// If the Activity object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ActivityMutation) OldActivity(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldActivity is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldActivity requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldActivity: %w", err)
	}
	return oldValue.Activity, nil
}

// AddActivity adds i to the "activity" field.
func (m *ActivityMutation) AddActivity(i int64) {
	if m.addactivity != nil {
		*m.addactivity += i
	} else {
		m.addactivity = &i
	}
}

// AddedActivity returns the value that was added to the "activity" field in this mutation.
func (m *ActivityMutation) AddedActivity() (r int64, exists bool) {
	v := m.addactivity
	if v == nil {
		return
	}
	return *v, true
}

// ResetActivity resets all changes to the "activity" field.
func (m *ActivityMutation) ResetActivity() {
	m.activity = nil
	m.addactivity = nil
}

// Where appends a list predicates to the ActivityMutation builder.
func (m *ActivityMutation) Where(ps ...predicate.Activity) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ActivityMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Activity).
func (m *ActivityMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ActivityMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.username != nil {
		fields = append(fields, activity.FieldUsername)
	}
	if m.activity != nil {
		fields = append(fields, activity.FieldActivity)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ActivityMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case activity.FieldUsername:
		return m.Username()
	case activity.FieldActivity:
		return m.Activity()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ActivityMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case activity.FieldUsername:
		return m.OldUsername(ctx)
	case activity.FieldActivity:
		return m.OldActivity(ctx)
	}
	return nil, fmt.Errorf("unknown Activity field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ActivityMutation) SetField(name string, value ent.Value) error {
	switch name {
	case activity.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case activity.FieldActivity:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetActivity(v)
		return nil
	}
	return fmt.Errorf("unknown Activity field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ActivityMutation) AddedFields() []string {
	var fields []string
	if m.addactivity != nil {
		fields = append(fields, activity.FieldActivity)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ActivityMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case activity.FieldActivity:
		return m.AddedActivity()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ActivityMutation) AddField(name string, value ent.Value) error {
	switch name {
	case activity.FieldActivity:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddActivity(v)
		return nil
	}
	return fmt.Errorf("unknown Activity numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ActivityMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ActivityMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ActivityMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Activity nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ActivityMutation) ResetField(name string) error {
	switch name {
	case activity.FieldUsername:
		m.ResetUsername()
		return nil
	case activity.FieldActivity:
		m.ResetActivity()
		return nil
	}
	return fmt.Errorf("unknown Activity field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ActivityMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ActivityMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ActivityMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ActivityMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ActivityMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ActivityMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ActivityMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Activity unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ActivityMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Activity edge %s", name)
}

// DayMutation represents an operation that mutates the Day nodes in the graph.
type DayMutation struct {
	config
	op              Op
	typ             string
	id              *int
	year            *int
	addyear         *int
	month           *int
	addmonth        *int
	day             *int
	addday          *int
	clearedFields   map[string]struct{}
	activity        map[int]struct{}
	removedactivity map[int]struct{}
	clearedactivity bool
	done            bool
	oldValue        func(context.Context) (*Day, error)
	predicates      []predicate.Day
}

var _ ent.Mutation = (*DayMutation)(nil)

// dayOption allows management of the mutation configuration using functional options.
type dayOption func(*DayMutation)

// newDayMutation creates new mutation for the Day entity.
func newDayMutation(c config, op Op, opts ...dayOption) *DayMutation {
	m := &DayMutation{
		config:        c,
		op:            op,
		typ:           TypeDay,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDayID sets the ID field of the mutation.
func withDayID(id int) dayOption {
	return func(m *DayMutation) {
		var (
			err   error
			once  sync.Once
			value *Day
		)
		m.oldValue = func(ctx context.Context) (*Day, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Day.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDay sets the old Day of the mutation.
func withDay(node *Day) dayOption {
	return func(m *DayMutation) {
		m.oldValue = func(context.Context) (*Day, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DayMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DayMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *DayMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *DayMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Day.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetYear sets the "year" field.
func (m *DayMutation) SetYear(i int) {
	m.year = &i
	m.addyear = nil
}

// Year returns the value of the "year" field in the mutation.
func (m *DayMutation) Year() (r int, exists bool) {
	v := m.year
	if v == nil {
		return
	}
	return *v, true
}

// OldYear returns the old "year" field's value of the Day entity.
// If the Day object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DayMutation) OldYear(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldYear is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldYear requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldYear: %w", err)
	}
	return oldValue.Year, nil
}

// AddYear adds i to the "year" field.
func (m *DayMutation) AddYear(i int) {
	if m.addyear != nil {
		*m.addyear += i
	} else {
		m.addyear = &i
	}
}

// AddedYear returns the value that was added to the "year" field in this mutation.
func (m *DayMutation) AddedYear() (r int, exists bool) {
	v := m.addyear
	if v == nil {
		return
	}
	return *v, true
}

// ResetYear resets all changes to the "year" field.
func (m *DayMutation) ResetYear() {
	m.year = nil
	m.addyear = nil
}

// SetMonth sets the "month" field.
func (m *DayMutation) SetMonth(i int) {
	m.month = &i
	m.addmonth = nil
}

// Month returns the value of the "month" field in the mutation.
func (m *DayMutation) Month() (r int, exists bool) {
	v := m.month
	if v == nil {
		return
	}
	return *v, true
}

// OldMonth returns the old "month" field's value of the Day entity.
// If the Day object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DayMutation) OldMonth(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMonth is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMonth requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMonth: %w", err)
	}
	return oldValue.Month, nil
}

// AddMonth adds i to the "month" field.
func (m *DayMutation) AddMonth(i int) {
	if m.addmonth != nil {
		*m.addmonth += i
	} else {
		m.addmonth = &i
	}
}

// AddedMonth returns the value that was added to the "month" field in this mutation.
func (m *DayMutation) AddedMonth() (r int, exists bool) {
	v := m.addmonth
	if v == nil {
		return
	}
	return *v, true
}

// ResetMonth resets all changes to the "month" field.
func (m *DayMutation) ResetMonth() {
	m.month = nil
	m.addmonth = nil
}

// SetDay sets the "day" field.
func (m *DayMutation) SetDay(i int) {
	m.day = &i
	m.addday = nil
}

// Day returns the value of the "day" field in the mutation.
func (m *DayMutation) Day() (r int, exists bool) {
	v := m.day
	if v == nil {
		return
	}
	return *v, true
}

// OldDay returns the old "day" field's value of the Day entity.
// If the Day object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DayMutation) OldDay(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDay is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDay requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDay: %w", err)
	}
	return oldValue.Day, nil
}

// AddDay adds i to the "day" field.
func (m *DayMutation) AddDay(i int) {
	if m.addday != nil {
		*m.addday += i
	} else {
		m.addday = &i
	}
}

// AddedDay returns the value that was added to the "day" field in this mutation.
func (m *DayMutation) AddedDay() (r int, exists bool) {
	v := m.addday
	if v == nil {
		return
	}
	return *v, true
}

// ResetDay resets all changes to the "day" field.
func (m *DayMutation) ResetDay() {
	m.day = nil
	m.addday = nil
}

// AddActivityIDs adds the "activity" edge to the Activity entity by ids.
func (m *DayMutation) AddActivityIDs(ids ...int) {
	if m.activity == nil {
		m.activity = make(map[int]struct{})
	}
	for i := range ids {
		m.activity[ids[i]] = struct{}{}
	}
}

// ClearActivity clears the "activity" edge to the Activity entity.
func (m *DayMutation) ClearActivity() {
	m.clearedactivity = true
}

// ActivityCleared reports if the "activity" edge to the Activity entity was cleared.
func (m *DayMutation) ActivityCleared() bool {
	return m.clearedactivity
}

// RemoveActivityIDs removes the "activity" edge to the Activity entity by IDs.
func (m *DayMutation) RemoveActivityIDs(ids ...int) {
	if m.removedactivity == nil {
		m.removedactivity = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.activity, ids[i])
		m.removedactivity[ids[i]] = struct{}{}
	}
}

// RemovedActivity returns the removed IDs of the "activity" edge to the Activity entity.
func (m *DayMutation) RemovedActivityIDs() (ids []int) {
	for id := range m.removedactivity {
		ids = append(ids, id)
	}
	return
}

// ActivityIDs returns the "activity" edge IDs in the mutation.
func (m *DayMutation) ActivityIDs() (ids []int) {
	for id := range m.activity {
		ids = append(ids, id)
	}
	return
}

// ResetActivity resets all changes to the "activity" edge.
func (m *DayMutation) ResetActivity() {
	m.activity = nil
	m.clearedactivity = false
	m.removedactivity = nil
}

// Where appends a list predicates to the DayMutation builder.
func (m *DayMutation) Where(ps ...predicate.Day) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *DayMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Day).
func (m *DayMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DayMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.year != nil {
		fields = append(fields, day.FieldYear)
	}
	if m.month != nil {
		fields = append(fields, day.FieldMonth)
	}
	if m.day != nil {
		fields = append(fields, day.FieldDay)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DayMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case day.FieldYear:
		return m.Year()
	case day.FieldMonth:
		return m.Month()
	case day.FieldDay:
		return m.Day()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DayMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case day.FieldYear:
		return m.OldYear(ctx)
	case day.FieldMonth:
		return m.OldMonth(ctx)
	case day.FieldDay:
		return m.OldDay(ctx)
	}
	return nil, fmt.Errorf("unknown Day field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DayMutation) SetField(name string, value ent.Value) error {
	switch name {
	case day.FieldYear:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetYear(v)
		return nil
	case day.FieldMonth:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMonth(v)
		return nil
	case day.FieldDay:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDay(v)
		return nil
	}
	return fmt.Errorf("unknown Day field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DayMutation) AddedFields() []string {
	var fields []string
	if m.addyear != nil {
		fields = append(fields, day.FieldYear)
	}
	if m.addmonth != nil {
		fields = append(fields, day.FieldMonth)
	}
	if m.addday != nil {
		fields = append(fields, day.FieldDay)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DayMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case day.FieldYear:
		return m.AddedYear()
	case day.FieldMonth:
		return m.AddedMonth()
	case day.FieldDay:
		return m.AddedDay()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DayMutation) AddField(name string, value ent.Value) error {
	switch name {
	case day.FieldYear:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddYear(v)
		return nil
	case day.FieldMonth:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddMonth(v)
		return nil
	case day.FieldDay:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDay(v)
		return nil
	}
	return fmt.Errorf("unknown Day numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DayMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DayMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DayMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Day nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DayMutation) ResetField(name string) error {
	switch name {
	case day.FieldYear:
		m.ResetYear()
		return nil
	case day.FieldMonth:
		m.ResetMonth()
		return nil
	case day.FieldDay:
		m.ResetDay()
		return nil
	}
	return fmt.Errorf("unknown Day field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DayMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.activity != nil {
		edges = append(edges, day.EdgeActivity)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DayMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case day.EdgeActivity:
		ids := make([]ent.Value, 0, len(m.activity))
		for id := range m.activity {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DayMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedactivity != nil {
		edges = append(edges, day.EdgeActivity)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DayMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case day.EdgeActivity:
		ids := make([]ent.Value, 0, len(m.removedactivity))
		for id := range m.removedactivity {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DayMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedactivity {
		edges = append(edges, day.EdgeActivity)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DayMutation) EdgeCleared(name string) bool {
	switch name {
	case day.EdgeActivity:
		return m.clearedactivity
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DayMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Day unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DayMutation) ResetEdge(name string) error {
	switch name {
	case day.EdgeActivity:
		m.ResetActivity()
		return nil
	}
	return fmt.Errorf("unknown Day edge %s", name)
}