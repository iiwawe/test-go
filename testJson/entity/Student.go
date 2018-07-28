package testJson

type Student struct {
	Name     string
	Age      int
	Guake    bool
	Classses []string
	Price    float32
}

func NewStudent(name string, age int, guake bool, classses []string, price float32) *Student {
	st := &Student{
		name, age, guake, classses, price,
	}
	return st
}

