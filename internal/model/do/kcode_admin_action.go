// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeAdminAction is the golang structure of table kcode_admin_action for DAO operations like Where/Data.
type KcodeAdminAction struct {
	g.Meta     `orm:"table:kcode_admin_action, do:true"`
	Id         interface{} //
	Module     interface{} // 所属模块名
	Name       interface{} // 行为唯一标识
	Title      interface{} // 行为标题
	Remark     interface{} // 行为描述
	Rule       interface{} // 行为规则
	Log        interface{} // 日志规则
	Status     interface{} // 状态
	CreateTime interface{} // 创建时间
	UpdateTime interface{} // 更新时间
}
