package vsc

import (
	"github.com/kercylan98/vivid/src/vivid"
	"time"
)

func newServiceAgent(ctx *ServiceContext, serviceType ServiceType, registry Registry, selector NodeSelector) *ServiceAgent {
	return &ServiceAgent{
		ctx:         ctx,
		serviceType: serviceType,
		registry:    registry,
		selector:    selector,
	}
}

type ServiceAgent struct {
	ctx         *ServiceContext
	registry    Registry
	selector    NodeSelector
	serviceType ServiceType
}

func (a *ServiceAgent) selectNode() vivid.ActorRef {
	nodes := a.registry.Get(a.serviceType)
	node, err := a.selector.Select(nodes)
	if err != nil {
		panic(err)
	}

	return node.Ref
}

func (a *ServiceAgent) Tell(message vivid.Message) {
	a.ctx.Tell(a.selectNode(), message)
}

func (a *ServiceAgent) Probe(message vivid.Message) {
	a.ctx.Probe(a.selectNode(), message)
}

func (a *ServiceAgent) Ask(message vivid.Message, timeout ...time.Duration) vivid.Future {
	return a.ctx.Ask(a.selectNode(), message, timeout...)
}
