package example

import (
	"fmt"
	"github.com/kunlun-qilian/confserver"
	"kunlun-qilian/server-example/cmd/server/global"
	"kunlun-qilian/server-example/internal/model"
	"net/http"

	"github.com/go-courier/sqlx/v2/builder"

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
// @Success 200 {object}  []model.Example
// @Router /car [get]
// @ID ListCar
func ListCar(ctx *gin.Context) {
	fmt.Println("业务处理")
	ex := model.Example{}

	cs := model.NewCondRules()
	where := builder.And(
		cs.When(true, ex.FieldName().Eq("bbb")),
	)
	where.Or(cs.When(true, ex.FieldCarType().Eq(123)))

	exList, err := ex.List(global.Config.DB, where)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, exList)
}

type CreateCarRequestBody struct {
	Name    string `json:"name"`
	CarType int    `json:"carType"`
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
// @Success 200 {object} model.Example 成功
// @Success 400 {object} ErrorResp 失败
// @Success 500 {object} ErrorResp 失败
// @Router  /car [post]
// @ID CreateCar
func CreateCar(ctx *gin.Context) {
	body := CreateCarRequestBody{}
	err := confserver.Bind(ctx, &body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResp{Msg: err.Error()})
		return
	}
	//
	m := model.Example{}
	m.CarType = body.CarType
	m.Name = body.Name

	err = m.Create(global.Config.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)

}
