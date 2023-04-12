package model

// RotationBase
// ===================================================================================================================================
// 基础结构
type RotationBase struct {
	PicUrl string
	Link   string
	Sort   int
}

// RotationCreateInput
// ===================================================================================================================================
// 创建输入
type RotationCreateInput struct {
	RotationBase
}

// RotationCreateOutput
// ===================================================================================================================================
// 创建输出
type RotationCreateOutput struct {
	RotationId int `json:"rotationId"`
}
