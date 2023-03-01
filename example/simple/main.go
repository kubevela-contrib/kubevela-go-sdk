package main

import (
	"context"
	"fmt"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	. "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/component/webservice"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/client"
)

func main() {
	application := common.New().
		Name("test-app").
		Namespace("default").
		SetComponents(
			Webservice("nginx").
				SetEnv(NewEnvList(
					NewEnv().SetName("test").SetValue("test"),
				)).
				SetImage("nginx:latest").
				SetCpu("500m"),
			//SetImagePullSecrets([]string{"test"}),
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
