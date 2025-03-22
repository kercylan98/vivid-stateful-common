package application

import (
	"github.com/kercylan98/go-log/log"
	"github.com/kercylan98/vivid-stateful-common/src/application/service"
	"github.com/kercylan98/vivid/src/vivid"
	"time"
)

func New(registry Registry) Application {
	app := &application{
		system: vivid.NewActorSystem(vivid.ActorSystemConfiguratorFN(func(config *vivid.ActorSystemConfig) {
			logger := log.GetBuilder().Production()
			config.WithLoggerProvider(log.ProviderFn(func() log.Logger {
				return logger
			}))
		})).StartP(),
	}

	app.ref = app.system.ActorOf(func() vivid.Actor {
		return newActor(registry)
	}, func(config *vivid.ActorConfig) {
		config.WithName("app")
	})

	return app
}

type Application interface {
	// Register 注册服务
	Register(provider service.Provider) error

	// RegisterP 注册服务
	RegisterP(provider service.Provider) Application

	// Run 运行应用
	Run(timeout ...time.Duration) error
}

type application struct {
	system vivid.ActorSystem // ActorSystem 实例
	ref    vivid.ActorRef    // application 自身的 ActorRef
}

func (a *application) Register(provider service.Provider) error {
	_, err := vivid.TypedFutureFrom[*registerServiceReply](a.system.Ask(a.ref, &registerServiceAsk{provider})).Result()
	if err != nil {
		return err
	}
	return nil
}

func (a *application) RegisterP(provider service.Provider) Application {
	if err := a.Register(provider); err != nil {
		panic(err)
	}
	return a
}

func (a *application) Run(timeout ...time.Duration) error {
	var d = time.Minute
	if len(timeout) > 0 {
		d = timeout[0]
	}
	_, err := a.system.Ask(a.ref, new(launchAsk), d).Result()
	return err
}
