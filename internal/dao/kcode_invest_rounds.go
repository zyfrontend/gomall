// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"learn_goframe01/internal/dao/internal"
)

// internalKcodeInvestRoundsDao is internal type for wrapping internal DAO implements.
type internalKcodeInvestRoundsDao = *internal.KcodeInvestRoundsDao

// kcodeInvestRoundsDao is the data access object for table kcode_invest_rounds.
// You can define custom methods on it to extend its functionality as you wish.
type kcodeInvestRoundsDao struct {
	internalKcodeInvestRoundsDao
}

var (
	// KcodeInvestRounds is globally public accessible object for table kcode_invest_rounds operations.
	KcodeInvestRounds = kcodeInvestRoundsDao{
		internal.NewKcodeInvestRoundsDao(),
	}
)

// Fill with you ideas below.