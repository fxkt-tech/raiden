// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDynamicHistory = "dynamic_history"

// DynamicHistory mapped from table <dynamic_history>
type DynamicHistory struct {
	DmcID       int64     `gorm:"column:dmc_id;type:bigint;primaryKey;autoIncrement:true" json:"dmc_id"`                                               // dmc_id
	DmcType     int32     `gorm:"column:dmc_type;type:tinyint;not null;default:1" json:"dmc_type"`                                                     // 动态类型 1图文 2视频 3专栏
	ByUID       int32     `gorm:"column:by_uid;type:int;not null;index:idx_order,priority:1;default:0" json:"by_uid"`                                  // 谁发布的动态
	Txt         string    `gorm:"column:txt;type:text" json:"txt"`                                                                                     // 文字
	Imgs        string    `gorm:"column:imgs;type:json" json:"imgs"`                                                                                   // 动态的图片数组
	Status      int32     `gorm:"column:status;type:tinyint;not null;default:1" json:"status"`                                                         // 状态 1可用
	PublishTime time.Time `gorm:"column:publish_time;type:datetime;not null;index:idx_order,priority:2;default:CURRENT_TIMESTAMP" json:"publish_time"` // 发布时间
}

// TableName DynamicHistory's table name
func (*DynamicHistory) TableName() string {
	return TableNameDynamicHistory
}
