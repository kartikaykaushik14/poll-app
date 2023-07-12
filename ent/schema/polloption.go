package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PollOption holds the schema definition for the PollOption entity.
type PollOption struct {
	ent.Schema
}

// Fields of the PollOption.
func (PollOption) Fields() []ent.Field {
	return []ent.Field{
		field.String("option"),
	}
}

// Edges of the PollOption.
func (PollOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("poll", Poll.Type).Ref("poll_options"),
		// edge.From("users", User.Type).Ref("votes"),
		// edge.To("users", User.Type),
	}
}

// // Mixin defines the manual edge methods for the PollOption entity.
// func (PollOption) Mixin() []ent.Mixin {
// 	return []ent.Mixin{
// 		PollOptionMixin{},
// 	}
// }

// // PollOptionMixin defines the manual edge methods for the PollOption entity.
// type PollOptionMixin struct {
// 	ent.Mixin
// }

// func (PollOptionMixin) SetPoll(p PollOption) PollOption {
// 	return p
// }
