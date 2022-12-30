package main

import (
	"fmt"
	"vela-go-sdk/app"
	. "vela-go-sdk/component"
	. "vela-go-sdk/trait"
)

func main() {
	// Build application with components and traits
	application := app.New().
		WithComponent(
			Webservice("backend").
				Image("nginx:latest").
				Traits(Resource().CPU(1.0)),
		).
		WithComponent(
			Webservice("frontend").
				Traits(
					Resource().Requests(Requests{CPU: 1.0, Memory: "1Gi"}),
				).
				DependsOn([]string{"backend"}),
		)

	// Create component and add traits.
	ws := Webservice("example").
		Image("nginx")

	ws.Traits(
		// Resource().CPU(1.0),
		Resource().Requests(Requests{CPU: 1.0, Memory: "1Gi"}),
	)
	application = app.New().WithComponent(ws)

	fmt.Println(application)
}
