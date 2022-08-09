package main

import (
	"context"
	"fmt"
	"kunlun-qilian/server-example/cmd/client/client_server"
	"kunlun-qilian/server-example/cmd/client/global"
	"time"

	"github.com/kunlun-qilian/confclient"
)

func main() {

	cli, err := client_server.NewClientWithResponses(global.Config.Cli.ApiServer(), client_server.WithRequestEditorFn(confclient.WithTrace))
	if err != nil {
		panic(err)
	}

	for {
		resp, err := cli.ListCarWithResponse(context.Background())
		if err != nil {
			fmt.Println("server err:", err.Error())
		} else {
			fmt.Println("orgin resp: ", resp)
			if resp.JSON200 != nil {
				fmt.Println("resp: ", resp.JSON200)
			}
		}
		time.Sleep(time.Second * 3)
	}

}
