//
// 指定した文字列を逆にする
//
package main

import "fmt"

//import "unicode/utf8"

const str string = "あいう　えお abc defg "

func ReverseString(str string) string {

	r := []rune(str)
	var reverse string
	for i := len(r) - 1; i >= 0; i-- {
		reverse += string(r[i : i+1])
	}

	return reverse
}

func main() {
	// 文字列を逆にする
	fmt.Println(ReverseString(str))
}
