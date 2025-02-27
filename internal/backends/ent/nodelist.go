// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/nodelist"
)

// NodeList is the model entity for the NodeList schema.
type NodeList struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RootElements holds the value of the "root_elements" field.
	RootElements []string `json:"root_elements,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NodeListQuery when eager-loading is set.
	Edges        NodeListEdges `json:"edges"`
	selectValues sql.SelectValues
}

// NodeListEdges holds the relations/edges for other nodes in the graph.
type NodeListEdges struct {
	// Nodes holds the value of the nodes edge.
	Nodes []*Node `json:"nodes,omitempty"`
	// Document holds the value of the document edge.
	Document *Document `json:"document,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// NodesOrErr returns the Nodes value or an error if the edge
// was not loaded in eager-loading.
func (e NodeListEdges) NodesOrErr() ([]*Node, error) {
	if e.loadedTypes[0] {
		return e.Nodes, nil
	}
	return nil, &NotLoadedError{edge: "nodes"}
}

// DocumentOrErr returns the Document value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NodeListEdges) DocumentOrErr() (*Document, error) {
	if e.Document != nil {
		return e.Document, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: document.Label}
	}
	return nil, &NotLoadedError{edge: "document"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NodeList) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case nodelist.FieldRootElements:
			values[i] = new([]byte)
		case nodelist.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NodeList fields.
func (nl *NodeList) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case nodelist.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			nl.ID = int(value.Int64)
		case nodelist.FieldRootElements:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field root_elements", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &nl.RootElements); err != nil {
					return fmt.Errorf("unmarshal field root_elements: %w", err)
				}
			}
		default:
			nl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the NodeList.
// This includes values selected through modifiers, order, etc.
func (nl *NodeList) Value(name string) (ent.Value, error) {
	return nl.selectValues.Get(name)
}

// QueryNodes queries the "nodes" edge of the NodeList entity.
func (nl *NodeList) QueryNodes() *NodeQuery {
	return NewNodeListClient(nl.config).QueryNodes(nl)
}

// QueryDocument queries the "document" edge of the NodeList entity.
func (nl *NodeList) QueryDocument() *DocumentQuery {
	return NewNodeListClient(nl.config).QueryDocument(nl)
}

// Update returns a builder for updating this NodeList.
// Note that you need to call NodeList.Unwrap() before calling this method if this NodeList
// was returned from a transaction, and the transaction was committed or rolled back.
func (nl *NodeList) Update() *NodeListUpdateOne {
	return NewNodeListClient(nl.config).UpdateOne(nl)
}

// Unwrap unwraps the NodeList entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nl *NodeList) Unwrap() *NodeList {
	_tx, ok := nl.config.driver.(*txDriver)
	if !ok {
		panic("ent: NodeList is not a transactional entity")
	}
	nl.config.driver = _tx.drv
	return nl
}

// String implements the fmt.Stringer.
func (nl *NodeList) String() string {
	var builder strings.Builder
	builder.WriteString("NodeList(")
	builder.WriteString(fmt.Sprintf("id=%v, ", nl.ID))
	builder.WriteString("root_elements=")
	builder.WriteString(fmt.Sprintf("%v", nl.RootElements))
	builder.WriteByte(')')
	return builder.String()
}

// NodeLists is a parsable slice of NodeList.
type NodeLists []*NodeList
