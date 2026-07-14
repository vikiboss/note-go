# 第 12 周：综合项目 III——可发布交付

目标：把正确的核心能力包装成用户能安全使用、维护者能诊断的工具。

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
go build ./cmd/filesync
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

