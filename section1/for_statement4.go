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