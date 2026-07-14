# 第 7 周：TCP 与协议分帧

目标：理解 TCP 是字节流，没有天然的“消息”边界。

## 精确阅读路线（100–140 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 精读 | [net package](https://pkg.go.dev/net) | Overview、Conn、Listener、Dial/Listen、Pipe、deadline | 画出 client/server 连接与关闭顺序 |
| 精读 | [io.ReadFull](https://pkg.go.dev/io#ReadFull) | 函数契约及 EOF/ErrUnexpectedEOF | 写分片读取和截断测试 |
| 精读 | [encoding/binary](https://pkg.go.dev/encoding/binary) | ByteOrder、BigEndian、PutUint32/Uint32 | 写出协议的 4-byte 长度字段布局 |
| 复习 | 《Learning Go, 2e》Ch.13 的 io 与 HTTP/network 相关内容 | 只读 I/O 边界与超时思想 | 能解释一次 Write 为何不保证一次 Read |
| 查阅 | [net.Conn](https://pkg.go.dev/net#Conn) | SetDeadline、Close、读写并发保证 | 用 net.Pipe 测 deadline 和退出 |

## 学、练、做

- 学：`net.Conn`、deadline、半关闭、长度前缀协议、最大消息限制、优雅关闭。
- 练：用 `net.Pipe` 测分片读取；测试截断帧和超大帧；加入读写 deadline。
- 做：基于 `frame` 包实现 echo server/client 或小型键值服务。
- 核心：永远限制帧大小；使用 `io.ReadFull`，不能假设一次 Read 返回完整消息。

## 命令

```bash
go run ./exercises/frame
go test ./...
go vet ./...
```

## Node/TS 迁移提示

Node socket 的 data event 同样不等于一条业务消息，只是 API 容易掩盖这一点。协议必须自己定义边界、编码、大小、超时和错误响应。

## 加餐

增加版本字节和消息类型；对未知版本返回协议错误。写一页协议说明，包含字节布局和最大长度。

完成标准：分片/截断/超大输入都有测试，连接双方能退出，能解释粘包为何不是 TCP bug。

## 任务梯度与证据

- 基础：验证分片读取、短写、截断和超大帧。
- 标准：用 `net.Pipe` 完成请求/响应，并给读写设置 deadline。
- 挑战：写协议 v1 文档，加入版本、消息类型和错误帧，再实现多客户端 TCP 服务。

必须提交：协议字节布局图；最大帧选择依据；客户端异常断开时服务端退出证据。

自检：一次 Write 对应一次 Read 吗？长度字段本身为什么也必须使用固定字节序和上限校验？
