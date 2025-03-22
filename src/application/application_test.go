package application_test

import (
	"github.com/kercylan98/go-log/log"
	"github.com/kercylan98/vivid-stateful-common/src/application"
	"github.com/kercylan98/vivid-stateful-common/src/application/service"
	"github.com/kercylan98/vivid-stateful-common/src/registry"
	"testing"
)

type TestService struct {
}

func (t *TestService) Type() string {
	return "Test"
}

func (t *TestService) OnLaunch(ctx service.Context) error {
	ctx.Logger().Info("test service launched")
	for _, ref := range ctx.SelectAll("Test") {
		ctx.Tell(ref, 1)
	}
	return nil
}

func (t *TestService) OnReceive(ctx service.Context) {
	switch m := ctx.Message().(type) {
	case int:
		ctx.Logger().Info("test service received", log.Any("message", m))
	}
}

func TestApplication_Run(t *testing.T) {
	if err := application.New(registry.NewMemoryRegistry()).RegisterP(service.ProviderFN(func() service.Service {
		return new(TestService)
	})).Run(); err != nil {
		t.Fatal(err)
	}
}
