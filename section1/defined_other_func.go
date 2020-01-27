// defined_other_func.go

package main

// X defined int type.
type X int

func (x *X) inc()  {  // 名称前的参数称作receiver，作用类似python self
	*x++
}

func main()  {
	var x X
	x.inc()
	print(x)
}