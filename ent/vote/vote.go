// Code generated by ent, DO NOT EDIT.

package vote

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the vote type in the database.
	Label = "vote"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldPollOptionId holds the string denoting the polloptionid field in the database.
	FieldPollOptionId = "poll_option_id"
	// Table holds the table name of the vote in the database.
	Table = "votes"
)

// Columns holds all SQL columns for vote fields.
var Columns = []string{
	FieldID,
	FieldUserId,
	FieldPollOptionId,
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

// OrderOption defines the ordering options for the Vote queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserId orders the results by the userId field.
func ByUserId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserId, opts...).ToFunc()
}

// ByPollOptionId orders the results by the pollOptionId field.
func ByPollOptionId(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPollOptionId, opts...).ToFunc()
}