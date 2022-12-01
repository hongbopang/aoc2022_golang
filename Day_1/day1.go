package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
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
	scanner.Split(bufio.ScanWords)
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
	ints, err := ReadInts(readFile)

	ans := 0
	running := 0

	for _, val := range ints {

		if val == 0 {

			if running > ans {
				ans = running
			}
			running = 0
		} else {
			running += val
		}

	}

	fmt.Println(ans)

	var totals []int

	running = 0

	for _, val := range ints {
		if val == 0 {

			totals = append(totals, running)
			running = 0
		} else {
			running += val
		}
	}

	sort.Ints(totals)
	fmt.Println(totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3])

}
