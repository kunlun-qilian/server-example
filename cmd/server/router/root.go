package router

import (
    "KunLunQiLian/server-example/cmd/server/router/example"
    "net/http"

    "github.com/gin-gonic/gin"
)

func Web(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func NewRooter(r *gin.Engine) {

    r.LoadHTMLGlob("../web/build/index.html")
    r.StaticFS("/static", http.Dir("../web/build/static"))
    r.GET("/", Web)
    // API
    v1 := r.Group("/api/v1")
    example.CarRouter(v1)
}
