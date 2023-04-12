// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"learn_goframe01/internal/dao/internal"
)

// internalKcodeUserGroupDao is internal type for wrapping internal DAO implements.
type internalKcodeUserGroupDao = *internal.KcodeUserGroupDao

// kcodeUserGroupDao is the data access object for table kcode_user_group.
// You can define custom methods on it to extend its functionality as you wish.
type kcodeUserGroupDao struct {
	internalKcodeUserGroupDao
}

var (
	// KcodeUserGroup is globally public accessible object for table kcode_user_group operations.
	KcodeUserGroup = kcodeUserGroupDao{
		internal.NewKcodeUserGroupDao(),
	}
)

// Fill with you ideas below.