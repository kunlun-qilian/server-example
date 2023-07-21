package global

import (
	"kunlun-qilian/server-example/internal/model"

	"github.com/kunlun-qilian/conflogger"
	"github.com/kunlun-qilian/confpostgres"
	"github.com/kunlun-qilian/confserver"
)

func init() {
	confserver.SetServiceName("example-server", "..")
	confserver.ConfP(&Config)

	confserver.AddCommand(Config.DB.Commands()...)
}

var Config = struct {
	DB         *confpostgres.Postgres
	Server     *confserver.Server
	Logger     *conflogger.Log
	TestEnvStr string `env:""`
}{
	Logger: &conflogger.Log{
		Level: "DEBUG",
	},
	Server: &confserver.Server{
		Mode: "debug",
	},

	DB: &confpostgres.Postgres{
		Host:     "127.0.0.1",
		User:     "postgres",
		Port:     55433,
		DBName:   "example",
		Password: "abc123",
		Database: model.DB,
	},
	TestEnvStr: "global.config",
}
