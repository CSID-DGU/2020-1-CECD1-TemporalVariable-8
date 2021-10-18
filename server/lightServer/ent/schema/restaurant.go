package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Restaurant holds the schema definition for the Restaurant entity.
type Restaurant struct {
	ent.Schema
}

// Fields of the Restaurant.
func (Restaurant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.String("uri").Optional(),
	}
}

// Edges of the Restaurant.
func (Restaurant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type).Unique(),
		edge.To("avatar", File.Type).Unique(),
		edge.To("root", Restaurant.Type).Unique(),
		edge.To("parent", Restaurant.Type).Unique().From("children"),
		edge.To("histories", History.Type),

		edge.From("categories", Category.Type).Ref("owner"),
		edge.From("orders", Order.Type).Ref("where"),
		edge.From("menus", Menu.Type).Ref("owner"),
	}
}
