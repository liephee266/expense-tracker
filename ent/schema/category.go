package schema

import ("entgo.io/ent"
	"entgo.io/ent/schema/field")

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

func (Category) Fields() []ent.Field {
    return []ent.Field{
        field.String("name"),
        field.String("icon").Optional(), // emoji ou nom d’icône
    }
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}
