package application

import (
	"fmt"
	"github.com/kercylan98/vivid-stateful-common/src/vsc/internal/core"
	"github.com/kercylan98/vivid-stateful-common/src/vsc/internal/service"
)

var _ core.Application = (*Application)(nil)

func NewApplication() *Application {
	return &Application{
		services: make(map[core.ServiceType]core.Service),
	}
}

type Application struct {
	services map[core.ServiceType]core.Service
}

func (a *Application) LoadService(service core.Service) {
	serviceType := service.Type()
	if _, registered := a.services[serviceType]; registered {
		panic(fmt.Errorf("service already registered, service: %T", service))
	}

	a.services[serviceType] = service
}

func (a *Application) Run() error {
	// 初始化服务，完成服务自身的初始化
	for _, s := range a.services {
		if initializer, ok := s.(core.ServiceInitializer); ok {
			if err := initializer.Init(); err != nil {
				return err
			}
		}
	}

	// 注入服务，完成服务之间的依赖注入
	serviceProvider := service.NewProvider()
	for _, s := range a.services {
		if injector, ok := s.(core.ServiceInjector); ok {
			injector.Inject(serviceProvider)
		}
	}

	return nil
}
