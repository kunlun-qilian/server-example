package global

import (
	"github.com/kunlun-qilian/confmysql"
	"github.com/kunlun-qilian/confserver"
	"kunlun-qilian/server-example/internal/model"
)

func init() {
	confserver.SetServiceName("example-server", "..")
	confserver.ConfP(&Config)

	confserver.AddCommand(Config.DB.Commands()...)
}

var Config = struct {
	DB     *confmysql.MySQL
	Server *confserver.Server

	TestEnvStr string `env:""`
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
