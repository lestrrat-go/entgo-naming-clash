// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserFoosColumns holds the columns for the "user_foos" table.
	UserFoosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "dummy", Type: field.TypeString},
	}
	// UserFoosTable holds the schema information for the "user_foos" table.
	UserFoosTable = &schema.Table{
		Name:       "user_foos",
		Columns:    UserFoosColumns,
		PrimaryKey: []*schema.Column{UserFoosColumns[0]},
	}
	// UserFoosColumns holds the columns for the "user_foos" table.
	UserFoosColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "user_foo_id", Type: field.TypeInt},
	}
	// UserFoosTable holds the schema information for the "user_foos" table.
	UserFoosTable = &schema.Table{
		Name:       "user_foos",
		Columns:    UserFoosColumns,
		PrimaryKey: []*schema.Column{UserFoosColumns[0], UserFoosColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_foos_user_id",
				Columns:    []*schema.Column{UserFoosColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_foos_user_foo_id",
				Columns:    []*schema.Column{UserFoosColumns[1]},
				RefColumns: []*schema.Column{UserFoosColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
		UserFoosTable,
		UserFoosTable,
	}
)

func init() {
	UserFoosTable.ForeignKeys[0].RefTable = UsersTable
	UserFoosTable.ForeignKeys[1].RefTable = UserFoosTable
}
