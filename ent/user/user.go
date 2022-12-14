// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeFoos holds the string denoting the foos edge name in mutations.
	EdgeFoos = "foos"
	// Table holds the table name of the user in the database.
	Table = "users"
	// FoosTable is the table that holds the foos relation/edge. The primary key declared below.
	FoosTable = "user_foos"
	// FoosInverseTable is the table name for the UserFoo entity.
	// It exists in this package in order to avoid circular dependency with the "userfoo" package.
	FoosInverseTable = "user_foos"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
}

var (
	// FoosPrimaryKey and FoosColumn2 are the table columns denoting the
	// primary key for the foos relation (M2M).
	FoosPrimaryKey = []string{"user_id", "user_foo_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
