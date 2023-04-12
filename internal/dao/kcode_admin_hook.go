// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"learn_goframe01/internal/dao/internal"
)

// internalKcodeAdminHookDao is internal type for wrapping internal DAO implements.
type internalKcodeAdminHookDao = *internal.KcodeAdminHookDao

// kcodeAdminHookDao is the data access object for table kcode_admin_hook.
// You can define custom methods on it to extend its functionality as you wish.
type kcodeAdminHookDao struct {
	internalKcodeAdminHookDao
}

var (
	// KcodeAdminHook is globally public accessible object for table kcode_admin_hook operations.
	KcodeAdminHook = kcodeAdminHookDao{
		internal.NewKcodeAdminHookDao(),
	}
)

// Fill with you ideas below.