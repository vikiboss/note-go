# 第 3 周：测试、竞态与基准

目标：把测试当成设计反馈，而不只是覆盖率数字。

## 精确阅读路线（100–140 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 精读 | 《Learning Go, 2e》Ch.15 | 单元测试、table tests、coverage、benchmarks 全部；fuzzing 可选读 | 自己写一个表驱动测试和 benchmark |
| 通读 | Ch.11 中 go test、code quality、lint/tooling 相关部分 | 跳过暂时不用的构建自动化 | 解释 test、vet、race 各自证明什么 |
| 精读 | [testing package](https://pkg.go.dev/testing) | Overview、Subtests、Benchmarks | benchmark 把 setup 移出计时区 |
| 精读 | [Using Subtests and Sub-benchmarks](https://go.dev/blog/subtests)、[Data Race Detector](https://go.dev/doc/articles/race_detector) | 两篇完整读完 | 临时制造竞态并保存 detector 报告 |

覆盖率只用于发现未执行路径，不作为质量分数。

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

## 任务梯度与证据

- 基础：运行测试、覆盖率、race 与 benchmark，保存原始输出。
- 标准：临时去掉锁，让 race detector 给出证据；恢复后验证 Value 与 Add 的锁粒度。
- 挑战：实现 mutex 与 atomic 两版计数器，用并行 benchmark 比较，说明结论适用的工作负载。

不要提交带竞态的最终代码。实验过程写进复盘即可。

自检：测试通过而 race 失败意味着什么？benchmark 更快是否足以证明生产环境更好？
