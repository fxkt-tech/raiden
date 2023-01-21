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
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/biz"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/predicate"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/rdmessage"
)

// RdMessageUpdate is the builder for updating RdMessage entities.
type RdMessageUpdate struct {
	config
	hooks    []Hook
	mutation *RdMessageMutation
}

// Where appends a list predicates to the RdMessageUpdate builder.
func (rmu *RdMessageUpdate) Where(ps ...predicate.RdMessage) *RdMessageUpdate {
	rmu.mutation.Where(ps...)
	return rmu
}

// SetSenderUID sets the "sender_uid" field.
func (rmu *RdMessageUpdate) SetSenderUID(i int32) *RdMessageUpdate {
	rmu.mutation.ResetSenderUID()
	rmu.mutation.SetSenderUID(i)
	return rmu
}

// SetNillableSenderUID sets the "sender_uid" field if the given value is not nil.
func (rmu *RdMessageUpdate) SetNillableSenderUID(i *int32) *RdMessageUpdate {
	if i != nil {
		rmu.SetSenderUID(*i)
	}
	return rmu
}

// AddSenderUID adds i to the "sender_uid" field.
func (rmu *RdMessageUpdate) AddSenderUID(i int32) *RdMessageUpdate {
	rmu.mutation.AddSenderUID(i)
	return rmu
}

// ClearSenderUID clears the value of the "sender_uid" field.
func (rmu *RdMessageUpdate) ClearSenderUID() *RdMessageUpdate {
	rmu.mutation.ClearSenderUID()
	return rmu
}

// SetRecverUID sets the "recver_uid" field.
func (rmu *RdMessageUpdate) SetRecverUID(i int32) *RdMessageUpdate {
	rmu.mutation.ResetRecverUID()
	rmu.mutation.SetRecverUID(i)
	return rmu
}

// SetNillableRecverUID sets the "recver_uid" field if the given value is not nil.
func (rmu *RdMessageUpdate) SetNillableRecverUID(i *int32) *RdMessageUpdate {
	if i != nil {
		rmu.SetRecverUID(*i)
	}
	return rmu
}

// AddRecverUID adds i to the "recver_uid" field.
func (rmu *RdMessageUpdate) AddRecverUID(i int32) *RdMessageUpdate {
	rmu.mutation.AddRecverUID(i)
	return rmu
}

// ClearRecverUID clears the value of the "recver_uid" field.
func (rmu *RdMessageUpdate) ClearRecverUID() *RdMessageUpdate {
	rmu.mutation.ClearRecverUID()
	return rmu
}

// SetContent sets the "content" field.
func (rmu *RdMessageUpdate) SetContent(b *biz.Content) *RdMessageUpdate {
	rmu.mutation.SetContent(b)
	return rmu
}

// SetCreateTime sets the "create_time" field.
func (rmu *RdMessageUpdate) SetCreateTime(t time.Time) *RdMessageUpdate {
	rmu.mutation.SetCreateTime(t)
	return rmu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rmu *RdMessageUpdate) SetNillableCreateTime(t *time.Time) *RdMessageUpdate {
	if t != nil {
		rmu.SetCreateTime(*t)
	}
	return rmu
}

// Mutation returns the RdMessageMutation object of the builder.
func (rmu *RdMessageUpdate) Mutation() *RdMessageMutation {
	return rmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rmu *RdMessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, RdMessageMutation](ctx, rmu.sqlSave, rmu.mutation, rmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rmu *RdMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := rmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rmu *RdMessageUpdate) Exec(ctx context.Context) error {
	_, err := rmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmu *RdMessageUpdate) ExecX(ctx context.Context) {
	if err := rmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rmu *RdMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rdmessage.Table,
			Columns: rdmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: rdmessage.FieldID,
			},
		},
	}
	if ps := rmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rmu.mutation.SenderUID(); ok {
		_spec.SetField(rdmessage.FieldSenderUID, field.TypeInt32, value)
	}
	if value, ok := rmu.mutation.AddedSenderUID(); ok {
		_spec.AddField(rdmessage.FieldSenderUID, field.TypeInt32, value)
	}
	if rmu.mutation.SenderUIDCleared() {
		_spec.ClearField(rdmessage.FieldSenderUID, field.TypeInt32)
	}
	if value, ok := rmu.mutation.RecverUID(); ok {
		_spec.SetField(rdmessage.FieldRecverUID, field.TypeInt32, value)
	}
	if value, ok := rmu.mutation.AddedRecverUID(); ok {
		_spec.AddField(rdmessage.FieldRecverUID, field.TypeInt32, value)
	}
	if rmu.mutation.RecverUIDCleared() {
		_spec.ClearField(rdmessage.FieldRecverUID, field.TypeInt32)
	}
	if value, ok := rmu.mutation.Content(); ok {
		_spec.SetField(rdmessage.FieldContent, field.TypeJSON, value)
	}
	if value, ok := rmu.mutation.CreateTime(); ok {
		_spec.SetField(rdmessage.FieldCreateTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rdmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rmu.mutation.done = true
	return n, nil
}

// RdMessageUpdateOne is the builder for updating a single RdMessage entity.
type RdMessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RdMessageMutation
}

