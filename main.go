package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	b[2] = 0
	fmt.Println(b)
	fmt.Println(a)
}
