package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type UserBase struct {
	Id          uint        `json:"id"          description:"用户 ID"`
	Address     string      `json:"address"     description:"用户钱包地址"`
	InviterId   string      `json:"inviterId"   description:"邀请人id"`
	InviterCode string      `json:"inviterCode" description:"邀请码"`
	CreateAt    *gtime.Time `json:"createAt"    description:"创建时间"`
	UpdateAt    *gtime.Time `json:"updateAt"    description:"更新时间"`
}

type UserCreateInput struct {
	Address  string `json:"address"`
	Inviters string `json:"inviter_id"`
}

type UserCreateOutput struct {
	*UserBase `json:"data"`
	Token     string
	Invite    string
}
