package vsc

import (
	"github.com/kercylan98/vivid/src/vivid"
)

type Node struct {
	Ref          vivid.ActorRef // 节点引用
	ServiceTypes []ServiceType  // 服务类型
}
