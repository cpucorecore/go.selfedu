package main

import "fmt"

type I interface {
	test()
}

type I2 interface {
	I2Test1()
	I2Test2()
}

type Impl1 int
type Impl2 int

// implementation of I
func (this *Impl1) test() {
	fmt.Println("test1")
	*this = 1
}

// implementation of I2
func (this *Impl2) test() {
	fmt.Println("test2")
	*this = 2
}

func (this *Impl2) hello(n int) {
	fmt.Println("hello")
	*this += Impl2(n)
}

func (this Impl2) I2Test1() {
	fmt.Println("I2Test1")
}

func (this Impl2) I2Test2() {
	fmt.Println("I2Test2")
}

//common function
func common(i I2) {
	i.I2Test1()
}

func main() {
	var t Impl1
	t.test()
	fmt.Println(t)

	var t2 Impl2
	t2.test()
	fmt.Println(t2)
	t2.hello(100)
	fmt.Println(t2)

	t2.I2Test1()

	common(t2)
}

