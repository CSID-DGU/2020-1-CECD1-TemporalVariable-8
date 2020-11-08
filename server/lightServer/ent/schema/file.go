package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/google/uuid"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Indexes returns the indexes of the schema.
func (File) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("file_id", "mime").Unique(),
	}
}
// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("file_id", uuid.UUID{}),
		field.String("mime").NotEmpty(),
		field.String("name").NotEmpty(),
		field.Time("create_at").Immutable(),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return nil
}
