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

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		panic("file open fail")
	}
	words, err := Readstrings(readFile)

	ans := 0
	var value int

	for _, word := range words {
		length := len(word)
		first := word[0 : length/2]
		second := word[length/2 : length]
		memmap := make(map[rune]int)

		for _, char := range first {
			asci := int(char)

			if asci >= 97 {
				value = asci - 96
			} else {
				value = asci - 64 + 26
			}
			memmap[char] = value
		}

		for _, char := range second {
			if val, ok := memmap[char]; ok {
				ans += val
				break
			}
		}
	}
	fmt.Println(ans)

	ptr := 0
	memmap2 := make(map[rune]int)
	ans2 := 0

	for _, word := range words {
		if ptr == 0 {
			for _, char := range word {
				memmap2[char] = 1
			}
		} else if ptr == 1 {
			for _, char := range word {
				if _, ok := memmap2[char]; ok {
					memmap2[char] = 2
				}
			}
		} else {
			for _, char := range word {
				if val, ok := memmap2[char]; ok && val == 2 {
					if int(char) >= 97 {
						ans2 += int(char) - 96
					} else {
						ans2 += int(char) - 64 + 26
					}
					for k := range memmap2 {
						delete(memmap2, k)
					}
					break
				}
			}
		}
		ptr = (ptr + 1) % 3
	}
	fmt.Println(ans2)
}
