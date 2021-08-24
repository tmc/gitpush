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
		{Name: "note", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"draft", "completed"}, Default: "draft"},
		{Name: "priority", Type: field.TypeInt, Default: 0},
		{Name: "widget_type", Type: field.TypeInt, Nullable: true},
	}
	// WidgetsTable holds the schema information for the "widgets" table.
	WidgetsTable = &schema.Table{
		Name:       "widgets",
		Columns:    WidgetsColumns,
		PrimaryKey: []*schema.Column{WidgetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "widgets_widget_types_type",
				Columns:    []*schema.Column{WidgetsColumns[5]},
				RefColumns: []*schema.Column{WidgetTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WidgetTypesColumns holds the columns for the "widget_types" table.
	WidgetTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// WidgetTypesTable holds the schema information for the "widget_types" table.
	WidgetTypesTable = &schema.Table{
		Name:       "widget_types",
		Columns:    WidgetTypesColumns,
		PrimaryKey: []*schema.Column{WidgetTypesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		WidgetsTable,
		WidgetTypesTable,
	}
)

func init() {
	WidgetsTable.ForeignKeys[0].RefTable = WidgetTypesTable
}
