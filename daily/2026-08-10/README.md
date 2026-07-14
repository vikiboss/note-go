# 第 4 周：goroutine、channel 与取消

目标：能画清每个 goroutine 的创建、阻塞和退出路径。

## 学、练、做

- 学：goroutine、无/有缓冲 channel、close、select、WaitGroup、context。
- 练：改变 worker 数；注入失败；取消长任务；确认结果 channel 只由发送方关闭。
- 做：扩展可取消 worker pool，支持错误收集和有序输出。
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

