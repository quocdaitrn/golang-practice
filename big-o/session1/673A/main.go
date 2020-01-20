package main

import "fmt"

func main() {
	var n int
	a, _ := fmt.Scan(n)
	fmt.Println(a)
	fmt.Println(n)
	start := 0
	next := 0
	minutes := 0
	for i := 0; i < n; i++ {
		b, _ := fmt.Scan(&next)
		fmt.Println(b)
		if next-start > 15 {
			minutes = start+15
			break
		}
		start = next
		minutes = start
		if i == n-1 {
			if 90-start > 15 {
				minutes = start+15
			} else {
				minutes = 90
			}
		}
	}
	fmt.Print(minutes)
}