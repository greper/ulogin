// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppDao is the data access object for the table app.
type AppDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of the current DAO.
	columns AppColumns // columns contains all the column names of Table for convenient usage.
}

// AppColumns defines and stores column names for the table app.
type AppColumns struct {
	Id        string // 主键ID
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	Logo      string // Logo
	Key       string // 应用code
	Title     string // 显示名称
}

// appColumns holds the columns for the table app.
var appColumns = AppColumns{
	Id:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Logo:      "logo",
	Key:       "key",
	Title:     "title",
}

// NewAppDao creates and returns a new DAO object for table data access.
func NewAppDao() *AppDao {
	return &AppDao{
		group:   "default",
		table:   "app",
		columns: appColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppDao) Columns() AppColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AppDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
