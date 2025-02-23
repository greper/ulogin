// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OauthClientDao is the data access object for the table oauth_client.
type OauthClientDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns OauthClientColumns // columns contains all the column names of Table for convenient usage.
}

// OauthClientColumns defines and stores column names for the table oauth_client.
type OauthClientColumns struct {
	Id           string // 主键ID
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 删除时间
	AppId        string //
	Key          string // code
	Title        string // 显示名称
	Type         string // 类型
	AccessId     string // appid
	AccessSecret string // 密钥
}

// oauthClientColumns holds the columns for the table oauth_client.
var oauthClientColumns = OauthClientColumns{
	Id:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	AppId:        "app_id",
	Key:          "key",
	Title:        "title",
	Type:         "type",
	AccessId:     "access_id",
	AccessSecret: "access_secret",
}

// NewOauthClientDao creates and returns a new DAO object for table data access.
func NewOauthClientDao() *OauthClientDao {
	return &OauthClientDao{
		group:   "default",
		table:   "oauth_client",
		columns: oauthClientColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OauthClientDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OauthClientDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OauthClientDao) Columns() OauthClientColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OauthClientDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OauthClientDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OauthClientDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
