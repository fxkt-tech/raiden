// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFollower = "followers"

// Follower mapped from table <followers>
type Follower struct {
	UID          int32     `gorm:"column:uid;primaryKey;default:0" json:"uid"`                               // uid
	FollowersUID int32     `gorm:"column:followers_uid;primaryKey;default:0" json:"followers_uid"`           // 粉丝uid
	Status       int32     `gorm:"column:status;not null;default:1" json:"status"`                           // 状态 1被关注 2被取消关注
	ActionTime   time.Time `gorm:"column:action_time;not null;default:CURRENT_TIMESTAMP" json:"action_time"` // 执行动作的时间
}

// TableName Follower's table name
func (*Follower) TableName() string {
	return TableNameFollower
}
