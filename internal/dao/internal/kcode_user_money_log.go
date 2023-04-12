// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeUserMoneyLogDao is the data access object for table kcode_user_money_log.
type KcodeUserMoneyLogDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns KcodeUserMoneyLogColumns // columns contains all the column names of Table for convenient usage.
}

// KcodeUserMoneyLogColumns defines and stores column names for table kcode_user_money_log.
type KcodeUserMoneyLogColumns struct {
	LogId         string // 自增编号
	UserId        string // 用户user_id
	MoneyType     string // 币种类型 关联money_type 表
	Money         string // 变化的金额
	MoneyAction   string // 增加还是减少 默认+
	Balance       string // 余下额度
	Messages      string // 交易信息
	LangVal       string // 语言键值
	MessagesLang  string // 语言内容
	AddTime       string // 添加时间
	ParentId      string // 父ID
	Status        string // 状态 0 未成功 1成功
	PayTime       string // 支付时间
	OperationType string // 操作类型
	BonusType     string // 奖金类型，1直推奖励 2团队奖励3全网加权分红4权益分红
	EndTime       string // 结算时间
	CoinPrice     string // 当前币价
	Usdt          string // usdt
}

// kcodeUserMoneyLogColumns holds the columns for table kcode_user_money_log.
var kcodeUserMoneyLogColumns = KcodeUserMoneyLogColumns{
	LogId:         "log_id",
	UserId:        "user_id",
	MoneyType:     "money_type",
	Money:         "money",
	MoneyAction:   "money_action",
	Balance:       "balance",
	Messages:      "messages",
	LangVal:       "lang_val",
	MessagesLang:  "messages_lang",
	AddTime:       "add_time",
	ParentId:      "parent_id",
	Status:        "status",
	PayTime:       "pay_time",
	OperationType: "operation_type",
	BonusType:     "bonus_type",
	EndTime:       "end_time",
	CoinPrice:     "coin_price",
	Usdt:          "usdt",
}

// NewKcodeUserMoneyLogDao creates and returns a new DAO object for table data access.
func NewKcodeUserMoneyLogDao() *KcodeUserMoneyLogDao {
	return &KcodeUserMoneyLogDao{
		group:   "default",
		table:   "kcode_user_money_log",
		columns: kcodeUserMoneyLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KcodeUserMoneyLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KcodeUserMoneyLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KcodeUserMoneyLogDao) Columns() KcodeUserMoneyLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KcodeUserMoneyLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KcodeUserMoneyLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KcodeUserMoneyLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
