// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/predicate"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/rdruserfollowers"
)

// RdRUserFollowersUpdate is the builder for updating RdRUserFollowers entities.
type RdRUserFollowersUpdate struct {
	config
	hooks    []Hook
	mutation *RdRUserFollowersMutation
}

// Where appends a list predicates to the RdRUserFollowersUpdate builder.
func (rrfu *RdRUserFollowersUpdate) Where(ps ...predicate.RdRUserFollowers) *RdRUserFollowersUpdate {
	rrfu.mutation.Where(ps...)
	return rrfu
}

// SetUID sets the "uid" field.
func (rrfu *RdRUserFollowersUpdate) SetUID(i int32) *RdRUserFollowersUpdate {
	rrfu.mutation.ResetUID()
	rrfu.mutation.SetUID(i)
	return rrfu
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (rrfu *RdRUserFollowersUpdate) SetNillableUID(i *int32) *RdRUserFollowersUpdate {
	if i != nil {
		rrfu.SetUID(*i)
	}
	return rrfu
}

// AddUID adds i to the "uid" field.
func (rrfu *RdRUserFollowersUpdate) AddUID(i int32) *RdRUserFollowersUpdate {
	rrfu.mutation.AddUID(i)
	return rrfu
}

// ClearUID clears the value of the "uid" field.
func (rrfu *RdRUserFollowersUpdate) ClearUID() *RdRUserFollowersUpdate {
	rrfu.mutation.ClearUID()
	return rrfu
}

// SetFollowersUID sets the "followers_uid" field.
func (rrfu *RdRUserFollowersUpdate) SetFollowersUID(i int32) *RdRUserFollowersUpdate {
	rrfu.mutation.ResetFollowersUID()
	rrfu.mutation.SetFollowersUID(i)
	return rrfu
}

// SetNillableFollowersUID sets the "followers_uid" field if the given value is not nil.
func (rrfu *RdRUserFollowersUpdate) SetNillableFollowersUID(i *int32) *RdRUserFollowersUpdate {
	if i != nil {
		rrfu.SetFollowersUID(*i)
	}
	return rrfu
}

// AddFollowersUID adds i to the "followers_uid" field.
func (rrfu *RdRUserFollowersUpdate) AddFollowersUID(i int32) *RdRUserFollowersUpdate {
	rrfu.mutation.AddFollowersUID(i)
	return rrfu
}

// ClearFollowersUID clears the value of the "followers_uid" field.
func (rrfu *RdRUserFollowersUpdate) ClearFollowersUID() *RdRUserFollowersUpdate {
	rrfu.mutation.ClearFollowersUID()
	return rrfu
}

// SetCreateTime sets the "create_time" field.
func (rrfu *RdRUserFollowersUpdate) SetCreateTime(t time.Time) *RdRUserFollowersUpdate {
	rrfu.mutation.SetCreateTime(t)
	return rrfu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rrfu *RdRUserFollowersUpdate) SetNillableCreateTime(t *time.Time) *RdRUserFollowersUpdate {
	if t != nil {
		rrfu.SetCreateTime(*t)
	}
	return rrfu
}

// Mutation returns the RdRUserFollowersMutation object of the builder.
func (rrfu *RdRUserFollowersUpdate) Mutation() *RdRUserFollowersMutation {
	return rrfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rrfu *RdRUserFollowersUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, RdRUserFollowersMutation](ctx, rrfu.sqlSave, rrfu.mutation, rrfu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rrfu *RdRUserFollowersUpdate) SaveX(ctx context.Context) int {
	affected, err := rrfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rrfu *RdRUserFollowersUpdate) Exec(ctx context.Context) error {
	_, err := rrfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rrfu *RdRUserFollowersUpdate) ExecX(ctx context.Context) {
	if err := rrfu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rrfu *RdRUserFollowersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rdruserfollowers.Table,
			Columns: rdruserfollowers.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: rdruserfollowers.FieldID,
			},
		},
	}
	if ps := rrfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rrfu.mutation.UID(); ok {
		_spec.SetField(rdruserfollowers.FieldUID, field.TypeInt32, value)
	}
	if value, ok := rrfu.mutation.AddedUID(); ok {
		_spec.AddField(rdruserfollowers.FieldUID, field.TypeInt32, value)
	}
	if rrfu.mutation.UIDCleared() {
		_spec.ClearField(rdruserfollowers.FieldUID, field.TypeInt32)
	}
	if value, ok := rrfu.mutation.FollowersUID(); ok {
		_spec.SetField(rdruserfollowers.FieldFollowersUID, field.TypeInt32, value)
	}
	if value, ok := rrfu.mutation.AddedFollowersUID(); ok {
		_spec.AddField(rdruserfollowers.FieldFollowersUID, field.TypeInt32, value)
	}
	if rrfu.mutation.FollowersUIDCleared() {
		_spec.ClearField(rdruserfollowers.FieldFollowersUID, field.TypeInt32)
	}
	if value, ok := rrfu.mutation.CreateTime(); ok {
		_spec.SetField(rdruserfollowers.FieldCreateTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rrfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rdruserfollowers.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rrfu.mutation.done = true
	return n, nil
}

// RdRUserFollowersUpdateOne is the builder for updating a single RdRUserFollowers entity.
type RdRUserFollowersUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RdRUserFollowersMutation
}

