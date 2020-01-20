package main

import "fmt"

func main() {
	var n, h int
	fmt.Scan(&n)
	fmt.Scan(&h)

	var a int
	w := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a > h {
			w += 2
		} else {
			w++
		}
	}
	fmt.Print(w)
}