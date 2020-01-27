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