package main

import (
    "fmt"
)


type Hero struct {
	Name string
	Ad int
	Level int
}

type SuperHero struct {
	Hero

	sex int
}

func (this Hero) show() {
	fmt.Println("Name = ", this.Name)
}

func (this *Hero) setName(name string) {
	this.Name = name
}

func (this *Hero) fly() {
	fmt.Println("Hero fly")
}

func (this *SuperHero) fly() {
	fmt.Println("SuperHero fly")
}

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)
}


func main() {

	fmt.Println("hello")

	hero := Hero { }

	hero.show()

	hero.setName("lisi")

	hero.show()

	hero.fly()

	superHero := SuperHero {
		Hero: Hero{
			Name: "lisi",
			Ad: 190,
			Level: 1,
		},
		sex: 1,
	}

	superHero.fly()
	superHero.show()
	myFunc(superHero)

	var a string
	a = "good"

	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}