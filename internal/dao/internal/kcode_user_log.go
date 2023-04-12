// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeUserLogDao is the data access object for table kcode_user_log.
type KcodeUserLogDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns KcodeUserLogColumns // columns contains all the column names of Table for convenient usage.
}

// KcodeUserLogColumns defines and stores column names for table kcode_user_log.
type KcodeUserLogColumns struct {
	Id      string //
	UserId  string //
	Msg     string // 用户操作信息
	AddTime string // 添加时间
}

// kcodeUserLogColumns holds the columns for table kcode_user_log.
var kcodeUserLogColumns = KcodeUserLogColumns{
	Id:      "id",
	UserId:  "user_id",
	Msg:     "msg",
	AddTime: "add_time",
}

// NewKcodeUserLogDao creates and returns a new DAO object for table data access.
func NewKcodeUserLogDao() *KcodeUserLogDao {
	return &KcodeUserLogDao{
		group:   "default",
		table:   "kcode_user_log",
		columns: kcodeUserLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KcodeUserLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KcodeUserLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KcodeUserLogDao) Columns() KcodeUserLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KcodeUserLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KcodeUserLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KcodeUserLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}