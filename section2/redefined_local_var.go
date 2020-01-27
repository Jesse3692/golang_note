// redefined_local_var.go
package main

var x = 100

func main()  {
	println(&x, x)  // &x 内存地址，这里的x是全局变量

	x := "abc"
	println(&x, x) // 这里是将全局变量重新定义和初始化
}