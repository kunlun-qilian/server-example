package router

import (
	"github.com/sirupsen/logrus"
	"kunlun-qilian/server-example/cmd/server/router/example"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Web(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

type AuthorizationParam struct {
	Authorization string `header:"Authorization" binding:"required" `
}

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		a := AuthorizationParam{}
		err := ctx.BindHeader(&a)
		if err != nil {
			logrus.Warn("Authorization is nil")
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

	}
}

func NewRooter(r *gin.Engine) {

	//r.LoadHTMLGlob("../web/build/index.html")
	//r.StaticFS("/static", http.Dir("../web/build/static"))
	//r.GET("/", Web)
	// API
	v1 := r.Group("/api/v1", Authorization())
	example.CarRouter(v1)
}
