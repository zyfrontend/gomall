// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeUserUpLogDao is the data access object for table kcode_user_up_log.
type KcodeUserUpLogDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns KcodeUserUpLogColumns // columns contains all the column names of Table for convenient usage.
}

// KcodeUserUpLogColumns defines and stores column names for table kcode_user_up_log.
type KcodeUserUpLogColumns struct {
	Id             string //
	OldGroupId     string //
	NewGroupId     string //
	UserId         string // 用户user_id
	ZtNum          string // 直推人数
	RefereeNum     string // 推荐的数量
	BuyNum         string // 个人业绩
	ZtBuyNum       string // 直推业绩
	CommunityMoney string // 小区业绩
	Title          string // 标题
	Content        string // 内容
	AddTime        string // 添加时间
}

// kcodeUserUpLogColumns holds the columns for table kcode_user_up_log.
var kcodeUserUpLogColumns = KcodeUserUpLogColumns{
	Id:             "id",
	OldGroupId:     "old_group_id",
	NewGroupId:     "new_group_id",
	UserId:         "user_id",
	ZtNum:          "zt_num",
	RefereeNum:     "referee_num",
	BuyNum:         "buy_num",
	ZtBuyNum:       "zt_buy_num",
	CommunityMoney: "community_money",
	Title:          "title",
	Content:        "content",
	AddTime:        "add_time",
}

// NewKcodeUserUpLogDao creates and returns a new DAO object for table data access.
func NewKcodeUserUpLogDao() *KcodeUserUpLogDao {
	return &KcodeUserUpLogDao{
		group:   "default",
		table:   "kcode_user_up_log",
		columns: kcodeUserUpLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KcodeUserUpLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KcodeUserUpLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KcodeUserUpLogDao) Columns() KcodeUserUpLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KcodeUserUpLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KcodeUserUpLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KcodeUserUpLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
