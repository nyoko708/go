package main

import "fmt"
import "strings"
import "sort"
import "reflect"

func _Sort(str string) []string {

	str = strings.Trim(str, " ")
	str = strings.Replace(str, " ", "", -1)

	strs := strings.Split(str, "")
	sort.Strings(strs)

	return strs
}

func IsAnagram(str1 string, str2 string) bool {

	strs1 := _Sort(str1)
	strs2 := _Sort(str2)

	if reflect.DeepEqual(strs1, strs2) == false {
		return false
	}

	return true
}

func main() {
	var str1 string = "anagrams"
	var str2 string = "ars magna"
	// 文字列を逆にする
	fmt.Println(IsAnagram(str1, str2))
}
