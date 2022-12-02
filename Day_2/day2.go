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

	for _, word := range words {
		letters := strings.Split(word, " ")
		opponent := letters[0]
		switch letters[1] {
		case "X":
			ans += 1
			if opponent == "A" {
				ans += 3
			} else if opponent == "C" {
				ans += 6
			}
		case "Y":
			ans += 2
			if opponent == "B" {
				ans += 3
			} else if opponent == "A" {
				ans += 6
			}
		default:
			ans += 3
			if opponent == "C" {
				ans += 3
			} else if opponent == "B" {
				ans += 6
			}

		}
	}

	fmt.Println(ans)
	ans2 := 0
	for _, word := range words {
		letters := strings.Split(word, " ")
		opponent := letters[0]

		switch letters[1] {
		case "X":
			if opponent == "A" {
				ans2 += 3
			} else if opponent == "B" {
				ans2 += 1
			} else {
				ans2 += 2
			}
		case "Y":
			ans2 += 3
			if opponent == "A" {
				ans2 += 1
			} else if opponent == "B" {
				ans2 += 2
			} else {
				ans2 += 3
			}
		default:
			ans2 += 6
			if opponent == "A" {
				ans2 += 2
			} else if opponent == "B" {
				ans2 += 3
			} else {
				ans2 += 1
			}
		}
	}
	fmt.Println(ans2)

}
