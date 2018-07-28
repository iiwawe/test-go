package main

import "fmt"

type Rect struct {
	width float64
	height float64
}

func (r *Rect) size() float64  {
	return r.width * r.height
}

func main() {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ",a)
	fmt.Println("&a = ",&a)
	fmt.Println("*&a = ",*&a)
	fmt.Println("b = ",b)
	fmt.Println("&b = ",&b)
	fmt.Println("*&b = ",*&b)
	fmt.Println("*b = ",*b)
	fmt.Println("c = ",c)
	fmt.Println("*c = ",*c)
	fmt.Println("&c = ",&c)
	fmt.Println("*&c = ",*&c)
	fmt.Println("**c = ",**c)
	fmt.Println("***&*&*&*&c = ",***&*&*&*&*&c)
	fmt.Println("x = ",x)
}