实现一个高性能的 goroutine_pool，支持特性如下：
- 可以指定协程最大数量
- 长时间（用户可以指定）空闲的协程会被关闭
- 优雅退出
- 优雅错误处理