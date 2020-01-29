package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
		field.String("url_shortened"),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return nil
}
