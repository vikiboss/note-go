# 第 1 周：类型、集合与值语义

目标：摆脱 JS 动态容器思维，理解 Go 的零值、slice 别名、map 查询与 string/rune。

## 精确阅读路线（100–140 分钟）

| 程度 | 阅读内容 | 读到哪里 | 完成证据 |
| --- | --- | --- | --- |
| 精读 | 《Learning Go, 2e》Ch.2–3 | Ch.2 的零值、字符串、类型转换；Ch.3 的 array/slice/map/struct 全部 | 画出 slice 的指针、长度、容量与底层数组 |
| 通读 | Ch.5 的“Go is Call by Value”和可变参数部分 | 不读闭包细节也可 | 能解释传入 slice 后哪些修改对调用方可见 |
| 精读 | [Go Slices: usage and internals](https://go.dev/blog/slices-intro) | 完整读完并运行 append/copy 小实验 | 写一个因共享底层数组失败的测试 |
| 精读 | [Strings, bytes, runes and characters](https://go.dev/blog/strings) | 读到 range 解码 UTF-8 的部分 | 预测中文字符串的 len、rune 数和 range 次数 |
| 查阅 | [strings](https://pkg.go.dev/strings)、[unicode/utf8](https://pkg.go.dev/unicode/utf8)、[sort](https://pkg.go.dev/sort) | 只查本周调用项 | 说明为何公开查询结果必须稳定排序 |

不要在本周提前抽象泛型容器；重点是值语义、别名和稳定 API。

## 学、练、做

- 学：Tour of Go 的 basics、slice、map；阅读 `strings`、`sort`、`unicode/utf8` 文档。
- 练：给目录查询增加“按最低库存过滤”和“按名称排序”；为中文名称补测试。
- 做：把 `exercises/catalog` 扩展成支持新增、查询、标签去重和防御性复制的内存目录。
- 核心：修改一次查询行为并补测试；解释为什么 `Items()` 不直接返回内部 slice。

## 日程

- 周一：运行测试，画出 slice 的指针、长度、容量。
- 周二：比较 byte、rune、string；完成中文测试。
- 周三：实现过滤与稳定排序。
- 周四：制造一次 slice 共享导致的修改，再用复制修复。
- 周五至周六：完成目录查询器和错误输入处理。
- 周日：验收并写 150–300 字复盘。

## 运行与验收

```bash
go run ./exercises/catalog
go test ./...
go vet ./...
gofmt -w .
```

## Node/TS 迁移提示

JS Array 是动态对象；Go slice 是指向底层数组的描述符。append 可能复用底层数组，也可能分配新数组。map 遍历顺序不可依赖，公开返回集合时要考虑调用方能否修改内部状态。

## 加餐

实现 `TopTags(n int)`，要求结果稳定；用 `go test -run TopTags -count=50` 验证没有依赖 map 随机顺序。

完成标准：能解释零值、`value, ok := m[key]`、byte/rune 区别、slice aliasing，并且新增至少两个边界测试。

## 任务梯度与证据

- 基础：给 `Add` 的名称、库存、标签规范化各补一个表驱动用例。
- 标准：实现 `TopTags(n)`，明确并列时排序规则，并证明调用方不能修改内部数据。
- 挑战：为查询增加分页，定义负数、越界和空结果语义，不要引入“万能 Options”。

必须提交：一张 slice 共享底层数组示意图；一个故意触发 aliasing 的失败测试；修复后的测试输出。

自检：什么时候返回 `nil slice`，什么时候返回空 slice？为什么 map 结果必须显式排序后才能成为稳定 API？