// SetSenderUID sets the "sender_uid" field.
func (rmuo *RdMessageUpdateOne) SetSenderUID(i int32) *RdMessageUpdateOne {
	rmuo.mutation.ResetSenderUID()
	rmuo.mutation.SetSenderUID(i)
	return rmuo
}

// SetNillableSenderUID sets the "sender_uid" field if the given value is not nil.
func (rmuo *RdMessageUpdateOne) SetNillableSenderUID(i *int32) *RdMessageUpdateOne {
	if i != nil {
		rmuo.SetSenderUID(*i)
	}
	return rmuo
}

// AddSenderUID adds i to the "sender_uid" field.
func (rmuo *RdMessageUpdateOne) AddSenderUID(i int32) *RdMessageUpdateOne {
	rmuo.mutation.AddSenderUID(i)
	return rmuo
}

// ClearSenderUID clears the value of the "sender_uid" field.
func (rmuo *RdMessageUpdateOne) ClearSenderUID() *RdMessageUpdateOne {
	rmuo.mutation.ClearSenderUID()
	return rmuo
}

// SetRecverUID sets the "recver_uid" field.
func (rmuo *RdMessageUpdateOne) SetRecverUID(i int32) *RdMessageUpdateOne {
	rmuo.mutation.ResetRecverUID()
	rmuo.mutation.SetRecverUID(i)
	return rmuo
}

// SetNillableRecverUID sets the "recver_uid" field if the given value is not nil.
func (rmuo *RdMessageUpdateOne) SetNillableRecverUID(i *int32) *RdMessageUpdateOne {
	if i != nil {
		rmuo.SetRecverUID(*i)
	}
	return rmuo
}

// AddRecverUID adds i to the "recver_uid" field.
func (rmuo *RdMessageUpdateOne) AddRecverUID(i int32) *RdMessageUpdateOne {
	rmuo.mutation.AddRecverUID(i)
	return rmuo
}

// ClearRecverUID clears the value of the "recver_uid" field.
func (rmuo *RdMessageUpdateOne) ClearRecverUID() *RdMessageUpdateOne {
	rmuo.mutation.ClearRecverUID()
	return rmuo
}

// SetContent sets the "content" field.
func (rmuo *RdMessageUpdateOne) SetContent(b *biz.Content) *RdMessageUpdateOne {
	rmuo.mutation.SetContent(b)
	return rmuo
}

// SetCreateTime sets the "create_time" field.
func (rmuo *RdMessageUpdateOne) SetCreateTime(t time.Time) *RdMessageUpdateOne {
	rmuo.mutation.SetCreateTime(t)
	return rmuo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rmuo *RdMessageUpdateOne) SetNillableCreateTime(t *time.Time) *RdMessageUpdateOne {
	if t != nil {
		rmuo.SetCreateTime(*t)
	}
	return rmuo
}

// Mutation returns the RdMessageMutation object of the builder.
func (rmuo *RdMessageUpdateOne) Mutation() *RdMessageMutation {
	return rmuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rmuo *RdMessageUpdateOne) Select(field string, fields ...string) *RdMessageUpdateOne {
	rmuo.fields = append([]string{field}, fields...)
	return rmuo
}

// Save executes the query and returns the updated RdMessage entity.
func (rmuo *RdMessageUpdateOne) Save(ctx context.Context) (*RdMessage, error) {
	return withHooks[*RdMessage, RdMessageMutation](ctx, rmuo.sqlSave, rmuo.mutation, rmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rmuo *RdMessageUpdateOne) SaveX(ctx context.Context) *RdMessage {
	node, err := rmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rmuo *RdMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := rmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmuo *RdMessageUpdateOne) ExecX(ctx context.Context) {
	if err := rmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rmuo *RdMessageUpdateOne) sqlSave(ctx context.Context) (_node *RdMessage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rdmessage.Table,
			Columns: rdmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: rdmessage.FieldID,
			},
		},
	}
	id, ok := rmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RdMessage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rdmessage.FieldID)
		for _, f := range fields {
			if !rdmessage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rdmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rmuo.mutation.SenderUID(); ok {
		_spec.SetField(rdmessage.FieldSenderUID, field.TypeInt32, value)
	}
	if value, ok := rmuo.mutation.AddedSenderUID(); ok {
		_spec.AddField(rdmessage.FieldSenderUID, field.TypeInt32, value)
	}
	if rmuo.mutation.SenderUIDCleared() {
		_spec.ClearField(rdmessage.FieldSenderUID, field.TypeInt32)
	}
	if value, ok := rmuo.mutation.RecverUID(); ok {
		_spec.SetField(rdmessage.FieldRecverUID, field.TypeInt32, value)
	}
	if value, ok := rmuo.mutation.AddedRecverUID(); ok {
		_spec.AddField(rdmessage.FieldRecverUID, field.TypeInt32, value)
	}
	if rmuo.mutation.RecverUIDCleared() {
		_spec.ClearField(rdmessage.FieldRecverUID, field.TypeInt32)
	}
	if value, ok := rmuo.mutation.Content(); ok {
		_spec.SetField(rdmessage.FieldContent, field.TypeJSON, value)
	}
	if value, ok := rmuo.mutation.CreateTime(); ok {
		_spec.SetField(rdmessage.FieldCreateTime, field.TypeTime, value)
	}
	_node = &RdMessage{config: rmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rdmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rmuo.mutation.done = true
	return _node, nil
}