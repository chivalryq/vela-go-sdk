package main

import (
	"context"
	"fmt"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	. "github.com/chivalryq/vela-go-sdk/pkg/apis/component/webservice"
	"github.com/chivalryq/vela-go-sdk/pkg/client"
)

func main() {
	application := common.New().
		Name("test-app").
		Namespace("default").
		AddComponents(
			Webservice("nginx").
				SetEnv(NewEnvs(
					NewEnv().SetName("test").SetValue("test"),
				)).
				SetImage("nginx:latest").
				SetCpu("500m"),
			//SetHostAliases(NewHostAliasess(
			//	NewHostAliases().SetHostnames([]string{"test"}),
			//)),
		)

	clt, err := client.NewDefault()
	if err != nil {
		fmt.Println("init client", err)
		return
	}

	err = clt.Create(context.Background(), application)
	if err != nil {
		fmt.Println("create fail")
		fmt.Println(err)
	}
	fmt.Println("create success")
}
