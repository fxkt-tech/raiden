package data

import (
	"context"
	"time"

	v1 "fxkt.tech/raiden/api/feed/service/v1"
	userv1 "fxkt.tech/raiden/api/user/service/v1"
	"fxkt.tech/raiden/app/feed/service/internal/biz"
	"fxkt.tech/raiden/app/feed/service/internal/data/db/model"
	"fxkt.tech/raiden/app/feed/service/internal/data/db/query"
	"fxkt.tech/raiden/pkg/json"
	"github.com/go-kratos/kratos/v2/log"
)

type feedSystemRepo struct {
	data *Data
	log  *log.Helper
}

func NewfeedSystemRepo(data *Data, logger log.Logger) biz.FeedSystemRepo {
	return &feedSystemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 发布视频时，首先插入自身的历史纪录中；
// 然后插入到自身的关注动态列表中；
// 最后获取粉丝列表，逐条插入他们的关注动态列表中。
func (r *feedSystemRepo) Publish(ctx context.Context, d *biz.Dynamic) error {
	return r.data.db.Transaction(func(tx *query.Query) error {
		var (
			dbdh    = tx.DynamicHistory
			dbdf    = tx.DynamicFollowing
			nowtime = time.Now()
		)

		poDmch := &model.DynamicHistory{
			DmcType:     d.DmcType,
			ByUID:       d.User.Uid,
			Txt:         d.Text,
			Imgs:        json.ToString(&d.Imgs),
			PublishTime: nowtime,
		}
		err := dbdh.WithContext(ctx).Create(poDmch)
		if err != nil {
			return v1.ErrorDatabase(err.Error())
		}
		d.DmcId = poDmch.DmcID

		poDmcf := &model.DynamicFollowing{
			DmcID:       poDmch.DmcID,
			DmcType:     d.DmcType,
			ForUID:      d.User.Uid,
			ByUID:       d.User.Uid,
			Txt:         d.Text,
			Imgs:        json.ToString(&d.Imgs),
			PublishTime: nowtime,
		}
		err = dbdf.WithContext(ctx).Create(poDmcf)
		if err != nil {
			return v1.ErrorDatabase(err.Error())
		}

		var page, count int32 = 1, 20
		for {
			followers, err := r.data.userService.Followers(ctx, &userv1.FollowersRequest{
				Uid:   d.User.Uid,
				Page:  page,
				Count: count,
			})
			if err != nil {
				return v1.ErrorExtApi(err.Error())
			}
			if len(followers.Users) == 0 {
				break
			}
			for _, user := range followers.Users {
				poDmcf := &model.DynamicFollowing{
					DmcID:       poDmch.DmcID,
					DmcType:     d.DmcType,
					ForUID:      user.Uid,
					ByUID:       d.User.Uid,
					Txt:         d.Text,
					Imgs:        json.ToString(&d.Imgs),
					PublishTime: nowtime,
				}
				err = dbdf.WithContext(ctx).Create(poDmcf)
				if err != nil {
					return v1.ErrorDatabase(err.Error())
				}
			}
			page = page + 1
		}
		return nil
	})

}
func (r *feedSystemRepo) Following(ctx context.Context, ds *biz.DynamicSearch) ([]*biz.Dynamic, error) {
	var (
		dbdf = r.data.db.DynamicFollowing
	)

	if ds.LastId == 0 {
		ds.LastId = int64(1 << 62)
	}
	poDmcf, err := dbdf.WithContext(ctx).
		Limit(int(ds.PageSize)).
		Order(dbdf.PublishTime.Desc()).
		Where(dbdf.ForUID.Eq(ds.Uid), dbdf.Status.Eq(1), dbdf.DmcID.Lt(ds.LastId)).
		Select(dbdf.DmcID, dbdf.DmcType, dbdf.ByUID, dbdf.Txt, dbdf.Imgs, dbdf.PublishTime).
		Find()
	if err != nil {
		return nil, v1.ErrorDatabase(err.Error())
	}

	dmcs := make([]*biz.Dynamic, len(poDmcf))
	for i, dmcf := range poDmcf {
		var user *biz.User
		reply, err := r.data.userService.Info(ctx, &userv1.InfoRequest{Uid: dmcf.ByUID})
		if err != nil {
			user = &biz.User{Uid: dmcf.ByUID}
		} else {
			user = &biz.User{
				Uid:  dmcf.ByUID,
				Nick: reply.User.Nick,
			}
		}
		var imgs []string
		json.ToObject(dmcf.Imgs, &imgs)
		dmcs[i] = &biz.Dynamic{
			DmcId:       dmcf.DmcID,
			DmcType:     dmcf.DmcType,
			User:        user,
			Text:        dmcf.Txt,
			Imgs:        imgs,
			PublishTime: dmcf.PublishTime,
		}
	}

	return dmcs, nil
}

func (r *feedSystemRepo) History(ctx context.Context, ds *biz.DynamicSearch) ([]*biz.Dynamic, error) {
	var (
		dbdh = r.data.db.DynamicHistory
	)
	if ds.LastId == 0 {
		ds.LastId = int64(1 << 62)
	}
	poDmch, err := dbdh.WithContext(ctx).
		Limit(int(ds.PageSize)).
		Order(dbdh.PublishTime.Desc()).
		Where(dbdh.ByUID.Eq(ds.Uid), dbdh.Status.Eq(1), dbdh.DmcID.Lt(ds.LastId)).
		Select(dbdh.DmcID, dbdh.DmcType, dbdh.ByUID, dbdh.Txt, dbdh.Imgs, dbdh.PublishTime).
		Find()
	if err != nil {
		return nil, v1.ErrorDatabase(err.Error())
	}

	dmcs := make([]*biz.Dynamic, len(poDmch))
	for i, dmch := range poDmch {
		var user *biz.User
		reply, err := r.data.userService.Info(ctx, &userv1.InfoRequest{Uid: dmch.ByUID})
		if err != nil {
			user = &biz.User{Uid: dmch.ByUID}
		} else {
			user = &biz.User{
				Uid:  dmch.ByUID,
				Nick: reply.User.Nick,
			}
		}
		var imgs []string
		json.ToObject(dmch.Imgs, &imgs)
		dmcs[i] = &biz.Dynamic{
			DmcId:       dmch.DmcID,
			DmcType:     dmch.DmcType,
			User:        user,
			Text:        dmch.Txt,
			Imgs:        imgs,
			PublishTime: dmch.PublishTime,
		}
	}

	return dmcs, nil
}
