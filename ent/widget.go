// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/tmc/moderncrud/ent/widget"
)

// Widget is the model entity for the Widget schema.
type Widget struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Widget) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case widget.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Widget", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Widget fields.
func (w *Widget) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case widget.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Widget.
// Note that you need to call Widget.Unwrap() before calling this method if this Widget
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Widget) Update() *WidgetUpdateOne {
	return (&WidgetClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Widget entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Widget) Unwrap() *Widget {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Widget is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Widget) String() string {
	var builder strings.Builder
	builder.WriteString("Widget(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Widgets is a parsable slice of Widget.
type Widgets []*Widget

func (w Widgets) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
