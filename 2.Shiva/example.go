/**
 * Accepted
 * 46 ms	100 KB
 * Jul/18/2022 17:22UTC+8
 */
package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanln(&t)

	for test := 0; test < t; test++ {
		var n, a, i int
		fmt.Scanln(&n)

		fmt.Scan(&a)
		odd := a % 2
		fmt.Scan(&a)
		even := a % 2

		flag := false
		for i = 2; i < n; i++ {
			fmt.Scan(&a)
			if flag {
				continue
			}

			if (i%2 == 0 && a%2 != odd) || (i%2 == 1 && a%2 != even) {
				fmt.Println("NO")
				flag = true
			}
		}

		if !flag {
			fmt.Println("YES")
		}

		fmt.Scanf("\n")
	}
}
