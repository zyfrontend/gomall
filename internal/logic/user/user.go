package user

import (
	"context"
	"errors"
	"fmt"
	"learn_goframe01/internal/dao"
	"learn_goframe01/internal/model"
	"learn_goframe01/internal/model/entity"
	"learn_goframe01/internal/service"
	"learn_goframe01/internal/utils"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (u *sUser) UserLogin(ctx context.Context, in model.UserCreateInput) (out model.UserCreateOutput, err error) {
	var userInfo *model.UserBase
	userId := 0
	pid := 0
	// 1. 验证用户是否存在
	user, err := GetUserByAccount(ctx, in.Address)
	if err != nil {
		return out, err
	}
	// 用户为空
	if user == nil {
		// 邀请人
		if in.Inviters != "" {
			pid64 := utils.NewCode().CodeToId(in.Inviters)
			pid = int(pid64) // 获得推荐人id
		}
		p, _ := dao.User.Ctx(ctx).Where("id=", pid).Count()
		// 推荐人也不存在
		if p <= 0 {
			fmt.Println(p)
			//return out, err
		}

		// 开始注册用户
		uid, err := u.Register(ctx, in.Address)
		if err != nil {
			return model.UserCreateOutput{}, err
		}
		userId = uid
	} else {
		userId = int(user.Id)
	}
	userInfo = (*model.UserBase)(user)
	// 生成token
	token, _ := utils.GenerateToken(uint(userId), user.Address)
	inviteCode := utils.NewCode().IdToCode(uint64(userId))
	return model.UserCreateOutput{
		UserBase: userInfo,
		Token:    token,
		Invite:   inviteCode,
	}, nil
}

func GetUserByAccount(ctx context.Context, address string) (*entity.User, error) {
	var info *entity.User
	err := dao.User.Ctx(ctx).Where("address = ?", address).Scan(&info)
	return info, err
}

func InsertUserData(ctx context.Context, address string) (int, error) {
	user := &entity.User{Address: address}
	res, err := dao.User.Ctx(ctx).Data(user).Save()
	if err != nil {
		return 0, err
	}
	id64, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	if id64 < 1 {
		return 0, errors.New("添加账户失败")
	}
	id := int(id64)
	return id, nil
}

func (u *sUser) Register(ctx context.Context, account string) (int, error) {
	uid, err := InsertUserData(ctx, account)
	if err != nil {
		return 0, err
	}
	return uid, nil
}
