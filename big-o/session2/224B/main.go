package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// =========== Better solution =========

func main() {
	var n int64
	var k int64
	var a []string
 
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadString('\n')
	tmp := strings.Split(strings.TrimSpace(line1), " ")
	n, _ = strconv.ParseInt(string(tmp[0]), 10, 64)
	k, _ = strconv.ParseInt(string(tmp[1]), 10, 64)
 
	line2, _ := reader.ReadString('\n')
	a = strings.Split(strings.TrimSpace(line2), " ")

	m := make(map[string]int)
	count := 0

	l := 0
	for r := int64(0); r < n; r++ {
		_, ok := m[a[r]]
		if !ok {
			count++
			m[a[r]] = 1
		} else {
			m[a[r]]++
		}

		for int64(count) == k {
			m[a[l]]--
			if m[a[l]] == 0 {
				fmt.Printf("%d %d", l+1, r+1)
				return
			}
			l++
		}
	}
	fmt.Printf("-1 -1")
}

// ========= Nice solution :)) =========

// func main() {
// 	var n, k int
// 	fmt.Scan(&n)
// 	fmt.Scan(&k)

// 	if k > n {
// 		fmt.Printf("-1 -1")
// 		return
// 	}

// 	var a []int
// 	var ai int
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&ai)
// 		a = append(a, ai)
// 	}

// 	l := k
// 	for ; l < n;{
// 		for i := 0; i < n-l+1; i++ {
// 			tmp := a[i:i+l]
// 			m := make(map[int]int)
// 			for j := 0; j < len(tmp); j++ {
// 				m[tmp[j]]++
// 			}
// 			if len(m) == k {
// 				fmt.Printf("%v %v", i+1, i+l)
// 				return
// 			}
// 		}
// 		l++
// 	}
// 	m := make(map[int]int)
// 	for i := 0; i < n; i++ {
// 		m[a[i]]++
// 	}
// 	if len(m) == k {
// 		fmt.Printf("%v %v", 1, n)
// 		return
// 	}

// 	fmt.Printf("-1 -1")
// }
