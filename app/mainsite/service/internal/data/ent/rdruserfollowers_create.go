// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/rdruserfollowers"
)

// RdRUserFollowersCreate is the builder for creating a RdRUserFollowers entity.
type RdRUserFollowersCreate struct {
	config
	mutation *RdRUserFollowersMutation
	hooks    []Hook
}

// SetUID sets the "uid" field.
func (rrfc *RdRUserFollowersCreate) SetUID(i int32) *RdRUserFollowersCreate {
	rrfc.mutation.SetUID(i)
	return rrfc
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (rrfc *RdRUserFollowersCreate) SetNillableUID(i *int32) *RdRUserFollowersCreate {
	if i != nil {
		rrfc.SetUID(*i)
	}
	return rrfc
}

// SetFollowersUID sets the "followers_uid" field.
func (rrfc *RdRUserFollowersCreate) SetFollowersUID(i int32) *RdRUserFollowersCreate {
	rrfc.mutation.SetFollowersUID(i)
	return rrfc
}

// SetNillableFollowersUID sets the "followers_uid" field if the given value is not nil.
func (rrfc *RdRUserFollowersCreate) SetNillableFollowersUID(i *int32) *RdRUserFollowersCreate {
	if i != nil {
		rrfc.SetFollowersUID(*i)
	}
	return rrfc
}

// SetCreateTime sets the "create_time" field.
func (rrfc *RdRUserFollowersCreate) SetCreateTime(t time.Time) *RdRUserFollowersCreate {
	rrfc.mutation.SetCreateTime(t)
	return rrfc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rrfc *RdRUserFollowersCreate) SetNillableCreateTime(t *time.Time) *RdRUserFollowersCreate {
	if t != nil {
		rrfc.SetCreateTime(*t)
	}
	return rrfc
}

// SetID sets the "id" field.
func (rrfc *RdRUserFollowersCreate) SetID(i int64) *RdRUserFollowersCreate {
	rrfc.mutation.SetID(i)
	return rrfc
}

// Mutation returns the RdRUserFollowersMutation object of the builder.
func (rrfc *RdRUserFollowersCreate) Mutation() *RdRUserFollowersMutation {
	return rrfc.mutation
}

// Save creates the RdRUserFollowers in the database.
func (rrfc *RdRUserFollowersCreate) Save(ctx context.Context) (*RdRUserFollowers, error) {
	rrfc.defaults()
	return withHooks[*RdRUserFollowers, RdRUserFollowersMutation](ctx, rrfc.sqlSave, rrfc.mutation, rrfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rrfc *RdRUserFollowersCreate) SaveX(ctx context.Context) *RdRUserFollowers {
	v, err := rrfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rrfc *RdRUserFollowersCreate) Exec(ctx context.Context) error {
	_, err := rrfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rrfc *RdRUserFollowersCreate) ExecX(ctx context.Context) {
	if err := rrfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rrfc *RdRUserFollowersCreate) defaults() {
	if _, ok := rrfc.mutation.CreateTime(); !ok {
		v := rdruserfollowers.DefaultCreateTime()
		rrfc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rrfc *RdRUserFollowersCreate) check() error {
	if _, ok := rrfc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "RdRUserFollowers.create_time"`)}
	}
	return nil
}

func (rrfc *RdRUserFollowersCreate) sqlSave(ctx context.Context) (*RdRUserFollowers, error) {
	if err := rrfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rrfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rrfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	rrfc.mutation.id = &_node.ID
	rrfc.mutation.done = true
	return _node, nil
}

func (rrfc *RdRUserFollowersCreate) createSpec() (*RdRUserFollowers, *sqlgraph.CreateSpec) {
	var (
		_node = &RdRUserFollowers{config: rrfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rdruserfollowers.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: rdruserfollowers.FieldID,
			},
		}
	)
	if id, ok := rrfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rrfc.mutation.UID(); ok {
		_spec.SetField(rdruserfollowers.FieldUID, field.TypeInt32, value)
		_node.UID = value
	}
	if value, ok := rrfc.mutation.FollowersUID(); ok {
		_spec.SetField(rdruserfollowers.FieldFollowersUID, field.TypeInt32, value)
		_node.FollowersUID = value
	}
	if value, ok := rrfc.mutation.CreateTime(); ok {
		_spec.SetField(rdruserfollowers.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	return _node, _spec
}

// RdRUserFollowersCreateBulk is the builder for creating many RdRUserFollowers entities in bulk.
type RdRUserFollowersCreateBulk struct {
	config
	builders []*RdRUserFollowersCreate
}

// Save creates the RdRUserFollowers entities in the database.
func (rrfcb *RdRUserFollowersCreateBulk) Save(ctx context.Context) ([]*RdRUserFollowers, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rrfcb.builders))
	nodes := make([]*RdRUserFollowers, len(rrfcb.builders))
	mutators := make([]Mutator, len(rrfcb.builders))
	for i := range rrfcb.builders {
		func(i int, root context.Context) {
			builder := rrfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RdRUserFollowersMutation)
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
					_, err = mutators[i+1].Mutate(root, rrfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rrfcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, rrfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rrfcb *RdRUserFollowersCreateBulk) SaveX(ctx context.Context) []*RdRUserFollowers {
	v, err := rrfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rrfcb *RdRUserFollowersCreateBulk) Exec(ctx context.Context) error {
	_, err := rrfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rrfcb *RdRUserFollowersCreateBulk) ExecX(ctx context.Context) {
	if err := rrfcb.Exec(ctx); err != nil {
		panic(err)
	}
}