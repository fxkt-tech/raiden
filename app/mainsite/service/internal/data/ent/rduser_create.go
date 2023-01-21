// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/rduser"
)

// RdUserCreate is the builder for creating a RdUser entity.
type RdUserCreate struct {
	config
	mutation *RdUserMutation
	hooks    []Hook
}

// SetNick sets the "nick" field.
func (ruc *RdUserCreate) SetNick(s string) *RdUserCreate {
	ruc.mutation.SetNick(s)
	return ruc
}

// SetStatus sets the "status" field.
func (ruc *RdUserCreate) SetStatus(i int32) *RdUserCreate {
	ruc.mutation.SetStatus(i)
	return ruc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruc *RdUserCreate) SetNillableStatus(i *int32) *RdUserCreate {
	if i != nil {
		ruc.SetStatus(*i)
	}
	return ruc
}

// SetCreateTime sets the "create_time" field.
func (ruc *RdUserCreate) SetCreateTime(t time.Time) *RdUserCreate {
	ruc.mutation.SetCreateTime(t)
	return ruc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ruc *RdUserCreate) SetNillableCreateTime(t *time.Time) *RdUserCreate {
	if t != nil {
		ruc.SetCreateTime(*t)
	}
	return ruc
}

// SetID sets the "id" field.
func (ruc *RdUserCreate) SetID(i int32) *RdUserCreate {
	ruc.mutation.SetID(i)
	return ruc
}

// Mutation returns the RdUserMutation object of the builder.
func (ruc *RdUserCreate) Mutation() *RdUserMutation {
	return ruc.mutation
}

// Save creates the RdUser in the database.
func (ruc *RdUserCreate) Save(ctx context.Context) (*RdUser, error) {
	ruc.defaults()
	return withHooks[*RdUser, RdUserMutation](ctx, ruc.sqlSave, ruc.mutation, ruc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ruc *RdUserCreate) SaveX(ctx context.Context) *RdUser {
	v, err := ruc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ruc *RdUserCreate) Exec(ctx context.Context) error {
	_, err := ruc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruc *RdUserCreate) ExecX(ctx context.Context) {
	if err := ruc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruc *RdUserCreate) defaults() {
	if _, ok := ruc.mutation.Status(); !ok {
		v := rduser.DefaultStatus
		ruc.mutation.SetStatus(v)
	}
	if _, ok := ruc.mutation.CreateTime(); !ok {
		v := rduser.DefaultCreateTime()
		ruc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruc *RdUserCreate) check() error {
	if _, ok := ruc.mutation.Nick(); !ok {
		return &ValidationError{Name: "nick", err: errors.New(`ent: missing required field "RdUser.nick"`)}
	}
	if v, ok := ruc.mutation.Nick(); ok {
		if err := rduser.NickValidator(v); err != nil {
			return &ValidationError{Name: "nick", err: fmt.Errorf(`ent: validator failed for field "RdUser.nick": %w`, err)}
		}
	}
	if _, ok := ruc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "RdUser.status"`)}
	}
	if _, ok := ruc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "RdUser.create_time"`)}
	}
	return nil
}

func (ruc *RdUserCreate) sqlSave(ctx context.Context) (*RdUser, error) {
	if err := ruc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ruc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ruc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	ruc.mutation.id = &_node.ID
	ruc.mutation.done = true
	return _node, nil
}

func (ruc *RdUserCreate) createSpec() (*RdUser, *sqlgraph.CreateSpec) {
	var (
		_node = &RdUser{config: ruc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rduser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: rduser.FieldID,
			},
		}
	)
	if id, ok := ruc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ruc.mutation.Nick(); ok {
		_spec.SetField(rduser.FieldNick, field.TypeString, value)
		_node.Nick = value
	}
	if value, ok := ruc.mutation.Status(); ok {
		_spec.SetField(rduser.FieldStatus, field.TypeInt32, value)
		_node.Status = value
	}
	if value, ok := ruc.mutation.CreateTime(); ok {
		_spec.SetField(rduser.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	return _node, _spec
}

// RdUserCreateBulk is the builder for creating many RdUser entities in bulk.
type RdUserCreateBulk struct {
	config
	builders []*RdUserCreate
}

// Save creates the RdUser entities in the database.
func (rucb *RdUserCreateBulk) Save(ctx context.Context) ([]*RdUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rucb.builders))
	nodes := make([]*RdUser, len(rucb.builders))
	mutators := make([]Mutator, len(rucb.builders))
	for i := range rucb.builders {
		func(i int, root context.Context) {
			builder := rucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RdUserMutation)
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
					_, err = mutators[i+1].Mutate(root, rucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rucb.driver, spec); err != nil {
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
					nodes[i].ID = int32(id)
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
		if _, err := mutators[0].Mutate(ctx, rucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rucb *RdUserCreateBulk) SaveX(ctx context.Context) []*RdUser {
	v, err := rucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rucb *RdUserCreateBulk) Exec(ctx context.Context) error {
	_, err := rucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rucb *RdUserCreateBulk) ExecX(ctx context.Context) {
	if err := rucb.Exec(ctx); err != nil {
		panic(err)
	}
}
