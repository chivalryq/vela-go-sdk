package trait

import (
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)

const resourceType = "resource"

// Requests Specify the properties in requests
type Requests struct {
	CPU    float64 `json:"cpu"`
	Memory string  `json:"memory"`
}

// Limits Specify the properties in limits
type Limits struct {
	CPU    float64 `json:"cpu"`
	Memory string  `json:"memory"`
}

type ResourceSpec struct {
	CPU      float64  `json:"cpu"`
	Memory   string   `json:"memory"`
	Requests Requests `json:"requests"`
	Limits   Limits   `json:"limits"`
}

type ResourceBuilder struct {
	properties ResourceSpec
}

func (b *ResourceBuilder) Build() common.ApplicationTrait {
	return common.ApplicationTrait{
		Type:       resourceType,
		Properties: util.Object2RawExtension(b.properties),
	}
}

type ResourceOps func(builder *ResourceBuilder) *ResourceBuilder

func Resource() *ResourceBuilder {
	return &ResourceBuilder{}
}

func (b *ResourceBuilder) CPU(cpu float64) *ResourceBuilder {
	b.properties.CPU = cpu
	return b
}

func (b *ResourceBuilder) Memory(memory string) *ResourceBuilder {
	b.properties.Memory = memory
	return b
}

func (b *ResourceBuilder) Requests(requests Requests) *ResourceBuilder {
	b.properties.Requests = requests
	return b
}

func (b *ResourceBuilder) Limits(limits Limits) *ResourceBuilder {
	b.properties.Limits = limits
	return b
}

func (b *ResourceBuilder) Do(ops ...ResourceOps) *ResourceBuilder {
	for _, op := range ops {
		op(b)
	}
	return b

}
