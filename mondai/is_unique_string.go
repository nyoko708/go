//
// 指定した文字列に重複があるかどうか判定する
//
package main

import "fmt"

const str string = "あいうえおabcdefg"

// 指定した文字列に重複があるかどうか判定する関数
// param string str
// return boolean
func IsUniqueString(str string) bool {

	uniqueFlags := make(map[string]bool)
	for _, c := range str {
		var val = string([]rune{c})
		var _, ok = uniqueFlags[val]
		if ok == true && uniqueFlags[val] == true {
			return false
		}
		uniqueFlags[val] = true
	}

	return true
}

// main 関数
func main() {
	// 判定
	fmt.Println(IsUniqueString(str))
}
