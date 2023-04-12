// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeAdminConfigDao is the data access object for table kcode_admin_config.
type KcodeAdminConfigDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns KcodeAdminConfigColumns // columns contains all the column names of Table for convenient usage.
}

// KcodeAdminConfigColumns defines and stores column names for table kcode_admin_config.
type KcodeAdminConfigColumns struct {
	Id         string //
	Name       string // 名称
	Title      string // 标题
	Group      string // 配置分组
	Type       string // 类型
	Value      string // 配置值
	Options    string // 配置项
	Tips       string // 配置提示
	AjaxUrl    string // 联动下拉框ajax地址
	NextItems  string // 联动下拉框的下级下拉框名，多个以逗号隔开
	Param      string // 联动下拉框请求参数名
	Format     string // 格式，用于格式文本
	Table      string // 表名，只用于快速联动类型
	Level      string // 联动级别，只用于快速联动类型
	Key        string // 键字段，只用于快速联动类型
	Option     string // 值字段，只用于快速联动类型
	Pid        string // 父级id字段，只用于快速联动类型
	Ak         string // 百度地图appkey
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	Sort       string // 排序
	Status     string // 状态：0禁用，1启用
}

// kcodeAdminConfigColumns holds the columns for table kcode_admin_config.
var kcodeAdminConfigColumns = KcodeAdminConfigColumns{
	Id:         "id",
	Name:       "name",
	Title:      "title",
	Group:      "group",
	Type:       "type",
	Value:      "value",
	Options:    "options",
	Tips:       "tips",
	AjaxUrl:    "ajax_url",
	NextItems:  "next_items",
	Param:      "param",
	Format:     "format",
	Table:      "table",
	Level:      "level",
	Key:        "key",
	Option:     "option",
	Pid:        "pid",
	Ak:         "ak",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Sort:       "sort",
	Status:     "status",
}

// NewKcodeAdminConfigDao creates and returns a new DAO object for table data access.
func NewKcodeAdminConfigDao() *KcodeAdminConfigDao {
	return &KcodeAdminConfigDao{
		group:   "default",
		table:   "kcode_admin_config",
		columns: kcodeAdminConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KcodeAdminConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KcodeAdminConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KcodeAdminConfigDao) Columns() KcodeAdminConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KcodeAdminConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KcodeAdminConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KcodeAdminConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
