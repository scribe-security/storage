// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package edgetype

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldLTE(FieldID, id))
}

// NodeID applies equality check predicate on the "node_id" field. It's identical to NodeIDEQ.
func NodeID(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldEQ(FieldNodeID, v))
}

// ToID applies equality check predicate on the "to_id" field. It's identical to ToIDEQ.
func ToID(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldEQ(FieldToID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldNotIn(FieldType, vs...))
}

// NodeIDEQ applies the EQ predicate on the "node_id" field.
func NodeIDEQ(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldEQ(FieldNodeID, v))
}

// NodeIDNEQ applies the NEQ predicate on the "node_id" field.
func NodeIDNEQ(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldNEQ(FieldNodeID, v))
}

// NodeIDIn applies the In predicate on the "node_id" field.
func NodeIDIn(vs ...string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldIn(FieldNodeID, vs...))
}

// NodeIDNotIn applies the NotIn predicate on the "node_id" field.
func NodeIDNotIn(vs ...string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldNotIn(FieldNodeID, vs...))
}

// NodeIDGT applies the GT predicate on the "node_id" field.
func NodeIDGT(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldGT(FieldNodeID, v))
}

// NodeIDGTE applies the GTE predicate on the "node_id" field.
func NodeIDGTE(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldGTE(FieldNodeID, v))
}

// NodeIDLT applies the LT predicate on the "node_id" field.
func NodeIDLT(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldLT(FieldNodeID, v))
}

// NodeIDLTE applies the LTE predicate on the "node_id" field.
func NodeIDLTE(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldLTE(FieldNodeID, v))
}

// NodeIDContains applies the Contains predicate on the "node_id" field.
func NodeIDContains(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldContains(FieldNodeID, v))
}

// NodeIDHasPrefix applies the HasPrefix predicate on the "node_id" field.
func NodeIDHasPrefix(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldHasPrefix(FieldNodeID, v))
}

// NodeIDHasSuffix applies the HasSuffix predicate on the "node_id" field.
func NodeIDHasSuffix(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldHasSuffix(FieldNodeID, v))
}

// NodeIDEqualFold applies the EqualFold predicate on the "node_id" field.
func NodeIDEqualFold(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldEqualFold(FieldNodeID, v))
}

// NodeIDContainsFold applies the ContainsFold predicate on the "node_id" field.
func NodeIDContainsFold(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldContainsFold(FieldNodeID, v))
}

// ToIDEQ applies the EQ predicate on the "to_id" field.
func ToIDEQ(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldEQ(FieldToID, v))
}

// ToIDNEQ applies the NEQ predicate on the "to_id" field.
func ToIDNEQ(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldNEQ(FieldToID, v))
}

// ToIDIn applies the In predicate on the "to_id" field.
func ToIDIn(vs ...string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldIn(FieldToID, vs...))
}

// ToIDNotIn applies the NotIn predicate on the "to_id" field.
func ToIDNotIn(vs ...string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldNotIn(FieldToID, vs...))
}

// ToIDGT applies the GT predicate on the "to_id" field.
func ToIDGT(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldGT(FieldToID, v))
}

// ToIDGTE applies the GTE predicate on the "to_id" field.
func ToIDGTE(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldGTE(FieldToID, v))
}

// ToIDLT applies the LT predicate on the "to_id" field.
func ToIDLT(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldLT(FieldToID, v))
}

// ToIDLTE applies the LTE predicate on the "to_id" field.
func ToIDLTE(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldLTE(FieldToID, v))
}

// ToIDContains applies the Contains predicate on the "to_id" field.
func ToIDContains(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldContains(FieldToID, v))
}

// ToIDHasPrefix applies the HasPrefix predicate on the "to_id" field.
func ToIDHasPrefix(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldHasPrefix(FieldToID, v))
}

// ToIDHasSuffix applies the HasSuffix predicate on the "to_id" field.
func ToIDHasSuffix(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldHasSuffix(FieldToID, v))
}

// ToIDEqualFold applies the EqualFold predicate on the "to_id" field.
func ToIDEqualFold(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldEqualFold(FieldToID, v))
}

// ToIDContainsFold applies the ContainsFold predicate on the "to_id" field.
func ToIDContainsFold(v string) predicate.EdgeType {
	return predicate.EdgeType(sql.FieldContainsFold(FieldToID, v))
}

// HasFrom applies the HasEdge predicate on the "from" edge.
func HasFrom() predicate.EdgeType {
	return predicate.EdgeType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FromTable, FromColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFromWith applies the HasEdge predicate on the "from" edge with a given conditions (other predicates).
func HasFromWith(preds ...predicate.Node) predicate.EdgeType {
	return predicate.EdgeType(func(s *sql.Selector) {
		step := newFromStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTo applies the HasEdge predicate on the "to" edge.
func HasTo() predicate.EdgeType {
	return predicate.EdgeType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ToTable, ToColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasToWith applies the HasEdge predicate on the "to" edge with a given conditions (other predicates).
func HasToWith(preds ...predicate.Node) predicate.EdgeType {
	return predicate.EdgeType(func(s *sql.Selector) {
		step := newToStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EdgeType) predicate.EdgeType {
	return predicate.EdgeType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EdgeType) predicate.EdgeType {
	return predicate.EdgeType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EdgeType) predicate.EdgeType {
	return predicate.EdgeType(sql.NotPredicates(p))
}
