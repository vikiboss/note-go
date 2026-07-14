# Go 学习：第 0 周

目标：完成 Go 环境验证，并用 Go 重写三个 Node.js 常见脚本。

## 本周任务

- [x] 安装并验证 Go（当前：Go 1.26.4）
- [ ] 通读 `exercises/01-file-reader/main.go`，运行并改成读取你自己的文件
- [ ] 为文件读取逻辑补充一个失败场景测试
- [ ] 完成 HTTP 请求练习中的超时、非 2xx 响应处理
- [ ] 完成 JSON 统计练习，并为它添加一个字段校验
- [ ] 每天写 5 分钟复盘：哪一点和 Node.js 的习惯不同？

## 运行

```bash
go run ./cmd/hello
go run ./exercises/01-file-reader -file ./exercises/01-file-reader/sample.txt
go run ./exercises/02-http-client -url https://api.github.com
go run ./exercises/03-json-summary -file ./exercises/03-json-summary/sample.json
go test ./...
go vet ./...
gofmt -w .
```

如果你的受限环境无法使用默认 Go 构建缓存，使用项目内缓存：

```bash
mkdir -p work/go-cache work/go-mod-cache
GOCACHE=$PWD/work/go-cache GOMODCACHE=$PWD/work/go-mod-cache go test ./...
```

## Node.js → Go 对照

| Node.js | Go |
| --- | --- |
| `package.json` | `go.mod` |
| `npm run` | `go run` / `go test` |
| `try/catch` / rejected Promise | 显式返回 `error` |
| `async/await` | 同步调用优先；并发使用 goroutine | 
| duck typing 常见 | 小 interface 的隐式实现 |

## 每日建议（60–90 分钟）

1. 20 分钟：阅读一个练习，先预测结果。
2. 30–45 分钟：修改并运行代码。
3. 10 分钟：加一个测试或错误分支。
4. 5 分钟：记录一次与 TypeScript/Node.js 不同的设计选择。

完成标准：三个练习均可运行，所有测试、`go vet` 通过，并能说清楚 `error`、`defer`、`io.Reader` 与模块的用途。
