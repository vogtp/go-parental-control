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
	"github.com/vogtp/go-win-session/ent/day"
)

// DayCreate is the builder for creating a Day entity.
type DayCreate struct {
	config
	mutation *DayMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetYear sets the "year" field.
func (dc *DayCreate) SetYear(i int) *DayCreate {
	dc.mutation.SetYear(i)
	return dc
}

// SetMonth sets the "month" field.
func (dc *DayCreate) SetMonth(i int) *DayCreate {
	dc.mutation.SetMonth(i)
	return dc
}

// SetDay sets the "day" field.
func (dc *DayCreate) SetDay(i int) *DayCreate {
	dc.mutation.SetDay(i)
	return dc
}

// AddActivityIDs adds the "activity" edge to the Activity entity by IDs.
func (dc *DayCreate) AddActivityIDs(ids ...int) *DayCreate {
	dc.mutation.AddActivityIDs(ids...)
	return dc
}

// AddActivity adds the "activity" edges to the Activity entity.
func (dc *DayCreate) AddActivity(a ...*Activity) *DayCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return dc.AddActivityIDs(ids...)
}

// Mutation returns the DayMutation object of the builder.
func (dc *DayCreate) Mutation() *DayMutation {
	return dc.mutation
}

// Save creates the Day in the database.
func (dc *DayCreate) Save(ctx context.Context) (*Day, error) {
	var (
		err  error
		node *Day
	)
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DayMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Day)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DayMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DayCreate) SaveX(ctx context.Context) *Day {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DayCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DayCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DayCreate) check() error {
	if _, ok := dc.mutation.Year(); !ok {
		return &ValidationError{Name: "year", err: errors.New(`ent: missing required field "Day.year"`)}
	}
	if _, ok := dc.mutation.Month(); !ok {
		return &ValidationError{Name: "month", err: errors.New(`ent: missing required field "Day.month"`)}
	}
	if _, ok := dc.mutation.Day(); !ok {
		return &ValidationError{Name: "day", err: errors.New(`ent: missing required field "Day.day"`)}
	}
	return nil
}

func (dc *DayCreate) sqlSave(ctx context.Context) (*Day, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dc *DayCreate) createSpec() (*Day, *sqlgraph.CreateSpec) {
	var (
		_node = &Day{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: day.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: day.FieldID,
			},
		}
	)
	_spec.OnConflict = dc.conflict
	if value, ok := dc.mutation.Year(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: day.FieldYear,
		})
		_node.Year = value
	}
	if value, ok := dc.mutation.Month(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: day.FieldMonth,
		})
		_node.Month = value
	}
	if value, ok := dc.mutation.Day(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: day.FieldDay,
		})
		_node.Day = value
	}
	if nodes := dc.mutation.ActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   day.ActivityTable,
			Columns: []string{day.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: activity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Day.Create().
//		SetYear(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DayUpsert) {
//			SetYear(v+v).
//		}).
//		Exec(ctx)
//
func (dc *DayCreate) OnConflict(opts ...sql.ConflictOption) *DayUpsertOne {
	dc.conflict = opts
	return &DayUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Day.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dc *DayCreate) OnConflictColumns(columns ...string) *DayUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DayUpsertOne{
		create: dc,
	}
}

type (
	// DayUpsertOne is the builder for "upsert"-ing
	//  one Day node.
	DayUpsertOne struct {
		create *DayCreate
	}

	// DayUpsert is the "OnConflict" setter.
	DayUpsert struct {
		*sql.UpdateSet
	}
)

// SetYear sets the "year" field.
func (u *DayUpsert) SetYear(v int) *DayUpsert {
	u.Set(day.FieldYear, v)
	return u
}

// UpdateYear sets the "year" field to the value that was provided on create.
func (u *DayUpsert) UpdateYear() *DayUpsert {
	u.SetExcluded(day.FieldYear)
	return u
}

// AddYear adds v to the "year" field.
func (u *DayUpsert) AddYear(v int) *DayUpsert {
	u.Add(day.FieldYear, v)
	return u
}

// SetMonth sets the "month" field.
func (u *DayUpsert) SetMonth(v int) *DayUpsert {
	u.Set(day.FieldMonth, v)
	return u
}

// UpdateMonth sets the "month" field to the value that was provided on create.
func (u *DayUpsert) UpdateMonth() *DayUpsert {
	u.SetExcluded(day.FieldMonth)
	return u
}

// AddMonth adds v to the "month" field.
func (u *DayUpsert) AddMonth(v int) *DayUpsert {
	u.Add(day.FieldMonth, v)
	return u
}

// SetDay sets the "day" field.
func (u *DayUpsert) SetDay(v int) *DayUpsert {
	u.Set(day.FieldDay, v)
	return u
}

// UpdateDay sets the "day" field to the value that was provided on create.
func (u *DayUpsert) UpdateDay() *DayUpsert {
	u.SetExcluded(day.FieldDay)
	return u
}

// AddDay adds v to the "day" field.
func (u *DayUpsert) AddDay(v int) *DayUpsert {
	u.Add(day.FieldDay, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Day.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *DayUpsertOne) UpdateNewValues() *DayUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Day.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *DayUpsertOne) Ignore() *DayUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DayUpsertOne) DoNothing() *DayUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DayCreate.OnConflict
// documentation for more info.
func (u *DayUpsertOne) Update(set func(*DayUpsert)) *DayUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DayUpsert{UpdateSet: update})
	}))
	return u
}

