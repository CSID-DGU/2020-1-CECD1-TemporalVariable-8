package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Time("order_at").Default(time.Now),
		field.Time("cooking_at").Optional().Nillable(),
		field.Time("deliver_at").Optional().Nillable(),
		field.Time("complete_at").Optional().Nillable(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("who", User.Type).Unique().Required(),
		edge.To("where", Restaurant.Type).Unique().Required(),
		edge.To("items", OrderField.Type),
	}
}
