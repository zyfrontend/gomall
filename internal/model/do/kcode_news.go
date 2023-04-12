// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeNews is the golang structure of table kcode_news for DAO operations like Where/Data.
type KcodeNews struct {
	g.Meta     `orm:"table:kcode_news, do:true"`
	Id         interface{} //
	TypeId     interface{} // 类型ID
	Title      interface{} // 标题
	TitleEn    interface{} //
	TitleTw    interface{} //
	Keysword   interface{} // 关键词
	Introduce  interface{} // 介绍
	ContentEn  interface{} // 英文介绍
	ContentTw  interface{} // 繁体介绍
	Content    interface{} // 新闻内容
	Sort       interface{} // 排序
	CreateTime interface{} // 添加日期
	UpdateTime interface{} // 修改日期
	Status     interface{} // 1开启 0关闭
	Addperson  interface{} // 添加者
	Flag       interface{} // 文章属性
	Click      interface{} // 点击数量
	Image      interface{} // 公共的图片地址
	PcImage    interface{} // pc缩略图
	WapImage   interface{} // 手机端缩略图
	JumpUrl    interface{} // 跳转的地址,如果不填，讲会跳转到平台自己的地址
	IsIndex    interface{} // 是否显示在首页，0=不显示，1=显示
}