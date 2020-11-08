package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"net/url"
)

// Restaurant holds the schema definition for the Restaurant entity.
type Restaurant struct {
	ent.Schema
}

// Fields of the Restaurant.
func (Restaurant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "sub_name"),
	}
}

// Fields of the Restaurant.
func (Restaurant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("sub_name").Optional(),
		field.JSON("uri", new(url.URL)).Optional(),
	}
}

// Edges of the Restaurant.
func (Restaurant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("avatar", File.Type).Unique(),
		edge.To("root", Restaurant.Type).Unique(),
		edge.To("parent", Restaurant.Type).Unique().From("children"),
		edge.To("histories", History.Type),

		edge.From("menus", Menu.Type).Ref("owner"),
	}
}
