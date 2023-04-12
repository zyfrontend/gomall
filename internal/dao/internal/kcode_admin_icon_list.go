// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeAdminIconListDao is the data access object for table kcode_admin_icon_list.
type KcodeAdminIconListDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns KcodeAdminIconListColumns // columns contains all the column names of Table for convenient usage.
}

// KcodeAdminIconListColumns defines and stores column names for table kcode_admin_icon_list.
type KcodeAdminIconListColumns struct {
	Id     string //
	IconId string // 所属图标id
	Title  string // 图标标题
	Class  string // 图标类名
	Code   string // 图标关键词
}

// kcodeAdminIconListColumns holds the columns for table kcode_admin_icon_list.
var kcodeAdminIconListColumns = KcodeAdminIconListColumns{
	Id:     "id",
	IconId: "icon_id",
	Title:  "title",
	Class:  "class",
	Code:   "code",
}

// NewKcodeAdminIconListDao creates and returns a new DAO object for table data access.
func NewKcodeAdminIconListDao() *KcodeAdminIconListDao {
	return &KcodeAdminIconListDao{
		group:   "default",
		table:   "kcode_admin_icon_list",
		columns: kcodeAdminIconListColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KcodeAdminIconListDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KcodeAdminIconListDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KcodeAdminIconListDao) Columns() KcodeAdminIconListColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KcodeAdminIconListDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KcodeAdminIconListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KcodeAdminIconListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
