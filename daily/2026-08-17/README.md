# 第 5 周：并发文件处理与背压

目标：把并发用于真实 I/O，同时限制资源、传播错误并支持取消。

## 学、练、做

- 学：fan-out/fan-in、并发上限、背压、mutex 与 channel 的选择、文件哈希。
- 练：生成重复文件；限制 worker；加入不存在路径；测试取消。
- 做：完善 `dedupe`，递归发现文件并输出重复内容组。
- 核心：错误不能静默丢弃；结果必须稳定排序；race 必须通过。

## 命令

```bash
go run ./exercises/dedupe file-a file-b
go test -race ./...
go vet ./...
```

## Node/TS 迁移提示

不要像 `Promise.all(十万个任务)` 一样为每个文件启动一个 goroutine。worker pool 表达资源上限；无缓冲或小缓冲 channel 会把下游速度反馈给上游。

## 加餐

递归遍历目录，忽略符号链接，并输出成功数、失败数、字节数和耗时。

完成标准：并发有上限、支持 context、错误可定位、结果稳定、race 通过。

