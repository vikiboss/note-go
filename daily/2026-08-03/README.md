# 第 3 周：测试、竞态与基准

目标：把测试当成设计反馈，而不只是覆盖率数字。

## 学、练、做

- 学：表驱动、子测试、`t.Helper`、`t.Cleanup`、benchmark、race detector。
- 练：给 Counter 补并发测试；先去掉锁观察 race，再恢复；比较批量与单次增加。
- 做：交付一个有正确性、并发和性能证据的安全计数器。
- 核心：亲自运行一次 `go test -race ./...` 和一次 benchmark。

## 命令

```bash
go test ./...
go test -race ./...
go test -cover ./...
go test -bench=. -benchmem ./...
go vet ./...
```

## Node/TS 迁移提示

Node 单线程事件循环常让共享内存竞态不显眼；Go 中两个 goroutine 同时读写 map/字段就是数据竞争。测试“跑过”不等于无竞态，race detector 是独立证据。

## 实验记录

在 README 末尾记录：机器、命令、ns/op、allocs/op，以及一次改动前后结果。不要只保留更快的数据。

完成标准：竞态检测通过；有表驱动、并发、benchmark；能解释 benchmark 为什么要避免把 setup 算进去。

