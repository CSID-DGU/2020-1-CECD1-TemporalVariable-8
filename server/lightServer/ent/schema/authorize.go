package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Authorize holds the schema definition for the Authorize entity.
type Authorize struct {
	ent.Schema
}

// Indexes of the schema.
func (Authorize) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("provider", "service_id").Unique(),
	}
}

// Fields of the Authorize.
func (Authorize) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("provider").Values("google", "facebook", "tweeter", "naver", "daum"),
		field.String("service_id").NotEmpty(),
	}
}

// Edges of the Authorize.
func (Authorize) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("auths").Unique().Required(),
	}
}
