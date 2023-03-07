package main

import (
	"fmt"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	. "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/component/webservice"
)

func main() {
	checkErr := func(err error) {
		if err != nil {
			fmt.Println("validate fail, err: ", err)
			return
		} else {
			fmt.Println("validate success!")
		}
	}

	application := common.New().
		Name("test-app").
		Namespace("default").
		SetComponents(
			Webservice("nginx").
				SetEnv(NewEnvList(
					NewEnvWithDefault().SetName("test").SetValue("test"),
				)).
				SetImage("nginx:latest").
				SetCpu("500m").
				SetExposeType("ClusterIP").
				SetAddRevisionLabel(true),
		)
	err := application.Validate()
	checkErr(err)
	comp := application.GetComponentByName("nginx").(*WebserviceComponent)
	// liveness probe have some required fields, so it will return error
	comp.SetLivenessProbe(*NewHealthProbeEmpty())
	application.SetComponents(comp)
	err = application.Validate()
	checkErr(err)
}
