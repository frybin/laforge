package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Tree holds the schema definition for the Tree entity.
type Tree struct {
	ent.Schema
}

// Fields of the Tree.
func (Tree) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
	}
}

// Edges of the Tree.
func (Tree) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("target", Tree.Type),
		edge.From("source", Tree.Type).Ref("target"),
	}
}
