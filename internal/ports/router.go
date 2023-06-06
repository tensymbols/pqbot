package ports

import (
	"github.com/gin-gonic/gin"
	"vkbot/internal/app"
)

func Router(r gin.IRouter, a app.App) {
	//	r.POST("/", returnInitString("abe03ede", 220458159))
	r.GET("/", hello(a))
	r.GET("/friends/:user_id", getFriends(a))
	r.POST("/", eventHandler(a))
}
