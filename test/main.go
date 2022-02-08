package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "asdtest"
	slice := strings.Split(str, ",")
	fmt.Println("slice: ", slice)
}
