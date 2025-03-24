package service

import "github.com/kercylan98/vivid-stateful-common/src/vsc/internal/core"

var _ core.ServiceProvider = (*Provider)(nil)

func NewProvider() *Provider {
	return &Provider{}
}

type Provider struct {
}

func (p *Provider) Provide(serviceType core.ServiceType) core.ServiceAgent {
	return NewAgent()
}
