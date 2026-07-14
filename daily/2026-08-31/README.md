# 第 7 周：TCP 与协议分帧

目标：理解 TCP 是字节流，没有天然的“消息”边界。

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

