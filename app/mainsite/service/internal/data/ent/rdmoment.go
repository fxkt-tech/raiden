// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/rdmoment"
)

// RdMoment is the model entity for the RdMoment schema.
type RdMoment struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// 动态类型 1图文 2视频 3专栏
	Type int32 `json:"type,omitempty"`
	// 谁发布的动态
	ByUID int32 `json:"by_uid,omitempty"`
	// 文字
	Txt string `json:"txt,omitempty"`
	// 动态的图片数组
	Imgs []string `json:"imgs,omitempty"`
	// 关注的人uid
	Status int32 `json:"status,omitempty"`
	// PublishTime holds the value of the "publish_time" field.
	PublishTime time.Time `json:"publish_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RdMoment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rdmoment.FieldImgs:
			values[i] = new([]byte)
		case rdmoment.FieldID, rdmoment.FieldType, rdmoment.FieldByUID, rdmoment.FieldStatus:
			values[i] = new(sql.NullInt64)
		case rdmoment.FieldTxt:
			values[i] = new(sql.NullString)
		case rdmoment.FieldPublishTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RdMoment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RdMoment fields.
func (rm *RdMoment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rdmoment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rm.ID = int64(value.Int64)
		case rdmoment.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				rm.Type = int32(value.Int64)
			}
		case rdmoment.FieldByUID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field by_uid", values[i])
			} else if value.Valid {
				rm.ByUID = int32(value.Int64)
			}
		case rdmoment.FieldTxt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field txt", values[i])
			} else if value.Valid {
				rm.Txt = value.String
			}
		case rdmoment.FieldImgs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field imgs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rm.Imgs); err != nil {
					return fmt.Errorf("unmarshal field imgs: %w", err)
				}
			}
		case rdmoment.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				rm.Status = int32(value.Int64)
			}
		case rdmoment.FieldPublishTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field publish_time", values[i])
			} else if value.Valid {
				rm.PublishTime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this RdMoment.
// Note that you need to call RdMoment.Unwrap() before calling this method if this RdMoment
// was returned from a transaction, and the transaction was committed or rolled back.
func (rm *RdMoment) Update() *RdMomentUpdateOne {
	return NewRdMomentClient(rm.config).UpdateOne(rm)
}

// Unwrap unwraps the RdMoment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rm *RdMoment) Unwrap() *RdMoment {
	_tx, ok := rm.config.driver.(*txDriver)
	if !ok {
		panic("ent: RdMoment is not a transactional entity")
	}
	rm.config.driver = _tx.drv
	return rm
}

// String implements the fmt.Stringer.
func (rm *RdMoment) String() string {
	var builder strings.Builder
	builder.WriteString("RdMoment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rm.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", rm.Type))
	builder.WriteString(", ")
	builder.WriteString("by_uid=")
	builder.WriteString(fmt.Sprintf("%v", rm.ByUID))
	builder.WriteString(", ")
	builder.WriteString("txt=")
	builder.WriteString(rm.Txt)
	builder.WriteString(", ")
	builder.WriteString("imgs=")
	builder.WriteString(fmt.Sprintf("%v", rm.Imgs))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", rm.Status))
	builder.WriteString(", ")
	builder.WriteString("publish_time=")
	builder.WriteString(rm.PublishTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// RdMoments is a parsable slice of RdMoment.
type RdMoments []*RdMoment

func (rm RdMoments) config(cfg config) {
	for _i := range rm {
		rm[_i].config = cfg
	}
}
