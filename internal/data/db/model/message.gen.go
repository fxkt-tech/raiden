// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMessage = "message"

// Message mapped from table <message>
type Message struct {
	MessageID int32     `gorm:"column:message_id;primaryKey;autoIncrement:true" json:"message_id"`
	SenderUID int32     `gorm:"column:sender_uid;not null;default:0" json:"sender_uid"`       // 发送者id
	RecverUID int32     `gorm:"column:recver_uid;not null;default:0" json:"recver_uid"`       // 接收者id
	Content   string    `gorm:"column:content" json:"content"`                                // 聊天内容
	Ctime     time.Time `gorm:"column:ctime;not null;default:CURRENT_TIMESTAMP" json:"ctime"` // 记录时间
}

// TableName Message's table name
func (*Message) TableName() string {
	return TableNameMessage
}