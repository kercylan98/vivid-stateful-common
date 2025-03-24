package vsc

// Service 是应用服务接口，它用于描述应用可提供的功能服务
type Service interface {
	// OnReceive 该函数将在服务接收到消息时进行调用
	OnReceive(ctx *ServiceContext)
}

// ServiceInitializer 是服务初始化接口，当 Service 实现了此接口时，将在服务启动时调用 Init 方法
type ServiceInitializer interface {
	Service

	// Init 该函数将在 Service 实例化后进行调用
	//
	// 在该函数中，通常是完成自身的初始化，而非产生对其他服务的依赖
	Init()
}

// ServiceInjector 是服务注入接口，当 Service 实现了此接口时，将在 ServiceInitializer.Init 后调用 Inject 方法
type ServiceInjector interface {
	Service

	// Inject 该函数将在 ServiceInitializer.Init 后进行调用，用于完成服务之间的依赖注入
	//
	// 该函数将传入 ServiceProvider 实例，通过 ServiceProvider 实例，可以获取到所需服务的实例
	Inject(provider *ServiceProvider)
}
