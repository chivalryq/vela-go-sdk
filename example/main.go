package main

import (
	"fmt"
	. "vela-go-sdk/apis/component/webservice"
	. "vela-go-sdk/apis/trait/resource"
	"vela-go-sdk/app"
)

func main() {
	// Build application with components and traits

	application := app.New().
		WithComponent(
			Webservice("nginx").
				Cmd([]string{"nginx", "-g", "daemon off;"}).
				Cpu("100m").
				Memory("128Mi"),
		)

	// Create component and add traits.
	ws := Webservice("example").
		Image("nginx")

	ws.AddTrait(
		Resource().Cpu(0.1).Memory("128Mi"),
	)
	application = app.New().WithComponent(ws)

	fmt.Println(application)

}
