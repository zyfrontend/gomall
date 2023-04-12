// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeAdminActionDao is the data access object for table kcode_admin_action.
type KcodeAdminActionDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns KcodeAdminActionColumns // columns contains all the column names of Table for convenient usage.
}

// KcodeAdminActionColumns defines and stores column names for table kcode_admin_action.
type KcodeAdminActionColumns struct {
	Id         string //
	Module     string // 所属模块名
	Name       string // 行为唯一标识
	Title      string // 行为标题
	Remark     string // 行为描述
	Rule       string // 行为规则
	Log        string // 日志规则
	Status     string // 状态
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// kcodeAdminActionColumns holds the columns for table kcode_admin_action.
var kcodeAdminActionColumns = KcodeAdminActionColumns{
	Id:         "id",
	Module:     "module",
	Name:       "name",
	Title:      "title",
	Remark:     "remark",
	Rule:       "rule",
	Log:        "log",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewKcodeAdminActionDao creates and returns a new DAO object for table data access.
func NewKcodeAdminActionDao() *KcodeAdminActionDao {
	return &KcodeAdminActionDao{
		group:   "default",
		table:   "kcode_admin_action",
		columns: kcodeAdminActionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KcodeAdminActionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KcodeAdminActionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KcodeAdminActionDao) Columns() KcodeAdminActionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KcodeAdminActionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KcodeAdminActionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KcodeAdminActionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}