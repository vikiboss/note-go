// 单行注释
/* 多行
   注释 */

// 导入包的子句在每个源文件的开头。
// main 比较特殊，它用来声明可执行文件，而不是一个库。
package main

// Import 语句声明了当前文件引用的包。
import (
	"fmt"       // Go 语言标准库中的包
	"io/ioutil" // 包含一些输入输出函数
	m "math"    // 数学标准库，在此文件中别名为m
	"net/http"  // 一个 web 服务器包
	"os"        // 系统底层函数，如文件读写
	"strconv"   // 字符串转换
)

// 函数声明：main 是程序执行的入口。
// 不管你喜欢还是不喜欢，反正 Go 就用了花括号来包住函数体。
func main() {
	// 往标准输出打印一行。
	// 用包名 fmt 限制打印函数。
	fmt.Println("你好世界，2 的立方是", m.Pow(2, 3))

	// 调用当前包的另一个函数。
	beyondHello()
}

// 函数可以在括号里加参数。
// 如果没有参数的话，也需要一个空括号。
func beyondHello() {
	var x int // 变量声明，变量必须在使用之前声明。
	x = 3     // 变量赋值。
	// 可以用 := 来偷懒，它自动把变量类型、声明和赋值都搞定了。
	y := 4
	sum, prod := learnMultiple(x, y)        // 返回多个变量的函数
	fmt.Println("sum:", sum, "prod:", prod) // 简单输出
	learnTypes()                            // 少于y分钟，学的更多！
}

/*
	<- 快看快看我是跨行注释 _(:з」∠)_

Go 语言的函数可以有多个参数和 *多个* 返回值。
在这个函数中， `x`、`y` 是参数，
`sum`、`prod` 是返回值的标识符（可以理解为名字）且类型为int
*/
func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y // 返回两个值
}

// 内置变量类型和关键词
func learnTypes() {
	// 短声明给你所想。
	str := "少说话多读书!" // String 类型

	s2 := `这是一个
可以换行的字符串` // 同样是 String 类型

	// 非 ascii 字符。Go 使用 UTF-8 编码。
	g := 'Σ' // rune 类型，int32 的别名，使用 UTF-8 编码

	f := 3.14195 // float64 类型，IEEE-754 64 位浮点数
	c := 3 + 4i  // complex128 类型，内部使用两个 float64 表示

	// var 变量可以直接初始化。
	var u uint = 7 // unsigned 无符号变量，但是实现依赖 int 型变量的长度
	var pi float32 = 22. / 7

	// 字符转换
	n := byte('\n') // byte 是 uint8 的别名

	// 数组（Array）类型的大小在编译时即确定
	var a4 [4]int           // 有 4 个 int 变量的数组，初始为0
	a3 := [...]int{3, 1, 5} // 有 3 个 int 变量的数组，同时进行了初始化

	// Array 和 slice 各有所长，但是 slice 可以动态的增删，所以更多时候还是使用 slice。
	s3 := []int{4, 5, 9}    // 回去看看 a3 ，是不是这里没有省略号？
	s4 := make([]int, 4)    // 分配 4 个 int 大小的内存并初始化为 0
	var d2 [][]float64      // 这里只是声明，并未分配内存空间
	bs := []byte("a slice") // 进行类型转换

	// 切片（Slice）的大小是动态的，它的长度可以按需增长
	// 用内置函数 append() 向切片末尾添加元素
	// 要增添到的目标是 append 函数第一个参数，
	// 多数时候数组在原内存处顺次增长，如
	s := []int{1, 2, 3}    // 这是个长度 3 的 slice
	s = append(s, 4, 5, 6) // 再加仨元素，长度变为 6 了
	fmt.Println(s)         // 更新后的数组是 [1 2 3 4 5 6]

	// 除了向 append() 提供一组原子元素（写死在代码里的）以外，我们
	// 还可以用如下方法传递一个 slice 常量或变量，并在后面加上省略号，
	// 用以表示我们将引用一个 slice、解包其中的元素并将其添加到 s 数组末尾。
	s = append(s, []int{7, 8, 9}...) // 第二个参数是一个 slice 常量
	fmt.Println(s)                   // 更新后的数组是 [1 2 3 4 5 6 7 8 9]

	p, q := learnMemory() // 声明 p, q 为 int 型变量的指针
	fmt.Println(*p, *q)   // * 取值

	// Map是动态可增长关联数组，和其他语言中的hash或者字典相似。
	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	// 在Go语言中未使用的变量在编译的时候会报错，而不是warning。
	// 下划线 _ 可以使你“使用”一个变量，但是丢弃它的值。
	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a3, s4, bs
	// 通常的用法是，在调用拥有多个返回值的函数时，
	// 用下划线抛弃其中的一个参数。下面的例子就是一个脏套路，
	// 调用os.Create并用下划线变量扔掉它的错误代码。
	// 因为我们觉得这个文件一定会成功创建。
	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "这句代码还示范了如何写入文件呢")
	file.Close()

	// 输出变量
	fmt.Println(s, c, a4, s3, d2, m)

	learnFlowControl() // 回到流程控制
}

