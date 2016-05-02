package main

import (
	"io/ioutil"
	"os"
)

func main() {
	content := []byte("hello world\n")
	ioutil.WriteFile("./tmp/test.txt", content, os.ModePerm)
}
