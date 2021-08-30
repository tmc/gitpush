// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tmc/moderncrud/ent/predicate"
	"github.com/tmc/moderncrud/ent/widget"
	"github.com/tmc/moderncrud/ent/widgettype"
)

// WidgetUpdate is the builder for updating Widget entities.
type WidgetUpdate struct {
	config
	hooks    []Hook
	mutation *WidgetMutation
}

// Where appends a list predicates to the WidgetUpdate builder.
func (wu *WidgetUpdate) Where(ps ...predicate.Widget) *WidgetUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetNote sets the "note" field.
func (wu *WidgetUpdate) SetNote(s string) *WidgetUpdate {
	wu.mutation.SetNote(s)
	return wu
}

// SetStatus sets the "status" field.
func (wu *WidgetUpdate) SetStatus(w widget.Status) *WidgetUpdate {
	wu.mutation.SetStatus(w)
	return wu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wu *WidgetUpdate) SetNillableStatus(w *widget.Status) *WidgetUpdate {
	if w != nil {
		wu.SetStatus(*w)
	}
	return wu
}

// SetPriority sets the "priority" field.
func (wu *WidgetUpdate) SetPriority(i int) *WidgetUpdate {
	wu.mutation.ResetPriority()
	wu.mutation.SetPriority(i)
	return wu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (wu *WidgetUpdate) SetNillablePriority(i *int) *WidgetUpdate {
	if i != nil {
		wu.SetPriority(*i)
	}
	return wu
}

// AddPriority adds i to the "priority" field.
func (wu *WidgetUpdate) AddPriority(i int) *WidgetUpdate {
	wu.mutation.AddPriority(i)
	return wu
}

// SetTestField sets the "test_field" field.
func (wu *WidgetUpdate) SetTestField(s string) *WidgetUpdate {
	wu.mutation.SetTestField(s)
	return wu
}

// SetTypeID sets the "type" edge to the WidgetType entity by ID.
func (wu *WidgetUpdate) SetTypeID(id int) *WidgetUpdate {
	wu.mutation.SetTypeID(id)
	return wu
}

// SetNillableTypeID sets the "type" edge to the WidgetType entity by ID if the given value is not nil.
func (wu *WidgetUpdate) SetNillableTypeID(id *int) *WidgetUpdate {
	if id != nil {
		wu = wu.SetTypeID(*id)
	}
	return wu
}

// SetType sets the "type" edge to the WidgetType entity.
func (wu *WidgetUpdate) SetType(w *WidgetType) *WidgetUpdate {
	return wu.SetTypeID(w.ID)
}

// Mutation returns the WidgetMutation object of the builder.
func (wu *WidgetUpdate) Mutation() *WidgetMutation {
	return wu.mutation
}

// ClearType clears the "type" edge to the WidgetType entity.
func (wu *WidgetUpdate) ClearType() *WidgetUpdate {
	wu.mutation.ClearType()
	return wu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WidgetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WidgetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WidgetUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WidgetUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WidgetUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WidgetUpdate) check() error {
	if v, ok := wu.mutation.Note(); ok {
		if err := widget.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	if v, ok := wu.mutation.Status(); ok {
		if err := widget.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (wu *WidgetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   widget.Table,
			Columns: widget.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: widget.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Note(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: widget.FieldNote,
		})
	}
	if value, ok := wu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: widget.FieldStatus,
		})
	}
	if value, ok := wu.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: widget.FieldPriority,
		})
	}
	if value, ok := wu.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: widget.FieldPriority,
		})
	}
	if value, ok := wu.mutation.TestField(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: widget.FieldTestField,
		})
	}
	if wu.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   widget.TypeTable,
			Columns: []string{widget.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: widgettype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   widget.TypeTable,
			Columns: []string{widget.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: widgettype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{widget.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WidgetUpdateOne is the builder for updating a single Widget entity.
type WidgetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WidgetMutation
}

// SetNote sets the "note" field.
func (wuo *WidgetUpdateOne) SetNote(s string) *WidgetUpdateOne {
	wuo.mutation.SetNote(s)
	return wuo
}

// SetStatus sets the "status" field.
func (wuo *WidgetUpdateOne) SetStatus(w widget.Status) *WidgetUpdateOne {
	wuo.mutation.SetStatus(w)
	return wuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wuo *WidgetUpdateOne) SetNillableStatus(w *widget.Status) *WidgetUpdateOne {
	if w != nil {
		wuo.SetStatus(*w)
	}
	return wuo
}

// SetPriority sets the "priority" field.
func (wuo *WidgetUpdateOne) SetPriority(i int) *WidgetUpdateOne {
	wuo.mutation.ResetPriority()
	wuo.mutation.SetPriority(i)
	return wuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (wuo *WidgetUpdateOne) SetNillablePriority(i *int) *WidgetUpdateOne {
	if i != nil {
		wuo.SetPriority(*i)
	}
	return wuo
}

// AddPriority adds i to the "priority" field.
func (wuo *WidgetUpdateOne) AddPriority(i int) *WidgetUpdateOne {
	wuo.mutation.AddPriority(i)
	return wuo
}

// SetTestField sets the "test_field" field.
func (wuo *WidgetUpdateOne) SetTestField(s string) *WidgetUpdateOne {
	wuo.mutation.SetTestField(s)
	return wuo
}

// SetTypeID sets the "type" edge to the WidgetType entity by ID.
func (wuo *WidgetUpdateOne) SetTypeID(id int) *WidgetUpdateOne {
	wuo.mutation.SetTypeID(id)
	return wuo
}

// SetNillableTypeID sets the "type" edge to the WidgetType entity by ID if the given value is not nil.
func (wuo *WidgetUpdateOne) SetNillableTypeID(id *int) *WidgetUpdateOne {
	if id != nil {
		wuo = wuo.SetTypeID(*id)
	}
	return wuo
}

// SetType sets the "type" edge to the WidgetType entity.
func (wuo *WidgetUpdateOne) SetType(w *WidgetType) *WidgetUpdateOne {
	return wuo.SetTypeID(w.ID)
}

// Mutation returns the WidgetMutation object of the builder.
func (wuo *WidgetUpdateOne) Mutation() *WidgetMutation {
	return wuo.mutation
}

// ClearType clears the "type" edge to the WidgetType entity.
func (wuo *WidgetUpdateOne) ClearType() *WidgetUpdateOne {
	wuo.mutation.ClearType()
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WidgetUpdateOne) Select(field string, fields ...string) *WidgetUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Widget entity.
func (wuo *WidgetUpdateOne) Save(ctx context.Context) (*Widget, error) {
	var (
		err  error
		node *Widget
	)
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WidgetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WidgetUpdateOne) SaveX(ctx context.Context) *Widget {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WidgetUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WidgetUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WidgetUpdateOne) check() error {
	if v, ok := wuo.mutation.Note(); ok {
		if err := widget.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	if v, ok := wuo.mutation.Status(); ok {
		if err := widget.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (wuo *WidgetUpdateOne) sqlSave(ctx context.Context) (_node *Widget, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   widget.Table,
			Columns: widget.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: widget.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Widget.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, widget.FieldID)
		for _, f := range fields {
			if !widget.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != widget.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Note(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: widget.FieldNote,
		})
	}
	if value, ok := wuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: widget.FieldStatus,
		})
	}
	if value, ok := wuo.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: widget.FieldPriority,
		})
	}
	if value, ok := wuo.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: widget.FieldPriority,
		})
	}
	if value, ok := wuo.mutation.TestField(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: widget.FieldTestField,
		})
	}
	if wuo.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   widget.TypeTable,
			Columns: []string{widget.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: widgettype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   widget.TypeTable,
			Columns: []string{widget.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: widgettype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Widget{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{widget.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
