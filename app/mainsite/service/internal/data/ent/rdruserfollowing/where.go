// Code generated by ent, DO NOT EDIT.

package rdruserfollowing

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldLTE(FieldID, id))
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldEQ(FieldUID, v))
}

// FollowingUID applies equality check predicate on the "following_uid" field. It's identical to FollowingUIDEQ.
func FollowingUID(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldEQ(FieldFollowingUID, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldEQ(FieldCreateTime, v))
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldEQ(FieldUID, v))
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldNEQ(FieldUID, v))
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldIn(FieldUID, vs...))
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldNotIn(FieldUID, vs...))
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldGT(FieldUID, v))
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldGTE(FieldUID, v))
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldLT(FieldUID, v))
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldLTE(FieldUID, v))
}

// UIDIsNil applies the IsNil predicate on the "uid" field.
func UIDIsNil() predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldIsNull(FieldUID))
}

// UIDNotNil applies the NotNil predicate on the "uid" field.
func UIDNotNil() predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldNotNull(FieldUID))
}

// FollowingUIDEQ applies the EQ predicate on the "following_uid" field.
func FollowingUIDEQ(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldEQ(FieldFollowingUID, v))
}

// FollowingUIDNEQ applies the NEQ predicate on the "following_uid" field.
func FollowingUIDNEQ(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldNEQ(FieldFollowingUID, v))
}

// FollowingUIDIn applies the In predicate on the "following_uid" field.
func FollowingUIDIn(vs ...int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldIn(FieldFollowingUID, vs...))
}

// FollowingUIDNotIn applies the NotIn predicate on the "following_uid" field.
func FollowingUIDNotIn(vs ...int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldNotIn(FieldFollowingUID, vs...))
}

// FollowingUIDGT applies the GT predicate on the "following_uid" field.
func FollowingUIDGT(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldGT(FieldFollowingUID, v))
}

// FollowingUIDGTE applies the GTE predicate on the "following_uid" field.
func FollowingUIDGTE(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldGTE(FieldFollowingUID, v))
}

// FollowingUIDLT applies the LT predicate on the "following_uid" field.
func FollowingUIDLT(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldLT(FieldFollowingUID, v))
}

// FollowingUIDLTE applies the LTE predicate on the "following_uid" field.
func FollowingUIDLTE(v int32) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldLTE(FieldFollowingUID, v))
}

// FollowingUIDIsNil applies the IsNil predicate on the "following_uid" field.
func FollowingUIDIsNil() predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldIsNull(FieldFollowingUID))
}

// FollowingUIDNotNil applies the NotNil predicate on the "following_uid" field.
func FollowingUIDNotNil() predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldNotNull(FieldFollowingUID))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(sql.FieldLTE(FieldCreateTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RdRUserFollowing) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RdRUserFollowing) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RdRUserFollowing) predicate.RdRUserFollowing {
	return predicate.RdRUserFollowing(func(s *sql.Selector) {
		p(s.Not())
	})
}
