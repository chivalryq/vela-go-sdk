package main

import (
	"bytes"
	json2 "encoding/json"
	"fmt"

	"k8s.io/apimachinery/pkg/util/json"

	"github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	. "github.com/chivalryq/vela-go-sdk/pkg/apis/component/webservice"
	applyonce "github.com/chivalryq/vela-go-sdk/pkg/apis/policy/apply-once"
	initcontainer "github.com/chivalryq/vela-go-sdk/pkg/apis/trait/init-container"
	. "github.com/chivalryq/vela-go-sdk/pkg/apis/trait/resource"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/utils"
	bu "github.com/chivalryq/vela-go-sdk/pkg/apis/workflow-step/build-push-image"
	notify "github.com/chivalryq/vela-go-sdk/pkg/apis/workflow-step/notification"
	stepgroup "github.com/chivalryq/vela-go-sdk/pkg/apis/workflow-step/step-group"
)

func main() {
	// Build application with components/trait/workflow/policy
	application := common.New().
		Name("test-app").
		Namespace("default").
		WithComponents(
			Webservice("nginx").
				Cpu("500m").
				Memory("128Mi").
				AddTrait(
					initcontainer.InitContainer().Image("busybox").Cmd([]string{"curl", "-s", "nginx.conf"}),
					Resource().Cpu(1.0).Memory("500Mi"),
				),
		).
		WithWorkflowSteps(
			stepgroup.StepGroup("group").
				AddSubStep(bu.BuildPushImage("bp").Image("my-image").Context(bu.Context{String: utils.PtrString("test")})).
				AddSubStep(notify.Notification("notify")),
			notify.Notification("single-step"),
		).
		WithPolicies(applyonce.ApplyOnce("once").Enable(true))

	// Create component and add traits.
	appObj := application.Build()
	marshal, err := json.Marshal(appObj)
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	_ = json2.Indent(buf, marshal, "", "  ")
	fmt.Println(buf.String())

}
