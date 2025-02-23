// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivationDao is the data access object for table activation.
type ActivationDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns ActivationColumns // columns contains all the column names of Table for convenient usage.
}

// ActivationColumns defines and stores column names for table activation.
type ActivationColumns struct {
	Id         string // 主键ID
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	DeletedAt  string // 删除时间
	Code       string // 激活码
	UserId     string // 用户id
	AppId      string // 应用id
	AppKey     string // AppKey
	Level      string // 会员等级,default:1
	Duration   string // 有效时长
	UseAt      string // 使用时间
	StartAt    string // 开始时间
	ExpiresAt  string // 到期时间
	SubjectId  string // 主体id
	Valid      string // 状态
	Remark     string // 备注
	Exported   string // 是否已导出
	Version    string // 版本
	VipType    string // 会员类型
	AppOwnerId string // 用户id
}

// activationColumns holds the columns for table activation.
var activationColumns = ActivationColumns{
	Id:         "id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	Code:       "code",
	UserId:     "user_id",
	AppId:      "app_id",
	AppKey:     "app_key",
	Level:      "level",
	Duration:   "duration",
	UseAt:      "use_at",
	StartAt:    "start_at",
	ExpiresAt:  "expires_at",
	SubjectId:  "subject_id",
	Valid:      "valid",
	Remark:     "remark",
	Exported:   "exported",
	Version:    "version",
	VipType:    "vip_type",
	AppOwnerId: "app_owner_id",
}

// NewActivationDao creates and returns a new DAO object for table data access.
func NewActivationDao() *ActivationDao {
	return &ActivationDao{
		group:   "default",
		table:   "activation",
		columns: activationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivationDao) Columns() ActivationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
