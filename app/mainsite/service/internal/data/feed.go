package data

import (
	"context"
	"time"

	errs "github.com/fxkt-tech/raiden/api/errs"

	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/biz"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/rdmoment"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/rdrmomentsfollowing"
	"github.com/go-kratos/kratos/v2/log"
)

type feedRepo struct {
	data *Data
	log  *log.Helper

	commonRepo *CommonRepo
	userRepo   biz.UserRepo
}

func NewFeedRepo(data *Data,
	commonRepo *CommonRepo,
	userRepo biz.UserRepo,
	logger log.Logger,
) biz.FeedRepo {
	return &feedRepo{
		data:       data,
		log:        log.NewHelper(logger),
		commonRepo: commonRepo,
		userRepo:   userRepo,
	}
}

// 发布视频时，首先插入自身的历史纪录中；
// 然后插入到自身的关注动态列表中；
// 最后获取粉丝列表，逐条插入他们的关注动态列表中。
func (r *feedRepo) Publish(ctx context.Context, m *biz.Moment) error {
	return r.commonRepo.Transaction(ctx, func(tx *ent.Tx) error {
		nowtime := time.Now()

		moment, err := tx.RdMoment.Create().
			SetType(m.MomentType).
			SetByUID(m.User.Uid).
			SetTxt(m.Text).
			SetImgs(m.Imgs).
			SetPublishTime(nowtime).
			Save(ctx)
		if err != nil {
			return errs.ErrorDatabase(err.Error())
		}
		m.MomentId = moment.ID

		err = tx.RdRMomentsFollowing.Create().
			SetMomentID(moment.ID).
			SetMomentType(m.MomentType).
			SetForUID(m.User.Uid).
			SetByUID(m.User.Uid).
			SetTxt(m.Text).
			SetImgs(m.Imgs).
			SetPublishTime(nowtime).
			Exec(ctx)
		if err != nil {
			return errs.ErrorDatabase(err.Error())
		}

		var page, count int32 = 1, 20
		for {
			followers, _, err := r.userRepo.Followers(ctx, &biz.UserSearch{
				Uid:   m.User.Uid,
				Page:  page,
				Count: count,
			})
			if err != nil {
				return errs.ErrorExtApi(err.Error())
			}
			if len(followers) == 0 {
				break
			}
			for _, user := range followers {
				err = tx.RdRMomentsFollowing.Create().
					SetMomentID(moment.ID).
					SetMomentType(m.MomentType).
					SetForUID(user.Uid).
					SetByUID(m.User.Uid).
					SetTxt(m.Text).
					SetImgs(m.Imgs).
					SetPublishTime(nowtime).
					Exec(ctx)
				if err != nil {
					return errs.ErrorDatabase(err.Error())
				}
			}
			page = page + 1
		}
		return nil
	})
}

func (r *feedRepo) Moments(ctx context.Context, ms *biz.MomentsSearch) ([]*biz.Moment, error) {
	if ms.LastId == 0 {
		ms.LastId = int64(1 << 62)
	}

	momentsFollowing, err := r.data.db.RdRMomentsFollowing.Query().
		Limit(int(ms.PageSize)).
		Order(ent.Desc(rdrmomentsfollowing.FieldPublishTime)).
		Where(rdrmomentsfollowing.ForUIDEQ(ms.Uid),
			rdrmomentsfollowing.StatusEQ(1),
			rdrmomentsfollowing.MomentIDLT(ms.LastId)).
		All(ctx)
	if err != nil {
		return nil, errs.ErrorDatabase(err.Error())
	}

	moments := make([]*biz.Moment, len(momentsFollowing))
	for i, momentFollowing := range momentsFollowing {
		user := &biz.User{Uid: momentFollowing.ByUID}
		userInfo, err := r.userRepo.Info(ctx, momentFollowing.ByUID)
		if err == nil {
			user = &biz.User{
				Uid:  momentFollowing.ByUID,
				Nick: userInfo.Nick,
			}
		}
		moments[i] = &biz.Moment{
			MomentId:    momentFollowing.MomentID,
			MomentType:  momentFollowing.MomentType,
			User:        user,
			Text:        momentFollowing.Txt,
			Imgs:        momentFollowing.Imgs,
			PublishTime: momentFollowing.PublishTime,
		}
	}

	return moments, nil
}

func (r *feedRepo) History(ctx context.Context, ms *biz.MomentsSearch) ([]*biz.Moment, error) {
	if ms.LastId == 0 {
		ms.LastId = int64(1 << 62)
	}

	rdMoments, err := r.data.db.RdMoment.Query().
		Limit(int(ms.PageSize)).
		Order(ent.Desc(rdmoment.FieldPublishTime)).
		Where(rdmoment.ByUIDEQ(ms.Uid),
			rdmoment.StatusEQ(1),
			rdmoment.IDLT(ms.LastId)).
		All(ctx)
	if err != nil {
		return nil, errs.ErrorDatabase(err.Error())
	}

	moments := make([]*biz.Moment, len(rdMoments))
	for i, moment := range rdMoments {
		user := &biz.User{Uid: moment.ByUID}
		userInfo, err := r.userRepo.Info(ctx, moment.ByUID)
		if err == nil {
			user = &biz.User{
				Uid:  moment.ByUID,
				Nick: userInfo.Nick,
			}
		}
		moments[i] = &biz.Moment{
			MomentId:    moment.ID,
			MomentType:  moment.Type,
			User:        user,
			Text:        moment.Txt,
			Imgs:        moment.Imgs,
			PublishTime: moment.PublishTime,
		}
	}

	return moments, nil
}
