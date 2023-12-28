package main

import "fmt"

type Person struct {
	name   string
	age    int
	weight float64
	height float64
}

func (p *Person) GetPersonInfo() string {
	return fmt.Sprintf("[Name - %s, Age - %d, Weight - %v, Height - %v]", p.name, p.age, p.weight, p.height)
}

type PersonBuilder struct {
	name   string
	age    int
	weight float64
	height float64
}

func (p *PersonBuilder) SetName(name string) *PersonBuilder {
	p.name = name
	return p
}
func (p *PersonBuilder) SetAge(age int) *PersonBuilder {
	p.age = age
	return p
}
func (p *PersonBuilder) SetWeight(weight float64) *PersonBuilder {
	p.weight = weight
	return p
}
func (p *PersonBuilder) SetHeight(height float64) *PersonBuilder {
	p.height = height
	return p
}
func (p *PersonBuilder) Build() *Person {
	return &Person{
		name:   p.name,
		age:    p.age,
		weight: p.weight,
		height: p.height,
	}
}

func main() {
	p := &PersonBuilder{}
	person1 := p.SetName("Oleg").SetAge(35).SetWeight(78.9).SetHeight(183.5).Build()
	fmt.Println(person1.GetPersonInfo())
}
