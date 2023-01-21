package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RdRUserFollowers holds the schema definition for the RdRUserFollowers entity.
type RdRUserFollowers struct {
	ent.Schema
}

func (RdRUserFollowers) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "rd_r_user_followers"},
	}
}

// Fields of the RdRUserFollowers.
func (RdRUserFollowers) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int32("uid").Optional().Comment("uid"),
		field.Int32("followers_uid").Optional().Comment("粉丝uid"),
		field.Time("create_time").Default(time.Now).SchemaType(map[string]string{dialect.MySQL: "datetime"}),
	}
}

// Edges of the RdRUserFollowers.
func (RdRUserFollowers) Edges() []ent.Edge {
	return nil
}
