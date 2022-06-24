package main

import (
    "fmt"
    "log"

    apiclient "KunLunQiLian/server-example/cmd/client/client/ex"
    httptransport "github.com/go-openapi/runtime/client"
    "github.com/go-openapi/strfmt"
)

func main() {

    transport := httptransport.New("127.0.0.1", "/api/v1", []string{"http"})
    cli := apiclient.New(transport, strfmt.Default)
    resp, err := cli.ListCar(apiclient.NewListCarParams())
    if err != nil {
        log.Fatal(err)
    }
    for _, m := range resp.Payload {
        fmt.Println(m)
    }
}
