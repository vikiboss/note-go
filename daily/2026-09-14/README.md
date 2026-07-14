# 第 9 周：性能、分配与可观测性

目标：用数据驱动优化，并让程序暴露足够的运行证据。

## 精确阅读路线（110–160 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 精读 | 《Learning Go, 2e》Ch.15 的 benchmarks；Ch.11 的 profiling/tooling 部分 | benchmark 完整；tooling 只读测量与分析 | 写可重复的基线命令和环境说明 |
| 精读 | [Diagnostics](https://go.dev/doc/diagnostics) | profiling、tracing、runtime statistics 的用途 | 为当前问题选择一种 profile，并说明为何不是另一种 |
| 精读 | [Profiling Go Programs](https://go.dev/blog/pprof) | 完整通读并实际运行 pprof top | 保存优化前热点证据 |
| 查阅 | [testing benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)、[runtime/pprof](https://pkg.go.dev/runtime/pprof)、[log/slog](https://pkg.go.dev/log/slog) | benchmark API、profile 类型、Handler/attributes | 报告 ns/op、B/op、allocs/op，并输出一条结构化事件 |

不要把示例中的最快实现当普遍结论；输入规模、Unicode 语义和内存峰值都属于基准契约。

## 学、练、做

- 学：benchmark、`-benchmem`、CPU/heap/block profile、逃逸分析、GC、`log/slog`。
- 练：比较 strings.Fields 与 Scanner；记录 allocs/op；为输入规模增加子基准。
- 做：给文本分析器建立基线、做一次单点优化，并写性能报告。
- 核心：提交优化前后数字；如果优化无效，也要保留结论。

## 命令

```bash
go test ./...
go test -bench=. -benchmem ./...
go test -run '^$' -bench=Analyze -cpuprofile cpu.out ./exercises/analyzer
go tool pprof cpu.out
go test -gcflags="-m" ./exercises/analyzer
```

## Node/TS 迁移提示

不要把 V8 的性能直觉直接搬到 Go。先用目标工作负载测量；关注吞吐、尾延迟、分配和 GC，而不是只看某段代码“更底层”。

## 性能报告模板

环境与命令；输入规模；基线；profile 证据；唯一改动；复测；正确性验证；结论与未解决问题。

完成标准：benchmark 可重复、优化不改语义、有结构化事件示例、报告能说明噪声和取舍。

## 任务梯度与证据

- 基础：固定输入，分别记录两个实现的 ns/op、B/op、allocs/op。
- 标准：用 CPU/heap profile 找到热点，只改一个变量后重复测量至少三次。
- 挑战：加入不同输入规模的子基准与真实文件测试，讨论吞吐、内存峰值和 Unicode 正确性。

必须提交：原始 benchmark 输出、profile 截图或 top 文本、优化前后表格、未优化的理由。

自检：0 allocs/op 为什么不自动代表更快？微基准如何避免编译器消除与不真实输入？
