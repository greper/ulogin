// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LoginLogDao is the data access object for the table login_log.
type LoginLogDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns LoginLogColumns // columns contains all the column names of Table for convenient usage.
}

// LoginLogColumns defines and stores column names for the table login_log.
type LoginLogColumns struct {
	Id        string //
	UserId    string //
	LoginTime string //
	AppId     string //
	ClientId  string //
	LoginIp   string //
}

// loginLogColumns holds the columns for the table login_log.
var loginLogColumns = LoginLogColumns{
	Id:        "id",
	UserId:    "user_id",
	LoginTime: "login_time",
	AppId:     "app_id",
	ClientId:  "client_id",
	LoginIp:   "login_ip",
}

// NewLoginLogDao creates and returns a new DAO object for table data access.
func NewLoginLogDao() *LoginLogDao {
	return &LoginLogDao{
		group:   "default",
		table:   "login_log",
		columns: loginLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LoginLogDao) Columns() LoginLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *LoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
