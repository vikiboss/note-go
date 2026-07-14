# 第 5 周：并发文件处理与背压

目标：把并发用于真实 I/O，同时限制资源、传播错误并支持取消。

## 精确阅读路线（90–130 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 复习精读 | 《Learning Go, 2e》Ch.12 的 channel/select、backpressure 相关内容 | 只重读能解释本周数据流的小节 | 写出最大 goroutine 数和最大同时打开文件数 |
| 精读 | [The Go Memory Model](https://go.dev/ref/mem) | 重点读 Advice、happens-before 与 channel/lock 同步；不背形式定义 | 用 race 解释“逻辑正确”与“无数据竞争”的差异 |
| 查阅 | [sync](https://pkg.go.dev/sync)、[filepath](https://pkg.go.dev/path/filepath)、[crypto/sha256](https://pkg.go.dev/crypto/sha256) | WaitGroup、WalkDir、Sum256 | 能说明为何 hash 可流式计算、为何 map 输出要排序 |
| 复习 | [Pipelines and cancellation](https://go.dev/blog/pipelines) | fan-out/fan-in 与停止短路部分 | 遇到首个错误后仍安全排空结果 |

## 学、练、做

- 学：fan-out/fan-in、并发上限、背压、mutex 与 channel 的选择、文件哈希。
- 练：生成重复文件；限制 worker；加入不存在路径；测试取消。
- 做：完善 `dedupe`，递归发现文件并输出重复内容组。
- 核心：错误不能静默丢弃；结果必须稳定排序；race 必须通过。

## 命令

```bash
go run ./exercises/dedupe file-a file-b
go test -race ./...
go vet ./...
```

## Node/TS 迁移提示

不要像 `Promise.all(十万个任务)` 一样为每个文件启动一个 goroutine。worker pool 表达资源上限；无缓冲或小缓冲 channel 会把下游速度反馈给上游。

## 加餐

递归遍历目录，忽略符号链接，并输出成功数、失败数、字节数和耗时。

完成标准：并发有上限、支持 context、错误可定位、结果稳定、race 通过。

## 任务梯度与证据

- 基础：重复、唯一、缺失文件各有测试；遇错后仍排空 worker 结果。
- 标准：递归发现普通文件、忽略符号链接、按 hash 与路径稳定输出。
- 挑战：边扫描边哈希，分别限制目录遍历和文件打开并发，输出成功/失败/字节/耗时指标。

必须提交：最大同时打开文件数的设计说明；取消或错误发生时所有 goroutine 如何退出的说明。

自检：一个文件一个 goroutine 会在哪些资源上失控？背压在这个程序中具体发生在哪个 channel？
