package global

import (
    "github.com/KunLunQiLian/confserver"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func init() {
    confserver.SetServiceName("example-server", ".")
    confserver.ConfP(&Config)

    db, err := gorm.Open(mysql.Open(Config.DSN), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    Config.DB = db
}

var Config = struct {
    DB     *gorm.DB `env:"-"`
    DSN    string   `env:""`
    Server *confserver.Server
}{
    Server: &confserver.Server{
        Mode: "debug",
    },

    DSN: "root:123456@tcp(127.0.0.1:33306)/example?charset=utf8mb4&parseTime=True&loc=Local",
}
