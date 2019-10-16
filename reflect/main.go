package main

import (
	"fmt"
	"reflect"
)

type I interface {
	test()
}

type Impl struct {
	n int
}

func (i Impl) test() {
	fmt.Println("test")
}

////
func test(v interface{}) {
	fmt.Println(v)
	fmt.Println(reflect.TypeOf(v))
}

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t)
	fmt.Println(t.String())
	x := Impl{100}
	i := reflect.TypeOf(x)
	fmt.Println(i)

////
	test(1)
	test(x)
	test("abc")
}
