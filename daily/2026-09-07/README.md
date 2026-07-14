# 第 8 周：持久化、原子写与幂等任务

目标：理解存储和后台任务的正确性边界，而不是只会调用 ORM。

## 精确阅读路线（120–170 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 精读 | [Opening a database handle](https://go.dev/doc/database/open-handle)、[Managing connections](https://go.dev/doc/database/manage-connections) | 两篇完整读完 | 解释 sql.DB 为什么不是单连接、何时调连接池参数 |
| 精读 | [Executing transactions](https://go.dev/doc/database/execute-transactions)、[Canceling operations](https://go.dev/doc/database/cancel-operations) | 完整读完 | 写出事务 commit/rollback 与 context 取消路径 |
| 复习 | 《Learning Go, 2e》Ch.8 的错误边界、Ch.14 的 context | 只重读 wrapping、取消传播 | 文件版和 SQL 版使用同一错误语义 |
| 查阅 | [database/sql](https://pkg.go.dev/database/sql)、[os.CreateTemp](https://pkg.go.dev/os#CreateTemp)、[os.Rename](https://pkg.go.dev/os#Rename) | DB/Tx、临时文件、rename | 画出写临时文件—sync—close—rename 的崩溃点 |

本周文件仓库用于理解语义，不代表它支持多进程事务；SQL 加餐必须用唯一约束保证幂等。

## 学、练、做

- 学：事务、连接池、context、迁移、幂等键、重试/退避、任务状态机。
- 练：重开仓库验证持久化；重复幂等键；模拟写失败；设计 pending/running/done/failed。
- 做：完善标准库文件任务仓库，再用 SQLite/PostgreSQL 实现同一语义作为加餐。
- 核心：写入采用临时文件 + rename；重复请求不会创建两个任务。

## 命令

```bash
go test ./...
go test -race ./...
go vet ./...
```

## Node/TS 迁移提示

ORM transaction callback 不是事务知识本身。先明确原子性、唯一约束、重试后果和崩溃恢复，再选数据库 API。幂等需要持久化约束，不能只靠进程内 Set。

## 加餐

定义最小 Repository 接口并实现 `database/sql` 版本；迁移中给 idempotency_key 建唯一索引；用 context timeout 测试取消。

完成标准：重启后数据存在、重复键稳定返回同一任务、损坏文件报错、race 通过。

## 任务梯度与证据

- 基础：覆盖幂等创建、重启恢复、损坏文件和合法/非法状态迁移。
- 标准：模拟持久化失败，证明内存状态不会领先磁盘；ID 不能因未来删除任务而复用。
- 挑战：用 `database/sql` 实现同一语义，以唯一约束而非“先查再插”保证跨进程幂等。

必须提交：任务状态图；崩溃可能发生的位置清单；文件版与数据库版原子性的差异说明。

自检：rename 原子是否等于整个事务持久可靠？两个进程同时写这个文件仓库会发生什么？
