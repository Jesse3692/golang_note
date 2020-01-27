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