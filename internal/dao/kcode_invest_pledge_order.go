// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"learn_goframe01/internal/dao/internal"
)

// internalKcodeInvestPledgeOrderDao is internal type for wrapping internal DAO implements.
type internalKcodeInvestPledgeOrderDao = *internal.KcodeInvestPledgeOrderDao

// kcodeInvestPledgeOrderDao is the data access object for table kcode_invest_pledge_order.
// You can define custom methods on it to extend its functionality as you wish.
type kcodeInvestPledgeOrderDao struct {
	internalKcodeInvestPledgeOrderDao
}

var (
	// KcodeInvestPledgeOrder is globally public accessible object for table kcode_invest_pledge_order operations.
	KcodeInvestPledgeOrder = kcodeInvestPledgeOrderDao{
		internal.NewKcodeInvestPledgeOrderDao(),
	}
)

// Fill with you ideas below.