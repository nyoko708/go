package main

import "fmt"
import "strconv"

type Car interface {
	run(int) string
	stop()
}

type MyCar struct {
	name  string
	speed int
}

func (u *MyCar) run(speed int) string {
	u.speed = speed
	return strconv.Itoa(speed) + "Kmで走ります"
}

func (u *MyCar) stop() {
	fmt.Println("停止します")
	u.speed = 0
}

func main() {

	myCar := &MyCar{name: "ジムニー", speed: 0}

	var i Car = myCar
	fmt.Println(i.run(50))
	i.stop()

}
