// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/protobom/storage/internal/backends/ent/identifiersentry"
	"github.com/protobom/storage/internal/backends/ent/node"
)

// IdentifiersEntryCreate is the builder for creating a IdentifiersEntry entity.
type IdentifiersEntryCreate struct {
	config
	mutation *IdentifiersEntryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetSoftwareIdentifierType sets the "software_identifier_type" field.
func (iec *IdentifiersEntryCreate) SetSoftwareIdentifierType(iit identifiersentry.SoftwareIdentifierType) *IdentifiersEntryCreate {
	iec.mutation.SetSoftwareIdentifierType(iit)
	return iec
}

// SetSoftwareIdentifierValue sets the "software_identifier_value" field.
func (iec *IdentifiersEntryCreate) SetSoftwareIdentifierValue(s string) *IdentifiersEntryCreate {
	iec.mutation.SetSoftwareIdentifierValue(s)
	return iec
}

// SetNodesID sets the "nodes" edge to the Node entity by ID.
func (iec *IdentifiersEntryCreate) SetNodesID(id string) *IdentifiersEntryCreate {
	iec.mutation.SetNodesID(id)
	return iec
}

// SetNillableNodesID sets the "nodes" edge to the Node entity by ID if the given value is not nil.
func (iec *IdentifiersEntryCreate) SetNillableNodesID(id *string) *IdentifiersEntryCreate {
	if id != nil {
		iec = iec.SetNodesID(*id)
	}
	return iec
}

// SetNodes sets the "nodes" edge to the Node entity.
func (iec *IdentifiersEntryCreate) SetNodes(n *Node) *IdentifiersEntryCreate {
	return iec.SetNodesID(n.ID)
}

// Mutation returns the IdentifiersEntryMutation object of the builder.
func (iec *IdentifiersEntryCreate) Mutation() *IdentifiersEntryMutation {
	return iec.mutation
}

// Save creates the IdentifiersEntry in the database.
func (iec *IdentifiersEntryCreate) Save(ctx context.Context) (*IdentifiersEntry, error) {
	return withHooks(ctx, iec.sqlSave, iec.mutation, iec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (iec *IdentifiersEntryCreate) SaveX(ctx context.Context) *IdentifiersEntry {
	v, err := iec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iec *IdentifiersEntryCreate) Exec(ctx context.Context) error {
	_, err := iec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iec *IdentifiersEntryCreate) ExecX(ctx context.Context) {
	if err := iec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iec *IdentifiersEntryCreate) check() error {
	if _, ok := iec.mutation.SoftwareIdentifierType(); !ok {
		return &ValidationError{Name: "software_identifier_type", err: errors.New(`ent: missing required field "IdentifiersEntry.software_identifier_type"`)}
	}
	if v, ok := iec.mutation.SoftwareIdentifierType(); ok {
		if err := identifiersentry.SoftwareIdentifierTypeValidator(v); err != nil {
			return &ValidationError{Name: "software_identifier_type", err: fmt.Errorf(`ent: validator failed for field "IdentifiersEntry.software_identifier_type": %w`, err)}
		}
	}
	if _, ok := iec.mutation.SoftwareIdentifierValue(); !ok {
		return &ValidationError{Name: "software_identifier_value", err: errors.New(`ent: missing required field "IdentifiersEntry.software_identifier_value"`)}
	}
	return nil
}

func (iec *IdentifiersEntryCreate) sqlSave(ctx context.Context) (*IdentifiersEntry, error) {
	if err := iec.check(); err != nil {
		return nil, err
	}
	_node, _spec := iec.createSpec()
	if err := sqlgraph.CreateNode(ctx, iec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	iec.mutation.id = &_node.ID
	iec.mutation.done = true
	return _node, nil
}

func (iec *IdentifiersEntryCreate) createSpec() (*IdentifiersEntry, *sqlgraph.CreateSpec) {
	var (
		_node = &IdentifiersEntry{config: iec.config}
		_spec = sqlgraph.NewCreateSpec(identifiersentry.Table, sqlgraph.NewFieldSpec(identifiersentry.FieldID, field.TypeInt))
	)
	_spec.OnConflict = iec.conflict
	if value, ok := iec.mutation.SoftwareIdentifierType(); ok {
		_spec.SetField(identifiersentry.FieldSoftwareIdentifierType, field.TypeEnum, value)
		_node.SoftwareIdentifierType = value
	}
	if value, ok := iec.mutation.SoftwareIdentifierValue(); ok {
		_spec.SetField(identifiersentry.FieldSoftwareIdentifierValue, field.TypeString, value)
		_node.SoftwareIdentifierValue = value
	}
	if nodes := iec.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   identifiersentry.NodesTable,
			Columns: []string{identifiersentry.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.node_identifiers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.IdentifiersEntry.Create().
//		SetSoftwareIdentifierType(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.IdentifiersEntryUpsert) {
//			SetSoftwareIdentifierType(v+v).
//		}).
//		Exec(ctx)
func (iec *IdentifiersEntryCreate) OnConflict(opts ...sql.ConflictOption) *IdentifiersEntryUpsertOne {
	iec.conflict = opts
	return &IdentifiersEntryUpsertOne{
		create: iec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.IdentifiersEntry.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (iec *IdentifiersEntryCreate) OnConflictColumns(columns ...string) *IdentifiersEntryUpsertOne {
	iec.conflict = append(iec.conflict, sql.ConflictColumns(columns...))
	return &IdentifiersEntryUpsertOne{
		create: iec,
	}
}

type (
	// IdentifiersEntryUpsertOne is the builder for "upsert"-ing
	//  one IdentifiersEntry node.
	IdentifiersEntryUpsertOne struct {
		create *IdentifiersEntryCreate
	}

	// IdentifiersEntryUpsert is the "OnConflict" setter.
	IdentifiersEntryUpsert struct {
		*sql.UpdateSet
	}
)

// SetSoftwareIdentifierType sets the "software_identifier_type" field.
func (u *IdentifiersEntryUpsert) SetSoftwareIdentifierType(v identifiersentry.SoftwareIdentifierType) *IdentifiersEntryUpsert {
	u.Set(identifiersentry.FieldSoftwareIdentifierType, v)
	return u
}

// UpdateSoftwareIdentifierType sets the "software_identifier_type" field to the value that was provided on create.
func (u *IdentifiersEntryUpsert) UpdateSoftwareIdentifierType() *IdentifiersEntryUpsert {
	u.SetExcluded(identifiersentry.FieldSoftwareIdentifierType)
	return u
}

// SetSoftwareIdentifierValue sets the "software_identifier_value" field.
func (u *IdentifiersEntryUpsert) SetSoftwareIdentifierValue(v string) *IdentifiersEntryUpsert {
	u.Set(identifiersentry.FieldSoftwareIdentifierValue, v)
	return u
}

// UpdateSoftwareIdentifierValue sets the "software_identifier_value" field to the value that was provided on create.
func (u *IdentifiersEntryUpsert) UpdateSoftwareIdentifierValue() *IdentifiersEntryUpsert {
	u.SetExcluded(identifiersentry.FieldSoftwareIdentifierValue)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.IdentifiersEntry.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *IdentifiersEntryUpsertOne) UpdateNewValues() *IdentifiersEntryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.IdentifiersEntry.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *IdentifiersEntryUpsertOne) Ignore() *IdentifiersEntryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *IdentifiersEntryUpsertOne) DoNothing() *IdentifiersEntryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the IdentifiersEntryCreate.OnConflict
// documentation for more info.
func (u *IdentifiersEntryUpsertOne) Update(set func(*IdentifiersEntryUpsert)) *IdentifiersEntryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&IdentifiersEntryUpsert{UpdateSet: update})
	}))
	return u
}

// SetSoftwareIdentifierType sets the "software_identifier_type" field.
func (u *IdentifiersEntryUpsertOne) SetSoftwareIdentifierType(v identifiersentry.SoftwareIdentifierType) *IdentifiersEntryUpsertOne {
	return u.Update(func(s *IdentifiersEntryUpsert) {
		s.SetSoftwareIdentifierType(v)
	})
}

// UpdateSoftwareIdentifierType sets the "software_identifier_type" field to the value that was provided on create.
func (u *IdentifiersEntryUpsertOne) UpdateSoftwareIdentifierType() *IdentifiersEntryUpsertOne {
	return u.Update(func(s *IdentifiersEntryUpsert) {
		s.UpdateSoftwareIdentifierType()
	})
}

// SetSoftwareIdentifierValue sets the "software_identifier_value" field.
func (u *IdentifiersEntryUpsertOne) SetSoftwareIdentifierValue(v string) *IdentifiersEntryUpsertOne {
	return u.Update(func(s *IdentifiersEntryUpsert) {
		s.SetSoftwareIdentifierValue(v)
	})
}

// UpdateSoftwareIdentifierValue sets the "software_identifier_value" field to the value that was provided on create.
func (u *IdentifiersEntryUpsertOne) UpdateSoftwareIdentifierValue() *IdentifiersEntryUpsertOne {
	return u.Update(func(s *IdentifiersEntryUpsert) {
		s.UpdateSoftwareIdentifierValue()
	})
}

// Exec executes the query.
func (u *IdentifiersEntryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for IdentifiersEntryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *IdentifiersEntryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *IdentifiersEntryUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *IdentifiersEntryUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// IdentifiersEntryCreateBulk is the builder for creating many IdentifiersEntry entities in bulk.
type IdentifiersEntryCreateBulk struct {
	config
	err      error
	builders []*IdentifiersEntryCreate
	conflict []sql.ConflictOption
}

// Save creates the IdentifiersEntry entities in the database.
func (iecb *IdentifiersEntryCreateBulk) Save(ctx context.Context) ([]*IdentifiersEntry, error) {
	if iecb.err != nil {
		return nil, iecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(iecb.builders))
	nodes := make([]*IdentifiersEntry, len(iecb.builders))
	mutators := make([]Mutator, len(iecb.builders))
	for i := range iecb.builders {
		func(i int, root context.Context) {
			builder := iecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IdentifiersEntryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, iecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = iecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, iecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, iecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (iecb *IdentifiersEntryCreateBulk) SaveX(ctx context.Context) []*IdentifiersEntry {
	v, err := iecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iecb *IdentifiersEntryCreateBulk) Exec(ctx context.Context) error {
	_, err := iecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iecb *IdentifiersEntryCreateBulk) ExecX(ctx context.Context) {
	if err := iecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.IdentifiersEntry.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.IdentifiersEntryUpsert) {
//			SetSoftwareIdentifierType(v+v).
//		}).
//		Exec(ctx)
func (iecb *IdentifiersEntryCreateBulk) OnConflict(opts ...sql.ConflictOption) *IdentifiersEntryUpsertBulk {
	iecb.conflict = opts
	return &IdentifiersEntryUpsertBulk{
		create: iecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.IdentifiersEntry.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (iecb *IdentifiersEntryCreateBulk) OnConflictColumns(columns ...string) *IdentifiersEntryUpsertBulk {
	iecb.conflict = append(iecb.conflict, sql.ConflictColumns(columns...))
	return &IdentifiersEntryUpsertBulk{
		create: iecb,
	}
}

// IdentifiersEntryUpsertBulk is the builder for "upsert"-ing
// a bulk of IdentifiersEntry nodes.
type IdentifiersEntryUpsertBulk struct {
	create *IdentifiersEntryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.IdentifiersEntry.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *IdentifiersEntryUpsertBulk) UpdateNewValues() *IdentifiersEntryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.IdentifiersEntry.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *IdentifiersEntryUpsertBulk) Ignore() *IdentifiersEntryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *IdentifiersEntryUpsertBulk) DoNothing() *IdentifiersEntryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the IdentifiersEntryCreateBulk.OnConflict
// documentation for more info.
func (u *IdentifiersEntryUpsertBulk) Update(set func(*IdentifiersEntryUpsert)) *IdentifiersEntryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&IdentifiersEntryUpsert{UpdateSet: update})
	}))
	return u
}

// SetSoftwareIdentifierType sets the "software_identifier_type" field.
func (u *IdentifiersEntryUpsertBulk) SetSoftwareIdentifierType(v identifiersentry.SoftwareIdentifierType) *IdentifiersEntryUpsertBulk {
	return u.Update(func(s *IdentifiersEntryUpsert) {
		s.SetSoftwareIdentifierType(v)
	})
}

// UpdateSoftwareIdentifierType sets the "software_identifier_type" field to the value that was provided on create.
func (u *IdentifiersEntryUpsertBulk) UpdateSoftwareIdentifierType() *IdentifiersEntryUpsertBulk {
	return u.Update(func(s *IdentifiersEntryUpsert) {
		s.UpdateSoftwareIdentifierType()
	})
}

// SetSoftwareIdentifierValue sets the "software_identifier_value" field.
func (u *IdentifiersEntryUpsertBulk) SetSoftwareIdentifierValue(v string) *IdentifiersEntryUpsertBulk {
	return u.Update(func(s *IdentifiersEntryUpsert) {
		s.SetSoftwareIdentifierValue(v)
	})
}

// UpdateSoftwareIdentifierValue sets the "software_identifier_value" field to the value that was provided on create.
func (u *IdentifiersEntryUpsertBulk) UpdateSoftwareIdentifierValue() *IdentifiersEntryUpsertBulk {
	return u.Update(func(s *IdentifiersEntryUpsert) {
		s.UpdateSoftwareIdentifierValue()
	})
}

// Exec executes the query.
func (u *IdentifiersEntryUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the IdentifiersEntryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for IdentifiersEntryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *IdentifiersEntryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
