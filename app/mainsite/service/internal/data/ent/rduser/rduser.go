// Code generated by ent, DO NOT EDIT.

package rduser

import (
	"time"
)

const (
	// Label holds the string label denoting the rduser type in the database.
	Label = "rd_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNick holds the string denoting the nick field in the database.
	FieldNick = "nick"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// Table holds the table name of the rduser in the database.
	Table = "rd_user"
)

// Columns holds all SQL columns for rduser fields.
var Columns = []string{
	FieldID,
	FieldNick,
	FieldStatus,
	FieldCreateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NickValidator is a validator for the "nick" field. It is called by the builders before save.
	NickValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int32
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
)
