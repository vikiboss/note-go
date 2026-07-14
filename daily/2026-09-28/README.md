# 第 11 周：综合项目 II——Diff、原子复制与取消

目标：把“发现差异”和“执行副作用”分开，让计划可审阅、可 dry-run、可重试。

## 精确阅读路线（90–130 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 复习精读 | 《Learning Go, 2e》Ch.14 Context；Ch.8 错误包装；Ch.13 io | context 完整重读，其余按 Apply 调用链查阅 | context 能在一个大文件内部终止复制 |
| 精读 | [context package](https://pkg.go.dev/context) | Overview、Context contract、WithCancel | 每个长操作都能说明取消来源和传播方向 |
| 查阅 | [io.Reader](https://pkg.go.dev/io#Reader)、[io.ErrShortWrite](https://pkg.go.dev/io#ErrShortWrite) | Reader 的 n/err 合法组合和短写 | 自定义 Reader 测中途取消与异常进度 |
| 查阅 | [os.CreateTemp](https://pkg.go.dev/os#CreateTemp)、[os.Rename](https://pkg.go.dev/os#Rename)、[filepath.IsLocal](https://pkg.go.dev/path/filepath#IsLocal) | 同目录临时文件、替换语义、路径安全 | 失败不会破坏旧目标，危险相对路径被拒绝 |

## 本周里程碑

- 学：纯计划函数、原子替换、context 取消、部分失败、校验与恢复。
- 练：新增/更新/不变文件；目标嵌套目录；复制中取消；目标原文件保护。
- 做：完成 `syncplan`，先生成 Action，再以临时文件 + rename 执行。
- 核心：默认不删除目标额外文件；计划排序稳定；复制失败不留下半文件。

## 命令

```bash
go test ./...
go test -race ./...
go vet ./...
```

## 教育重点

纯 `Plan` 函数像前端状态 reducer：输入清单，输出动作，没有 I/O，容易穷举测试。执行器单独处理权限、磁盘空间、取消和崩溃一致性。

## 加餐

给 Action 增加预期 hash，复制完成后校验；实现失败报告并允许安全重试。删除功能必须显式 flag 且先进入 trash。

完成标准：计划和执行分离、写入原子、取消可传播、集成测试覆盖新增与更新、race 通过。

## 任务梯度与证据

- 基础：Plan 覆盖 copy/update/unchanged/target-extra，并保证稳定顺序。
- 标准：复制使用同目录临时文件和 rename；context 能在单个大文件复制中生效。
- 挑战：复制后校验 hash，返回包含已完成与失败动作的报告，使重试可解释。

必须提交：一次 dry-run 计划、一条中途取消测试、一个证明旧目标文件未被半成品破坏的失败测试。

自检：为什么临时文件必须在目标同一文件系统？纯 Plan 与有副作用 Apply 分离后，哪些测试变简单了？
