// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tmc/moderncrud/ent/predicate"
	"github.com/tmc/moderncrud/ent/widgettype"
)

// WidgetTypeUpdate is the builder for updating WidgetType entities.
type WidgetTypeUpdate struct {
	config
	hooks    []Hook
	mutation *WidgetTypeMutation
}

// Where appends a list predicates to the WidgetTypeUpdate builder.
func (wtu *WidgetTypeUpdate) Where(ps ...predicate.WidgetType) *WidgetTypeUpdate {
	wtu.mutation.Where(ps...)
	return wtu
}

// Mutation returns the WidgetTypeMutation object of the builder.
func (wtu *WidgetTypeUpdate) Mutation() *WidgetTypeMutation {
	return wtu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wtu *WidgetTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wtu.hooks) == 0 {
		affected, err = wtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WidgetTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wtu.mutation = mutation
			affected, err = wtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wtu.hooks) - 1; i >= 0; i-- {
			if wtu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wtu *WidgetTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := wtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wtu *WidgetTypeUpdate) Exec(ctx context.Context) error {
	_, err := wtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wtu *WidgetTypeUpdate) ExecX(ctx context.Context) {
	if err := wtu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wtu *WidgetTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   widgettype.Table,
			Columns: widgettype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: widgettype.FieldID,
			},
		},
	}
	if ps := wtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{widgettype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WidgetTypeUpdateOne is the builder for updating a single WidgetType entity.
type WidgetTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WidgetTypeMutation
}

// Mutation returns the WidgetTypeMutation object of the builder.
func (wtuo *WidgetTypeUpdateOne) Mutation() *WidgetTypeMutation {
	return wtuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wtuo *WidgetTypeUpdateOne) Select(field string, fields ...string) *WidgetTypeUpdateOne {
	wtuo.fields = append([]string{field}, fields...)
	return wtuo
}

// Save executes the query and returns the updated WidgetType entity.
func (wtuo *WidgetTypeUpdateOne) Save(ctx context.Context) (*WidgetType, error) {
	var (
		err  error
		node *WidgetType
	)
	if len(wtuo.hooks) == 0 {
		node, err = wtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WidgetTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wtuo.mutation = mutation
			node, err = wtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wtuo.hooks) - 1; i >= 0; i-- {
			if wtuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wtuo *WidgetTypeUpdateOne) SaveX(ctx context.Context) *WidgetType {
	node, err := wtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wtuo *WidgetTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := wtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wtuo *WidgetTypeUpdateOne) ExecX(ctx context.Context) {
	if err := wtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wtuo *WidgetTypeUpdateOne) sqlSave(ctx context.Context) (_node *WidgetType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   widgettype.Table,
			Columns: widgettype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: widgettype.FieldID,
			},
		},
	}
	id, ok := wtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing WidgetType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := wtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, widgettype.FieldID)
		for _, f := range fields {
			if !widgettype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != widgettype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &WidgetType{config: wtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{widgettype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}