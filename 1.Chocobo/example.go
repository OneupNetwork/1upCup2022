/**
 * Accepted
 * 15 ms	0 KB
 * Jul/21/2022 17:32UTC+8
 */
package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanln(&t)

	for test := 0; test < t; test++ {
		lines := make([]string, 8)

		fmt.Scanln()
		for i := 0; i < 8; i++ {
			fmt.Scanln(&lines[i])
		}

	out:
		for i := 1; i < 7; i++ {
			for j := 1; j < 7; j++ {
				if lines[i][j] == '#' && lines[i-1][j-1] == '#' &&
					lines[i-1][j+1] == '#' && lines[i+1][j-1] == '#' && lines[i+1][j+1] == '#' {
					fmt.Printf("%d %d\n", i+1, j+1)
					break out
				}
			}
		}
	}
}
