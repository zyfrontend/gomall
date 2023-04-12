// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeAdminMenu is the golang structure of table kcode_admin_menu for DAO operations like Where/Data.
type KcodeAdminMenu struct {
	g.Meta     `orm:"table:kcode_admin_menu, do:true"`
	Id         interface{} //
	Pid        interface{} // 上级菜单id
	Module     interface{} // 模块名称
	Title      interface{} // 菜单标题
	Icon       interface{} // 菜单图标
	UrlType    interface{} // 链接类型（link：外链，module：模块）
	UrlValue   interface{} // 链接地址
	UrlTarget  interface{} // 链接打开方式：_blank,_self
	OnlineHide interface{} // 网站上线后是否隐藏
	CreateTime interface{} // 创建时间
	UpdateTime interface{} // 更新时间
	Sort       interface{} // 排序
	SystemMenu interface{} // 是否为系统菜单，系统菜单不可删除
	Status     interface{} // 状态
	Params     interface{} // 参数
}