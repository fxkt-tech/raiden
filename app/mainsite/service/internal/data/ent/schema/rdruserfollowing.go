package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RdRUserFollowing holds the schema definition for the RdRUserFollowing entity.
type RdRUserFollowing struct {
	ent.Schema
}

func (RdRUserFollowing) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "rd_r_user_following"},
	}
}

// Fields of the RdRUserFollowing.
func (RdRUserFollowing) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int32("uid").Optional().Comment("uid"),
		field.Int32("following_uid").Optional().Comment("关注的人uid"),
		field.Time("create_time").Default(time.Now).SchemaType(map[string]string{dialect.MySQL: "datetime"}),
	}
}

// Edges of the RdRUserFollowing.
func (RdRUserFollowing) Edges() []ent.Edge {
	return nil
}
