// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/MQEnergy/go-skeleton/internal/app/model"
)

func newSysUser(db *gorm.DB, opts ...gen.DOOption) sysUser {
	_sysUser := sysUser{}

	_sysUser.sysUserDo.UseDB(db, opts...)
	_sysUser.sysUserDo.UseModel(&model.SysUser{})

	tableName := _sysUser.sysUserDo.TableName()
	_sysUser.ALL = field.NewAsterisk(tableName)
	_sysUser.ID = field.NewInt64(tableName, "id")
	_sysUser.UUID = field.NewString(tableName, "uuid")
	_sysUser.Username = field.NewString(tableName, "username")
	_sysUser.Password = field.NewString(tableName, "password")
	_sysUser.Phone = field.NewString(tableName, "phone")
	_sysUser.Avatar = field.NewString(tableName, "avatar")
	_sysUser.Salt = field.NewString(tableName, "salt")
	_sysUser.RealName = field.NewString(tableName, "real_name")
	_sysUser.RegisterTime = field.NewInt64(tableName, "register_time")
	_sysUser.RegisterIP = field.NewString(tableName, "register_ip")
	_sysUser.LoginTime = field.NewInt64(tableName, "login_time")
	_sysUser.LoginIP = field.NewString(tableName, "login_ip")
	_sysUser.RoleIds = field.NewString(tableName, "role_ids")
	_sysUser.Status = field.NewInt8(tableName, "status")
	_sysUser.CreatedAt = field.NewInt64(tableName, "created_at")
	_sysUser.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_sysUser.fillFieldMap()

	return _sysUser
}

// sysUser 后台管理员表
type sysUser struct {
	sysUserDo

	ALL          field.Asterisk
	ID           field.Int64
	UUID         field.String // 唯一id号
	Username     field.String // 账号
	Password     field.String // 密码
	Phone        field.String // 手机号
	Avatar       field.String // 头像
	Salt         field.String // 密码
	RealName     field.String // 真实姓名
	RegisterTime field.Int64  // 注册时间
	RegisterIP   field.String // 注册ip
	LoginTime    field.Int64  // 登录时间
	LoginIP      field.String // 登录ip
	RoleIds      field.String // 角色IDs
	Status       field.Int8   // 状态 1：正常 2：禁用
	CreatedAt    field.Int64
	UpdatedAt    field.Int64

	fieldMap map[string]field.Expr
}

func (s sysUser) Table(newTableName string) *sysUser {
	s.sysUserDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysUser) As(alias string) *sysUser {
	s.sysUserDo.DO = *(s.sysUserDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysUser) updateTableName(table string) *sysUser {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.UUID = field.NewString(table, "uuid")
	s.Username = field.NewString(table, "username")
	s.Password = field.NewString(table, "password")
	s.Phone = field.NewString(table, "phone")
	s.Avatar = field.NewString(table, "avatar")
	s.Salt = field.NewString(table, "salt")
	s.RealName = field.NewString(table, "real_name")
	s.RegisterTime = field.NewInt64(table, "register_time")
	s.RegisterIP = field.NewString(table, "register_ip")
	s.LoginTime = field.NewInt64(table, "login_time")
	s.LoginIP = field.NewString(table, "login_ip")
	s.RoleIds = field.NewString(table, "role_ids")
	s.Status = field.NewInt8(table, "status")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *sysUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysUser) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 16)
	s.fieldMap["id"] = s.ID
	s.fieldMap["uuid"] = s.UUID
	s.fieldMap["username"] = s.Username
	s.fieldMap["password"] = s.Password
	s.fieldMap["phone"] = s.Phone
	s.fieldMap["avatar"] = s.Avatar
	s.fieldMap["salt"] = s.Salt
	s.fieldMap["real_name"] = s.RealName
	s.fieldMap["register_time"] = s.RegisterTime
	s.fieldMap["register_ip"] = s.RegisterIP
	s.fieldMap["login_time"] = s.LoginTime
	s.fieldMap["login_ip"] = s.LoginIP
	s.fieldMap["role_ids"] = s.RoleIds
	s.fieldMap["status"] = s.Status
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s sysUser) clone(db *gorm.DB) sysUser {
	s.sysUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysUser) replaceDB(db *gorm.DB) sysUser {
	s.sysUserDo.ReplaceDB(db)
	return s
}

type sysUserDo struct{ gen.DO }

