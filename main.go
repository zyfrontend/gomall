package main

import (
	"github.com/gin-gonic/gin"
	serviceAccount "server/service/account/account"
	serviceOrder "server/service/account/order"
	"server/service/server"
)

func main() {
	g := gin.New()

	// 登陆 注册 注销 修改密码
	g.POST("/login", server.Login)
	g.POST("/register", server.Register)
	g.POST("/logout", server.Logout)

	account := g.Group("/account")
	order := account.Group("/order")
	order.POST("/list", serviceOrder.List)
	account.POST("/profile", serviceAccount.Profile)

	g.Run()
}
