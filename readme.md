# Start

# It is [Gin](https://github.com/gin-gonic/gin)!

[confserver](https://github.com/kunlun-qilian/confserver)  provide a new function named Bind() to validate and bind values in http request,this function is more convenient than gin's function

```
func CarRouter(r *gin.RouterGroup) {
	r.GET("/car/:id", ListCar)
	r.POST("/car", CreateCar)
}

type Car struct {
	Name    string `json:"name"`
	CarType int    `json:"carType"`
}

type ListCarParam struct {
	// in path
	ID string `in:"path" name:"id"`         <-----  path 
	// in query
	Name string `in:"query" name:"name"`    <-------  query
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
	err := confserver.Bind(ctx, &param)        <-------  Validate  params both in query and path
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

```

## Inject Env Value
Read the env value and overwrite the value
### Env Tag
cmd/server/global/config.go
```
package global

import (
	"kunlun-qilian/server-example/internal/model"

	"github.com/kunlun-qilian/confmysql"
	"github.com/kunlun-qilian/confserver"
)

func init() {
	confserver.SetServiceName("example-server", "..")
	confserver.ConfP(&Config)

	confserver.AddCommand(Config.DB.Commands()...)
}

var Config = struct {
	DB     *confmysql.MySQL
	Server *confserver.Server

	TestEnvStr string `env:""`     <------   mark the env value
}{
	Server: &confserver.Server{
		Mode: "debug",
		LogOption: confserver.LogOption{
			LogLevel: "debug",
		},
	},

	DB: &confmysql.MySQL{
		Host:     "127.0.0.1",
		User:     "root",
		Port:     33306,
		DBName:   "example",
		Password: "123456",
		Database: model.DB,
	},
	TestEnvStr: "global.config",
}

```

### Auto generate default.yml file,when started
cmd/server/config/default.yml
```
EXAMPLE_SERVER__DB_DBName: example
EXAMPLE_SERVER__DB_Extra: autocommit=true&charset=utf8mb4&interpolateParams=true&loc=Local&parseTime=true
EXAMPLE_SERVER__DB_Host: 127.0.0.1
EXAMPLE_SERVER__DB_Password: "123456"
EXAMPLE_SERVER__DB_PoolSize: "10"
EXAMPLE_SERVER__DB_Port: "33306"
EXAMPLE_SERVER__DB_User: root
EXAMPLE_SERVER__Server_LogFormatter: json
EXAMPLE_SERVER__Server_LogLevel: debug
EXAMPLE_SERVER__Server_Mode: debug
EXAMPLE_SERVER__TestEnvStr: global.config
GOENV: DEV

```

### local env file
cmd/server/config/local.yml

this file will overwrite the value in default.yml, it just could be used at local machine!!!
this file should be ignored in your project, put it into your .gitignore !!!!!!
```
EXAMPLE_SERVER__TestEnvStr: from.local.yml
```


## ORM
[Sqlx](https://github.com/go-courier/sqlx)
use [klctl](https://github.com/kunlun-qilian/klctl) generate model
inerternal/model/example.go
```
//go:generate klctl gen model2 Example --database DB
// @def primary ID
// @def unique_index I_name_id Name
// @def index I_ff_user UserID FF      <--------  union index, this is more convenient than gorm
type Example struct {
	PrimaryID
	Name    string `db:"F_name,default='',size=100" json:"name"` // Name
	CarType int    `db:"F_car_type,default=0" json:"CarType"`
	FF      string `db:"f_FF,size=100"`
	RefUser
}

```

## Openapi and sdk
use gin swagger [swag](https://github.com/swaggo/swag) [swagger](https://github.com/go-swagger/go-swagger)
use [klctl](https://github.com/kunlun-qilian/klctl) convert swagger2.0 to openapi3.0

### generate openapi3.0
```
make openapi
```

### generate sdk

start the server and use  "make client"


see more in makefile

```
PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
NAME = $(shell basename $(PKG))
VERSION = v$(shell cat .version)
COMMIT_SHA ?= $(shell git rev-parse --short HEAD)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CGO_ENABLED ?= 0

GOBUILD=CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"

WORKSPACE ?= name

build:
	cd ./cmd/$(WORKSPACE) && $(GOBUILD)

# swagger
swagger:
	swag init --pd -d ./cmd/server -o ./cmd/server/docs

# openapi 3.0
openapi: swagger
	cd ./cmd/server && klctl openapi  -f ./docs/swagger.json

# client
client:
	klctl gen client server  --output ./cmd/client --spec-url http://127.0.0.1/example-server


gen-web:
	npx create-react-app web --template typescript


gen-web-client:
	restful-react import --file ./cmd/server/docs/swagger.json  --output ./cmd/web/src/client-bff.ts
```


# Tools
1. [swag](https://github.com/swaggo/swag)
2. [swagger](https://github.com/go-swagger/go-swagger)
3. [klctl](https://github.com/kunlun-qilian/klctl)
4. [confserver](https://github.com/kunlun-qilian/confserver)
5. [restful-react](https://github.com/contiamo/restful-react)