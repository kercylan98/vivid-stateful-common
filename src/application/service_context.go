package application

import (
	"github.com/kercylan98/vivid-stateful-common/src/application/service"
	"github.com/kercylan98/vivid/src/vivid"
)

func newServiceContext(registry Registry, ctx vivid.ActorContext) service.Context {
	return &serviceContext{
		ActorContext: ctx,
		registry:     registry,
	}
}

type serviceContext struct {
	vivid.ActorContext
	registry Registry
}

func (s *serviceContext) Select(serviceType string) (vivid.ActorRef, error) {
	return s.registry.Select(serviceType)
}

func (s *serviceContext) SelectAll(serviceType string) []vivid.ActorRef {
	return s.registry.SelectAll(serviceType)
}
