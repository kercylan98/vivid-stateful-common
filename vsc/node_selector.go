package vsc

type NodeSelector interface {
	// Select 选择服务
	Select(nodes []*Node) (*Node, error)
}
