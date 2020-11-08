package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// OrderField holds the schema definition for the OrderField entity.
type OrderField struct {
	ent.Schema
}

// Fields of the OrderField.
func (OrderField) Fields() []ent.Field {
	return []ent.Field{
		field.Uint16("count").Positive(),
	}
}

// Edges of the OrderField.
func (OrderField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("menu", Menu.Type).Required(),
	}
}
