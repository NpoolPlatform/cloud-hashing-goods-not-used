// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/apptargetarea"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
)

// AppTargetAreaDelete is the builder for deleting a AppTargetArea entity.
type AppTargetAreaDelete struct {
	config
	hooks    []Hook
	mutation *AppTargetAreaMutation
}

// Where appends a list predicates to the AppTargetAreaDelete builder.
func (atad *AppTargetAreaDelete) Where(ps ...predicate.AppTargetArea) *AppTargetAreaDelete {
	atad.mutation.Where(ps...)
	return atad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atad *AppTargetAreaDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atad.hooks) == 0 {
		affected, err = atad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppTargetAreaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atad.mutation = mutation
			affected, err = atad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atad.hooks) - 1; i >= 0; i-- {
			if atad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (atad *AppTargetAreaDelete) ExecX(ctx context.Context) int {
	n, err := atad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atad *AppTargetAreaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: apptargetarea.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: apptargetarea.FieldID,
			},
		},
	}
	if ps := atad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, atad.driver, _spec)
}

// AppTargetAreaDeleteOne is the builder for deleting a single AppTargetArea entity.
type AppTargetAreaDeleteOne struct {
	atad *AppTargetAreaDelete
}

// Exec executes the deletion query.
func (atado *AppTargetAreaDeleteOne) Exec(ctx context.Context) error {
	n, err := atado.atad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{apptargetarea.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atado *AppTargetAreaDeleteOne) ExecX(ctx context.Context) {
	atado.atad.ExecX(ctx)
}
