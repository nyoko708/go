package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

type Cat struct {
	Animal
	ServantName string
}

func main() {

	c := Cat{}

	c.Name = "Kuro"
	fmt.Println(c.Name)

	a := Animal{"hoge", 5}
	fmt.Println(a.Name)
}
