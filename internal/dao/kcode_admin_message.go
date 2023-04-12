// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"learn_goframe01/internal/dao/internal"
)

// internalKcodeAdminMessageDao is internal type for wrapping internal DAO implements.
type internalKcodeAdminMessageDao = *internal.KcodeAdminMessageDao

// kcodeAdminMessageDao is the data access object for table kcode_admin_message.
// You can define custom methods on it to extend its functionality as you wish.
type kcodeAdminMessageDao struct {
	internalKcodeAdminMessageDao
}

var (
	// KcodeAdminMessage is globally public accessible object for table kcode_admin_message operations.
	KcodeAdminMessage = kcodeAdminMessageDao{
		internal.NewKcodeAdminMessageDao(),
	}
)

// Fill with you ideas below.
