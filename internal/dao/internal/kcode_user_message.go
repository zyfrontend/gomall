// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeUserMessageDao is the data access object for table kcode_user_message.
type KcodeUserMessageDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns KcodeUserMessageColumns // columns contains all the column names of Table for convenient usage.
}

// KcodeUserMessageColumns defines and stores column names for table kcode_user_message.
type KcodeUserMessageColumns struct {
	Id         string //
	UidReceive string // 接收消息的用户id
	UidSend    string // 发送消息的用户id
	Type       string // 消息分类
	Title      string // 发送的标题
	Content    string // 消息内容
	Status     string // 状态
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	ReadTime   string // 阅读时间
}

// kcodeUserMessageColumns holds the columns for table kcode_user_message.
var kcodeUserMessageColumns = KcodeUserMessageColumns{
	Id:         "id",
	UidReceive: "uid_receive",
	UidSend:    "uid_send",
	Type:       "type",
	Title:      "title",
	Content:    "content",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	ReadTime:   "read_time",
}

// NewKcodeUserMessageDao creates and returns a new DAO object for table data access.
func NewKcodeUserMessageDao() *KcodeUserMessageDao {
	return &KcodeUserMessageDao{
		group:   "default",
		table:   "kcode_user_message",
		columns: kcodeUserMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KcodeUserMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KcodeUserMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KcodeUserMessageDao) Columns() KcodeUserMessageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KcodeUserMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KcodeUserMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KcodeUserMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
