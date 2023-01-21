package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RdMoment holds the schema definition for the RdMoment entity.
type RdMoment struct {
	ent.Schema
}

func (RdMoment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "rd_moment"},
	}
}

// Fields of the RdMoment.
func (RdMoment) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int32("type").Optional().Default(1).Comment("动态类型 1图文 2视频 3专栏"),
		field.Int32("by_uid").Optional().Comment("谁发布的动态"),
		field.Text("txt").Comment("文字"),
		field.JSON("imgs", []string{}).Comment("动态的图片数组"),
		field.Int32("status").Optional().Comment("关注的人uid"),
		field.Time("publish_time").Default(time.Now).SchemaType(map[string]string{dialect.MySQL: "datetime"}),
	}
}

// Edges of the RdMoment.
func (RdMoment) Edges() []ent.Edge {
	return nil
}
