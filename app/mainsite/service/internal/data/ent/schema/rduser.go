package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RdUser holds the schema definition for the RdUser entity.
type RdUser struct {
	ent.Schema
}

func (RdUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "rd_user"},
	}
}

// Fields of the RdUser.
func (RdUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").Comment("用户ID"),
		field.String("nick").NotEmpty().Comment("昵称"),
		field.Int32("status").Default(1).Comment("状态 1可用 2不可用"),
		field.Time("create_time").Default(time.Now).SchemaType(map[string]string{dialect.MySQL: "datetime"}),
	}
}

// Edges of the RdUser.
func (RdUser) Edges() []ent.Edge {
	return nil
}
