package main

import (
	"fmt"
	"github.com/ohhyel/stringutil_go"
)

func main() {
	fmt.Println("hello world!!")
	str := stringutil_go.Reverse("TEST REVERSE!!")
	fmt.Println(str)
}
