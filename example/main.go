package main

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"github.com/chivalryq/vela-go-sdk/pkg/apis/common"
	"k8s.io/apimachinery/pkg/util/json"
	. "vela-go-sdk/pkg/apis/component/webservice"
	applyonce "vela-go-sdk/pkg/apis/policy/apply-once"
	initcontainer "vela-go-sdk/pkg/apis/trait/init-container"
	. "vela-go-sdk/pkg/apis/trait/resource"
	"vela-go-sdk/pkg/apis/utils"
	. "vela-go-sdk/pkg/apis/workflow-step/build-push-image"
	. "vela-go-sdk/pkg/apis/workflow-step/notification"
	. "vela-go-sdk/pkg/apis/workflow-step/step-group"
)

func main() {
	// Build application with components/trait/workflow/policy
	application := common.New().
		Name("test-app").
		Namespace("default").
		WithComponents(
			Webservice("nginx").
				Cmd([]string{"nginx", "-g", "daemon off;"}).
				Cpu("100m").
				Memory("128Mi").
				AddTrait(
					initcontainer.InitContainer().Image("busybox").Cmd([]string{"curl", "-s", "nginx.conf"}),
					Resource().Cpu(1.0).Memory("500Mi"),
				),
		).
		WithWorkflowSteps(
			StepGroup("group").
				AddSubStep(BuildPushImage("bp").Image("my-image").Context(Context{String: utils.PtrString("test")})).
				AddSubStep(Notification("notify")),
			Notification("single-step"),
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
