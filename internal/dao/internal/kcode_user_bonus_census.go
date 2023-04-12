// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeUserBonusCensusDao is the data access object for table kcode_user_bonus_census.
type KcodeUserBonusCensusDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns KcodeUserBonusCensusColumns // columns contains all the column names of Table for convenient usage.
}

// KcodeUserBonusCensusColumns defines and stores column names for table kcode_user_bonus_census.
type KcodeUserBonusCensusColumns struct {
	Id           string // 主键
	UserId       string // 用户id
	MoneyType    string // 币种类型
	ZhituiBonus  string // 直推奖励
	TeamBonus    string // 团队奖励
	WholeBonus   string // 全网加权分红
	EquityBonus  string // 权益分红
	ReleaseBonus string // 释放收益
	PjBonus      string // 平级奖
}

// kcodeUserBonusCensusColumns holds the columns for table kcode_user_bonus_census.
var kcodeUserBonusCensusColumns = KcodeUserBonusCensusColumns{
	Id:           "id",
	UserId:       "user_id",
	MoneyType:    "money_type",
	ZhituiBonus:  "zhitui_bonus",
	TeamBonus:    "team_bonus",
	WholeBonus:   "whole_bonus",
	EquityBonus:  "equity_bonus",
	ReleaseBonus: "release_bonus",
	PjBonus:      "pj_bonus",
}

// NewKcodeUserBonusCensusDao creates and returns a new DAO object for table data access.
func NewKcodeUserBonusCensusDao() *KcodeUserBonusCensusDao {
	return &KcodeUserBonusCensusDao{
		group:   "default",
		table:   "kcode_user_bonus_census",
		columns: kcodeUserBonusCensusColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KcodeUserBonusCensusDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KcodeUserBonusCensusDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KcodeUserBonusCensusDao) Columns() KcodeUserBonusCensusColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KcodeUserBonusCensusDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KcodeUserBonusCensusDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KcodeUserBonusCensusDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}