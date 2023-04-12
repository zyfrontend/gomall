package controller

import (
	"context"
	"learn_goframe01/api/backend"
	"learn_goframe01/internal/model"
	"learn_goframe01/internal/service"
)

var Rotation = cRotation{}

type cRotation struct {
}

func (c *cRotation) Create(ctx context.Context, req *backend.RotationCreateReq) (res *backend.RotationCreateRes, err error) {
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		RotationBase: model.RotationBase{
			PicUrl: req.PicUrl,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RotationCreateRes{RotationId: out.RotationId}, nil
}
