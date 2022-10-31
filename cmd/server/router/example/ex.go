package example

import (
	"kunlun-qilian/server-example/cmd/server/global"
	"kunlun-qilian/server-example/internal/model"
	"net/http"

	"github.com/kunlun-qilian/confserver"

	"github.com/go-courier/sqlx/v2/builder"

	"github.com/gin-gonic/gin"
)

func CarRouter(r *gin.RouterGroup) {
	r.GET("/car/:id", ListCar)
	r.POST("/car/:userID", CreateCar)
}

type Car struct {
	Name    string `json:"name"`
	CarType int    `json:"carType"`
}

type ListCarParam struct {
	// in path
	ID string `in:"path" name:"id"`
	// in query
	Name string `in:"query" name:"name"`
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

	param := ListCarParam{}
	err := confserver.Bind(ctx, &param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

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

type CreateCarReqeust struct {
	FF                   string `in:"query" name:"ff"`
	UserID               string `in:"path" name:"userID"`
	CanBeEmpty           string `in:"query" name:"canBeEmpty,omitempty"`
	CreateCarRequestBody `in:"body"`
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
// @Param ReqeustBody body CreateCarReqeust true "Create Car"
// @Success 200 {object} model.Example OK
// @Success 400 {object} ErrorResp Error
// @Success 500 {object} ErrorResp Error
// @Router  /car [post]
// @ID CreateCar
func CreateCar(ctx *gin.Context) {
	req := CreateCarReqeust{}
	err := confserver.Bind(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorResp{Msg: err.Error()})
		return
	}

	//
	m := model.Example{}
	m.FF = req.FF
	m.CarType = req.CarType
	m.Name = req.Name
	m.UserID = req.UserID
	m.SetNowForCreate()

	err = m.Create(global.Config.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)

}
