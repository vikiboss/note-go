# 第 4 周：goroutine、channel 与取消

目标：能画清每个 goroutine 的创建、阻塞和退出路径。

## 精确阅读路线（120–160 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 精读 | 《Learning Go, 2e》Ch.12 | goroutine、channel、select、WaitGroup、并发实践全部 | 为 worker pool 画出启动、阻塞、关闭、回收路径 |
| 精读 | Ch.14 Context | 完整读完，特别关注取消传播和“不把 context 存进 struct” | 写预先取消和执行中取消测试 |
| 精读 | [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines) | 完整读完 | 指出谁关闭 jobs/results，为什么 |
| 通读 | [Go Concurrency Patterns: Context](https://go.dev/blog/context) | 完整通读，API 细节查 package 文档 | 能解释 Done、Err、deadline 的职责 |
| 查阅 | [context](https://pkg.go.dev/context)、[sync](https://pkg.go.dev/sync) | Context、WithCancel、WaitGroup | race 连续运行仍通过 |

## 学、练、做

- 学：goroutine、无/有缓冲 channel、close、select、WaitGroup、context。
- 练：改变 worker 数；取消长任务；确认结果 channel 只由发送方关闭。
- 做：扩展可取消 worker pool，在并发执行时保持输入顺序。
- 核心：新增“取消后快速返回”测试，并用 race detector 验收。

## 命令

```bash
go run ./exercises/pipeline
go test -race ./...
go vet ./...
```

## Node/TS 迁移提示

goroutine 不是 Promise；它没有自动 join、异常传播或取消。启动前必须决定谁等待、谁关闭 channel、谁触发取消。context 传取消信号，不用来装随意的业务参数。

## 加餐

让处理函数可注入，并实现 fail-fast 与 collect-all 两种错误策略，对比 API 和退出路径。

完成标准：无泄漏迹象、race 通过、取消测试稳定，并能解释 channel 由谁关闭。

## 任务梯度与证据

- 基础：画出 jobs、results、WaitGroup 与三个 goroutine 角色的生命周期。
- 标准：证明 worker 数变化不改变输出顺序；覆盖 0 worker、空输入、预先取消。
- 挑战：让处理函数可注入并返回 error，实现 fail-fast；取消后仍要回收所有已启动 goroutine。

必须提交：一张退出路径图和一条 `go test -race` 输出。不要用 sleep 让测试“碰巧通过”。

自检：为什么接收方通常不应关闭 channel？有缓冲 channel 会解决泄漏问题吗？
