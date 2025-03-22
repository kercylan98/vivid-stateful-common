package service

// Service 是应用服务接口，它用于描述应用可提供的服务功能
type Service interface {
	// Type 获取服务类型名称
	Type() string

	// OnLaunch 当服务启动时，将调用此方法
	OnLaunch(ctx Context) error

	// OnReceive 当服务接收到消息时，将调用此方法
	OnReceive(ctx Context)
}

type Provider interface {
	// Provide 提供服务实例
	Provide() Service
}

type ProviderFN func() Service

func (fn ProviderFN) Provide() Service {
	return fn()
}
