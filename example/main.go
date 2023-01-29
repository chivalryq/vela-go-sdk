package main

import (
	"vela-go-sdk/pkg/apis/app"
	. "vela-go-sdk/pkg/apis/component/webservice"
	applyonce "vela-go-sdk/pkg/apis/policy/apply-once"
	initcontainer "vela-go-sdk/pkg/apis/trait/init-container"
	. "vela-go-sdk/pkg/apis/trait/resource"
	. "vela-go-sdk/pkg/apis/workflow-step/build-push-image"
	. "vela-go-sdk/pkg/apis/workflow-step/notification"
	. "vela-go-sdk/pkg/apis/workflow-step/step-group"
)

func main() {
	// Build application with components/trait/workflow/policy
	application := app.New().
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
				AddSubStep(BuildPushImage("bp").Image("my-image")).
				AddSubStep(Notification("notify")),
			Notification("single-step"),
		).
		WithPolicies(applyonce.ApplyOnce("once").Enable(true))

	// Create component and add traits.

}
