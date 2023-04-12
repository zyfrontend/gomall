package rotation

import (
	"context"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"learn_goframe01/internal/dao"
	"learn_goframe01/internal/model"
	"learn_goframe01/internal/service"
)

type sRotation struct {
}

func init() {
	service.RegisterRotation(New())
}
func New() *sRotation {
	return &sRotation{}
}

func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	Id, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: int(Id)}, err
}