// 和其他编程语言不同的是，go支持有名称的变量返回值。
// 声明返回值时带上一个名字允许我们在函数内的不同位置
// 只用写return一个词就能将函数内指定名称的变量返回
func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return // 隐式返回z，因为前面指定了它。
}

// Go 全面支持垃圾回收。Go 有指针，但是不支持指针运算。
// 你会因为空指针而犯错，但是不会因为增加指针而犯错。
func learnMemory() (p, q *int) {
	// 返回 int 型变量指针 p 和 q
	p = new(int) // 内置函数 new 分配内存
	// 自动将分配的 int 赋值 0，p 不再是空的了。
	s := make([]int, 20) // 给 20 个 int 变量分配一块内存
	s[3] = 7             // 赋值
	r := -2              // 声明另一个局部变量
	return &s[3], &r     // & 取地址
}

func expensiveComputation() int {
	return 1e6
}

func learnFlowControl() {
	// if 需要花括号，括号就免了
	if true {
		fmt.Println("这句话肯定被执行")
	}
	// 用 go fmt 命令可以帮你格式化代码，所以不用怕被人吐槽代码风格了，
	// 也不用容忍别人的代码风格。
	if false {
		// pout（生气、撅嘴 😠）
	} else {
		// gloat（窃喜，偷笑 😏）
	}
	// 如果太多嵌套的 if 语句，推荐使用 switch
	x := 1
	switch x {
	case 0:
	case 1:
		// 隐式调用 break 语句，匹配上一个即停止
	case 2:
		// 不会运行
	}
	// 和 if 一样，for 也不用括号
	for x := 0; x < 3; x++ { // ++ 自增
		fmt.Println("遍历", x)
	}
	// x 在这里还是 1，为什么？

	// for 是 go 里唯一的循环关键字，不过它有很多变种
	for { // 死循环
		break    // 骗你的
		continue // 不会运行的
	}

	// 用 range 可以枚举 array、slice、string、map、channel 等不同类型
	// 对于 channel 来说，range 会返回一个值，
	// 而 array、slice、string、map 等其他类型返回一对儿
	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		// 打印 map 中的每一个键值对
		fmt.Printf("索引：%s, 值为：%d\n", key, value)
	}
	// 如果你只想要值，那就用前面讲的下划线扔掉没用的
	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("你是。。 %s\n", name)
	}

	// 和 for 一样，if 中的 := 先给 y 赋值，然后再和 x 作比较。
	if y := expensiveComputation(); y > x {
		x = y
	}
	// 闭包函数
	xBig := func() bool {
		return x > 100 // x是上面声明的变量引用
	}
	fmt.Println("xBig:", xBig()) // true （上面把 y 赋给 x 了）
	x /= 1e5                     // x 变成 10
	fmt.Println("xBig:", xBig()) // 现在是 false

	// 除此之外，函数体可以在其他函数中定义并调用，
	// 满足下列条件时，也可以作为参数传递给其他函数：
	//   a) 定义的函数被立即调用
	//   b) 函数返回值符合调用者对类型的要求
	fmt.Println("两数相加乘二: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2)) // Called with args 10 and 2
	// => Add + double two numbers: 24

	// 当你需要 goto 的时候，你会爱死它的！
	goto love
love:

	learnFunctionFactory() // 返回函数的函数多棒啊
	learnDefer()      // 对 defer 关键字的简单介绍
	learnInterfaces() // 好东西来了！
}

func learnFunctionFactory() {
	// 空行分割的两个写法是相同的，不过第二个写法比较实用
	fmt.Println(sentenceFactory("原谅")("当然选择", "她！"))

	d := sentenceFactory("原谅")
	fmt.Println(d("当然选择", "她！"))
	fmt.Println(d("你怎么可以", "她？"))
}

// Decorator 在一些语言中很常见，在 go 语言中，
// 接受参数作为其定义的一部分的函数是修饰符的替代品
func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, mystring, after) // new string
	}
}

func learnDefer() (ok bool) {
	// defer 表达式在函数返回的前一刻执行
	defer fmt.Println("defer 表达式执行顺序为后进先出（LIFO）")
	defer fmt.Println("\n这句话比上句话先输出，因为")
	// 关于 defer 的用法，例如用 defer 关闭一个文件，
	// 就可以让关闭操作与打开操作的代码更近一些
	return true
}

