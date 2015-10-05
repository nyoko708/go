package main

import "fmt"

func replaceSpaces(str string) string {
	var rs string = ""
	for _, s := range str {
		if (string(s)) == " " {
			rs += "%20"
		} else {
			rs += string(s)
		}
	}
	return rs
}

func main() {
	var str string = "a b c d e f"
	fmt.Println(replaceSpaces(str))
}