type ISysUserDo interface {
	gen.SubQuery
	Debug() ISysUserDo
	WithContext(ctx context.Context) ISysUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysUserDo
	WriteDB() ISysUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysUserDo
	Not(conds ...gen.Condition) ISysUserDo
	Or(conds ...gen.Condition) ISysUserDo
	Select(conds ...field.Expr) ISysUserDo
	Where(conds ...gen.Condition) ISysUserDo
	Order(conds ...field.Expr) ISysUserDo
	Distinct(cols ...field.Expr) ISysUserDo
	Omit(cols ...field.Expr) ISysUserDo
	Join(table schema.Tabler, on ...field.Expr) ISysUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysUserDo
	Group(cols ...field.Expr) ISysUserDo
	Having(conds ...gen.Condition) ISysUserDo
	Limit(limit int) ISysUserDo
	Offset(offset int) ISysUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysUserDo
	Unscoped() ISysUserDo
	Create(values ...*model.SysUser) error
	CreateInBatches(values []*model.SysUser, batchSize int) error
	Save(values ...*model.SysUser) error
	First() (*model.SysUser, error)
	Take() (*model.SysUser, error)
	Last() (*model.SysUser, error)
	Find() ([]*model.SysUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysUser, err error)
	FindInBatches(result *[]*model.SysUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysUserDo
	Assign(attrs ...field.AssignExpr) ISysUserDo
	Joins(fields ...field.RelationField) ISysUserDo
	Preload(fields ...field.RelationField) ISysUserDo
	FirstOrInit() (*model.SysUser, error)
	FirstOrCreate() (*model.SysUser, error)
	FindByPage(offset int, limit int) (result []*model.SysUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id int) (result model.SysUser, err error)
	FindAll() (result []model.SysUser, err error)
	FindOne() (result model.SysUser)
	GetByUserName(username string) (result *model.SysUser, err error)
}
// SELECT * FROM @@table WHERE id = @id
func (s sysUserDo) GetByID(id int) (result model.SysUser, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM sys_user WHERE id = ? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
func (s sysUserDo) FindAll() (result []model.SysUser, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM sys_user ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table LIMIT 1
func (s sysUserDo) FindOne() (result model.SysUser) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM sys_user LIMIT 1 ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	_ = executeSQL

	return
}

// SELECT * FROM @@table WHERE username = @username
func (s sysUserDo) GetByUserName(account string) (result *model.SysUser, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, account)
	generateSQL.WriteString("SELECT * FROM sys_user WHERE username = ? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (s sysUserDo) Debug() ISysUserDo {
	return s.withDO(s.DO.Debug())
}

func (s sysUserDo) WithContext(ctx context.Context) ISysUserDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysUserDo) ReadDB() ISysUserDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysUserDo) WriteDB() ISysUserDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysUserDo) Session(config *gorm.Session) ISysUserDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysUserDo) Clauses(conds ...clause.Expression) ISysUserDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysUserDo) Returning(value interface{}, columns ...string) ISysUserDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysUserDo) Not(conds ...gen.Condition) ISysUserDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysUserDo) Or(conds ...gen.Condition) ISysUserDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysUserDo) Select(conds ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysUserDo) Where(conds ...gen.Condition) ISysUserDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysUserDo) Order(conds ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysUserDo) Distinct(cols ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysUserDo) Omit(cols ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysUserDo) Join(table schema.Tabler, on ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysUserDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysUserDo) Group(cols ...field.Expr) ISysUserDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysUserDo) Having(conds ...gen.Condition) ISysUserDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysUserDo) Limit(limit int) ISysUserDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysUserDo) Offset(offset int) ISysUserDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysUserDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysUserDo) Unscoped() ISysUserDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysUserDo) Create(values ...*model.SysUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysUserDo) CreateInBatches(values []*model.SysUser, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysUserDo) Save(values ...*model.SysUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysUserDo) First() (*model.SysUser, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) Take() (*model.SysUser, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) Last() (*model.SysUser, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) Find() ([]*model.SysUser, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysUser), err
}

func (s sysUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysUser, err error) {
	buf := make([]*model.SysUser, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysUserDo) FindInBatches(result *[]*model.SysUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysUserDo) Attrs(attrs ...field.AssignExpr) ISysUserDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysUserDo) Assign(attrs ...field.AssignExpr) ISysUserDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysUserDo) Joins(fields ...field.RelationField) ISysUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysUserDo) Preload(fields ...field.RelationField) ISysUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysUserDo) FirstOrInit() (*model.SysUser, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) FirstOrCreate() (*model.SysUser, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) FindByPage(offset int, limit int) (result []*model.SysUser, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysUserDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysUserDo) Delete(models ...*model.SysUser) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysUserDo) withDO(do gen.Dao) *sysUserDo {
	s.DO = *do.(*gen.DO)
	return s
}
