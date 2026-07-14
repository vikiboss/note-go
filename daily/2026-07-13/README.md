# 第 0 周：工具链与 Node → Go 认知迁移

这不是语法速成周，而是一次入门诊断：用三个你在 Node.js 中熟悉的任务，观察 Go 如何表达数据、错误、资源生命周期与可测试边界。

预计投入 6–9 小时。完成后再进入第 1 周。

## 精确阅读路线（90–120 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 精读 | [Tutorial: Getting started](https://go.dev/doc/tutorial/getting-started)、[Create a module](https://go.dev/doc/tutorial/create-module) | 完整读完并亲手执行命令 | 能从空目录创建 module，解释 module、package、main 的关系 |
| 精读 | 《Learning Go, 2e》Ch.1；Ch.9 的 module/package/import 基础；Ch.11 的 go run/build/fmt/vet | Ch.1 完整；其余只读本周会用到的命令和包规则 | 不看 README 写出本周五条验收命令 |
| 精读 | [Error handling and Go](https://go.dev/blog/error-handling-and-go)、[Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors) | 前者完整；后者读 wrapping、Is、As | 用一句话区分“添加上下文”和“改变错误语义” |
| 查阅 | [os](https://pkg.go.dev/os)、[net/http](https://pkg.go.dev/net/http)、[encoding/json](https://pkg.go.dev/encoding/json)、[io](https://pkg.go.dev/io) | 只查样例实际调用的类型和函数 | 为每个练习标出资源所有者和可能返回的错误 |

阅读纪律与教材版本见 [../reading-guide.md](../reading-guide.md)。本周不通读 Ch.13；只在练习需要时查标准库。

## 本周交付

- 文件统计器：区分 byte、rune、word、line，并为缺失文件保留路径上下文；
- HTTP 客户端：设置超时、限制响应体、关闭 Body、把非 2xx 表达为可判断错误；
- JSON 汇总器：解码、验证字段、输出完成/未完成数量；
- 一份 150–300 字的 Node/TS → Go 迁移复盘。

## 开始前诊断

先写下你的答案，周末再修正：

1. `len("你好")` 是 2 还是 6？为什么？
2. `defer resp.Body.Close()` 应该放在检查哪个错误之后？
3. HTTP 404 时，`client.Get` 本身会返回 error 吗？
4. 为什么测试 HTTP 客户端不应该依赖真实公网？
5. Go 的 `error` 为什么更像返回值协议，而不是 exception？

## 学：只学够支撑三个任务的内容

- module 与 package：`go.mod`、`package main`、导出规则；
- 工具链：`go run`、`go test`、`go vet`、`gofmt`、`go build`；
- 基础控制流、struct、多返回值、`if err != nil`、`defer`；
- 文件与文本：`os.ReadFile`、byte、rune、UTF-8、`strings.Fields`；
- HTTP：`http.Client`、Transport、Body 所有权、状态码与大小限制；
- JSON：struct tag、`json.Unmarshal`、语法错误与业务校验的区别；
- 测试：`testing.T`、`t.TempDir`、依赖注入一个假的 RoundTripper。

不要在本周提前学习框架、ORM、泛型或复杂并发。

## 练：按顺序完成

### 1. 文件读取与 UTF-8

```bash
go run ./exercises/01-file-reader -file ./exercises/01-file-reader/sample.txt
go test ./exercises/01-file-reader -v
```

先预测输出，再把 sample 改成含中文和末尾换行的文本。解释 bytes 与 runes 为什么不同。新增一个空文件或多行边界测试。

### 2. 可测试的 HTTP 客户端

```bash
go test ./exercises/02-http-client -v
go run ./exercises/02-http-client -url https://example.com
```

先读测试中的 `roundTripFunc`：它替换的是网络边界，不需要真的请求公网。增加一个超时或响应体读取失败测试，并用 `errors.As` 判断 `HTTPStatusError`。

### 3. JSON 解码与业务验证

```bash
go run ./exercises/03-json-summary -file ./exercises/03-json-summary/sample.json
go test ./exercises/03-json-summary -v
```

分别输入 malformed JSON、空标题和合法任务。说清“解析失败”与“字段不合法”为什么是两个层次。

## 做：三级任务

- 基础：不改实现，运行全部样例；新增 1 个边界测试；写出 5 条 Node/Go 差异。
- 标准：三个练习各新增 1 个测试；文件统计加入空输入语义；HTTP 加入一种底层失败；JSON 报告第几个任务字段非法。
- 挑战：把三个 main 的“解析参数/执行/输出”拆开，使退出行为可做集成测试；不要引入第三方包。

## Node.js → Go 对照

| Node.js / TypeScript 习惯 | 本周要建立的 Go 模型 |
| --- | --- |
| `package.json` 与 scripts | `go.mod` 与统一 go 命令 |
| exception / rejected Promise | error 是显式返回值，调用方选择处理方式 |
| 默认围绕异步 I/O | 先写清晰同步流程，需要吞吐时再引入并发 |
| `Buffer.length` / JS string length | byte 与 UTF-8 rune 必须明确区分 |
| mock 整个 HTTP 库 | 在 Transport / Reader / Writer 等小边界替换依赖 |
| npm 包优先 | 先理解标准库能力与资源所有权 |

## 一周节奏

- 周一：安装验证、运行 hello、完成开始前诊断；
- 周二：文件练习，补失败路径；
- 周三：HTTP 成功路径与 Body 生命周期；
- 周四：HTTP 非 2xx、底层错误与离线测试；
- 周五：JSON 解码、校验与表驱动测试；
- 周六：完成标准档任务，统一格式化；
- 周日：全量验收、复盘，并列出第 1 周最想验证的问题。

## 验收

```bash
go run ./cmd/hello
gofmt -w .
go test ./...
go vet ./...
go build ./...
```

通过标准：

- 命令全部成功，且没有依赖真实网络的测试；
- 至少新增 3 个有教育意义的测试，包含错误或边界路径；
- 能解释 `%w`、`errors.Is/As`、`defer`、`io.Reader` 的用途；
- 能指出每个练习的资源所有者、错误边界和可测试边界；
- 复盘不是罗列 API，而是至少记录一次“预测错误 → 实验 → 新模型”。

## 周末自检

1. 哪个错误应该包装，哪个错误应该转成领域错误？
2. 为什么限制响应体大小不仅是性能问题，也是安全边界？
3. 如果文件非常大，`os.ReadFile` 为什么不再合适？下一步会换成什么抽象？
4. 哪段代码最像你原来会写的 Node.js，哪段最体现 Go？
