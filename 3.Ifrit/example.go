/**
 * Accepted
 * 218 ms	4300 KB
 * Jul/21/2022 15:48UTC+8
 */
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var input string
	var strList []string
	if scanner.Scan() {
		input = scanner.Text()
	}

	t, _ := strconv.Atoi(input)

	for test := 0; test < t; test++ {
		if scanner.Scan() {
			input = scanner.Text()
		}

		strList = strings.Split(input, " ")

		hC, _ := strconv.ParseFloat(strList[0], 64)
		dC, _ := strconv.ParseFloat(strList[1], 64)

		if scanner.Scan() {
			input = scanner.Text()
		}

		strList = strings.Split(input, " ")

		hM, _ := strconv.ParseFloat(strList[0], 64)
		dM, _ := strconv.ParseFloat(strList[1], 64)

		if scanner.Scan() {
			input = scanner.Text()
		}

		strList = strings.Split(input, " ")

		k, _ := strconv.ParseFloat(strList[0], 64)
		w, _ := strconv.ParseFloat(strList[1], 64)
		a, _ := strconv.ParseFloat(strList[2], 64)

		slay := false
		for coin := 0.0; coin <= k; coin++ {
			hCu := hC + coin*a
			dCu := dC + (k-coin)*w

			if math.Ceil(hM/dCu) <= math.Ceil(hCu/dM) {
				slay = true
				break
			}
		}

		if slay {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
