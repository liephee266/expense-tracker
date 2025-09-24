// ent/schema/expense.go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "time"
    "entgo.io/ent/schema/edge"
)

// Expense définit un schéma de dépense
type Expense struct {
    ent.Schema
}

func (Expense) Fields() []ent.Field {
    return []ent.Field{
        field.String("title").NotEmpty(),
        field.Float("amount").Positive(),
        field.Time("date").Default(time.Now),
    }
}

func (Expense) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("category", Category.Type).Unique(),
        edge.To("user", User.Type).Unique(),
    }
}
