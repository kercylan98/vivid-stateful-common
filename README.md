# Vivid Stateful Common 

基于 [Vivid](https://github.com/kercylan98/vivid) 实现的轻量化分布式有状态服务开发套件，为分布式系统中的有状态服务提供简洁的生命周期管理与跨节点通信方案。

## 核心机制

- 通过 `Service` 抽象定义节点能力；
- 使用可自定义的服务注册表统一管理各节点的 `Service` 元数据
- 基于 Actor 消息机制实现透明 RPC 调用
- 内置服务激活 / 注销的完整生命周期管理

