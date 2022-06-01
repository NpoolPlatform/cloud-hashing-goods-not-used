// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgood"
	"github.com/google/uuid"
)

// AppGoodCreate is the builder for creating a AppGood entity.
type AppGoodCreate struct {
	config
	mutation *AppGoodMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (agc *AppGoodCreate) SetAppID(u uuid.UUID) *AppGoodCreate {
	agc.mutation.SetAppID(u)
	return agc
}

// SetGoodID sets the "good_id" field.
func (agc *AppGoodCreate) SetGoodID(u uuid.UUID) *AppGoodCreate {
	agc.mutation.SetGoodID(u)
	return agc
}

// SetOnline sets the "online" field.
func (agc *AppGoodCreate) SetOnline(b bool) *AppGoodCreate {
	agc.mutation.SetOnline(b)
	return agc
}

// SetNillableOnline sets the "online" field if the given value is not nil.
func (agc *AppGoodCreate) SetNillableOnline(b *bool) *AppGoodCreate {
	if b != nil {
		agc.SetOnline(*b)
	}
	return agc
}

// SetInitAreaStrategy sets the "init_area_strategy" field.
func (agc *AppGoodCreate) SetInitAreaStrategy(aas appgood.InitAreaStrategy) *AppGoodCreate {
	agc.mutation.SetInitAreaStrategy(aas)
	return agc
}

// SetPrice sets the "price" field.
func (agc *AppGoodCreate) SetPrice(u uint64) *AppGoodCreate {
	agc.mutation.SetPrice(u)
	return agc
}

// SetDisplayIndex sets the "display_index" field.
func (agc *AppGoodCreate) SetDisplayIndex(u uint32) *AppGoodCreate {
	agc.mutation.SetDisplayIndex(u)
	return agc
}

// SetVisible sets the "visible" field.
func (agc *AppGoodCreate) SetVisible(b bool) *AppGoodCreate {
	agc.mutation.SetVisible(b)
	return agc
}

// SetNillableVisible sets the "visible" field if the given value is not nil.
func (agc *AppGoodCreate) SetNillableVisible(b *bool) *AppGoodCreate {
	if b != nil {
		agc.SetVisible(*b)
	}
	return agc
}

// SetCreateAt sets the "create_at" field.
func (agc *AppGoodCreate) SetCreateAt(u uint32) *AppGoodCreate {
	agc.mutation.SetCreateAt(u)
	return agc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (agc *AppGoodCreate) SetNillableCreateAt(u *uint32) *AppGoodCreate {
	if u != nil {
		agc.SetCreateAt(*u)
	}
	return agc
}

// SetUpdateAt sets the "update_at" field.
func (agc *AppGoodCreate) SetUpdateAt(u uint32) *AppGoodCreate {
	agc.mutation.SetUpdateAt(u)
	return agc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (agc *AppGoodCreate) SetNillableUpdateAt(u *uint32) *AppGoodCreate {
	if u != nil {
		agc.SetUpdateAt(*u)
	}
	return agc
}

// SetDeleteAt sets the "delete_at" field.
func (agc *AppGoodCreate) SetDeleteAt(u uint32) *AppGoodCreate {
	agc.mutation.SetDeleteAt(u)
	return agc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (agc *AppGoodCreate) SetNillableDeleteAt(u *uint32) *AppGoodCreate {
	if u != nil {
		agc.SetDeleteAt(*u)
	}
	return agc
}

// SetID sets the "id" field.
func (agc *AppGoodCreate) SetID(u uuid.UUID) *AppGoodCreate {
	agc.mutation.SetID(u)
	return agc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (agc *AppGoodCreate) SetNillableID(u *uuid.UUID) *AppGoodCreate {
	if u != nil {
		agc.SetID(*u)
	}
	return agc
}

// Mutation returns the AppGoodMutation object of the builder.
func (agc *AppGoodCreate) Mutation() *AppGoodMutation {
	return agc.mutation
}

// Save creates the AppGood in the database.
func (agc *AppGoodCreate) Save(ctx context.Context) (*AppGood, error) {
	var (
		err  error
		node *AppGood
	)
	agc.defaults()
	if len(agc.hooks) == 0 {
		if err = agc.check(); err != nil {
			return nil, err
		}
		node, err = agc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = agc.check(); err != nil {
				return nil, err
			}
			agc.mutation = mutation
			if node, err = agc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(agc.hooks) - 1; i >= 0; i-- {
			if agc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, agc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (agc *AppGoodCreate) SaveX(ctx context.Context) *AppGood {
	v, err := agc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (agc *AppGoodCreate) Exec(ctx context.Context) error {
	_, err := agc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agc *AppGoodCreate) ExecX(ctx context.Context) {
	if err := agc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (agc *AppGoodCreate) defaults() {
	if _, ok := agc.mutation.Online(); !ok {
		v := appgood.DefaultOnline
		agc.mutation.SetOnline(v)
	}
	if _, ok := agc.mutation.Visible(); !ok {
		v := appgood.DefaultVisible
		agc.mutation.SetVisible(v)
	}
	if _, ok := agc.mutation.CreateAt(); !ok {
		v := appgood.DefaultCreateAt()
		agc.mutation.SetCreateAt(v)
	}
	if _, ok := agc.mutation.UpdateAt(); !ok {
		v := appgood.DefaultUpdateAt()
		agc.mutation.SetUpdateAt(v)
	}
	if _, ok := agc.mutation.DeleteAt(); !ok {
		v := appgood.DefaultDeleteAt()
		agc.mutation.SetDeleteAt(v)
	}
	if _, ok := agc.mutation.ID(); !ok {
		v := appgood.DefaultID()
		agc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (agc *AppGoodCreate) check() error {
	if _, ok := agc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AppGood.app_id"`)}
	}
	if _, ok := agc.mutation.GoodID(); !ok {
		return &ValidationError{Name: "good_id", err: errors.New(`ent: missing required field "AppGood.good_id"`)}
	}
	if _, ok := agc.mutation.Online(); !ok {
		return &ValidationError{Name: "online", err: errors.New(`ent: missing required field "AppGood.online"`)}
	}
	if _, ok := agc.mutation.InitAreaStrategy(); !ok {
		return &ValidationError{Name: "init_area_strategy", err: errors.New(`ent: missing required field "AppGood.init_area_strategy"`)}
	}
	if v, ok := agc.mutation.InitAreaStrategy(); ok {
		if err := appgood.InitAreaStrategyValidator(v); err != nil {
			return &ValidationError{Name: "init_area_strategy", err: fmt.Errorf(`ent: validator failed for field "AppGood.init_area_strategy": %w`, err)}
		}
	}
	if _, ok := agc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "AppGood.price"`)}
	}
	if _, ok := agc.mutation.DisplayIndex(); !ok {
		return &ValidationError{Name: "display_index", err: errors.New(`ent: missing required field "AppGood.display_index"`)}
	}
	if _, ok := agc.mutation.Visible(); !ok {
		return &ValidationError{Name: "visible", err: errors.New(`ent: missing required field "AppGood.visible"`)}
	}
	if _, ok := agc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "AppGood.create_at"`)}
	}
	if _, ok := agc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "AppGood.update_at"`)}
	}
	if _, ok := agc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "AppGood.delete_at"`)}
	}
	return nil
}

func (agc *AppGoodCreate) sqlSave(ctx context.Context) (*AppGood, error) {
	_node, _spec := agc.createSpec()
	if err := sqlgraph.CreateNode(ctx, agc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (agc *AppGoodCreate) createSpec() (*AppGood, *sqlgraph.CreateSpec) {
	var (
		_node = &AppGood{config: agc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appgood.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appgood.FieldID,
			},
		}
	)
	_spec.OnConflict = agc.conflict
	if id, ok := agc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := agc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgood.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := agc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgood.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := agc.mutation.Online(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: appgood.FieldOnline,
		})
		_node.Online = value
	}
	if value, ok := agc.mutation.InitAreaStrategy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: appgood.FieldInitAreaStrategy,
		})
		_node.InitAreaStrategy = value
	}
	if value, ok := agc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: appgood.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := agc.mutation.DisplayIndex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgood.FieldDisplayIndex,
		})
		_node.DisplayIndex = value
	}
	if value, ok := agc.mutation.Visible(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: appgood.FieldVisible,
		})
		_node.Visible = value
	}
	if value, ok := agc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgood.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := agc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgood.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := agc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgood.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppGood.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppGoodUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (agc *AppGoodCreate) OnConflict(opts ...sql.ConflictOption) *AppGoodUpsertOne {
	agc.conflict = opts
	return &AppGoodUpsertOne{
		create: agc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppGood.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (agc *AppGoodCreate) OnConflictColumns(columns ...string) *AppGoodUpsertOne {
	agc.conflict = append(agc.conflict, sql.ConflictColumns(columns...))
	return &AppGoodUpsertOne{
		create: agc,
	}
}

type (
	// AppGoodUpsertOne is the builder for "upsert"-ing
	//  one AppGood node.
	AppGoodUpsertOne struct {
		create *AppGoodCreate
	}

	// AppGoodUpsert is the "OnConflict" setter.
	AppGoodUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *AppGoodUpsert) SetAppID(v uuid.UUID) *AppGoodUpsert {
	u.Set(appgood.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppGoodUpsert) UpdateAppID() *AppGoodUpsert {
	u.SetExcluded(appgood.FieldAppID)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *AppGoodUpsert) SetGoodID(v uuid.UUID) *AppGoodUpsert {
	u.Set(appgood.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *AppGoodUpsert) UpdateGoodID() *AppGoodUpsert {
	u.SetExcluded(appgood.FieldGoodID)
	return u
}

// SetOnline sets the "online" field.
func (u *AppGoodUpsert) SetOnline(v bool) *AppGoodUpsert {
	u.Set(appgood.FieldOnline, v)
	return u
}

// UpdateOnline sets the "online" field to the value that was provided on create.
func (u *AppGoodUpsert) UpdateOnline() *AppGoodUpsert {
	u.SetExcluded(appgood.FieldOnline)
	return u
}

// SetInitAreaStrategy sets the "init_area_strategy" field.
func (u *AppGoodUpsert) SetInitAreaStrategy(v appgood.InitAreaStrategy) *AppGoodUpsert {
	u.Set(appgood.FieldInitAreaStrategy, v)
	return u
}

// UpdateInitAreaStrategy sets the "init_area_strategy" field to the value that was provided on create.
func (u *AppGoodUpsert) UpdateInitAreaStrategy() *AppGoodUpsert {
	u.SetExcluded(appgood.FieldInitAreaStrategy)
	return u
}

// SetPrice sets the "price" field.
func (u *AppGoodUpsert) SetPrice(v uint64) *AppGoodUpsert {
	u.Set(appgood.FieldPrice, v)
	return u
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *AppGoodUpsert) UpdatePrice() *AppGoodUpsert {
	u.SetExcluded(appgood.FieldPrice)
	return u
}

// AddPrice adds v to the "price" field.
func (u *AppGoodUpsert) AddPrice(v uint64) *AppGoodUpsert {
	u.Add(appgood.FieldPrice, v)
	return u
}

// SetDisplayIndex sets the "display_index" field.
func (u *AppGoodUpsert) SetDisplayIndex(v uint32) *AppGoodUpsert {
	u.Set(appgood.FieldDisplayIndex, v)
	return u
}

// UpdateDisplayIndex sets the "display_index" field to the value that was provided on create.
func (u *AppGoodUpsert) UpdateDisplayIndex() *AppGoodUpsert {
	u.SetExcluded(appgood.FieldDisplayIndex)
	return u
}

// AddDisplayIndex adds v to the "display_index" field.
func (u *AppGoodUpsert) AddDisplayIndex(v uint32) *AppGoodUpsert {
	u.Add(appgood.FieldDisplayIndex, v)
	return u
}

// SetVisible sets the "visible" field.
func (u *AppGoodUpsert) SetVisible(v bool) *AppGoodUpsert {
	u.Set(appgood.FieldVisible, v)
	return u
}

// UpdateVisible sets the "visible" field to the value that was provided on create.
func (u *AppGoodUpsert) UpdateVisible() *AppGoodUpsert {
	u.SetExcluded(appgood.FieldVisible)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *AppGoodUpsert) SetCreateAt(v uint32) *AppGoodUpsert {
	u.Set(appgood.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppGoodUpsert) UpdateCreateAt() *AppGoodUpsert {
	u.SetExcluded(appgood.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppGoodUpsert) AddCreateAt(v uint32) *AppGoodUpsert {
	u.Add(appgood.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *AppGoodUpsert) SetUpdateAt(v uint32) *AppGoodUpsert {
	u.Set(appgood.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppGoodUpsert) UpdateUpdateAt() *AppGoodUpsert {
	u.SetExcluded(appgood.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppGoodUpsert) AddUpdateAt(v uint32) *AppGoodUpsert {
	u.Add(appgood.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppGoodUpsert) SetDeleteAt(v uint32) *AppGoodUpsert {
	u.Set(appgood.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppGoodUpsert) UpdateDeleteAt() *AppGoodUpsert {
	u.SetExcluded(appgood.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppGoodUpsert) AddDeleteAt(v uint32) *AppGoodUpsert {
	u.Add(appgood.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppGood.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appgood.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppGoodUpsertOne) UpdateNewValues() *AppGoodUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appgood.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppGood.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppGoodUpsertOne) Ignore() *AppGoodUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppGoodUpsertOne) DoNothing() *AppGoodUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppGoodCreate.OnConflict
// documentation for more info.
func (u *AppGoodUpsertOne) Update(set func(*AppGoodUpsert)) *AppGoodUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppGoodUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppGoodUpsertOne) SetAppID(v uuid.UUID) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppGoodUpsertOne) UpdateAppID() *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateAppID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *AppGoodUpsertOne) SetGoodID(v uuid.UUID) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *AppGoodUpsertOne) UpdateGoodID() *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateGoodID()
	})
}

// SetOnline sets the "online" field.
func (u *AppGoodUpsertOne) SetOnline(v bool) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetOnline(v)
	})
}

// UpdateOnline sets the "online" field to the value that was provided on create.
func (u *AppGoodUpsertOne) UpdateOnline() *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateOnline()
	})
}

