// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/recommend"
	"github.com/google/uuid"
)

// RecommendUpdate is the builder for updating Recommend entities.
type RecommendUpdate struct {
	config
	hooks    []Hook
	mutation *RecommendMutation
}

// Where appends a list predicates to the RecommendUpdate builder.
func (ru *RecommendUpdate) Where(ps ...predicate.Recommend) *RecommendUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetAppID sets the "app_id" field.
func (ru *RecommendUpdate) SetAppID(u uuid.UUID) *RecommendUpdate {
	ru.mutation.SetAppID(u)
	return ru
}

// SetGoodID sets the "good_id" field.
func (ru *RecommendUpdate) SetGoodID(u uuid.UUID) *RecommendUpdate {
	ru.mutation.SetGoodID(u)
	return ru
}

// SetRecommenderID sets the "recommender_id" field.
func (ru *RecommendUpdate) SetRecommenderID(u uuid.UUID) *RecommendUpdate {
	ru.mutation.SetRecommenderID(u)
	return ru
}

// SetMessage sets the "message" field.
func (ru *RecommendUpdate) SetMessage(s string) *RecommendUpdate {
	ru.mutation.SetMessage(s)
	return ru
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (ru *RecommendUpdate) SetNillableMessage(s *string) *RecommendUpdate {
	if s != nil {
		ru.SetMessage(*s)
	}
	return ru
}

// SetCreateAt sets the "create_at" field.
func (ru *RecommendUpdate) SetCreateAt(i int64) *RecommendUpdate {
	ru.mutation.ResetCreateAt()
	ru.mutation.SetCreateAt(i)
	return ru
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ru *RecommendUpdate) SetNillableCreateAt(i *int64) *RecommendUpdate {
	if i != nil {
		ru.SetCreateAt(*i)
	}
	return ru
}

// AddCreateAt adds i to the "create_at" field.
func (ru *RecommendUpdate) AddCreateAt(i int64) *RecommendUpdate {
	ru.mutation.AddCreateAt(i)
	return ru
}

// SetUpdateAt sets the "update_at" field.
func (ru *RecommendUpdate) SetUpdateAt(i int64) *RecommendUpdate {
	ru.mutation.ResetUpdateAt()
	ru.mutation.SetUpdateAt(i)
	return ru
}

// AddUpdateAt adds i to the "update_at" field.
func (ru *RecommendUpdate) AddUpdateAt(i int64) *RecommendUpdate {
	ru.mutation.AddUpdateAt(i)
	return ru
}

// SetDeleteAt sets the "delete_at" field.
func (ru *RecommendUpdate) SetDeleteAt(i int64) *RecommendUpdate {
	ru.mutation.ResetDeleteAt()
	ru.mutation.SetDeleteAt(i)
	return ru
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ru *RecommendUpdate) SetNillableDeleteAt(i *int64) *RecommendUpdate {
	if i != nil {
		ru.SetDeleteAt(*i)
	}
	return ru
}

// AddDeleteAt adds i to the "delete_at" field.
func (ru *RecommendUpdate) AddDeleteAt(i int64) *RecommendUpdate {
	ru.mutation.AddDeleteAt(i)
	return ru
}

// Mutation returns the RecommendMutation object of the builder.
func (ru *RecommendUpdate) Mutation() *RecommendMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RecommendUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecommendMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RecommendUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecommendUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecommendUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RecommendUpdate) defaults() {
	if _, ok := ru.mutation.UpdateAt(); !ok {
		v := recommend.UpdateDefaultUpdateAt()
		ru.mutation.SetUpdateAt(v)
	}
}

func (ru *RecommendUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recommend.Table,
			Columns: recommend.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: recommend.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: recommend.FieldAppID,
		})
	}
	if value, ok := ru.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: recommend.FieldGoodID,
		})
	}
	if value, ok := ru.mutation.RecommenderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: recommend.FieldRecommenderID,
		})
	}
	if value, ok := ru.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recommend.FieldMessage,
		})
	}
	if value, ok := ru.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommend.FieldCreateAt,
		})
	}
	if value, ok := ru.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommend.FieldCreateAt,
		})
	}
	if value, ok := ru.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommend.FieldUpdateAt,
		})
	}
	if value, ok := ru.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommend.FieldUpdateAt,
		})
	}
	if value, ok := ru.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommend.FieldDeleteAt,
		})
	}
	if value, ok := ru.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommend.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recommend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RecommendUpdateOne is the builder for updating a single Recommend entity.
type RecommendUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecommendMutation
}

// SetAppID sets the "app_id" field.
func (ruo *RecommendUpdateOne) SetAppID(u uuid.UUID) *RecommendUpdateOne {
	ruo.mutation.SetAppID(u)
	return ruo
}

