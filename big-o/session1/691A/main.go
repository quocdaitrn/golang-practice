package main

import(
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var a int
	var s int

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		s += a
	}
	if (n > 1 && s == n-1) || (n == 1 && s == 1) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}