// SetInitAreaStrategy sets the "init_area_strategy" field.
func (u *AppGoodUpsertOne) SetInitAreaStrategy(v appgood.InitAreaStrategy) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetInitAreaStrategy(v)
	})
}

// UpdateInitAreaStrategy sets the "init_area_strategy" field to the value that was provided on create.
func (u *AppGoodUpsertOne) UpdateInitAreaStrategy() *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateInitAreaStrategy()
	})
}

// SetPrice sets the "price" field.
func (u *AppGoodUpsertOne) SetPrice(v uint64) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetPrice(v)
	})
}

// AddPrice adds v to the "price" field.
func (u *AppGoodUpsertOne) AddPrice(v uint64) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.AddPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *AppGoodUpsertOne) UpdatePrice() *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdatePrice()
	})
}

// SetDisplayIndex sets the "display_index" field.
func (u *AppGoodUpsertOne) SetDisplayIndex(v uint32) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetDisplayIndex(v)
	})
}

// AddDisplayIndex adds v to the "display_index" field.
func (u *AppGoodUpsertOne) AddDisplayIndex(v uint32) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.AddDisplayIndex(v)
	})
}

// UpdateDisplayIndex sets the "display_index" field to the value that was provided on create.
func (u *AppGoodUpsertOne) UpdateDisplayIndex() *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateDisplayIndex()
	})
}

