package common

import "github.com/chivalryq/vela-go-sdk/pkg/apis"

type ComponentConstructor func(name string) apis.Component
type TraitConstructor func() apis.Trait
type WorkflowStepConstructor func(name string) apis.WorkflowStep
type PolicyConstructor func(name string) apis.Policy

var Components = make(map[string]ComponentConstructor, 0)

func RegisterComponent(_type string, c ComponentConstructor) {
	Components[_type] = c
}
