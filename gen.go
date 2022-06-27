package main

import (
	"KunLunQiLian/server-example/cmd/server/global"

	"gorm.io/gen"
)

func genModel() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
	})
	g.UseDB(global.Config.DB)
	g.GenerateAllTable()
	g.Execute()
}

func genQuery() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/query",
	})
	g.UseDB(global.Config.DB)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

func main() {
	genModel()
	genQuery()
}
