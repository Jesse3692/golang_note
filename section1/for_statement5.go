// for_statement5.go
package main

func main()  {
	x := []int{100, 101, 102}

	for i, n := range x {
		println(i, ":", n)
	}
}