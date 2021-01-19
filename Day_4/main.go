package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	if initialAge < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		p.age = 0
	}
	//Add some more code to run some checks on initialAge
	p.age = initialAge
	return p
}

func (p person) amIOld() {
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age < 18 && p.age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}
}

func (p person) yearPasses() person {
	p.age = p.age + 1
	return p
}

func main() {
	age := 10

	person1 := person{age: age}
	person1 = person1.NewPerson(age)
	person1.amIOld()
	person1 = person1.yearPasses()
	person1 = person1.yearPasses()
	person1 = person1.yearPasses()
	person1.amIOld()

}
