package ports

import (
	"github.com/gin-gonic/gin"
	"vkbot/internal/app"
)

func Router(r gin.IRouter, app app.App) {
	//	r.POST("/", returnInitString("abe03ede", 220458159))
	r.GET("/", hello(app))
	r.POST("/", eventHandler(app))
}
