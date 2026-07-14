# 第 12 周：综合项目 III——可发布交付

目标：把正确的核心能力包装成用户能安全使用、维护者能诊断的工具。

## 精确阅读路线（100–150 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 复习精读 | 《Learning Go, 2e》Ch.9、Ch.11、Ch.15 | module/package、build/tooling、integration tests/benchmarks | 从干净目录完成 test、race、vet、build |
| 精读 | [Compile and install](https://go.dev/doc/tutorial/compile-install)、[Module release workflow](https://go.dev/doc/modules/release-workflow) | 完整读完 | 写版本号、构建命令、发布与回滚说明 |
| 通读 | [Fuzzing tutorial](https://go.dev/doc/tutorial/fuzz) | 完整通读；至少给路径或协议边界写一个 fuzz target 作为挑战 | 保存一个 seed corpus 和 fuzz 命令 |
| 查阅 | [log/slog](https://pkg.go.dev/log/slog)、[signal.NotifyContext](https://pkg.go.dev/os/signal#NotifyContext) | 结构化字段、Handler、信号取消 | 演示成功、失败、中断三类日志 |
| 通读 | [actions/setup-go](https://github.com/actions/setup-go) 官方 README | Usage、版本选择、缓存与工作流示例 | CI 在干净环境运行 fmt 检查、test、race、vet |

发布前回到 [../reading-guide.md](../reading-guide.md)，确认你使用的是指定教材版本，并能用官方资料解释所有关键设计。

## 本周里程碑

- 学：CLI 契约、dry-run、结构化日志、集成测试、CI、版本与跨平台构建。
- 练：空目录、嵌套目录、更新、dry-run、取消、错误退出；运行 race/vet。
- 做：完善 `filesync`，写最终 README、架构说明、限制、性能记录和发布说明。
- 核心：dry-run 绝不改目标；实际执行可重复；错误指出具体路径。

## 命令

```bash
go run ./cmd/filesync -source ./source -target ./target -dry-run
go test ./...
go test -race ./...
go vet ./...
go build -o "${TMPDIR:-/tmp}/filesync" ./cmd/filesync
```

## 发布清单

- `--help` 可理解，破坏性行为默认关闭；
- core 与 main 分离；主流程有集成测试；
- CI 跑 gofmt 检查、test、race、vet；
- README 有安装、示例、非目标、故障恢复；
- 记录一次 benchmark/pprof 或大目录实测；
- 打版本标签前从干净目录走一遍教程。

## 最终复盘

对照 `../../plan.md` 的完成定义评分。写下：最重要的三次认知迁移、仍不确定的两个并发/存储问题、下一阶段要读的一个开源项目，以及未来四周的维护计划。

完成标准：陌生用户仅看 README 能安全运行；测试、race、vet、build 全通过；你能用 10 分钟讲清数据流、错误边界和崩溃恢复。

## 最终演示脚本

1. 创建 source/target，演示新增、更新、不变文件；
2. 先 dry-run，证明目标未改变；
3. 正式同步并再次运行，证明第二次为零动作；
4. 中断一个大文件同步，证明不会留下目标半文件；
5. 展示结构化日志、测试、race、vet 与跨平台构建结果。

最终提交应包含：用户 README、架构/数据流图、限制与恢复说明、测试证据、性能记录、发布说明。代码完成只是交付的一部分。

自检：如果目标目录位于另一种文件系统、磁盘空间不足或进程在 rename 前崩溃，用户会观察到什么？下一版本最值得增加的保障是什么？
