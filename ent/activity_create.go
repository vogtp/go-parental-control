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
)

// ActivityCreate is the builder for creating a Activity entity.
type ActivityCreate struct {
	config
	mutation *ActivityMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUsername sets the "username" field.
func (ac *ActivityCreate) SetUsername(s string) *ActivityCreate {
	ac.mutation.SetUsername(s)
	return ac
}

// SetActivity sets the "activity" field.
func (ac *ActivityCreate) SetActivity(i int64) *ActivityCreate {
	ac.mutation.SetActivity(i)
	return ac
}

// SetNillableActivity sets the "activity" field if the given value is not nil.
func (ac *ActivityCreate) SetNillableActivity(i *int64) *ActivityCreate {
	if i != nil {
		ac.SetActivity(*i)
	}
	return ac
}

// Mutation returns the ActivityMutation object of the builder.
func (ac *ActivityCreate) Mutation() *ActivityMutation {
	return ac.mutation
}

// Save creates the Activity in the database.
func (ac *ActivityCreate) Save(ctx context.Context) (*Activity, error) {
	var (
		err  error
		node *Activity
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActivityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (ac *ActivityCreate) SaveX(ctx context.Context) *Activity {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ActivityCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ActivityCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ActivityCreate) defaults() {
	if _, ok := ac.mutation.Activity(); !ok {
		v := activity.DefaultActivity
		ac.mutation.SetActivity(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ActivityCreate) check() error {
	if _, ok := ac.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Activity.username"`)}
	}
	if _, ok := ac.mutation.Activity(); !ok {
		return &ValidationError{Name: "activity", err: errors.New(`ent: missing required field "Activity.activity"`)}
	}
	return nil
}

func (ac *ActivityCreate) sqlSave(ctx context.Context) (*Activity, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *ActivityCreate) createSpec() (*Activity, *sqlgraph.CreateSpec) {
	var (
		_node = &Activity{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: activity.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: activity.FieldID,
			},
		}
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: activity.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := ac.mutation.Activity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: activity.FieldActivity,
		})
		_node.Activity = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Activity.Create().
//		SetUsername(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ActivityUpsert) {
//			SetUsername(v+v).
//		}).
//		Exec(ctx)
//
func (ac *ActivityCreate) OnConflict(opts ...sql.ConflictOption) *ActivityUpsertOne {
	ac.conflict = opts
	return &ActivityUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Activity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ac *ActivityCreate) OnConflictColumns(columns ...string) *ActivityUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &ActivityUpsertOne{
		create: ac,
	}
}

type (
	// ActivityUpsertOne is the builder for "upsert"-ing
	//  one Activity node.
	ActivityUpsertOne struct {
		create *ActivityCreate
	}

	// ActivityUpsert is the "OnConflict" setter.
	ActivityUpsert struct {
		*sql.UpdateSet
	}
)

// SetUsername sets the "username" field.
func (u *ActivityUpsert) SetUsername(v string) *ActivityUpsert {
	u.Set(activity.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *ActivityUpsert) UpdateUsername() *ActivityUpsert {
	u.SetExcluded(activity.FieldUsername)
	return u
}

// SetActivity sets the "activity" field.
func (u *ActivityUpsert) SetActivity(v int64) *ActivityUpsert {
	u.Set(activity.FieldActivity, v)
	return u
}

// UpdateActivity sets the "activity" field to the value that was provided on create.
func (u *ActivityUpsert) UpdateActivity() *ActivityUpsert {
	u.SetExcluded(activity.FieldActivity)
	return u
}

// AddActivity adds v to the "activity" field.
func (u *ActivityUpsert) AddActivity(v int64) *ActivityUpsert {
	u.Add(activity.FieldActivity, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Activity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ActivityUpsertOne) UpdateNewValues() *ActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Activity.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ActivityUpsertOne) Ignore() *ActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ActivityUpsertOne) DoNothing() *ActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ActivityCreate.OnConflict
// documentation for more info.
func (u *ActivityUpsertOne) Update(set func(*ActivityUpsert)) *ActivityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ActivityUpsert{UpdateSet: update})
	}))
	return u
}

