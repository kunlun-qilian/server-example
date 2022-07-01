package global

import (
	"kunlun-qilian/server-example/internal/model"

	"github.com/kunlun-qilian/confmysql"
	"github.com/kunlun-qilian/confserver"
)

func init() {
	confserver.SetServiceName("example-server", "..")
	confserver.ConfP(&Config)

	// migrate 数据库 命令
	confserver.AddCommand(Config.DB.Commands()...)
}

// 需要migrate 到数据库的信息
var AutoMigrateModelList = []interface{}{
	model.User{},
	model.TExample{},
}

var Config = struct {
	DB     *confmysql.MySQL
	Server *confserver.Server

	TestEnvStr string `env:""`
}{
	Server: &confserver.Server{
		Mode: "debug",
	},

	DB: &confmysql.MySQL{
		DSN: "root:123456@tcp(127.0.0.1:33306)/example?charset=utf8mb4&parseTime=True&loc=Local",
		AutoMigrateConfig: &confmysql.AutoMigrateConfig{
			Models:    AutoMigrateModelList,
			ModelPath: "./internal/model",
			QueryPath: "./internal/query",
		},
	},
	TestEnvStr: "global.config",
}
