# Go语言学习笔记（读书笔记）

## 运行环境

- OS: CentOS Linux release 7.7.1908 (Core)
- Golang: go version go1.13.3 linux/amd64
- gcc (GCC) 4.8.5 20150623 (Red Hat 4.8.5-39)

## 第一章 概述

源文件使用UTF-8编码。每个源文件都属于包的一部分，在文件头部用package声明所属包名称。入口函数main没有参数，且必须放在main包中。
```go
// test.go

package main

func main()  {
	println("Hello, World")
}
```

使用import导入标准库或第三方包

```go
// import_package.go
package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Hello, World")
}
```

### 变量

使用var定义变量，支持类型推断。基础数据类型划分清晰明确，有助于编写跨平台应用。编译器确保变量总是被初始化为零值，避免出现意外情况。

```go
// defined_var.go

package main

func main()  {
	var x int32
	var s = "Hello, world"

	println(x, s)
}

-> 0 Hello, world
```

在函数内部，还可以省略var关键字，使用更简单的定义模式。

```go
// simple_defined.go

package main

func main()  {
	x := 100
	println(x)
}
```

还需要注意的是，编译器将未使用的局部变量定义当作错误。

### 表达式

Go仅有三种流程控制语句，比大多数语言简单。

#### if语句

```go
// if_statement.go

package main

func main()  {
	x := 100
	if x > 0 {
		println("x")
	}else if x < 0 {
		println("-x")
	}else {
		println("0")
	}

}
```

#### switch语句

```go
// switch_statement.go

package main

func main()  {
	x := 100

	switch  {
	case x > 0:
		println("x")
	case x < 0:
		println("-x")
	default:
		println("0")
	}
}
```

#### for语句

```go
// for_statement1.go
package main

func main()  {
	for i := 0; i < 5; i++ {
		println(i)
	}
}

// for_statement2.go
package main

func main()  {
	for i := 4; i >= 0; i-- {
		println(i)
	}
}


// for_statement3.go
package main

func main()  {
	x := 0

	// 相当于while (x < 5)
	for x < 5 {
		println(x)
		x++
	}
}

// for_statement4.go
package main

func main()  {
	x := 4
	// 相当于while (true)
	for {
		println(x)
		x--

		if x < 0 {
			break
		}
	}

}

// for_statement5.go
package main

func main()  {
	x := []int{100, 101, 102}

	for i, n := range x {
		println(i, ":", n)
	}
}
```

### 函数

函数可定义多个返回值，甚至对其命名。

```go
// defined_return_val.go
package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error)  {
	if b == 0 {
		return 0, errors.New("divsion by zero")
	}

	return a / b, nil
}

func main()  {
	a, b := 10, 2  // 定义多个变量
	c, err := div(a, b)  // 接收多返回值

	fmt.Println(c, err)  // -> 5 <nil>
}
```

## 第二章 类型

从计算机系统的角度看，变量是一段或多段用来存储数据的内存。

通过类型转换或指针操作，我们可用不同饭食修改变量值，但这并不意味着改变了变量类型。

关键字`var`用于定义变量，和C不同，类型被放在变量名后。另外，运行时内存分配操作会确保变量自动初始化为二进制零值（zero value），避免出现不可预测行为。如显式提供初始化值，可省略变量类型，由编译器推断。

```go
var x int // 自动初始化为零值
var y = false // 自动推断为bool类型
var x, y int // 定义多个相同类型的变量
var z, a = 100, "xyz" // 定义多个不同类型

// 以组方式整理多行变量定义
var (
    x ,y int
    z, a = 100, "xyz"
)

// 使用简短模式进行变量定义和初始化语法
func main() {
    x := 100
    z, a := 100, "xyz"
}
```

虽然简短模式这么方便但是它还有一些限制：

- 定义变量，同时显式初始化

- 不能提供数据类型

- 只能用在函数内部

简短模式的错误操作：

1. 在局部作用域重新定义和初始化全局变量

```go
// redefined_local_var.go
package main

var x = 100

func main()  {
	println(&x, x)  // &x 内存地址，这里的x是全局变量

	x := "abc"
	println(&x, x) // 这里是将全局变量重新定义和初始化
}
```

2. 简短模式并不总是重新定义变量，也可能是部分退化的赋值操作。

```go
// degrade_assgin.go

package main

func main()  {
	x := 100
	println(&x)  // 0xc000036748

	x, y := 200, "abc"

	println(&x, x)  // 0xc000036748 200
	println(y)
}
```

退化赋值的前提条件：最少一个新变量被重新被定义，且必须是同一作用域。而不同的作用域，全部是新变量定义。

```go
// redefined_diffscope.go

package main

func main()  {
	x := 100
	println(&x, x)  // 0xc000036748 100

	{
		x, y := 200, 300
		println(&x, x, y)  // 0xc000036740 200 300
	}
}
```

在处理函数错误返回值时，退化赋值允许我们重复使用err变量

```go
// error_var.go（不能正常运行）

package main

import (
	"log"
	"os"
)

func main()  {
	f, err := os.Open("/dev/random")
	...
	
	buf := make([]byte, 1024)
	n, err := f.Read(buf)  // err退化赋值，n重新定义
	...
}
```

### 多变量赋值

在进行多变量赋值操作时，首先计算出所有右值，然后再依次完成赋值操作。

```go
// mutil_assign.go

package main

func main()  {
	x, y := 1, 2
	x, y = y + 3, x + 2

	println(x, y)
}
```



