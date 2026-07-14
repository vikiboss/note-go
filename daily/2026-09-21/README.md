# 第 10 周：综合项目 I——Manifest 与边界

目标：先定义问题和稳定数据模型，再写同步执行逻辑。

## 精确阅读路线（90–130 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 复习 | 《Learning Go, 2e》Ch.5 的 API 边界、Ch.7 的类型设计、Ch.13 的 io | 只重读与 Scan/Entry 设计相关部分 | 写一页 ADR，说明为什么 Scan 返回稳定 Entry |
| 精读 | [filepath.WalkDir](https://pkg.go.dev/path/filepath#WalkDir)、[filepath.Rel](https://pkg.go.dev/path/filepath#Rel) | 回调错误、SkipDir、路径相对化 | 覆盖嵌套、.git、空目录和错误路径 |
| 查阅 | [io.Copy](https://pkg.go.dev/io#Copy)、[hash.Hash](https://pkg.go.dev/hash#Hash)、[crypto/sha256](https://pkg.go.dev/crypto/sha256) | 流式 hash 接口 | 说明为何不把整文件读入内存 |
| 查阅 | [encoding/json](https://pkg.go.dev/encoding/json) | Encoder 与 map key 稳定性、struct tags | 连续两次扫描输出字节级一致 |

## 本周里程碑

- 学：需求切片、非目标、包边界、相对路径安全、文件哈希、稳定序列化。
- 练：扫描嵌套目录；忽略 `.git`；验证同内容同哈希；处理不可读文件。
- 做：完成 `manifest` 生成器，输出按相对路径排序的 JSON 清单。
- 核心：清单不能包含绝对路径；相同目录重复扫描结果一致。

## 先写设计

在复盘中明确：同步方向、冲突策略、是否删除目标多余文件、符号链接策略、权限/mtime 是否保留、最大文件假设。默认采用单向复制且不删除目标额外文件。

## 命令

```bash
go run ./exercises/manifest .
go test ./...
go vet ./...
```

## 教育重点

这一周刻意不做复制。稳定 manifest 是后续 diff、缓存、重试和审计的契约。过早把扫描、比较、复制塞进一个函数会让错误恢复和测试都变困难。

完成标准：输出稳定、路径相对且安全、错误带路径、测试覆盖嵌套与忽略规则，并写架构决策记录。

## 任务梯度与证据

- 基础：嵌套文件、`.git`、空目录、同内容 hash 与稳定排序都有测试。
- 标准：定义符号链接、权限、mtime、大小变化中的明确策略，并写入 manifest schema 版本。
- 挑战：并发哈希但保持稳定输出；比较小文件很多与大文件很少两种负载。

必须提交：一页 ADR，明确目标、非目标、路径模型和冲突策略；两个相同目录 manifest 完全一致的证据。

自检：为什么 manifest 只保存相对路径？只比较 size/mtime 与比较内容 hash 各有什么风险？
