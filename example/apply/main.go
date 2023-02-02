package main

import (
	"bytes"
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/trait/resource"
	"github.com/chivalryq/vela-go-sdk/pkg/client"
	"k8s.io/apimachinery/pkg/util/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	. "github.com/chivalryq/vela-go-sdk/pkg/apis/component/webservice"
	initcontainer "github.com/chivalryq/vela-go-sdk/pkg/apis/trait/init-container"
)

func main() {
	// Build application with components/trait/workflow/policy
	application :=
		common.New().
			Name("test-app").
			Namespace("default").
			WithComponents(
				Webservice("nginx").
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
			)
	//WithWorkflowSteps(
	//	stepgroup.StepGroup("group").
	//		AddSubStep(bu.BuildPushImage("bp").Image("my-image").Context(bu.Context{String: utils.PtrString("test")})).
	//		AddSubStep(notify.Notification("notify")),
	//	notify.Notification("single-step"),
	//).
	//WithPolicies(applyonce.ApplyOnce("once").Enable(true))

	appObj := application.Build()
	marshal, err := json.Marshal(appObj)
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	_ = json2.Indent(buf, marshal, "", "  ")
	fmt.Println(buf.String())

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

}
