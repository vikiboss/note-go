## Go 学习笔记

- go 的类型声明放在变量之后
- 返回值可在函数的参数后被定义并直接 return 返回
- := 结构不能出现在函数外
- defer 前缀关键字的语句会将函数推迟到外层函数返回之后执行(但参数的求值会立即完成)

### Go 的数据结构

- bool
- int (int, int32, int64)
- float (float32, float64)
- string

- byte(unit8 alias)
- unit(uint uint8 uint16 uint32 uint64 uintptr)
- complex(complex64, complex128)

### 零值

- bool 为 false
- int 为 0
- string 为 ""

### 类型转换

```go
var num int = 32
var fnum float32 = float32(num)
// 或者使用 := 自动推断
fnum := float32(num)
```

### for

Go 中只有 for 一种循环 且条件不需要小括号括起来 (if 也是如此)

```go
// for 循环
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// "while" 循环
for count < 100 { }

// 无限循环
for { }
```

### if

if 语句可以在条件表达式前执行一个简单的语句

但该语句声明的变量作用域仅在 if 与 else 的 { }块之内

```go
if v := math.Pow(x,n); v < lim { }
```

### switch

go 的 switch 不需要在每个 case 块内 break (即默认只会执行一个 case)

相反的 若要依次执行 case 需要使用 fallthrough 关键字

没有条件的 switch 同 switch true 一样

```go
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
```
