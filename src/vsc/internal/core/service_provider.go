package core

type ServiceProvider interface {
	// Provide 获取指定类型的服务代理
	Provide(serviceType ServiceType) ServiceAgent
}
