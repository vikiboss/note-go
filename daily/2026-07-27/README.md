# 第 2 周：struct、方法、接口与错误

目标：用小接口、明确领域错误和组合设计一个可演进的库存库。

## 精确阅读路线（110–150 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 精读 | 《Learning Go, 2e》Ch.6–8 | Ch.6 指针语义；Ch.7 types/methods/interfaces；Ch.8 wrapping、Is/As、自定义错误 | 为 Inventory 画出方法集和公开错误表 |
| 精读 | [Effective Go: Methods](https://go.dev/doc/effective_go#methods)、[Interfaces and other types](https://go.dev/doc/effective_go#interfaces_and_types) | 读完对应小节，跳过与练习无关的格式建议 | 说明接口为什么由消费者定义、何时选指针接收者 |
| 精读 | [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors) | wrapping、Unwrap、Is、As | 写一个跨两层包装仍能 errors.Is 的测试 |
| 查阅 | [errors](https://pkg.go.dev/errors)、[encoding/json](https://pkg.go.dev/encoding/json) | Is/As/Join 与 Marshal/Unmarshal | 区分 JSON 语法错误、字段校验错误、NotFound |
| 选读 | Ch.10 Generics 的“何时使用”与 constraints 基础 | 只读概念，不改造库存库 | 写下为什么当前项目暂不需要泛型 |

## 学、练、做

- 学：值/指针接收者、隐式接口实现、`errors.Is/As`、JSON tag、包可见性。
- 练：增加 Update/Delete；区分“不存在”“非法数量”；验证导入数据。
- 做：扩展 `exercises/inventory`，实现可导入导出的库存服务。
- 核心：禁止通过错误字符串判断类型；测试必须覆盖 `ErrNotFound`。

## 日程

周一读 API；周二练接收者；周三练 error wrapping；周四完成 JSON 边界；周五至周六完成 CRUD；周日验收复盘。

## 运行与验收

```bash
go run ./exercises/inventory
go test ./...
go vet ./...
gofmt -w .
```

## Node/TS 迁移提示

Go interface 不需要 `implements`，通常由消费者定义。不要先造一个含十几个方法的 Repository；先从调用方真正需要的 `Get/Put` 开始。错误是 API 的一部分，应保留可判断的语义。

## 加餐

增加只读 `Finder` 接口，并写一个函数只依赖它；再实现第二个 fake，检验接口是否足够小。

完成标准：能说明指针接收者选择、小接口归属、`%w` 与 `errors.Is`，并新增一个 JSON 非法输入测试。

## 任务梯度与证据

- 基础：完成 Put/Get/Update/Delete 的正常与错误路径。
- 标准：证明 Import 具有“全有或全无”语义：中间一项非法时旧库存不改变。
- 挑战：在调用方定义只含 `Get` 的 Finder 接口，用内存实现和 fake 各运行同一组契约测试。

必须提交：一张公开 API 错误表（条件、返回错误、能否 `errors.Is/As`）；一个验证导入原子性的测试。

自检：为什么接口应放在消费者附近？如果只有一个实现，什么时候仍值得定义接口？
