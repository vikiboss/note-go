# 第 9 周：性能、分配与可观测性

目标：用数据驱动优化，并让程序暴露足够的运行证据。

## 学、练、做

- 学：benchmark、`-benchmem`、CPU/heap/block profile、逃逸分析、GC、`log/slog`。
- 练：比较 strings.Fields 与 Scanner；记录 allocs/op；为输入规模增加子基准。
- 做：给文本分析器建立基线、做一次单点优化，并写性能报告。
- 核心：提交优化前后数字；如果优化无效，也要保留结论。

## 命令

```bash
go test ./...
go test -bench=. -benchmem ./...
go test -cpuprofile cpu.out ./exercises/analyzer
go tool pprof cpu.out
go test -gcflags="-m" ./exercises/analyzer
```

## Node/TS 迁移提示

不要把 V8 的性能直觉直接搬到 Go。先用目标工作负载测量；关注吞吐、尾延迟、分配和 GC，而不是只看某段代码“更底层”。

## 性能报告模板

环境与命令；输入规模；基线；profile 证据；唯一改动；复测；正确性验证；结论与未解决问题。

完成标准：benchmark 可重复、优化不改语义、有结构化事件示例、报告能说明噪声和取舍。

