package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []int
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			result = append(result, 0)
		} else {
			x, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}
			result = append(result, x)
		}
	}
	result = append(result, 0)
	return result, scanner.Err()
}

func Readstrings(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		panic("file open fail")
	}
	words, err := Readstrings(readFile)

	ans := 0
	ans2 := 0

	for _, word := range words {
		tags := strings.Split(word, ",")
		first := strings.Split(tags[0], "-")
		second := strings.Split(tags[1], "-")

		start_1, _ := strconv.Atoi(first[0])
		end_1, _ := strconv.Atoi(first[1])

		start_2, _ := strconv.Atoi(second[0])
		end_2, _ := strconv.Atoi(second[1])

		if (start_1 <= start_2 && end_1 >= end_2) || (start_1 >= start_2 && end_1 <= end_2) {
			ans++
		}

		ans2++

		if start_1 > end_2 || start_2 > end_1 {
			ans2--
		}

	}

	fmt.Println(ans)
	fmt.Println(ans2)

}
