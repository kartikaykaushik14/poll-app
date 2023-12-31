// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"poll-app/ent/vote"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Vote is the model entity for the Vote schema.
type Vote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserId holds the value of the "userId" field.
	UserId string `json:"userId,omitempty"`
	// PollOptionId holds the value of the "pollOptionId" field.
	PollOptionId string `json:"pollOptionId,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Vote) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vote.FieldID:
			values[i] = new(sql.NullInt64)
		case vote.FieldUserId, vote.FieldPollOptionId:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vote fields.
func (v *Vote) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case vote.FieldUserId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				v.UserId = value.String
			}
		case vote.FieldPollOptionId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pollOptionId", values[i])
			} else if value.Valid {
				v.PollOptionId = value.String
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Vote.
// This includes values selected through modifiers, order, etc.
func (v *Vote) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// Update returns a builder for updating this Vote.
// Note that you need to call Vote.Unwrap() before calling this method if this Vote
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vote) Update() *VoteUpdateOne {
	return NewVoteClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Vote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vote) Unwrap() *Vote {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Vote is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vote) String() string {
	var builder strings.Builder
	builder.WriteString("Vote(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("userId=")
	builder.WriteString(v.UserId)
	builder.WriteString(", ")
	builder.WriteString("pollOptionId=")
	builder.WriteString(v.PollOptionId)
	builder.WriteByte(')')
	return builder.String()
}

// Votes is a parsable slice of Vote.
type Votes []*Vote
