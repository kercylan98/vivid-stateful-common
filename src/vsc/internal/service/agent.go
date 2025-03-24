package service

import "github.com/kercylan98/vivid-stateful-common/src/vsc/internal/core"

var _ core.ServiceAgent = (*Agent)(nil)

func NewAgent() *Agent {
	return &Agent{}
}

type Agent struct {
}
