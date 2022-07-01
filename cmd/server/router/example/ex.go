package example

import (
	"fmt"
	"kunlun-qilian/server-example/cmd/server/global"
	"kunlun-qilian/server-example/internal/model"
	"kunlun-qilian/server-example/internal/query"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CarRouter(r *gin.RouterGroup) {
	r.GET("/car", SelfCheckCar(), ListCar)
	r.POST("/car", CreateCar)
}

func SelfCheckCar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("前置校验")
		ctx.Next()
		fmt.Println("后置处理")

	}
}

type Car struct {
	Name    string `json:"name"`
	CarType int    `json:"carType"`
}

// @BasePath /api/v1
// PingExample godoc
// @Summary ListCar
// @Schemes
// @Description List car
// @Tags ex
// @Accept json
// @Produce json
// @Success 200 {object}  []model.TExample
// @Router /car [get]
// @ID ListCar
func ListCar(ctx *gin.Context) {
	fmt.Println("业务处理")
	q := query.Use(global.Config.DB.DB()).TExample
	carList, err := q.WithContext(ctx).Find()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, carList)
}

type CreateCarRequestBody struct {
	Name    string `json:"name" binding:"required"`
	CarType int32  `json:"carType" binding:"required"`
}

type ErrorResp struct {
	Msg string `json:"msg"`
}

// @BasePath /api/v1
// PingExample godoc
// @Summary CreateCar
// @Schemes
// @Description Create Car
// @Tags ex
// @Accept json
// @Produce json
// @Param ReqeustBody body CreateCarRequestBody true "Create Car"
// @Success 200 {object} model.TExample 成功
// @Success 400 {object} ErrorResp 失败
// @Success 500 {object} ErrorResp 失败
// @Router  /car [post]
// @ID CreateCar
func CreateCar(ctx *gin.Context) {
	body := CreateCarRequestBody{}
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResp{Msg: err.Error()})
		return
	}

	m := model.TExample{}
	m.CarType = body.CarType
	m.Name = body.Name

	q := query.Use(global.Config.DB.DB()).TExample
	err = q.WithContext(ctx).Create(&m)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, m)
}
