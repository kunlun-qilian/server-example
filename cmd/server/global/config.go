package global

import (
	"KunLunQiLian/server-example/internal/model"

	"github.com/KunLunQiLian/confserver"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
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

	// migrate 数据库
	confserver.AddCommand(
		&cobra.Command{
			Use: "migrate",
			Run: func(cmd *cobra.Command, args []string) {
				if err := Config.DB.AutoMigrate(AutoMigrateModelList...); err != nil {
					panic(err)
				}
			},
		},
	)

	// 根据数据库表 生成模型
	confserver.AddCommand(
		&cobra.Command{
			Use: "gen-model",
			Run: func(cmd *cobra.Command, args []string) {
				g := gen.NewGenerator(gen.Config{
					OutPath: "./internal/model",
				})
				g.UseDB(Config.DB)
				g.GenerateAllTable()
				g.Execute()
			},
		},
	)

	// 根据模型生成query
	confserver.AddCommand(
		&cobra.Command{
			Use: "gen-query",
			Run: func(cmd *cobra.Command, args []string) {
				g := gen.NewGenerator(gen.Config{
					OutPath: "./internal/query",
				})
				g.UseDB(Config.DB)
				g.ApplyBasic(AutoMigrateModelList...)
				g.Execute()
			},
		},
	)

}

// 需要migrate 到数据库的信息
var AutoMigrateModelList = []interface{}{
	model.User{},
	model.TExample{},
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
