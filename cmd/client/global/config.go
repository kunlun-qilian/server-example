package global

import (
	"github.com/kunlun-qilian/confclient"
	"github.com/kunlun-qilian/confserver"
)

func init() {
	confserver.ConfP(&Config)
}

var Config = struct {
	Cli     *confclient.RestClient
	TestEnv string `env:""`
}{
	Cli: &confclient.RestClient{
		Host:     "127.0.0.1",
		Port:     80,
		Protocol: "http",
	},
	TestEnv: "123",
}
