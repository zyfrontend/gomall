// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"learn_goframe01/internal/dao/internal"
)

// internalKcodeUserLogDao is internal type for wrapping internal DAO implements.
type internalKcodeUserLogDao = *internal.KcodeUserLogDao

// kcodeUserLogDao is the data access object for table kcode_user_log.
// You can define custom methods on it to extend its functionality as you wish.
type kcodeUserLogDao struct {
	internalKcodeUserLogDao
}

var (
	// KcodeUserLog is globally public accessible object for table kcode_user_log operations.
	KcodeUserLog = kcodeUserLogDao{
		internal.NewKcodeUserLogDao(),
	}
)

// Fill with you ideas below.