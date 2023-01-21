package data

import (
	"context"

	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent"
	"github.com/go-kratos/kratos/v2/log"
)

type CommonRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommonRepo(data *Data, logger log.Logger) *CommonRepo {
	return &CommonRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *CommonRepo) Transaction(ctx context.Context, f func(tx *ent.Tx) error) error {
	tx, err := r.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	err = f(tx)
	if err != nil {
		e := tx.Rollback()
		if e != nil {
			return e
		}
		return err
	}
	return tx.Commit()
}
