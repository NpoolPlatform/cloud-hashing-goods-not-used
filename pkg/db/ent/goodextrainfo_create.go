// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodextrainfo"
	"github.com/google/uuid"
)

// GoodExtraInfoCreate is the builder for creating a GoodExtraInfo entity.
type GoodExtraInfoCreate struct {
	config
	mutation *GoodExtraInfoMutation
	hooks    []Hook
}

// SetGoodID sets the "good_id" field.
func (geic *GoodExtraInfoCreate) SetGoodID(u uuid.UUID) *GoodExtraInfoCreate {
	geic.mutation.SetGoodID(u)
	return geic
}

// SetPosters sets the "posters" field.
func (geic *GoodExtraInfoCreate) SetPosters(s []string) *GoodExtraInfoCreate {
	geic.mutation.SetPosters(s)
	return geic
}

// SetLabels sets the "labels" field.
func (geic *GoodExtraInfoCreate) SetLabels(s []string) *GoodExtraInfoCreate {
	geic.mutation.SetLabels(s)
	return geic
}

// SetOutSale sets the "out_sale" field.
func (geic *GoodExtraInfoCreate) SetOutSale(b bool) *GoodExtraInfoCreate {
	geic.mutation.SetOutSale(b)
	return geic
}

// SetPreSale sets the "pre_sale" field.
func (geic *GoodExtraInfoCreate) SetPreSale(b bool) *GoodExtraInfoCreate {
	geic.mutation.SetPreSale(b)
	return geic
}

// SetVoteCount sets the "vote_count" field.
func (geic *GoodExtraInfoCreate) SetVoteCount(u uint32) *GoodExtraInfoCreate {
	geic.mutation.SetVoteCount(u)
	return geic
}

// SetRating sets the "rating" field.
func (geic *GoodExtraInfoCreate) SetRating(f float32) *GoodExtraInfoCreate {
	geic.mutation.SetRating(f)
	return geic
}

