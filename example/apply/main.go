package main

import (
	"fmt"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/component/helm"
	. "github.com/chivalryq/vela-go-sdk/pkg/apis/component/webservice"
	initcontainer "github.com/chivalryq/vela-go-sdk/pkg/apis/trait/init-container"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/trait/resource"
	bu "github.com/chivalryq/vela-go-sdk/pkg/apis/workflow-step/build-push-image"
	notify "github.com/chivalryq/vela-go-sdk/pkg/apis/workflow-step/notification"
	stepgroup "github.com/chivalryq/vela-go-sdk/pkg/apis/workflow-step/step-group"
)

func main() {
	// Build application with components/trait/workflow/policy
	application := common.New().
		Name("test-app").
		Namespace("default").
		AddComponents(
			Webservice("nginx").
				SetEnv(NewEnvs(
					NewEnv().SetName("test").SetValue("test"),
				)).
				SetImage("nginx:latest").
				SetCpu("500m").
				AddTrait(
					initcontainer.InitContainer().
						SetName("init").
						SetImage("busybox").
						SetCmd([]string{"echo", "hello", "vela"}).
						SetAppMountPath("/workdir").
						SetInitMountPath("/app/dir"),
					resource.Resource().SetMemory("256Mi"),
				),
			helm.Helm("helm").
				SetChart("some-chart").
				SetRepoType("git").
				SetGit(*helm.NewGit().SetBranch("branch")).
				SetUrl("https://github.com/xxx/xxx.git"),
		).
		AddWorkflowSteps(
			stepgroup.StepGroup("group").
				AddSubStep(bu.BuildPushImage("bp").SetImage("my-image").SetContext(bu.GitAsContext(bu.NewGit().SetBranch("111")))).
				AddSubStep(notify.Notification("notify")),
			notify.Notification("single-step"),
		)
	//WithPolicies(applyonce.ApplyOnce("once").Enable(true))

	yaml, err := application.ToYAML()
	if err != nil {
		fmt.Println("convert to yaml err", err)
	}
	fmt.Println(yaml)

	//clt, err := client.NewDefault()
	//if err != nil {
	//	fmt.Println("init client", err)
	//	return
	//}
	//
	//err = clt.Create(context.Background(), application)
	//if err != nil {
	//	fmt.Println("create fail")
	//	fmt.Println(err)
	//}

}
