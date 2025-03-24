package vsc

import (
	"fmt"
	"github.com/kercylan98/vivid/src/vivid"
)

func NewApplication(system vivid.ActorSystem, config Config) *Application {
	return &Application{
		system:   system,
		config:   &config,
		services: make(map[ServiceType]ServiceFactory),
	}
}

type Application struct {
	system   vivid.ActorSystem
	config   *Config
	services map[ServiceType]ServiceFactory
}

func (a *Application) LoadService(serviceType ServiceType, serviceFactory ServiceFactory) {
	if _, registered := a.services[serviceType]; registered {
		panic(fmt.Errorf("service already registered, service: %s", serviceType))
	}

	a.services[serviceType] = serviceFactory
}

func (a *Application) Run() error {

	for _, factory := range a.services {
		factory := factory
		a.system.ActorOf(func() vivid.Actor {
			var ctx *ServiceContext
			var actor = factory()
			return vivid.ActorFN(func(c vivid.ActorContext) {
				switch c.Message().(type) {
				case *vivid.OnLaunch:
					ctx = newServiceContext(c)
					if initializer, ok := actor.(ServiceInitializer); ok {
						initializer.Init()
					}

					serviceProvider := newServiceProvider(ctx, a.config.Registry)
					if injector, ok := actor.(ServiceInjector); ok {
						injector.Inject(serviceProvider)
					}
				default:
					actor.OnReceive(ctx)
				}
			})
		})
	}

	return nil
}
