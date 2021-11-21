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
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/currency"
	"github.com/google/uuid"
)

// CurrencyCreate is the builder for creating a Currency entity.
type CurrencyCreate struct {
	config
	mutation *CurrencyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (cc *CurrencyCreate) SetName(s string) *CurrencyCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetUnit sets the "unit" field.
func (cc *CurrencyCreate) SetUnit(s string) *CurrencyCreate {
	cc.mutation.SetUnit(s)
	return cc
}

// SetSymbol sets the "symbol" field.
func (cc *CurrencyCreate) SetSymbol(s string) *CurrencyCreate {
	cc.mutation.SetSymbol(s)
	return cc
}

// SetCreateAt sets the "create_at" field.
func (cc *CurrencyCreate) SetCreateAt(u uint32) *CurrencyCreate {
	cc.mutation.SetCreateAt(u)
	return cc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cc *CurrencyCreate) SetNillableCreateAt(u *uint32) *CurrencyCreate {
	if u != nil {
		cc.SetCreateAt(*u)
	}
	return cc
}

// SetUpdateAt sets the "update_at" field.
func (cc *CurrencyCreate) SetUpdateAt(u uint32) *CurrencyCreate {
	cc.mutation.SetUpdateAt(u)
	return cc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (cc *CurrencyCreate) SetNillableUpdateAt(u *uint32) *CurrencyCreate {
	if u != nil {
		cc.SetUpdateAt(*u)
	}
	return cc
}

// SetDeleteAt sets the "delete_at" field.
func (cc *CurrencyCreate) SetDeleteAt(u uint32) *CurrencyCreate {
	cc.mutation.SetDeleteAt(u)
	return cc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cc *CurrencyCreate) SetNillableDeleteAt(u *uint32) *CurrencyCreate {
	if u != nil {
		cc.SetDeleteAt(*u)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CurrencyCreate) SetID(u uuid.UUID) *CurrencyCreate {
	cc.mutation.SetID(u)
	return cc
}

// Mutation returns the CurrencyMutation object of the builder.
func (cc *CurrencyCreate) Mutation() *CurrencyMutation {
	return cc.mutation
}

// Save creates the Currency in the database.
func (cc *CurrencyCreate) Save(ctx context.Context) (*Currency, error) {
	var (
		err  error
		node *Currency
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CurrencyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CurrencyCreate) SaveX(ctx context.Context) *Currency {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CurrencyCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CurrencyCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CurrencyCreate) defaults() {
	if _, ok := cc.mutation.CreateAt(); !ok {
		v := currency.DefaultCreateAt()
		cc.mutation.SetCreateAt(v)
	}
	if _, ok := cc.mutation.UpdateAt(); !ok {
		v := currency.DefaultUpdateAt()
		cc.mutation.SetUpdateAt(v)
	}
	if _, ok := cc.mutation.DeleteAt(); !ok {
		v := currency.DefaultDeleteAt()
		cc.mutation.SetDeleteAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := currency.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CurrencyCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := cc.mutation.Unit(); !ok {
		return &ValidationError{Name: "unit", err: errors.New(`ent: missing required field "unit"`)}
	}
	if _, ok := cc.mutation.Symbol(); !ok {
		return &ValidationError{Name: "symbol", err: errors.New(`ent: missing required field "symbol"`)}
	}
	if _, ok := cc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := cc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := cc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (cc *CurrencyCreate) sqlSave(ctx context.Context) (*Currency, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *CurrencyCreate) createSpec() (*Currency, *sqlgraph.CreateSpec) {
	var (
		_node = &Currency{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: currency.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: currency.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: currency.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Unit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: currency.FieldUnit,
		})
		_node.Unit = value
	}
	if value, ok := cc.mutation.Symbol(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: currency.FieldSymbol,
		})
		_node.Symbol = value
	}
	if value, ok := cc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currency.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := cc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currency.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := cc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currency.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Currency.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CurrencyUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (cc *CurrencyCreate) OnConflict(opts ...sql.ConflictOption) *CurrencyUpsertOne {
	cc.conflict = opts
	return &CurrencyUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Currency.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *CurrencyCreate) OnConflictColumns(columns ...string) *CurrencyUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CurrencyUpsertOne{
		create: cc,
	}
}

type (
	// CurrencyUpsertOne is the builder for "upsert"-ing
	//  one Currency node.
	CurrencyUpsertOne struct {
		create *CurrencyCreate
	}

	// CurrencyUpsert is the "OnConflict" setter.
	CurrencyUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *CurrencyUpsert) SetName(v string) *CurrencyUpsert {
	u.Set(currency.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CurrencyUpsert) UpdateName() *CurrencyUpsert {
	u.SetExcluded(currency.FieldName)
	return u
}

// SetUnit sets the "unit" field.
func (u *CurrencyUpsert) SetUnit(v string) *CurrencyUpsert {
	u.Set(currency.FieldUnit, v)
	return u
}

// UpdateUnit sets the "unit" field to the value that was provided on create.
func (u *CurrencyUpsert) UpdateUnit() *CurrencyUpsert {
	u.SetExcluded(currency.FieldUnit)
	return u
}

// SetSymbol sets the "symbol" field.
func (u *CurrencyUpsert) SetSymbol(v string) *CurrencyUpsert {
	u.Set(currency.FieldSymbol, v)
	return u
}

// UpdateSymbol sets the "symbol" field to the value that was provided on create.
func (u *CurrencyUpsert) UpdateSymbol() *CurrencyUpsert {
	u.SetExcluded(currency.FieldSymbol)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *CurrencyUpsert) SetCreateAt(v uint32) *CurrencyUpsert {
	u.Set(currency.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CurrencyUpsert) UpdateCreateAt() *CurrencyUpsert {
	u.SetExcluded(currency.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *CurrencyUpsert) SetUpdateAt(v uint32) *CurrencyUpsert {
	u.Set(currency.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CurrencyUpsert) UpdateUpdateAt() *CurrencyUpsert {
	u.SetExcluded(currency.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *CurrencyUpsert) SetDeleteAt(v uint32) *CurrencyUpsert {
	u.Set(currency.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CurrencyUpsert) UpdateDeleteAt() *CurrencyUpsert {
	u.SetExcluded(currency.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Currency.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(currency.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CurrencyUpsertOne) UpdateNewValues() *CurrencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(currency.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Currency.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CurrencyUpsertOne) Ignore() *CurrencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CurrencyUpsertOne) DoNothing() *CurrencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CurrencyCreate.OnConflict
// documentation for more info.
func (u *CurrencyUpsertOne) Update(set func(*CurrencyUpsert)) *CurrencyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CurrencyUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *CurrencyUpsertOne) SetName(v string) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CurrencyUpsertOne) UpdateName() *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateName()
	})
}

// SetUnit sets the "unit" field.
func (u *CurrencyUpsertOne) SetUnit(v string) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetUnit(v)
	})
}

// UpdateUnit sets the "unit" field to the value that was provided on create.
func (u *CurrencyUpsertOne) UpdateUnit() *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateUnit()
	})
}

