// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vogtp/go-win-session/ent/activity"
	"github.com/vogtp/go-win-session/ent/predicate"
)

// ActivityUpdate is the builder for updating Activity entities.
type ActivityUpdate struct {
	config
	hooks    []Hook
	mutation *ActivityMutation
}

// Where appends a list predicates to the ActivityUpdate builder.
func (au *ActivityUpdate) Where(ps ...predicate.Activity) *ActivityUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUsername sets the "username" field.
func (au *ActivityUpdate) SetUsername(s string) *ActivityUpdate {
	au.mutation.SetUsername(s)
	return au
}

// SetActivity sets the "activity" field.
func (au *ActivityUpdate) SetActivity(i int64) *ActivityUpdate {
	au.mutation.ResetActivity()
	au.mutation.SetActivity(i)
	return au
}

// SetNillableActivity sets the "activity" field if the given value is not nil.
func (au *ActivityUpdate) SetNillableActivity(i *int64) *ActivityUpdate {
	if i != nil {
		au.SetActivity(*i)
	}
	return au
}

// AddActivity adds i to the "activity" field.
func (au *ActivityUpdate) AddActivity(i int64) *ActivityUpdate {
	au.mutation.AddActivity(i)
	return au
}

// Mutation returns the ActivityMutation object of the builder.
func (au *ActivityUpdate) Mutation() *ActivityMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ActivityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ActivityUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ActivityUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ActivityUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ActivityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   activity.Table,
			Columns: activity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: activity.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activity.FieldUsername,
		})
	}
	if value, ok := au.mutation.Activity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: activity.FieldActivity,
		})
	}
	if value, ok := au.mutation.AddedActivity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: activity.FieldActivity,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ActivityUpdateOne is the builder for updating a single Activity entity.
type ActivityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ActivityMutation
}

// SetUsername sets the "username" field.
func (auo *ActivityUpdateOne) SetUsername(s string) *ActivityUpdateOne {
	auo.mutation.SetUsername(s)
	return auo
}

// SetActivity sets the "activity" field.
func (auo *ActivityUpdateOne) SetActivity(i int64) *ActivityUpdateOne {
	auo.mutation.ResetActivity()
	auo.mutation.SetActivity(i)
	return auo
}

// SetNillableActivity sets the "activity" field if the given value is not nil.
func (auo *ActivityUpdateOne) SetNillableActivity(i *int64) *ActivityUpdateOne {
	if i != nil {
		auo.SetActivity(*i)
	}
	return auo
}

// AddActivity adds i to the "activity" field.
func (auo *ActivityUpdateOne) AddActivity(i int64) *ActivityUpdateOne {
	auo.mutation.AddActivity(i)
	return auo
}

// Mutation returns the ActivityMutation object of the builder.
func (auo *ActivityUpdateOne) Mutation() *ActivityMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ActivityUpdateOne) Select(field string, fields ...string) *ActivityUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Activity entity.
func (auo *ActivityUpdateOne) Save(ctx context.Context) (*Activity, error) {
	var (
		err  error
		node *Activity
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Activity)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ActivityMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ActivityUpdateOne) SaveX(ctx context.Context) *Activity {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ActivityUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ActivityUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ActivityUpdateOne) sqlSave(ctx context.Context) (_node *Activity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   activity.Table,
			Columns: activity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: activity.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Activity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, activity.FieldID)
		for _, f := range fields {
			if !activity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != activity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activity.FieldUsername,
		})
	}
	if value, ok := auo.mutation.Activity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: activity.FieldActivity,
		})
	}
	if value, ok := auo.mutation.AddedActivity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: activity.FieldActivity,
		})
	}
	_node = &Activity{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}