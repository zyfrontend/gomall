// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeAdminHookPlugin is the golang structure of table kcode_admin_hook_plugin for DAO operations like Where/Data.
type KcodeAdminHookPlugin struct {
	g.Meta     `orm:"table:kcode_admin_hook_plugin, do:true"`
	Id         interface{} //
	Hook       interface{} // 钩子id
	Plugin     interface{} // 插件标识
	CreateTime interface{} // 添加时间
	UpdateTime interface{} // 更新时间
	Sort       interface{} // 排序
	Status     interface{} // 状态
}