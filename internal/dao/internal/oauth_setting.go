// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OauthSettingDao is the data access object for the table oauth_setting.
type OauthSettingDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns OauthSettingColumns // columns contains all the column names of Table for convenient usage.
}

// OauthSettingColumns defines and stores column names for the table oauth_setting.
type OauthSettingColumns struct {
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

// oauthSettingColumns holds the columns for the table oauth_setting.
var oauthSettingColumns = OauthSettingColumns{
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

// NewOauthSettingDao creates and returns a new DAO object for table data access.
func NewOauthSettingDao() *OauthSettingDao {
	return &OauthSettingDao{
		group:   "default",
		table:   "oauth_setting",
		columns: oauthSettingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OauthSettingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OauthSettingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OauthSettingDao) Columns() OauthSettingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OauthSettingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OauthSettingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OauthSettingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
