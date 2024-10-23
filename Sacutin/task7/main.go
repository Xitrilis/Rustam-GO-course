package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int
	var c []string
	fmt.Println("Введите число")
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		var b []rune
		for j := 0; j < i; j++ {
			b = append(b, '*')
		}
		c = append(c, string(b))
	}
	d := strings.Join(c, "\n")
	fmt.Println(d)
}