// SetUsername sets the "username" field.
func (u *ActivityUpsertOne) SetUsername(v string) *ActivityUpsertOne {
	return u.Update(func(s *ActivityUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *ActivityUpsertOne) UpdateUsername() *ActivityUpsertOne {
	return u.Update(func(s *ActivityUpsert) {
		s.UpdateUsername()
	})
}

// SetActivity sets the "activity" field.
func (u *ActivityUpsertOne) SetActivity(v int64) *ActivityUpsertOne {
	return u.Update(func(s *ActivityUpsert) {
		s.SetActivity(v)
	})
}

// AddActivity adds v to the "activity" field.
func (u *ActivityUpsertOne) AddActivity(v int64) *ActivityUpsertOne {
	return u.Update(func(s *ActivityUpsert) {
		s.AddActivity(v)
	})
}

// UpdateActivity sets the "activity" field to the value that was provided on create.
func (u *ActivityUpsertOne) UpdateActivity() *ActivityUpsertOne {
	return u.Update(func(s *ActivityUpsert) {
		s.UpdateActivity()
	})
}

// Exec executes the query.
func (u *ActivityUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ActivityCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ActivityUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ActivityUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ActivityUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ActivityCreateBulk is the builder for creating many Activity entities in bulk.
type ActivityCreateBulk struct {
	config
	builders []*ActivityCreate
	conflict []sql.ConflictOption
}

// Save creates the Activity entities in the database.
func (acb *ActivityCreateBulk) Save(ctx context.Context) ([]*Activity, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Activity, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActivityMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ActivityCreateBulk) SaveX(ctx context.Context) []*Activity {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ActivityCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ActivityCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Activity.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ActivityUpsert) {
//			SetUsername(v+v).
//		}).
//		Exec(ctx)
//
func (acb *ActivityCreateBulk) OnConflict(opts ...sql.ConflictOption) *ActivityUpsertBulk {
	acb.conflict = opts
	return &ActivityUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Activity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (acb *ActivityCreateBulk) OnConflictColumns(columns ...string) *ActivityUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &ActivityUpsertBulk{
		create: acb,
	}
}

// ActivityUpsertBulk is the builder for "upsert"-ing
// a bulk of Activity nodes.
type ActivityUpsertBulk struct {
	create *ActivityCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Activity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ActivityUpsertBulk) UpdateNewValues() *ActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Activity.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ActivityUpsertBulk) Ignore() *ActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ActivityUpsertBulk) DoNothing() *ActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ActivityCreateBulk.OnConflict
// documentation for more info.
func (u *ActivityUpsertBulk) Update(set func(*ActivityUpsert)) *ActivityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ActivityUpsert{UpdateSet: update})
	}))
	return u
}

// SetUsername sets the "username" field.
func (u *ActivityUpsertBulk) SetUsername(v string) *ActivityUpsertBulk {
	return u.Update(func(s *ActivityUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *ActivityUpsertBulk) UpdateUsername() *ActivityUpsertBulk {
	return u.Update(func(s *ActivityUpsert) {
		s.UpdateUsername()
	})
}

// SetActivity sets the "activity" field.
func (u *ActivityUpsertBulk) SetActivity(v int64) *ActivityUpsertBulk {
	return u.Update(func(s *ActivityUpsert) {
		s.SetActivity(v)
	})
}

// AddActivity adds v to the "activity" field.
func (u *ActivityUpsertBulk) AddActivity(v int64) *ActivityUpsertBulk {
	return u.Update(func(s *ActivityUpsert) {
		s.AddActivity(v)
	})
}

// UpdateActivity sets the "activity" field to the value that was provided on create.
func (u *ActivityUpsertBulk) UpdateActivity() *ActivityUpsertBulk {
	return u.Update(func(s *ActivityUpsert) {
		s.UpdateActivity()
	})
}

// Exec executes the query.
func (u *ActivityUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ActivityCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ActivityCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ActivityUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
