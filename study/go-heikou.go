package main

import (
	"fmt"
	"time"
)

func repeat(say string, count int) {
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(say, fmt.Sprintf("%d回目", i))
	}
}
func main() {
	go repeat("平沢\t", 2) // 2秒かかるgoroutineをつくった
	go repeat("秋山\t", 5) // 5秒かかるgoroutineをつくった
	go repeat("琴吹\t", 3) // 3秒かかるgoroutineをつくった
	repeat("田井中\t", 6)   // 'main'というgoroutineが終わるのに6秒かかる
}