// SetUID sets the "uid" field.
func (rrfuo *RdRUserFollowersUpdateOne) SetUID(i int32) *RdRUserFollowersUpdateOne {
	rrfuo.mutation.ResetUID()
	rrfuo.mutation.SetUID(i)
	return rrfuo
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (rrfuo *RdRUserFollowersUpdateOne) SetNillableUID(i *int32) *RdRUserFollowersUpdateOne {
	if i != nil {
		rrfuo.SetUID(*i)
	}
	return rrfuo
}

// AddUID adds i to the "uid" field.
func (rrfuo *RdRUserFollowersUpdateOne) AddUID(i int32) *RdRUserFollowersUpdateOne {
	rrfuo.mutation.AddUID(i)
	return rrfuo
}

// ClearUID clears the value of the "uid" field.
func (rrfuo *RdRUserFollowersUpdateOne) ClearUID() *RdRUserFollowersUpdateOne {
	rrfuo.mutation.ClearUID()
	return rrfuo
}

// SetFollowersUID sets the "followers_uid" field.
func (rrfuo *RdRUserFollowersUpdateOne) SetFollowersUID(i int32) *RdRUserFollowersUpdateOne {
	rrfuo.mutation.ResetFollowersUID()
	rrfuo.mutation.SetFollowersUID(i)
	return rrfuo
}

// SetNillableFollowersUID sets the "followers_uid" field if the given value is not nil.
func (rrfuo *RdRUserFollowersUpdateOne) SetNillableFollowersUID(i *int32) *RdRUserFollowersUpdateOne {
	if i != nil {
		rrfuo.SetFollowersUID(*i)
	}
	return rrfuo
}

// AddFollowersUID adds i to the "followers_uid" field.
func (rrfuo *RdRUserFollowersUpdateOne) AddFollowersUID(i int32) *RdRUserFollowersUpdateOne {
	rrfuo.mutation.AddFollowersUID(i)
	return rrfuo
}

// ClearFollowersUID clears the value of the "followers_uid" field.
func (rrfuo *RdRUserFollowersUpdateOne) ClearFollowersUID() *RdRUserFollowersUpdateOne {
	rrfuo.mutation.ClearFollowersUID()
	return rrfuo
}

// SetCreateTime sets the "create_time" field.
func (rrfuo *RdRUserFollowersUpdateOne) SetCreateTime(t time.Time) *RdRUserFollowersUpdateOne {
	rrfuo.mutation.SetCreateTime(t)
	return rrfuo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rrfuo *RdRUserFollowersUpdateOne) SetNillableCreateTime(t *time.Time) *RdRUserFollowersUpdateOne {
	if t != nil {
		rrfuo.SetCreateTime(*t)
	}
	return rrfuo
}

// Mutation returns the RdRUserFollowersMutation object of the builder.
func (rrfuo *RdRUserFollowersUpdateOne) Mutation() *RdRUserFollowersMutation {
	return rrfuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rrfuo *RdRUserFollowersUpdateOne) Select(field string, fields ...string) *RdRUserFollowersUpdateOne {
	rrfuo.fields = append([]string{field}, fields...)
	return rrfuo
}

// Save executes the query and returns the updated RdRUserFollowers entity.
func (rrfuo *RdRUserFollowersUpdateOne) Save(ctx context.Context) (*RdRUserFollowers, error) {
	return withHooks[*RdRUserFollowers, RdRUserFollowersMutation](ctx, rrfuo.sqlSave, rrfuo.mutation, rrfuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rrfuo *RdRUserFollowersUpdateOne) SaveX(ctx context.Context) *RdRUserFollowers {
	node, err := rrfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rrfuo *RdRUserFollowersUpdateOne) Exec(ctx context.Context) error {
	_, err := rrfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rrfuo *RdRUserFollowersUpdateOne) ExecX(ctx context.Context) {
	if err := rrfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rrfuo *RdRUserFollowersUpdateOne) sqlSave(ctx context.Context) (_node *RdRUserFollowers, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rdruserfollowers.Table,
			Columns: rdruserfollowers.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: rdruserfollowers.FieldID,
			},
		},
	}
	id, ok := rrfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RdRUserFollowers.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rrfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rdruserfollowers.FieldID)
		for _, f := range fields {
			if !rdruserfollowers.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rdruserfollowers.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rrfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rrfuo.mutation.UID(); ok {
		_spec.SetField(rdruserfollowers.FieldUID, field.TypeInt32, value)
	}
	if value, ok := rrfuo.mutation.AddedUID(); ok {
		_spec.AddField(rdruserfollowers.FieldUID, field.TypeInt32, value)
	}
	if rrfuo.mutation.UIDCleared() {
		_spec.ClearField(rdruserfollowers.FieldUID, field.TypeInt32)
	}
	if value, ok := rrfuo.mutation.FollowersUID(); ok {
		_spec.SetField(rdruserfollowers.FieldFollowersUID, field.TypeInt32, value)
	}
	if value, ok := rrfuo.mutation.AddedFollowersUID(); ok {
		_spec.AddField(rdruserfollowers.FieldFollowersUID, field.TypeInt32, value)
	}
	if rrfuo.mutation.FollowersUIDCleared() {
		_spec.ClearField(rdruserfollowers.FieldFollowersUID, field.TypeInt32)
	}
	if value, ok := rrfuo.mutation.CreateTime(); ok {
		_spec.SetField(rdruserfollowers.FieldCreateTime, field.TypeTime, value)
	}
	_node = &RdRUserFollowers{config: rrfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rrfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rdruserfollowers.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rrfuo.mutation.done = true
	return _node, nil
}
