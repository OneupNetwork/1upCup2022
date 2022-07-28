/**
 * Accepted
 * 46 ms	6500 KB
 * Jul/21/2022 14:07UTC+8
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 13*200000)
	scanner.Buffer(buf, 13*200000)

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

		n, _ := strconv.Atoi(input)

		if scanner.Scan() {
			input = scanner.Text()
		}

		strList = strings.Split(input, " ")

		countList := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for i := 0; i < n; i++ {
			a, _ := strconv.Atoi(strList[i])
			countList[a%10]++
		}

		var nums []int
		for i, count := range countList {
			for j := 0; j < count; j++ {
				nums = append(nums, i)
				if j == 3 {
					break
				}
			}
		}

		size := len(nums)

		flag := false
		for i := 0; i < size; i++ {
			for j := i + 1; j < size; j++ {
				for k := j + 1; k < size; k++ {
					if (nums[i]+nums[j]+nums[k])%10 == 3 {
						flag = true
						break
					}
				}
			}
		}

		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
