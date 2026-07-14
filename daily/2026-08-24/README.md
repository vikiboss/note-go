# 第 6 周：CLI、流式 I/O 与系统边界

目标：设计一个能进入 shell pipeline 的工具，而不是把逻辑都塞进 main。

## 精确阅读路线（90–130 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 精读 | 《Learning Go, 2e》Ch.13 的 io interfaces、Reader/Writer、buffered I/O、JSON 部分 | 只读与流和编码相关的小节 | 把核心 Run 保持为 Reader/Writer 边界 |
| 精读 | [io](https://pkg.go.dev/io) | Reader、Writer、Copy、LimitReader、ReadAll | 为每个 I/O 函数写清所有权和内存特征 |
| 精读 | [bufio.Scanner](https://pkg.go.dev/bufio#Scanner) | Split、Buffer、Err 与 token 限制 | 用测试证明 70 KiB 合法行不会被默认限制误伤 |
| 查阅 | [flag](https://pkg.go.dev/flag)、[os](https://pkg.go.dev/os)、[os/signal](https://pkg.go.dev/os/signal) | Parse、Stdin/Stdout/Stderr、NotifyContext | 写 stdout/stderr/exit code 契约表 |

读完后应能回答：何时 Scanner 合适，何时应改用 Reader；为什么业务逻辑不应直接调用 os.Exit。

## 学、练、做

- 学：`io.Reader/Writer`、bufio、flag、stdin/stdout/stderr、退出码、环境变量和信号。
- 练：从文件与 stdin 读取 JSONL；无效行带行号；输出稳定 JSON。
- 做：完成日志级别汇总 CLI，并增加 `--strict` 或过滤参数。
- 核心：核心逻辑只依赖 Reader/Writer，测试不启动子进程。

## 命令

```bash
printf '{"level":"info"}\n{"level":"error"}\n' | go run ./exercises/jsonl-summary
go test ./...
go vet ./...
```

## Node/TS 迁移提示

`io.Reader/io.Writer` 类似 Node stream 的最小同步接口。先围绕流设计核心逻辑，再让 main 负责 flags 和退出码，会自然得到可测试、可组合的程序。

## 加餐

正确区分 usage error（退出码 2）、数据错误（退出码 1）和成功（0），并用子进程集成测试验证。

完成标准：支持 stdin/stdout、错误含行号、无全局状态、测试覆盖坏 JSON 与空输入。

## 任务梯度与证据

- 基础：覆盖合法、坏 JSON、缺失 level、空输入和超过默认 Scanner 大小的合法行。
- 标准：增加 `-input` 与 level 过滤，但核心 `Run` 仍只依赖 Reader/Writer。
- 挑战：增加 `--strict=false`，跳过坏行并把诊断写 stderr；定义部分成功的退出码。

必须提交：CLI 契约表（stdout、stderr、退出码）；一个用 bytes.Buffer 完成的纯内存测试。

自检：为什么日志数据不能写 stderr？Scanner 默认 token 限制如果不显式处理，会造成什么生产问题？