// SetVisible sets the "visible" field.
func (u *AppGoodUpsertOne) SetVisible(v bool) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetVisible(v)
	})
}

// UpdateVisible sets the "visible" field to the value that was provided on create.
func (u *AppGoodUpsertOne) UpdateVisible() *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateVisible()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *AppGoodUpsertOne) SetCreateAt(v uint32) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppGoodUpsertOne) AddCreateAt(v uint32) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppGoodUpsertOne) UpdateCreateAt() *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppGoodUpsertOne) SetUpdateAt(v uint32) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppGoodUpsertOne) AddUpdateAt(v uint32) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppGoodUpsertOne) UpdateUpdateAt() *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppGoodUpsertOne) SetDeleteAt(v uint32) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppGoodUpsertOne) AddDeleteAt(v uint32) *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppGoodUpsertOne) UpdateDeleteAt() *AppGoodUpsertOne {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *AppGoodUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppGoodCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppGoodUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppGoodUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppGoodUpsertOne.ID is not supported by MySQL driver. Use AppGoodUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppGoodUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppGoodCreateBulk is the builder for creating many AppGood entities in bulk.
type AppGoodCreateBulk struct {
	config
	builders []*AppGoodCreate
	conflict []sql.ConflictOption
}

// Save creates the AppGood entities in the database.
func (agcb *AppGoodCreateBulk) Save(ctx context.Context) ([]*AppGood, error) {
	specs := make([]*sqlgraph.CreateSpec, len(agcb.builders))
	nodes := make([]*AppGood, len(agcb.builders))
	mutators := make([]Mutator, len(agcb.builders))
	for i := range agcb.builders {
		func(i int, root context.Context) {
			builder := agcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppGoodMutation)
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
					_, err = mutators[i+1].Mutate(root, agcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = agcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, agcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, agcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (agcb *AppGoodCreateBulk) SaveX(ctx context.Context) []*AppGood {
	v, err := agcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (agcb *AppGoodCreateBulk) Exec(ctx context.Context) error {
	_, err := agcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agcb *AppGoodCreateBulk) ExecX(ctx context.Context) {
	if err := agcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppGood.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppGoodUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (agcb *AppGoodCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppGoodUpsertBulk {
	agcb.conflict = opts
	return &AppGoodUpsertBulk{
		create: agcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppGood.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (agcb *AppGoodCreateBulk) OnConflictColumns(columns ...string) *AppGoodUpsertBulk {
	agcb.conflict = append(agcb.conflict, sql.ConflictColumns(columns...))
	return &AppGoodUpsertBulk{
		create: agcb,
	}
}

// AppGoodUpsertBulk is the builder for "upsert"-ing
// a bulk of AppGood nodes.
type AppGoodUpsertBulk struct {
	create *AppGoodCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppGood.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appgood.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppGoodUpsertBulk) UpdateNewValues() *AppGoodUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appgood.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppGood.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppGoodUpsertBulk) Ignore() *AppGoodUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppGoodUpsertBulk) DoNothing() *AppGoodUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppGoodCreateBulk.OnConflict
// documentation for more info.
func (u *AppGoodUpsertBulk) Update(set func(*AppGoodUpsert)) *AppGoodUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppGoodUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppGoodUpsertBulk) SetAppID(v uuid.UUID) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppGoodUpsertBulk) UpdateAppID() *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateAppID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *AppGoodUpsertBulk) SetGoodID(v uuid.UUID) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *AppGoodUpsertBulk) UpdateGoodID() *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateGoodID()
	})
}

