/**
 * Accepted
 * 592 ms	9300 KB
 * Jul/19/2022 15:46UTC+8
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 13*200000)
	scanner.Buffer(buf, 13*200000)

	var line string
	var strList []string
	if scanner.Scan() {
		line = scanner.Text()
	}

	t, _ := strconv.Atoi(line)

	for test := 0; test < t; test++ {
		if scanner.Scan() {
			line = scanner.Text()
		}

		strList = strings.Split(line, " ")
		n, _ := strconv.Atoi(strList[0])
		q, _ := strconv.Atoi(strList[1])

		if scanner.Scan() {
			line = scanner.Text()
		}

		strList = strings.Split(line, " ")

		list := make([]int64, n)
		for i := 0; i < n; i++ {
			list[i], _ = strconv.ParseInt(strList[i], 10, 64)
		}

		sort.Slice(list, func(i, j int) bool { return list[i] > list[j] })

		sum := make([]int64, n)
		for i := 0; i < n; i++ {
			var prev int
			if i == 0 {
				prev = 0
			} else {
				prev = i - 1
			}

			sum[i] = sum[prev] + list[i]
		}

		for i := 0; i < q; i++ {
			var quantityString string
			var quantity int64
			if scanner.Scan() {
				quantityString = scanner.Text()
				quantity, _ = strconv.ParseInt(quantityString, 10, 64)
			}

			candies := -1
			left := 1
			right := n
			for left <= right {
				mid := (left + right) / 2
				if sum[mid-1] >= quantity {
					candies = mid
					right = mid - 1
				} else {
					left = mid + 1
				}
			}

			fmt.Println(candies)
		}
	}
}
