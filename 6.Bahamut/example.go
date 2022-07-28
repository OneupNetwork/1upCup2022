/**
 * Accepted
 * 904 ms	5200 KB
 * Jul/22/2022 11:33UTC+8
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	fmt.Scanln(&t)

	for test := 0; test < t; test++ {
		var n, k int
		fmt.Scanf("%d %d\n", &n, &k)

		var str string
		fmt.Scanln(&str)

		s := []byte(str)

		f := make([]int, n)
		remainK := k
		for i := 0; i < n && remainK > 0; i++ {
			if (k%2 == 0 && s[i] == '0') || (k%2 == 1 && s[i] == '1') {
				f[i] = 1
				remainK--
			}
		}

		f[n-1] += remainK

		for i := 0; i < n; i++ {
			if (k-f[i])%2 == 1 {
				if s[i] == '1' {
					s[i] = '0'
				} else {
					s[i] = '1'
				}
			}
		}

		fmt.Println(string(s))

		output := fmt.Sprint(f)
		output = strings.Replace(output, "[", "", -1)
		output = strings.Replace(output, "]", "", -1)
		fmt.Println(output)
	}
}
