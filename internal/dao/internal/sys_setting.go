// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysSettingDao is the data access object for the table sys_setting.
type SysSettingDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns SysSettingColumns // columns contains all the column names of Table for convenient usage.
}

// SysSettingColumns defines and stores column names for the table sys_setting.
type SysSettingColumns struct {
	Id        string // 主键ID
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	Key       string // 应用code
	Title     string // 显示名称
	Setting   string // Logo
}

// sysSettingColumns holds the columns for the table sys_setting.
var sysSettingColumns = SysSettingColumns{
	Id:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Key:       "key",
	Title:     "title",
	Setting:   "setting",
}

// NewSysSettingDao creates and returns a new DAO object for table data access.
func NewSysSettingDao() *SysSettingDao {
	return &SysSettingDao{
		group:   "default",
		table:   "sys_setting",
		columns: sysSettingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysSettingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysSettingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysSettingDao) Columns() SysSettingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysSettingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysSettingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysSettingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