// SetOnline sets the "online" field.
func (u *AppGoodUpsertBulk) SetOnline(v bool) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetOnline(v)
	})
}

// UpdateOnline sets the "online" field to the value that was provided on create.
func (u *AppGoodUpsertBulk) UpdateOnline() *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateOnline()
	})
}

// SetInitAreaStrategy sets the "init_area_strategy" field.
func (u *AppGoodUpsertBulk) SetInitAreaStrategy(v appgood.InitAreaStrategy) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetInitAreaStrategy(v)
	})
}

// UpdateInitAreaStrategy sets the "init_area_strategy" field to the value that was provided on create.
func (u *AppGoodUpsertBulk) UpdateInitAreaStrategy() *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateInitAreaStrategy()
	})
}

// SetPrice sets the "price" field.
func (u *AppGoodUpsertBulk) SetPrice(v uint64) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetPrice(v)
	})
}

// AddPrice adds v to the "price" field.
func (u *AppGoodUpsertBulk) AddPrice(v uint64) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.AddPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *AppGoodUpsertBulk) UpdatePrice() *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdatePrice()
	})
}

// SetDisplayIndex sets the "display_index" field.
func (u *AppGoodUpsertBulk) SetDisplayIndex(v uint32) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetDisplayIndex(v)
	})
}

