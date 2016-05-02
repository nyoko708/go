package main

import "os"
import "fmt"

func main() {
	os.Mkdir("./tmp", 0777)

	for i := 1; i <= 10; i++ {
		os.MkdirAll(fmt.Sprintf("./tmp/%02d", i), os.ModePerm)
	}
}
