// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"fxkt.tech/raiden/app/user/service/internal/data/db/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newUserFollowing(db *gorm.DB) userFollowing {
	_userFollowing := userFollowing{}

	_userFollowing.userFollowingDo.UseDB(db)
	_userFollowing.userFollowingDo.UseModel(&model.UserFollowing{})

	tableName := _userFollowing.userFollowingDo.TableName()
	_userFollowing.ALL = field.NewField(tableName, "*")
	_userFollowing.UID = field.NewInt32(tableName, "uid")
	_userFollowing.FollowingUID = field.NewInt32(tableName, "following_uid")
	_userFollowing.Status = field.NewInt32(tableName, "status")
	_userFollowing.ActionTime = field.NewTime(tableName, "action_time")

	_userFollowing.fillFieldMap()

	return _userFollowing
}

type userFollowing struct {
	userFollowingDo userFollowingDo

	ALL          field.Field
	UID          field.Int32
	FollowingUID field.Int32
	Status       field.Int32
	ActionTime   field.Time

	fieldMap map[string]field.Expr
}

func (u userFollowing) Table(newTableName string) *userFollowing {
	u.userFollowingDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userFollowing) As(alias string) *userFollowing {
	u.userFollowingDo.DO = *(u.userFollowingDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userFollowing) updateTableName(table string) *userFollowing {
	u.ALL = field.NewField(table, "*")
	u.UID = field.NewInt32(table, "uid")
	u.FollowingUID = field.NewInt32(table, "following_uid")
	u.Status = field.NewInt32(table, "status")
	u.ActionTime = field.NewTime(table, "action_time")

	u.fillFieldMap()

	return u
}

func (u *userFollowing) WithContext(ctx context.Context) *userFollowingDo {
	return u.userFollowingDo.WithContext(ctx)
}

func (u userFollowing) TableName() string { return u.userFollowingDo.TableName() }

func (u *userFollowing) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userFollowing) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 4)
	u.fieldMap["uid"] = u.UID
	u.fieldMap["following_uid"] = u.FollowingUID
	u.fieldMap["status"] = u.Status
	u.fieldMap["action_time"] = u.ActionTime
}

func (u userFollowing) clone(db *gorm.DB) userFollowing {
	u.userFollowingDo.ReplaceDB(db)
	return u
}

type userFollowingDo struct{ gen.DO }

func (u userFollowingDo) Debug() *userFollowingDo {
	return u.withDO(u.DO.Debug())
}

func (u userFollowingDo) WithContext(ctx context.Context) *userFollowingDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userFollowingDo) Clauses(conds ...clause.Expression) *userFollowingDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userFollowingDo) Returning(value interface{}, columns ...string) *userFollowingDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userFollowingDo) Not(conds ...gen.Condition) *userFollowingDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userFollowingDo) Or(conds ...gen.Condition) *userFollowingDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userFollowingDo) Select(conds ...field.Expr) *userFollowingDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userFollowingDo) Where(conds ...gen.Condition) *userFollowingDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userFollowingDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userFollowingDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userFollowingDo) Order(conds ...field.Expr) *userFollowingDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userFollowingDo) Distinct(cols ...field.Expr) *userFollowingDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userFollowingDo) Omit(cols ...field.Expr) *userFollowingDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userFollowingDo) Join(table schema.Tabler, on ...field.Expr) *userFollowingDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userFollowingDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userFollowingDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userFollowingDo) RightJoin(table schema.Tabler, on ...field.Expr) *userFollowingDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userFollowingDo) Group(cols ...field.Expr) *userFollowingDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userFollowingDo) Having(conds ...gen.Condition) *userFollowingDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userFollowingDo) Limit(limit int) *userFollowingDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userFollowingDo) Offset(offset int) *userFollowingDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userFollowingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userFollowingDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userFollowingDo) Unscoped() *userFollowingDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userFollowingDo) Create(values ...*model.UserFollowing) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userFollowingDo) CreateInBatches(values []*model.UserFollowing, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userFollowingDo) Save(values ...*model.UserFollowing) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userFollowingDo) First() (*model.UserFollowing, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFollowing), nil
	}
}

func (u userFollowingDo) Take() (*model.UserFollowing, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFollowing), nil
	}
}

func (u userFollowingDo) Last() (*model.UserFollowing, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFollowing), nil
	}
}

func (u userFollowingDo) Find() ([]*model.UserFollowing, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserFollowing), err
}

func (u userFollowingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserFollowing, err error) {
	buf := make([]*model.UserFollowing, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userFollowingDo) FindInBatches(result *[]*model.UserFollowing, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userFollowingDo) Attrs(attrs ...field.AssignExpr) *userFollowingDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userFollowingDo) Assign(attrs ...field.AssignExpr) *userFollowingDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userFollowingDo) Joins(field field.RelationField) *userFollowingDo {
	return u.withDO(u.DO.Joins(field))
}

func (u userFollowingDo) Preload(field field.RelationField) *userFollowingDo {
	return u.withDO(u.DO.Preload(field))
}

func (u userFollowingDo) FirstOrInit() (*model.UserFollowing, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFollowing), nil
	}
}

func (u userFollowingDo) FirstOrCreate() (*model.UserFollowing, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserFollowing), nil
	}
}

func (u userFollowingDo) FindByPage(offset int, limit int) (result []*model.UserFollowing, count int64, err error) {
	if limit <= 0 {
		count, err = u.Count()
		return
	}

	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userFollowingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u *userFollowingDo) withDO(do gen.Dao) *userFollowingDo {
	u.DO = *do.(*gen.DO)
	return u
}
