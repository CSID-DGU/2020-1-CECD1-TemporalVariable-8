package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Indexes of the schema.
func (Category) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("owner"),
	}
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", Restaurant.Type).Unique().Required(),
		edge.To("menus", Menu.Type),
	}
}