// AddDisplayIndex adds v to the "display_index" field.
func (u *AppGoodUpsertBulk) AddDisplayIndex(v uint32) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.AddDisplayIndex(v)
	})
}

// UpdateDisplayIndex sets the "display_index" field to the value that was provided on create.
func (u *AppGoodUpsertBulk) UpdateDisplayIndex() *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateDisplayIndex()
	})
}

// SetVisible sets the "visible" field.
func (u *AppGoodUpsertBulk) SetVisible(v bool) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetVisible(v)
	})
}

// UpdateVisible sets the "visible" field to the value that was provided on create.
func (u *AppGoodUpsertBulk) UpdateVisible() *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateVisible()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *AppGoodUpsertBulk) SetCreateAt(v uint32) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppGoodUpsertBulk) AddCreateAt(v uint32) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppGoodUpsertBulk) UpdateCreateAt() *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppGoodUpsertBulk) SetUpdateAt(v uint32) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppGoodUpsertBulk) AddUpdateAt(v uint32) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppGoodUpsertBulk) UpdateUpdateAt() *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppGoodUpsertBulk) SetDeleteAt(v uint32) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppGoodUpsertBulk) AddDeleteAt(v uint32) *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppGoodUpsertBulk) UpdateDeleteAt() *AppGoodUpsertBulk {
	return u.Update(func(s *AppGoodUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *AppGoodUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppGoodCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppGoodCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppGoodUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
