package schema

import (
	"github.com/Rhymond/go-money"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description"),
		field.Bool("isOption"),
		field.JSON("price", new(money.Money)).Optional(),
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", Restaurant.Type).Required().Unique(),
		edge.From("category", Category.Type).Ref("menus"),
		edge.To("images", File.Type).Unique(),
		edge.To("options", Menu.Type),
	}
}
