package data

import (
	"context"

	errs "github.com/fxkt-tech/raiden/api/errs"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/biz"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/rdruserfollowers"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/rdruserfollowing"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/rduser"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper

	commonRepo *CommonRepo
}

func NewUserRepo(data *Data, commonRepo *CommonRepo, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data:       data,
		log:        log.NewHelper(logger),
		commonRepo: commonRepo,
	}
}

func (r *userRepo) Register(ctx context.Context, u *biz.User) error {
	user, err := r.data.db.RdUser.Create().SetNick(u.Nick).Save(ctx)
	if err != nil {
		return errs.ErrorDatabase(err.Error())
	}
	u.Uid = user.ID
	return nil
}

// 只能查询正常状态的用户信息
func (r *userRepo) Info(ctx context.Context, uid int32) (*biz.User, error) {
	user, err := r.data.db.RdUser.Query().Where(rduser.IDEQ(uid), rduser.StatusEQ(1)).First(ctx)
	if err != nil {
		return nil, errs.ErrorDatabase(err.Error())
	}
	return &biz.User{
		Uid:  user.ID,
		Nick: user.Nick,
	}, nil
}

func (r *userRepo) Followers(ctx context.Context, us *biz.UserSearch) ([]*biz.User, int32, error) {
	userFollowers, err := r.data.db.RdRUserFollowers.Query().
		Limit(int(us.Count)).
		Offset(int((us.Page - 1) * us.Count)).
		Where(rdruserfollowers.UIDEQ(us.Uid)).All(ctx)
	if err != nil {
		return nil, 0, errs.ErrorDatabase(err.Error())
	}

	users := make([]*biz.User, len(userFollowers))
	for i, userFollower := range userFollowers {
		users[i] = &biz.User{
			Uid: userFollower.UID,
		}
	}

	return users, 0, nil
}

func (r *userRepo) Following(ctx context.Context, us *biz.UserSearch) ([]*biz.User, int32, error) {
	userFollowing, err := r.data.db.RdRUserFollowers.Query().
		Limit(int(us.Count)).
		Offset(int((us.Page - 1) * us.Count)).
		Where(rdruserfollowers.UIDEQ(us.Uid)).All(ctx)
	if err != nil {
		return nil, 0, errs.ErrorDatabase(err.Error())
	}

	users := make([]*biz.User, len(userFollowing))
	for i, userFollowing := range userFollowing {
		users[i] = &biz.User{
			Uid: userFollowing.UID,
		}
	}

	return users, 0, nil
}

// 点击关注触发的事件
// 1. 添加到主动人的关注列表里
// 2. 添加到被动人的粉丝列表里
// 取关同理
func (r *userRepo) Relation(ctx context.Context, ur *biz.UserRelation) error {
	switch ur.Action {
	case 1: // 关注
		exist, err := r.data.db.RdRUserFollowing.Query().
			Where(rdruserfollowing.UIDEQ(ur.ActiveUid),
				rdruserfollowing.FollowingUIDEQ(ur.PassiveUid)).
			Exist(ctx)
		switch {
		case err != nil:
			return errs.ErrorDatabase(err.Error())
		case exist:
			return errs.ErrorForbidden("重复操作")
		}

		exist, err = r.data.db.RdRUserFollowers.Query().
			Where(rdruserfollowers.UIDEQ(ur.ActiveUid),
				rdruserfollowers.FollowersUIDEQ(ur.PassiveUid)).
			Exist(ctx)
		switch {
		case err != nil:
			return errs.ErrorDatabase(err.Error())
		case exist:
			return errs.ErrorForbidden("重复操作")
		}

		return r.commonRepo.Transaction(ctx, func(tx *ent.Tx) error {
			err = tx.RdRUserFollowing.Create().
				SetUID(ur.ActiveUid).
				SetFollowingUID(ur.PassiveUid).
				Exec(ctx)
			if err != nil {
				return errs.ErrorDatabase(err.Error())
			}

			err = tx.RdRUserFollowers.Create().
				SetUID(ur.PassiveUid).
				SetFollowersUID(ur.ActiveUid).
				Exec(ctx)
			if err != nil {
				return errs.ErrorDatabase(err.Error())
			}

			return nil
		})
	case 2: // 取消关注
		return r.commonRepo.Transaction(ctx, func(tx *ent.Tx) error {
			_, err := tx.RdRUserFollowing.Delete().
				Where(rdruserfollowing.UIDEQ(ur.ActiveUid),
					rdruserfollowing.FollowingUIDEQ(ur.PassiveUid)).
				Exec(ctx)
			if err != nil {
				return errs.ErrorDatabase(err.Error())
			}

			_, err = tx.RdRUserFollowers.Delete().
				Where(rdruserfollowers.UIDEQ(ur.PassiveUid),
					rdruserfollowers.FollowersUIDEQ(ur.ActiveUid)).
				Exec(ctx)
			if err != nil {
				return errs.ErrorDatabase(err.Error())
			}

			return nil
		})
	default:
		return errs.ErrorValidator("action is not allowed")
	}
}
