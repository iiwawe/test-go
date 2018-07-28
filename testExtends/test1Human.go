package main

import "fmt"

func main() {
	var h Human

	s := Student{
		Grade: 1,
		Major: "English",
		Human: Human{
			Name: "Jason",
			Age:  12,
			Being: Being{
				IsLive: true}}}
	fmt.Println("student:", s)
	fmt.Println("student:", s.Name, ", isLive:", s.IsLive, ", age:", s.Age, ", grade:", s.Grade, ", major:", s.Major)

	//h = s // cannot use s (type Student) as type Human in assignment
	fmt.Println("------human------")
	fmt.Println(h)

	//Heal(s) // cannot use s (type Student) as type Being in argument to Heal
	fmt.Println("------heal------")
	Heal(s.Human.Being) // true

	s.Human.Drink()
	//s.Eat()
}

type Car struct {
	Color     string
	SeatCount int
}

type Being struct {
	IsLive bool
}

type Human struct {
	Being
	Name string
	Age  int
}

func (h Human) Eat() {
	fmt.Println("human eating...")
	h.Drink()
}

func (h Human) Drink() {
	fmt.Println("human drinking...")
}

func (h Human) Move() {
	fmt.Println("human moving...")
}

type Student struct {
	Human
	Grade int
	Major string
}

func (s Student) Drink() {
	fmt.Println("student drinking...")
}

type Teacher struct {
	Human
	School string
	Major  string
	Grade  int
	Salary int
}

func (s Teacher) Drink() {
	fmt.Println("teacher drinking...")
}

type IEat interface {
	Eat()
}

type IMove interface {
	Move()
}

type IDrink interface {
	Drink()
}

func Heal(b Being) {
	fmt.Println(b.IsLive)
}
