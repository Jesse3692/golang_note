// defer_call_func.go
package main

func test(a, b int)  {
	defer println("dispose...")

	println(a / b)
}

func main()  {
	test(10, 0)
}