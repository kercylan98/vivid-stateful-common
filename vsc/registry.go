package vsc

// Registry 是服务注册表接口
type Registry interface {
	// Register 注册服务
	Register(node *Node)

	// Unregister 注销服务
	Unregister(node *Node)

	// Get 获取服务
	Get(serviceType ServiceType) []*Node
}
