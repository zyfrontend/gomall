package serviceOrder

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 输入
type ListRequest struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type ListResponseItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// 输出
type ListResponse struct {
	List []*ListResponseItem `json:"list"`
}

func List(c *gin.Context) {
	req := ListRequest{}
	c.ShouldBindJSON(&req)
	fmt.Println(req)
	// 增删改查之后
	resp := ListResponse{}
	for i := 0; i < 5; i++ {
		resp.List = append(resp.List, &ListResponseItem{
			Id:   fmt.Sprintf("%d", i),
			Name: "name",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    &resp,
		"code":    200,
	})
}
