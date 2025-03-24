package vsc

func newServiceProvider(ctx *ServiceContext, registry Registry) *ServiceProvider {
	return &ServiceProvider{
		ctx:      ctx,
		registry: registry,
	}
}

type ServiceProvider struct {
	ctx      *ServiceContext
	registry Registry
}

func (p *ServiceProvider) Provide(serviceType ServiceType, selector NodeSelector) *ServiceAgent {
	return newServiceAgent(p.ctx, serviceType, p.registry, selector)
}
