package backend

import "github.com/gogf/gf/v2/frame/g"

type RotationCreateReq struct {
	g.Meta `path:"/backend/rotation/add" tag:"轮播图" method:"post" summary:"轮播图添加接口"`
	PicUrl string `json:"picUrl"`
	Link   string `json:"link"`
	Sort   int    `json:"sort"`
}
type RotationCreateRes struct {
	RotationId int `json:"rotationId"`
}
