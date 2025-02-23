// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccessDao is the data access object for the table access.
type AccessDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of the current DAO.
	columns AccessColumns // columns contains all the column names of Table for convenient usage.
}

// AccessColumns defines and stores column names for the table access.
type AccessColumns struct {
	Id               string // 主键ID
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
	DeletedAt        string // 删除时间
	Key              string // code
	Title            string // 显示名称
	Type             string // 类型
	Setting          string // 配置
	EncryptedSetting string // 加密配置
}

// accessColumns holds the columns for the table access.
var accessColumns = AccessColumns{
	Id:               "id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
	Key:              "key",
	Title:            "title",
	Type:             "type",
	Setting:          "setting",
	EncryptedSetting: "encrypted_setting",
}

// NewAccessDao creates and returns a new DAO object for table data access.
func NewAccessDao() *AccessDao {
	return &AccessDao{
		group:   "default",
		table:   "access",
		columns: accessColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AccessDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AccessDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AccessDao) Columns() AccessColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AccessDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AccessDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AccessDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
