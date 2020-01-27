// degrade_assgin.go

package main

func main()  {
	x := 100
	println(&x)

	x, y := 200, "abc"

	println(&x, x)
	println(y)
}