// SetYear sets the "year" field.
func (u *DayUpsertOne) SetYear(v int) *DayUpsertOne {
	return u.Update(func(s *DayUpsert) {
		s.SetYear(v)
	})
}

// AddYear adds v to the "year" field.
func (u *DayUpsertOne) AddYear(v int) *DayUpsertOne {
	return u.Update(func(s *DayUpsert) {
		s.AddYear(v)
	})
}

// UpdateYear sets the "year" field to the value that was provided on create.
func (u *DayUpsertOne) UpdateYear() *DayUpsertOne {
	return u.Update(func(s *DayUpsert) {
		s.UpdateYear()
	})
}

// SetMonth sets the "month" field.
func (u *DayUpsertOne) SetMonth(v int) *DayUpsertOne {
	return u.Update(func(s *DayUpsert) {
		s.SetMonth(v)
	})
}

// AddMonth adds v to the "month" field.
func (u *DayUpsertOne) AddMonth(v int) *DayUpsertOne {
	return u.Update(func(s *DayUpsert) {
		s.AddMonth(v)
	})
}

// UpdateMonth sets the "month" field to the value that was provided on create.
func (u *DayUpsertOne) UpdateMonth() *DayUpsertOne {
	return u.Update(func(s *DayUpsert) {
		s.UpdateMonth()
	})
}

// SetDay sets the "day" field.
func (u *DayUpsertOne) SetDay(v int) *DayUpsertOne {
	return u.Update(func(s *DayUpsert) {
		s.SetDay(v)
	})
}

// AddDay adds v to the "day" field.
func (u *DayUpsertOne) AddDay(v int) *DayUpsertOne {
	return u.Update(func(s *DayUpsert) {
		s.AddDay(v)
	})
}

// UpdateDay sets the "day" field to the value that was provided on create.
func (u *DayUpsertOne) UpdateDay() *DayUpsertOne {
	return u.Update(func(s *DayUpsert) {
		s.UpdateDay()
	})
}

// Exec executes the query.
func (u *DayUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DayCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DayUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DayUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DayUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DayCreateBulk is the builder for creating many Day entities in bulk.
type DayCreateBulk struct {
	config
	builders []*DayCreate
	conflict []sql.ConflictOption
}

// Save creates the Day entities in the database.
func (dcb *DayCreateBulk) Save(ctx context.Context) ([]*Day, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Day, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DayMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DayCreateBulk) SaveX(ctx context.Context) []*Day {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DayCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DayCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Day.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DayUpsert) {
//			SetYear(v+v).
//		}).
//		Exec(ctx)
//
func (dcb *DayCreateBulk) OnConflict(opts ...sql.ConflictOption) *DayUpsertBulk {
	dcb.conflict = opts
	return &DayUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Day.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dcb *DayCreateBulk) OnConflictColumns(columns ...string) *DayUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DayUpsertBulk{
		create: dcb,
	}
}

// DayUpsertBulk is the builder for "upsert"-ing
// a bulk of Day nodes.
type DayUpsertBulk struct {
	create *DayCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Day.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *DayUpsertBulk) UpdateNewValues() *DayUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Day.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *DayUpsertBulk) Ignore() *DayUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DayUpsertBulk) DoNothing() *DayUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DayCreateBulk.OnConflict
// documentation for more info.
func (u *DayUpsertBulk) Update(set func(*DayUpsert)) *DayUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DayUpsert{UpdateSet: update})
	}))
	return u
}

// SetYear sets the "year" field.
func (u *DayUpsertBulk) SetYear(v int) *DayUpsertBulk {
	return u.Update(func(s *DayUpsert) {
		s.SetYear(v)
	})
}

// AddYear adds v to the "year" field.
func (u *DayUpsertBulk) AddYear(v int) *DayUpsertBulk {
	return u.Update(func(s *DayUpsert) {
		s.AddYear(v)
	})
}

// UpdateYear sets the "year" field to the value that was provided on create.
func (u *DayUpsertBulk) UpdateYear() *DayUpsertBulk {
	return u.Update(func(s *DayUpsert) {
		s.UpdateYear()
	})
}

// SetMonth sets the "month" field.
func (u *DayUpsertBulk) SetMonth(v int) *DayUpsertBulk {
	return u.Update(func(s *DayUpsert) {
		s.SetMonth(v)
	})
}

// AddMonth adds v to the "month" field.
func (u *DayUpsertBulk) AddMonth(v int) *DayUpsertBulk {
	return u.Update(func(s *DayUpsert) {
		s.AddMonth(v)
	})
}

// UpdateMonth sets the "month" field to the value that was provided on create.
func (u *DayUpsertBulk) UpdateMonth() *DayUpsertBulk {
	return u.Update(func(s *DayUpsert) {
		s.UpdateMonth()
	})
}

// SetDay sets the "day" field.
func (u *DayUpsertBulk) SetDay(v int) *DayUpsertBulk {
	return u.Update(func(s *DayUpsert) {
		s.SetDay(v)
	})
}

// AddDay adds v to the "day" field.
func (u *DayUpsertBulk) AddDay(v int) *DayUpsertBulk {
	return u.Update(func(s *DayUpsert) {
		s.AddDay(v)
	})
}

// UpdateDay sets the "day" field to the value that was provided on create.
func (u *DayUpsertBulk) UpdateDay() *DayUpsertBulk {
	return u.Update(func(s *DayUpsert) {
		s.UpdateDay()
	})
}

// Exec executes the query.
func (u *DayUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DayCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DayCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DayUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}