package main

import (
	"fmt"
	"strings"
)

func main() {

	someText := `
	this is the first line
	this is the second line
	`
	reader := strings.NewReader(someText)

	// read a fix number of bytes
	var res []byte = make([]byte, 10)
	_, _ = reader.Read(res)
	fmt.Println("next -> ", string(res))
	_, _ = reader.Read(res)
	fmt.Println("next -> ", string(res))

}
