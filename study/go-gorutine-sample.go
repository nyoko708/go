package main

import (
	"log"
	"time"
)

func f(msg string) {
	log.Println(msg)
}

func main() {
	go f("hello")
	go func() {
		log.Println("gopher")
	}()

	time.Sleep(1 * time.Second)

}
