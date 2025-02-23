// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SchemaMigrationsDao is the data access object for the table schema_migrations.
type SchemaMigrationsDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns SchemaMigrationsColumns // columns contains all the column names of Table for convenient usage.
}

// SchemaMigrationsColumns defines and stores column names for the table schema_migrations.
type SchemaMigrationsColumns struct {
	Version string //
	Dirty   string //
}

// schemaMigrationsColumns holds the columns for the table schema_migrations.
var schemaMigrationsColumns = SchemaMigrationsColumns{
	Version: "version",
	Dirty:   "dirty",
}

// NewSchemaMigrationsDao creates and returns a new DAO object for table data access.
func NewSchemaMigrationsDao() *SchemaMigrationsDao {
	return &SchemaMigrationsDao{
		group:   "default",
		table:   "schema_migrations",
		columns: schemaMigrationsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SchemaMigrationsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SchemaMigrationsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SchemaMigrationsDao) Columns() SchemaMigrationsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SchemaMigrationsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SchemaMigrationsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SchemaMigrationsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
