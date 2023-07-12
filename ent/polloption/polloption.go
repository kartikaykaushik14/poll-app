// Code generated by ent, DO NOT EDIT.

package polloption

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the polloption type in the database.
	Label = "poll_option"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOption holds the string denoting the option field in the database.
	FieldOption = "option"
	// EdgePoll holds the string denoting the poll edge name in mutations.
	EdgePoll = "poll"
	// Table holds the table name of the polloption in the database.
	Table = "poll_options"
	// PollTable is the table that holds the poll relation/edge. The primary key declared below.
	PollTable = "poll_poll_options"
	// PollInverseTable is the table name for the Poll entity.
	// It exists in this package in order to avoid circular dependency with the "poll" package.
	PollInverseTable = "polls"
)

// Columns holds all SQL columns for polloption fields.
var Columns = []string{
	FieldID,
	FieldOption,
}

var (
	// PollPrimaryKey and PollColumn2 are the table columns denoting the
	// primary key for the poll relation (M2M).
	PollPrimaryKey = []string{"poll_id", "poll_option_id"}
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

// OrderOption defines the ordering options for the PollOption queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOption orders the results by the option field.
func ByOption(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOption, opts...).ToFunc()
}

// ByPollCount orders the results by poll count.
func ByPollCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPollStep(), opts...)
	}
}

// ByPoll orders the results by poll terms.
func ByPoll(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPollStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPollStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PollInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PollTable, PollPrimaryKey...),
	)
}
