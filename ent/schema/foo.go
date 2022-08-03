package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserFoo struct {
	ent.Schema
}

func (UserFoo) Fields() []ent.Field {
	return []ent.Field{
		field.String(`dummy`),
	}
}

func (UserFoo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(`parent`, User.Type).Ref(`foos`),
	}
}
