package schema

import ("entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field")

// Budget holds the schema definition for the Budget entity.
type Budget struct {
	ent.Schema
}

func (Budget) Fields() []ent.Field {
    return []ent.Field{
        field.Float("amount"),
        field.Time("start_date"),
        field.Time("end_date").Optional(),
    }
}

func (Budget) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("category", Category.Type).Unique(),
        edge.To("user", User.Type).Unique(),
    }
}
