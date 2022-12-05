package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}
func Readday5(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var stacks [8][]byte
	var ptr int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		for i := 0; i < 8; i++ {
			ptr = 4 * i
			if line[ptr] == '[' {
				letter := line[ptr+1]
				stacks[i] = append(stacks[i], letter)
			}
		}
	}

}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		panic("file open fail")
	}
	words, err := Readstrings(readFile)

	fmt.Println(ans)
	fmt.Println(ans2)

}
