package main

import "fmt"

type Duck interface {
	Walk() string
}

type Dog interface {
	Bark() string
}

type Robot struct {
	walkStr string
	barkStr string
}

func (r Robot) Walk() string {
	return r.walkStr
}

func (r Robot) Bark() string {
	return r.barkStr
}

func LetDuckWalk(duck Duck) string {
	return duck.Walk()
}

func LetDogBark(dog Dog) string {
	return dog.Bark()
}

func main() {

	robot := &Robot{"歩く", "吠える"}

	fmt.Println(
		LetDuckWalk(robot),
		LetDogBark(robot),
	)
}
