// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/predicate"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/rdmoment"
)

// RdMomentDelete is the builder for deleting a RdMoment entity.
type RdMomentDelete struct {
	config
	hooks    []Hook
	mutation *RdMomentMutation
}

// Where appends a list predicates to the RdMomentDelete builder.
func (rmd *RdMomentDelete) Where(ps ...predicate.RdMoment) *RdMomentDelete {
	rmd.mutation.Where(ps...)
	return rmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rmd *RdMomentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, RdMomentMutation](ctx, rmd.sqlExec, rmd.mutation, rmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rmd *RdMomentDelete) ExecX(ctx context.Context) int {
	n, err := rmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rmd *RdMomentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: rdmoment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: rdmoment.FieldID,
			},
		},
	}
	if ps := rmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rmd.mutation.done = true
	return affected, err
}

// RdMomentDeleteOne is the builder for deleting a single RdMoment entity.
type RdMomentDeleteOne struct {
	rmd *RdMomentDelete
}

// Exec executes the deletion query.
func (rmdo *RdMomentDeleteOne) Exec(ctx context.Context) error {
	n, err := rmdo.rmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rdmoment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rmdo *RdMomentDeleteOne) ExecX(ctx context.Context) {
	rmdo.rmd.ExecX(ctx)
}