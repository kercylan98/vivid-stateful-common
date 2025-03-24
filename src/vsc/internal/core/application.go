package core

type Application interface {
	// LoadService 加载服务
	LoadService(service Service)

	// Run 运行应用程序
	Run() error
}
