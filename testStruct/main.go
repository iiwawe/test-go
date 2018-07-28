package main


type ss1 struct {
	name string
	ss2 *ss2
}

func NewSs1() {
	s1 := &ss1{
		name:"sss1",
	}
	s1.ss2 = NewSs2(s1)
}

type ss2 struct {
	name string
	ss1 *ss1
}

func NewSs2(s1 *ss1) *ss2  {
	s2 := &ss2{
		name: "sss2",
	}

	s2.ss1 = s1
	return s2
}

func main()  {
	NewSs1()
}