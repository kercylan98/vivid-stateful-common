package application

import (
	"fmt"
	"github.com/kercylan98/go-log/log"
	"github.com/kercylan98/vivid-stateful-common/src/application/service"
	"github.com/kercylan98/vivid/src/vivid"
)

var _ vivid.Actor = (*actor)(nil)

func newActor(registry Registry) vivid.Actor {
	return &actor{
		registry: registry,
	}
}

type actor struct {
	registry Registry // 注册表

	services   map[string]service.Provider // Actor 支持的服务集合
	serviceIns map[string]vivid.ActorRef   // Actor 实例化的服务集合
}

func (a *actor) OnReceive(ctx vivid.ActorContext) {
	switch m := ctx.Message().(type) {
	case *vivid.OnLaunch:
		a.onLaunch(ctx)
	case *registerServiceAsk:
		a.onRegisterServiceAsk(ctx, m)
	case *launchAsk:
		a.onLaunchAsk(ctx, m)
	}
}

func (a *actor) onLaunch(ctx vivid.ActorContext) {
	a.services = make(map[string]service.Provider)
	a.serviceIns = make(map[string]vivid.ActorRef)
}

// onRegisterServiceAsk 注册服务
func (a *actor) onRegisterServiceAsk(ctx vivid.ActorContext, m *registerServiceAsk) {
	logger := ctx.Logger()

	// 生成一个校验用途的服务实例
	s := m.Provider.Provide()
	if s == nil {
		ctx.Reply(ErrorInvalidService)
		return
	}

	// 检查服务类型是否为空
	if s.Type() == "" {
		ctx.Reply(fmt.Errorf("%w, service: %T", ErrorServiceTypeEmpty, s))
		return
	}

	// 检查服务是否已注册
	if _, ok := a.services[s.Type()]; ok {
		ctx.Reply(fmt.Errorf("%w, service: %T", ErrorServiceRegistered, s))
		return
	}

	a.services[s.Type()] = m.Provider

	ctx.Reply(&registerServiceReply{})

	logger.Info("service", log.String("event", "registered"), log.String("type", s.Type()))
}

func (a *actor) onLaunchAsk(ctx vivid.ActorContext, m *launchAsk) {
	// 实例化服务
	for typ, provider := range a.services {
		a.serviceIns[typ] = ctx.ActorOf(func() vivid.Actor {
			var serviceIns = provider.Provide()
			var ctx service.Context
			return vivid.ActorFN(func(c vivid.ActorContext) {
				switch c.Message().(type) {
				case *vivid.OnLaunch:
					ctx = newServiceContext(a.registry, c)
				case *launchAsk:
					ctx.Reply(serviceIns.OnLaunch(ctx))
				default:
					serviceIns.OnReceive(ctx)
				}
			})
		})
	}

	// 向注册表注册服务
	for typ, ref := range a.serviceIns {
		a.registry.Register(typ, ref)
	}

	// 全部启动完毕后初始化服务
	for _, ref := range a.serviceIns {
		_, err := ctx.Ask(ref, m).Result()
		if err != nil {
			ctx.Reply(err)
			return
		}
	}

	ctx.Reply(new(launchReply))
}
