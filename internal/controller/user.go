package controller

import (
	"context"
	"learn_goframe01/api/app"
	"learn_goframe01/internal/model"
	"learn_goframe01/internal/service"
)

var User = vUser{}

type vUser struct {
}

func (v *vUser) UserCreate(ctx context.Context, req *app.LoginReq) (res *app.LoginRes, err error) {
	out, err := service.User().UserLogin(ctx, model.UserCreateInput{
		Address:  req.Address,
		Inviters: req.Inviters,
	})
	if err != nil {
		return nil, err
	}
	return &app.LoginRes{Data: out.UserBase, Token: out.Token, Invite: out.Invite}, nil
}
