package main

import (
	"fmt"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	. "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/component/webservice"
	apply_once "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/policy/apply-once"
	initcontainer "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/trait/init-container"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/trait/resource"
	bu "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/workflow-step/build-push-image"
	notify "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/workflow-step/notification"
	stepgroup "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/workflow-step/step-group"
)

func main() {
	// Build application with components/trait/workflow/policy
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
				SetTraits(
					initcontainer.InitContainer().
						SetName("init").
						SetImage("busybox").
						SetCmd([]string{"echo", "hello", "vela"}).
						SetAppMountPath("/workdir").
						SetInitMountPath("/app/dir"),
					resource.Resource().SetMemory("256Mi"),
				),
		).
		SetWorkflowSteps(
			stepgroup.StepGroup("group").
				AddSubStep(bu.BuildPushImage("bp").SetImage("my-image").SetContext(bu.GitAsContext(bu.NewGitWithDefault().SetBranch("111")))).
				AddSubStep(notify.Notification("notify")),
			notify.Notification("single-step"),
		).
		SetPolicies(apply_once.ApplyOnce("apply-po"))

	yaml, err := application.ToYAML()
	if err != nil {
		fmt.Println("convert to yaml err", err)
	}
	fmt.Println(yaml)

}
