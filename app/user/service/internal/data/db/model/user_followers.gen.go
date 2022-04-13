// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserFollower = "user_followers"

// UserFollower mapped from table <user_followers>
type UserFollower struct {
	UID          int32     `gorm:"column:uid;type:int;primaryKey;index:idx_flist,priority:1;default:0" json:"uid"`                                    // uid
	FollowersUID int32     `gorm:"column:followers_uid;type:int;primaryKey;default:0" json:"followers_uid"`                                           // 粉丝uid
	Status       int32     `gorm:"column:status;type:tinyint;not null;index:idx_flist,priority:2;default:1" json:"status"`                            // 状态 1被关注 2被取消关注
	ActionTime   time.Time `gorm:"column:action_time;type:datetime;not null;index:idx_flist,priority:3;default:CURRENT_TIMESTAMP" json:"action_time"` // 执行动作的时间
}

// TableName UserFollower's table name
func (*UserFollower) TableName() string {
	return TableNameUserFollower
}