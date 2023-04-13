package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"learn_goframe01/internal/model"
)

type LoginReq struct {
	g.Meta   `path:"/login" tags:"用户模块" method:"post" summary:"用户登陆操作"`
	Address  string `json:"address" v:"required#钱包地址不能为空" dc:"用户登陆钱包地址"`
	Inviters string `json:"inviters" dc:"用户推荐人"`
}

type LoginRes struct {
	Data   *model.UserBase `json:"data"`
	Token  string          `json:"token"`
	Invite string          `json:"inviteCode"`
}
