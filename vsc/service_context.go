package vsc

import (
	"github.com/kercylan98/vivid/src/vivid"
)

func newServiceContext(ctx vivid.ActorContext) *ServiceContext {
	return &ServiceContext{
		ActorContext: ctx,
	}
}

type ServiceContext struct {
	vivid.ActorContext
}