// SetCreateAt sets the "create_at" field.
func (geic *GoodExtraInfoCreate) SetCreateAt(i int64) *GoodExtraInfoCreate {
	geic.mutation.SetCreateAt(i)
	return geic
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (geic *GoodExtraInfoCreate) SetNillableCreateAt(i *int64) *GoodExtraInfoCreate {
	if i != nil {
		geic.SetCreateAt(*i)
	}
	return geic
}

// SetUpdateAt sets the "update_at" field.
func (geic *GoodExtraInfoCreate) SetUpdateAt(i int64) *GoodExtraInfoCreate {
	geic.mutation.SetUpdateAt(i)
	return geic
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (geic *GoodExtraInfoCreate) SetNillableUpdateAt(i *int64) *GoodExtraInfoCreate {
	if i != nil {
		geic.SetUpdateAt(*i)
	}
	return geic
}

// SetDeleteAt sets the "delete_at" field.
func (geic *GoodExtraInfoCreate) SetDeleteAt(i int64) *GoodExtraInfoCreate {
	geic.mutation.SetDeleteAt(i)
	return geic
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (geic *GoodExtraInfoCreate) SetNillableDeleteAt(i *int64) *GoodExtraInfoCreate {
	if i != nil {
		geic.SetDeleteAt(*i)
	}
	return geic
}

// SetID sets the "id" field.
func (geic *GoodExtraInfoCreate) SetID(u uuid.UUID) *GoodExtraInfoCreate {
	geic.mutation.SetID(u)
	return geic
}

// Mutation returns the GoodExtraInfoMutation object of the builder.
func (geic *GoodExtraInfoCreate) Mutation() *GoodExtraInfoMutation {
	return geic.mutation
}

// Save creates the GoodExtraInfo in the database.
func (geic *GoodExtraInfoCreate) Save(ctx context.Context) (*GoodExtraInfo, error) {
	var (
		err  error
		node *GoodExtraInfo
	)
	geic.defaults()
	if len(geic.hooks) == 0 {
		if err = geic.check(); err != nil {
			return nil, err
		}
		node, err = geic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodExtraInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = geic.check(); err != nil {
				return nil, err
			}
			geic.mutation = mutation
			if node, err = geic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(geic.hooks) - 1; i >= 0; i-- {
			if geic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = geic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, geic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (geic *GoodExtraInfoCreate) SaveX(ctx context.Context) *GoodExtraInfo {
	v, err := geic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (geic *GoodExtraInfoCreate) Exec(ctx context.Context) error {
	_, err := geic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (geic *GoodExtraInfoCreate) ExecX(ctx context.Context) {
	if err := geic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (geic *GoodExtraInfoCreate) defaults() {
	if _, ok := geic.mutation.CreateAt(); !ok {
		v := goodextrainfo.DefaultCreateAt()
		geic.mutation.SetCreateAt(v)
	}
	if _, ok := geic.mutation.UpdateAt(); !ok {
		v := goodextrainfo.DefaultUpdateAt()
		geic.mutation.SetUpdateAt(v)
	}
	if _, ok := geic.mutation.DeleteAt(); !ok {
		v := goodextrainfo.DefaultDeleteAt()
		geic.mutation.SetDeleteAt(v)
	}
	if _, ok := geic.mutation.ID(); !ok {
		v := goodextrainfo.DefaultID()
		geic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (geic *GoodExtraInfoCreate) check() error {
	if _, ok := geic.mutation.GoodID(); !ok {
		return &ValidationError{Name: "good_id", err: errors.New(`ent: missing required field "good_id"`)}
	}
	if _, ok := geic.mutation.Posters(); !ok {
		return &ValidationError{Name: "posters", err: errors.New(`ent: missing required field "posters"`)}
	}
	if _, ok := geic.mutation.Labels(); !ok {
		return &ValidationError{Name: "labels", err: errors.New(`ent: missing required field "labels"`)}
	}
	if _, ok := geic.mutation.OutSale(); !ok {
		return &ValidationError{Name: "out_sale", err: errors.New(`ent: missing required field "out_sale"`)}
	}
	if _, ok := geic.mutation.PreSale(); !ok {
		return &ValidationError{Name: "pre_sale", err: errors.New(`ent: missing required field "pre_sale"`)}
	}
	if _, ok := geic.mutation.VoteCount(); !ok {
		return &ValidationError{Name: "vote_count", err: errors.New(`ent: missing required field "vote_count"`)}
	}
	if _, ok := geic.mutation.Rating(); !ok {
		return &ValidationError{Name: "rating", err: errors.New(`ent: missing required field "rating"`)}
	}
	if _, ok := geic.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := geic.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := geic.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (geic *GoodExtraInfoCreate) sqlSave(ctx context.Context) (*GoodExtraInfo, error) {
	_node, _spec := geic.createSpec()
	if err := sqlgraph.CreateNode(ctx, geic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (geic *GoodExtraInfoCreate) createSpec() (*GoodExtraInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &GoodExtraInfo{config: geic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: goodextrainfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodextrainfo.FieldID,
			},
		}
	)
	if id, ok := geic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := geic.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodextrainfo.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := geic.mutation.Posters(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: goodextrainfo.FieldPosters,
		})
		_node.Posters = value
	}
	if value, ok := geic.mutation.Labels(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: goodextrainfo.FieldLabels,
		})
		_node.Labels = value
	}
	if value, ok := geic.mutation.OutSale(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: goodextrainfo.FieldOutSale,
		})
		_node.OutSale = value
	}
	if value, ok := geic.mutation.PreSale(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: goodextrainfo.FieldPreSale,
		})
		_node.PreSale = value
	}
	if value, ok := geic.mutation.VoteCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodextrainfo.FieldVoteCount,
		})
		_node.VoteCount = value
	}
	if value, ok := geic.mutation.Rating(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: goodextrainfo.FieldRating,
		})
		_node.Rating = value
	}
	if value, ok := geic.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodextrainfo.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := geic.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodextrainfo.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := geic.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodextrainfo.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// GoodExtraInfoCreateBulk is the builder for creating many GoodExtraInfo entities in bulk.
type GoodExtraInfoCreateBulk struct {
	config
	builders []*GoodExtraInfoCreate
}

// Save creates the GoodExtraInfo entities in the database.
func (geicb *GoodExtraInfoCreateBulk) Save(ctx context.Context) ([]*GoodExtraInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(geicb.builders))
	nodes := make([]*GoodExtraInfo, len(geicb.builders))
	mutators := make([]Mutator, len(geicb.builders))
	for i := range geicb.builders {
		func(i int, root context.Context) {
			builder := geicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoodExtraInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, geicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, geicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, geicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (geicb *GoodExtraInfoCreateBulk) SaveX(ctx context.Context) []*GoodExtraInfo {
	v, err := geicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (geicb *GoodExtraInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := geicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (geicb *GoodExtraInfoCreateBulk) ExecX(ctx context.Context) {
	if err := geicb.Exec(ctx); err != nil {
		panic(err)
	}
}
