// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OauthUserDao is the data access object for the table oauth_user.
type OauthUserDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns OauthUserColumns // columns contains all the column names of Table for convenient usage.
}

// OauthUserColumns defines and stores column names for the table oauth_user.
type OauthUserColumns struct {
	Id        string // 主键ID
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	AppId     string //
	UserId    string // 用户id
	Provider  string // 第三方类型
	OpenId    string // 第三方id
}

// oauthUserColumns holds the columns for the table oauth_user.
var oauthUserColumns = OauthUserColumns{
	Id:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	AppId:     "app_id",
	UserId:    "user_id",
	Provider:  "provider",
	OpenId:    "open_id",
}

// NewOauthUserDao creates and returns a new DAO object for table data access.
func NewOauthUserDao() *OauthUserDao {
	return &OauthUserDao{
		group:   "default",
		table:   "oauth_user",
		columns: oauthUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OauthUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OauthUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OauthUserDao) Columns() OauthUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OauthUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OauthUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OauthUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
