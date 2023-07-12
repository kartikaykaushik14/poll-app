// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PollsColumns holds the columns for the "polls" table.
	PollsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "question", Type: field.TypeString},
	}
	// PollsTable holds the schema information for the "polls" table.
	PollsTable = &schema.Table{
		Name:       "polls",
		Columns:    PollsColumns,
		PrimaryKey: []*schema.Column{PollsColumns[0]},
	}
	// PollOptionsColumns holds the columns for the "poll_options" table.
	PollOptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "option", Type: field.TypeString},
	}
	// PollOptionsTable holds the schema information for the "poll_options" table.
	PollOptionsTable = &schema.Table{
		Name:       "poll_options",
		Columns:    PollOptionsColumns,
		PrimaryKey: []*schema.Column{PollOptionsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VotesColumns holds the columns for the "votes" table.
	VotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeString},
		{Name: "poll_option_id", Type: field.TypeString},
	}
	// VotesTable holds the schema information for the "votes" table.
	VotesTable = &schema.Table{
		Name:       "votes",
		Columns:    VotesColumns,
		PrimaryKey: []*schema.Column{VotesColumns[0]},
	}
	// PollPollOptionsColumns holds the columns for the "poll_poll_options" table.
	PollPollOptionsColumns = []*schema.Column{
		{Name: "poll_id", Type: field.TypeInt},
		{Name: "poll_option_id", Type: field.TypeInt},
	}
	// PollPollOptionsTable holds the schema information for the "poll_poll_options" table.
	PollPollOptionsTable = &schema.Table{
		Name:       "poll_poll_options",
		Columns:    PollPollOptionsColumns,
		PrimaryKey: []*schema.Column{PollPollOptionsColumns[0], PollPollOptionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "poll_poll_options_poll_id",
				Columns:    []*schema.Column{PollPollOptionsColumns[0]},
				RefColumns: []*schema.Column{PollsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "poll_poll_options_poll_option_id",
				Columns:    []*schema.Column{PollPollOptionsColumns[1]},
				RefColumns: []*schema.Column{PollOptionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PollsTable,
		PollOptionsTable,
		UsersTable,
		VotesTable,
		PollPollOptionsTable,
	}
)

func init() {
	PollPollOptionsTable.ForeignKeys[0].RefTable = PollsTable
	PollPollOptionsTable.ForeignKeys[1].RefTable = PollOptionsTable
}
