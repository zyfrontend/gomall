// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeUser is the golang structure of table kcode_user for DAO operations like Where/Data.
type KcodeUser struct {
	g.Meta          `orm:"table:kcode_user, do:true"`
	UserId          interface{} // 用户user_id
	JobId           interface{} // 节点级别
	GroupId         interface{} // 会员级别
	InviterCode     interface{} // 邀请码
	UserType        interface{} // 用户类型
	Address         interface{} // 用户钱包地址
	Username        interface{} // 用户名
	Nickname        interface{} // 昵称
	Salt            interface{} // 密码加密字符
	Password        interface{} // 密码
	Payword         interface{} // 支付密码
	Mobile          interface{} // 手机号码
	MobileBind      interface{} // 是否绑定手机号码
	Email           interface{} // 用户邮箱地址
	EmailBind       interface{} // 是否绑定邮箱地址
	Avatar          interface{} // 头像
	RegIp           interface{} // 注册ip
	CreateTime      interface{} // 注册时间
	UpdateTime      interface{} // 最后更新时间
	LastLoginIp     interface{} // 最后登录IP
	LastLoginTime   interface{} // 最后登录时间
	OpenTime        interface{} // 会员开通时间
	OpenStatus      interface{} // 会员开通状态 0 未开通 1 已经开通
	InviterUsername interface{} // 邀请人用户名
	Inviters        interface{} // 邀请人集合
	InviterId       interface{} // 邀请人id
	IsAtm           interface{} // 当前帐号是否可以提现 1 可以提现 0 不可以提现
	IsLogin         interface{} // 1 可以登录 0冻结禁止登录
	LoginNum        interface{} // 登录次数
	Age             interface{} // 年龄
	Sex             interface{} // 性别 0 保密 1 男 2 女
	AvatarImg       interface{} // 图片URL地址
	IsFrozen        interface{} // 0 默认 1 冻结
	IsDel           interface{} // 注销 0 是否删除  1 删除注销
	IsTeam          interface{} // 是否绑定微信
	Floor           interface{} // 层数
	ZtNum           interface{} // 直推人数
	TeamNum         interface{} // 团队人数
	BuyNum          interface{} // 累计投资USDT数量
	TeamBuyNum      interface{} // 团队累计投资USDT数量
	BuyNumUsdt      interface{} // 个人业绩（U）
	TeamBuyNumUsdt  interface{} // 团队业绩（U）
	IfEffer         interface{} // 是否是有效用户
	IsAd            interface{} // 1今日看过广告
	TokenBuyNum     interface{} // 累计投资换算hvt数量
	Suanli          interface{} // 用户算力
}
