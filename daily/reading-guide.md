# 阅读指南：如何把官方文档、书籍与实验结合

本指南服务于 [13 周实验路线](README.md)。课程不要求靠 README 独立承担全部知识讲解；官方资料负责定义事实与语义，主教材负责连续叙事，daily 实验负责把理解变成证据。

## 固定教材版本

主教材统一采用：

> Jon Bodner, *Learning Go, 2nd Edition*, O'Reilly, 2024。

章节映射只对第 2 版有效：

1. Setting Up Your Go Environment
2. Predeclared Types and Declarations
3. Composite Types
4. Blocks, Shadows, and Control Structures
5. Functions
6. Pointers
7. Types, Methods, and Interfaces
8. Errors
9. Modules, Packages, and Imports
10. Generics
11. Go Tooling
12. Concurrency in Go
13. The Standard Library
14. Context: It's for Cancellation
15. Writing Tests
16. Here There Be Dragons: Reflection, Unsafe, and cgo

没有这本书时，直接完成每周 README 中的官方必读即可。《Go 程序设计语言》可以作为原理补充，但不要强行按章号替换，因为它早于 modules、context、errors.Is/As、slog 等现代实践。

## 三种阅读程度

- 精读：逐段阅读并运行关键示例；能不看原文解释核心规则。通常 30–60 分钟。
- 通读：理解文章论点、边界和典型模式；不记忆所有 API。通常 15–30 分钟。
- 查阅：围绕当周代码定位类型、函数、错误和示例；用完后能找到同一信息。通常 5–15 分钟。

每周总阅读时间控制在 90–150 分钟。超过三小时仍未开始实验，通常说明阅读范围过大。

## 正确顺序

1. 回答当周 README 的诊断问题。
2. 精读第一份官方资料，并记录一个仍不确定的问题。
3. 阅读对应书章，建立连续结构。
4. 运行实验，先预测再验证。
5. 只有遇到具体问题时查阅 package 文档和 Spec。
6. 周末回看原文，修正最初答案并写三条“规则 + 证据”。

不要先把一整本书读完再开始写 Go，也不要只看视频或二手博客。语言语义冲突时，以当前 [Go Specification](https://go.dev/ref/spec) 和标准库文档为准。

## 全局映射

| 周 | 教材章节 | 官方阅读主线 |
| ---: | --- | --- |
| 0 | Ch.1、Ch.8/9/11/13 选段 | Getting Started、modules、errors、os/http/json |
| 1 | Ch.2–3、Ch.5 选段 | Tour、Slices、Strings/UTF-8 |
| 2 | Ch.6–8，Ch.10 选读 | methods/interfaces、Go 1.13 errors |
| 3 | Ch.15，Ch.11 选段 | testing、subtests、race detector |
| 4 | Ch.12、Ch.14 | pipelines、context |
| 5 | Ch.12 复习 | memory model、sync、filesystem/hash |
| 6 | Ch.13 的 io 部分 | io、bufio、flag、signals |
| 7 | Ch.13 相关部分 | net、binary、ReadFull、deadlines |
| 8 | Ch.8/14 复习 | database/sql、transactions、cancellation |
| 9 | Ch.11、Ch.15 benchmark 部分 | diagnostics、pprof、benchmark、slog |
| 10 | Ch.5/7/13 复习 | WalkDir、Rel、hash、stable JSON |
| 11 | Ch.8/13/14 复习 | context-aware I/O、rename、recovery |
| 12 | Ch.9/11/15 | build/install、module release、CI/testing |

具体链接、阅读边界和完成证据位于每个日期目录的“精确阅读路线”。

## 阅读笔记模板

```markdown
## 阅读记录

- 资料：
- 阅读程度：精读 / 通读 / 查阅
- 一句话规则：
- 它解决的问题：
- 容易从 Node/TS 错迁移的点：
- 验证它的测试或命令：
- 仍未解决的问题：
```

判断是否读懂的标准不是划线数量，而是能否写出一个可能失败的测试，区分规则的适用与不适用场景。