// SetSymbol sets the "symbol" field.
func (u *CurrencyUpsertOne) SetSymbol(v string) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetSymbol(v)
	})
}

// UpdateSymbol sets the "symbol" field to the value that was provided on create.
func (u *CurrencyUpsertOne) UpdateSymbol() *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateSymbol()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *CurrencyUpsertOne) SetCreateAt(v uint32) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CurrencyUpsertOne) UpdateCreateAt() *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *CurrencyUpsertOne) SetUpdateAt(v uint32) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CurrencyUpsertOne) UpdateUpdateAt() *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *CurrencyUpsertOne) SetDeleteAt(v uint32) *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CurrencyUpsertOne) UpdateDeleteAt() *CurrencyUpsertOne {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *CurrencyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CurrencyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CurrencyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CurrencyUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CurrencyUpsertOne.ID is not supported by MySQL driver. Use CurrencyUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CurrencyUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CurrencyCreateBulk is the builder for creating many Currency entities in bulk.
type CurrencyCreateBulk struct {
	config
	builders []*CurrencyCreate
	conflict []sql.ConflictOption
}

// Save creates the Currency entities in the database.
func (ccb *CurrencyCreateBulk) Save(ctx context.Context) ([]*Currency, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Currency, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CurrencyMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CurrencyCreateBulk) SaveX(ctx context.Context) []*Currency {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CurrencyCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CurrencyCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Currency.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CurrencyUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *CurrencyCreateBulk) OnConflict(opts ...sql.ConflictOption) *CurrencyUpsertBulk {
	ccb.conflict = opts
	return &CurrencyUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Currency.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *CurrencyCreateBulk) OnConflictColumns(columns ...string) *CurrencyUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CurrencyUpsertBulk{
		create: ccb,
	}
}

// CurrencyUpsertBulk is the builder for "upsert"-ing
// a bulk of Currency nodes.
type CurrencyUpsertBulk struct {
	create *CurrencyCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Currency.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(currency.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CurrencyUpsertBulk) UpdateNewValues() *CurrencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(currency.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Currency.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CurrencyUpsertBulk) Ignore() *CurrencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CurrencyUpsertBulk) DoNothing() *CurrencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CurrencyCreateBulk.OnConflict
// documentation for more info.
func (u *CurrencyUpsertBulk) Update(set func(*CurrencyUpsert)) *CurrencyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CurrencyUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *CurrencyUpsertBulk) SetName(v string) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CurrencyUpsertBulk) UpdateName() *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateName()
	})
}

// SetUnit sets the "unit" field.
func (u *CurrencyUpsertBulk) SetUnit(v string) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetUnit(v)
	})
}

// UpdateUnit sets the "unit" field to the value that was provided on create.
func (u *CurrencyUpsertBulk) UpdateUnit() *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateUnit()
	})
}

// SetSymbol sets the "symbol" field.
func (u *CurrencyUpsertBulk) SetSymbol(v string) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetSymbol(v)
	})
}

// UpdateSymbol sets the "symbol" field to the value that was provided on create.
func (u *CurrencyUpsertBulk) UpdateSymbol() *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateSymbol()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *CurrencyUpsertBulk) SetCreateAt(v uint32) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CurrencyUpsertBulk) UpdateCreateAt() *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *CurrencyUpsertBulk) SetUpdateAt(v uint32) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CurrencyUpsertBulk) UpdateUpdateAt() *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *CurrencyUpsertBulk) SetDeleteAt(v uint32) *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CurrencyUpsertBulk) UpdateDeleteAt() *CurrencyUpsertBulk {
	return u.Update(func(s *CurrencyUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *CurrencyUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CurrencyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CurrencyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CurrencyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
