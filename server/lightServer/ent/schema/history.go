package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// History holds the schema definition for the History entity.
type History struct {
	ent.Schema
}

// Fields of the History.
func (History) Fields() []ent.Field {
	return []ent.Field{
		field.Time("open_at").Immutable(),
		field.Time("close_at").Immutable(),
	}
}

// Edges of the History.
func (History) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("of", Restaurant.Type).Ref("histories"),
	}
}
