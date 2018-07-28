package main

import "fmt"

func f1() {
	i := 0
	defer fmt.Println(i) //实际上是将fmt.Println(0)加载到内存
	i++
	return
}

func f2() (i int) {
	var a int = 1
	defer func() {
		a++
		fmt.Println("defer内部", a)
	}()
	return a
}

func f3() (i int) {
	defer func() {
		i++
	}()
	return 1
}
func main() {
	f1()
	//fmt.Println("main中", f2())
	//fmt.Println(f3())
}