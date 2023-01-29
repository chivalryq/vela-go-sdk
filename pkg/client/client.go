package client

import (
	"context"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	apiapp "vela-go-sdk/pkg/apis/app"

	"vela-go-sdk/pkg/apis"
)

type Client interface {
	Get(ctx context.Context, key client.ObjectKey) (apis.Application, error)
	List(ctx context.Context, opts client.ListOption) ([]apis.Application, error)
	Create(ctx context.Context, app apis.Application) error
	Delete(ctx context.Context, app apis.Application) error
	Update(ctx context.Context, app apis.Application) error
	Patch(ctx context.Context, app apis.Application, patch client.Patch) error
}

type clientImpl struct {
	clt client.Client
}

func NewFromClient(clt client.Client) Client {
	return &clientImpl{clt: clt}
}

func New(config *rest.Config) (Client, error) {
	clt, err := client.New(config, client.Options{})
	if err != nil {
		return nil, err
	}
	return NewFromClient(clt), nil
}

func NewFromConfig(config *rest.Config) (Client, error) {
	return New(config)
}

func NewFromConfigWithOptions(config *rest.Config, options client.Options) (Client, error) {
	clt, err := client.New(config, options)
	if err != nil {
		return nil, err
	}
	return NewFromClient(clt), nil
}

func (c clientImpl) Get(ctx context.Context, key client.ObjectKey) (apis.Application, error) {
	_app := v1beta1.Application{}
	err := c.clt.Get(ctx, key, &_app)
	if err != nil {
		return nil, err
	}
	return apiapp.FromK8sObject(&_app)
}

func (c clientImpl) List(ctx context.Context, opts client.ListOption) ([]apis.Application, error) {
	//TODO implement me
	panic("implement me")
}

func (c clientImpl) Create(ctx context.Context, app apis.Application) error {
	appObj := app.Build()
	return c.clt.Create(ctx, &appObj)
}

func (c clientImpl) Delete(ctx context.Context, app apis.Application) error {
	appObj := app.Build()
	return c.clt.Delete(ctx, &appObj)
}

func (c clientImpl) Update(ctx context.Context, app apis.Application) error {
	appObj := app.Build()
	return c.clt.Update(ctx, &appObj)
}

func (c clientImpl) Patch(ctx context.Context, app apis.Application, patch client.Patch) error {
	appObj := app.Build()
	return c.clt.Patch(ctx, &appObj, patch)
}
