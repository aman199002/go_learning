package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person{
	p := person{name: name}
	p.age = 42
	return &p
}

func main(){
	fmt.Println(person{"Bob",50})
	fmt.Println(person{name: "Allen",age: 40})
	fmt.Println(person{name: "Fred"})
	fmt.Println(newPerson("Jon"))
	
	s := person{name: "Jack", age: 30}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 31
	fmt.Println(sp.age)
}
