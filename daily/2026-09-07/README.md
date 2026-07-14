# 第 8 周：持久化、原子写与幂等任务

目标：理解存储和后台任务的正确性边界，而不是只会调用 ORM。

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