// 定义 Stringer 为一个接口类型，有一个方法 String
type Stringer interface {
	String() string
}

// 定义 pair 为一个结构体，有 x 和 y 两个 int 型变量。
type pair struct {
	x, y int
}

// 定义 pair 类型的方法，实现 Stringer 接口。
func (p pair) String() string { // p 被叫做“接收器”
	// Sprintf 是 fmt 包中的另一个公有函数。
	// 用 . 调用 p 中的元素。
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func learnInterfaces() {
	// 花括号用来定义结构体变量，:= 在这里将一个结构体变量赋值给 p。
	p := pair{3, 4}
	fmt.Println(p.String()) // 调用 pair 类型 p 的 String 方法
	var i Stringer          // 声明 i 为 Stringer 接口类型
	i = p                   // 有效！因为 p 实现了 Stringer 接口（类似 Java 中的塑型）
	// 调用 i 的 String 方法，输出和上面一样
	fmt.Println(i.String())

	// fmt 包中的 Println 函数向对象要它们的 string 输出，实现了 String 方法就可以这样使用了。
	// （类似 Java 中的序列化）
	fmt.Println(p) // 输出和上面一样，自动调用 String 函数。
	fmt.Println(i) // 输出和上面一样。

	learnVariadicParams("great", "learning", "here!")
}

// 有变长参数列表的函数
func learnVariadicParams(myStrings ...interface{}) {
	// 枚举变长参数列表的每个参数值
	// 下划线在这里用来抛弃枚举时返回的数组索引值
	for _, param := range myStrings {
		fmt.Println("param:", param)
	}

	// 将可变参数列表作为其他函数的参数列表
	fmt.Println("params:", fmt.Sprintln(myStrings...))

	learnErrorHandling()
}

func learnErrorHandling() {
	// ", ok"用来判断有没有正常工作
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[1]; !ok { // ok 为 false，因为 m 中没有 1
		fmt.Println("别找了真没有")
	} else {
		fmt.Print(x) // 如果 x 在 map 中的话，x 就是那个值喽。
	}
	// 错误可不只是 ok，它还可以给出关于问题的更多细节。
	if _, err := strconv.Atoi("non-int"); err != nil { // _ discards value
		// 输出 "strconv.ParseInt: parsing "non-int": invalid syntax"
		fmt.Println(err)
	}
	// 待会再说接口吧。同时，
	learnConcurrency()
}

// c 是 channel 类型，一个并发安全的通信对象。
func inc(i int, c chan int) {
	c <- i + 1 // <- 把右边的发送到左边的 channel。
}

// 我们将用inc函数来并发地增加一些数字。
func learnConcurrency() {
	// 用 make 来声明一个 slice，make 会分配和初始化 slice，map 和 channel。
	c := make(chan int)
	// 用 go 关键字开始三个并发的 goroutine，如果机器支持的话，还可能是并行执行。
	// 三个都被发送到同一个 channel。
	go inc(0, c) // go is a statement that starts a new goroutine.
	go inc(10, c)
	go inc(-805, c)
	// 从 channel 中读取结果并打印。
	// 打印出什么东西是不可预知的。
	fmt.Println(<-c, <-c, <-c) // channel 在右边的时候，<-是读操作。

	cs := make(chan string)       // 操作 string 的 channel
	cc := make(chan chan string)  // 操作 channel 的 channel
	go func() { c <- 84 }()       // 开始一个 goroutine 来发送一个新的数字
	go func() { cs <- "wordy" }() // 发送给 cs
	// Select 类似于 switch，但是每个 case 包括一个 channel 操作。
	// 它随机选择一个准备好通讯的 case。
	select {
	case i := <-c: // 从 channel 接收的值可以赋给其他变量
		fmt.Println("这是……", i)
	case <-cs: // 或者直接丢弃
		fmt.Println("这是个字符串！")
	case <-cc: // 空的，还没作好通讯的准备
		fmt.Println("别瞎想")
	}
	// 上面 c 或者 cs 的值被取到，其中一个 goroutine 结束，另外一个一直阻塞。

	learnWebProgramming() // Go 很适合 web 编程，我知道你也想学！
}

// http 包中的一个简单的函数就可以开启 web 服务器。
func learnWebProgramming() {
	// ListenAndServe 第一个参数指定了监听端口，第二个参数是一个接口，特定是 http.Handler。
	go func() {
		err := http.ListenAndServe(":8080", pair{})
		fmt.Println(err) // 不要无视错误。
	}()

	requestServer()
}

// 使 pair 实现 http.Handler 接口的 ServeHTTP 方法。
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 使用 http.ResponseWriter 返回数据
	w.Write([]byte("Y分钟golang速成!"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("\n服务器消息： `%s`", string(body))
}
