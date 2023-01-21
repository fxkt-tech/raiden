package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/biz"
)

// RdMessage holds the schema definition for the RdMessage entity.
type RdMessage struct {
	ent.Schema
}

func (RdMessage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "rd_message"},
	}
}

// Fields of the RdMessage.
func (RdMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").Comment("消息ID"),
		field.Int32("sender_uid").Optional().Comment("发送者id"),
		field.Int32("recver_uid").Optional().Comment("接收者id"),
		field.JSON("content", &biz.Content{}).Comment("聊天内容"),
		field.Time("create_time").Default(time.Now).SchemaType(map[string]string{dialect.MySQL: "datetime"}),
	}
}

// Edges of the RdMessage.
func (RdMessage) Edges() []ent.Edge {
	return nil
}
