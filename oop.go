package main

import "fmt"

func main()  {

	man1 := Man{person: Person{name: "肖战",age: 100},school: "幼儿园"}
	fmt.Println(man1)
	fmt.Println(man1.person.name)

	woman1 := Woman{
		Person: Person{name: "lady-gaga",age:10000},
		school: "小学",
	}
	fmt.Println(woman1)
	fmt.Println(woman1.name)

	fmt.Println(woman1.changeName())
	fmt.Println(woman1)
	fmt.Println(woman1.changeAge())
	fmt.Println(woman1)
}

type Person struct {
	name string
	age int

}

type Man struct {
	person Person
	school string
}

type Woman struct {
	Person
	school string
}

func (p Person) changeName() string {
	p.name = "qqq"
	return p.name
}

func (p *Person) changeAge() int {
	p.age = p.age + 1
	return p.age
}