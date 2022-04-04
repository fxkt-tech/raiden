package data

import (
	"context"

	v1 "fxkt.tech/raiden/api/user/service/v1"
	"fxkt.tech/raiden/app/user/service/internal/biz"
	"fxkt.tech/raiden/app/user/service/internal/data/db/model"
	"fxkt.tech/raiden/app/user/service/internal/data/db/query"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type userSystemRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserSystemRepo(data *Data, logger log.Logger) biz.UserSystemRepo {
	return &userSystemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userSystemRepo) Register(ctx context.Context, u *biz.User) error {
	dbu := r.data.db.User
	poUser := &model.User{
		Nick: u.Nick,
	}
	err := dbu.WithContext(ctx).Create(poUser)
	if err != nil {
		return v1.ErrorDatabase(err.Error())
	}
	u.Uid = poUser.UID
	return nil
}
func (r *userSystemRepo) Followers(ctx context.Context, us *biz.UserSearch) ([]*biz.User, int32, error) {
	var (
		dbf     = r.data.db.Follower
		dbu     = r.data.db.User
		poUsers []*model.User
	)
	total, err := dbf.WithContext(ctx).
		Order(dbf.ActionTime.Desc()).
		Where(dbf.UID.Eq(us.Uid), dbf.Status.Eq(1)).
		Join(dbu, dbf.FollowersUID.EqCol(dbu.UID)).
		Select(dbu.UID, dbu.Nick).
		ScanByPage(&poUsers, int((us.Page-1)*us.Count), int(us.Count))
	if err != nil {
		return nil, 0, v1.ErrorDatabase(err.Error())
	}

	users := make([]*biz.User, len(poUsers))
	for i, poUser := range poUsers {
		users[i] = &biz.User{
			Uid:  poUser.UID,
			Nick: poUser.Nick,
		}
	}

	return users, int32(total), nil
}
func (r *userSystemRepo) Following(ctx context.Context, us *biz.UserSearch) ([]*biz.User, int32, error) {
	var (
		dbf     = r.data.db.Following
		dbu     = r.data.db.User
		poUsers []*model.User
	)
	total, err := dbf.WithContext(ctx).
		Order(dbf.ActionTime.Desc()).
		Where(dbf.UID.Eq(us.Uid), dbf.Status.Eq(1)).
		Join(dbu, dbf.FollowingUID.EqCol(dbu.UID)).
		Select(dbu.UID, dbu.Nick).
		ScanByPage(&poUsers, int((us.Page-1)*us.Count), int(us.Count))
	if err != nil {
		return nil, 0, v1.ErrorDatabase(err.Error())
	}

	users := make([]*biz.User, len(poUsers))
	for i, poUser := range poUsers {
		users[i] = &biz.User{
			Uid:  poUser.UID,
			Nick: poUser.Nick,
		}
	}

	return users, int32(total), nil
}
func (r *userSystemRepo) Relation(ctx context.Context, ur *biz.UserRelation) error {
	var (
		q      = r.data.db
		dbfg   = r.data.db.Following
		dbfr   = r.data.db.Follower
		status int32
	)

	switch ur.Action {
	case 1: // 关注
		status = 1
	case 2: // 取消关注
		status = 2
	default:
		return v1.ErrorValidator("action is not allowed.")
	}

	poFollowing := &model.Following{
		UID:          ur.ActiveUid,
		FollowingUID: ur.PassiveUid,
		Status:       status,
	}
	poFollwers := &model.Follower{
		UID:          ur.PassiveUid,
		FollowersUID: ur.ActiveUid,
		Status:       status,
	}

	err := q.Transaction(func(tx *query.Query) error {
		err := dbfg.WithContext(ctx).Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: dbfg.UID.ColumnName().String()},
				{Name: dbfg.FollowingUID.ColumnName().String()},
			},
			DoUpdates: clause.AssignmentColumns([]string{dbfg.Status.ColumnName().String()}),
		}).Create(poFollowing)
		if err != nil {
			return err
		}
		return dbfr.WithContext(ctx).Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: dbfr.UID.ColumnName().String()},
				{Name: dbfr.FollowersUID.ColumnName().String()},
			},
			DoUpdates: clause.AssignmentColumns([]string{dbfg.Status.ColumnName().String()}),
		}).Create(poFollwers)
	})
	if err != nil {
		return v1.ErrorDatabase(err.Error())
	}

	return nil
}