// SetGoodID sets the "good_id" field.
func (ruo *RecommendUpdateOne) SetGoodID(u uuid.UUID) *RecommendUpdateOne {
	ruo.mutation.SetGoodID(u)
	return ruo
}

// SetRecommenderID sets the "recommender_id" field.
func (ruo *RecommendUpdateOne) SetRecommenderID(u uuid.UUID) *RecommendUpdateOne {
	ruo.mutation.SetRecommenderID(u)
	return ruo
}

// SetMessage sets the "message" field.
func (ruo *RecommendUpdateOne) SetMessage(s string) *RecommendUpdateOne {
	ruo.mutation.SetMessage(s)
	return ruo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (ruo *RecommendUpdateOne) SetNillableMessage(s *string) *RecommendUpdateOne {
	if s != nil {
		ruo.SetMessage(*s)
	}
	return ruo
}

// SetCreateAt sets the "create_at" field.
func (ruo *RecommendUpdateOne) SetCreateAt(i int64) *RecommendUpdateOne {
	ruo.mutation.ResetCreateAt()
	ruo.mutation.SetCreateAt(i)
	return ruo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ruo *RecommendUpdateOne) SetNillableCreateAt(i *int64) *RecommendUpdateOne {
	if i != nil {
		ruo.SetCreateAt(*i)
	}
	return ruo
}

// AddCreateAt adds i to the "create_at" field.
func (ruo *RecommendUpdateOne) AddCreateAt(i int64) *RecommendUpdateOne {
	ruo.mutation.AddCreateAt(i)
	return ruo
}

// SetUpdateAt sets the "update_at" field.
func (ruo *RecommendUpdateOne) SetUpdateAt(i int64) *RecommendUpdateOne {
	ruo.mutation.ResetUpdateAt()
	ruo.mutation.SetUpdateAt(i)
	return ruo
}

// AddUpdateAt adds i to the "update_at" field.
func (ruo *RecommendUpdateOne) AddUpdateAt(i int64) *RecommendUpdateOne {
	ruo.mutation.AddUpdateAt(i)
	return ruo
}

// SetDeleteAt sets the "delete_at" field.
func (ruo *RecommendUpdateOne) SetDeleteAt(i int64) *RecommendUpdateOne {
	ruo.mutation.ResetDeleteAt()
	ruo.mutation.SetDeleteAt(i)
	return ruo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ruo *RecommendUpdateOne) SetNillableDeleteAt(i *int64) *RecommendUpdateOne {
	if i != nil {
		ruo.SetDeleteAt(*i)
	}
	return ruo
}

// AddDeleteAt adds i to the "delete_at" field.
func (ruo *RecommendUpdateOne) AddDeleteAt(i int64) *RecommendUpdateOne {
	ruo.mutation.AddDeleteAt(i)
	return ruo
}

// Mutation returns the RecommendMutation object of the builder.
func (ruo *RecommendUpdateOne) Mutation() *RecommendMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RecommendUpdateOne) Select(field string, fields ...string) *RecommendUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Recommend entity.
func (ruo *RecommendUpdateOne) Save(ctx context.Context) (*Recommend, error) {
	var (
		err  error
		node *Recommend
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecommendMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecommendUpdateOne) SaveX(ctx context.Context) *Recommend {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RecommendUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecommendUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RecommendUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateAt(); !ok {
		v := recommend.UpdateDefaultUpdateAt()
		ruo.mutation.SetUpdateAt(v)
	}
}

func (ruo *RecommendUpdateOne) sqlSave(ctx context.Context) (_node *Recommend, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recommend.Table,
			Columns: recommend.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: recommend.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Recommend.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recommend.FieldID)
		for _, f := range fields {
			if !recommend.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != recommend.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: recommend.FieldAppID,
		})
	}
	if value, ok := ruo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: recommend.FieldGoodID,
		})
	}
	if value, ok := ruo.mutation.RecommenderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: recommend.FieldRecommenderID,
		})
	}
	if value, ok := ruo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recommend.FieldMessage,
		})
	}
	if value, ok := ruo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommend.FieldCreateAt,
		})
	}
	if value, ok := ruo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommend.FieldCreateAt,
		})
	}
	if value, ok := ruo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommend.FieldUpdateAt,
		})
	}
	if value, ok := ruo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommend.FieldUpdateAt,
		})
	}
	if value, ok := ruo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommend.FieldDeleteAt,
		})
	}
	if value, ok := ruo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommend.FieldDeleteAt,
		})
	}
	_node = &Recommend{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recommend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
