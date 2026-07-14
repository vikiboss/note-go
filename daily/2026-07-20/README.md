# 第 1 周：类型、集合与值语义

目标：摆脱 JS 动态容器思维，理解 Go 的零值、slice 别名、map 查询与 string/rune。

## 学、练、做

- 学：Tour of Go 的 basics、slice、map；阅读 `strings`、`sort`、`unicode/utf8` 文档。
- 练：给目录查询增加“按最低库存过滤”和“按名称排序”；为中文名称补测试。
- 做：把 `exercises/catalog` 扩展成支持新增、查询、标签去重和防御性复制的内存目录。
- 核心：完成所有 TODO，并解释为什么 `Items()` 不直接返回内部 slice。

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

