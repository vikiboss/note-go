# 第 6 周：CLI、流式 I/O 与系统边界

目标：设计一个能进入 shell pipeline 的工具，而不是把逻辑都塞进 main。

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

