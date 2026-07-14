# 第 2 周：struct、方法、接口与错误

目标：用小接口、明确领域错误和组合设计一个可演进的库存库。

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

