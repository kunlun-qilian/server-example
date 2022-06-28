package example

import (
	"KunLunQiLian/server-example/cmd/server/global"
	"KunLunQiLian/server-example/internal/model"
	"KunLunQiLian/server-example/internal/query"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CarRouter(r *gin.RouterGroup) {
	r.GET("/car", ListCar)
	r.POST("/car", CreateCar)
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
