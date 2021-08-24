// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// WidgetsColumns holds the columns for the "widgets" table.
	WidgetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// WidgetsTable holds the schema information for the "widgets" table.
	WidgetsTable = &schema.Table{
		Name:       "widgets",
		Columns:    WidgetsColumns,
		PrimaryKey: []*schema.Column{WidgetsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		WidgetsTable,
	}
)

func init() {
}
