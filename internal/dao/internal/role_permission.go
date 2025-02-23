// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RolePermissionDao is the data access object for the table role_permission.
type RolePermissionDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns RolePermissionColumns // columns contains all the column names of Table for convenient usage.
}

// RolePermissionColumns defines and stores column names for the table role_permission.
type RolePermissionColumns struct {
	RoleId       string // 主键ID
	PermissionId string // 主键ID
}

// rolePermissionColumns holds the columns for the table role_permission.
var rolePermissionColumns = RolePermissionColumns{
	RoleId:       "role_id",
	PermissionId: "permission_id",
}

// NewRolePermissionDao creates and returns a new DAO object for table data access.
func NewRolePermissionDao() *RolePermissionDao {
	return &RolePermissionDao{
		group:   "default",
		table:   "role_permission",
		columns: rolePermissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RolePermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RolePermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RolePermissionDao) Columns() RolePermissionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RolePermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RolePermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *RolePermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
