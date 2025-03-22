package application

import "github.com/kercylan98/vivid/src/vivid"

type Registry interface {
	// Register 注册服务，内部应非阻塞的持续重试失败的情况
	Register(serviceType string, ref vivid.ActorRef)

	// Select 选择服务，返回服务的 ActorRef
	Select(serviceType string) (vivid.ActorRef, error)

	// SelectAll 选择所有服务，返回服务的 ActorRef 列表
	SelectAll(serviceType string) []vivid.ActorRef
}
