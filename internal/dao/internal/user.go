// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for the table user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of the current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for the table user.
type UserColumns struct {
	Id           string // 主键ID
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 删除时间
	AppId        string //
	Avatar       string // 用户头像
	Username     string // 用户登录名
	Password     string // 用户登录密码
	NickName     string // 用户昵称
	PhoneCode    string // 手机区号
	Mobile       string // 手机号
	Email        string // 邮箱
	Gender       string // 性别
	UserType     string // 用户类型
	InviteCode   string // 邀请码
	InviteUserId string // 邀请人id
	OpenId       string //
}

// userColumns holds the columns for the table user.
var userColumns = UserColumns{
	Id:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	AppId:        "app_id",
	Avatar:       "avatar",
	Username:     "username",
	Password:     "password",
	NickName:     "nick_name",
	PhoneCode:    "phone_code",
	Mobile:       "mobile",
	Email:        "email",
	Gender:       "gender",
	UserType:     "user_type",
	InviteCode:   "invite_code",
	InviteUserId: "invite_user_id",
	OpenId:       "open_id",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
