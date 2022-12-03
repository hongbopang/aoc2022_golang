package main

import (
	"fmt"
)

func main() {
	memmap := make(map[int]int)

	for i := 0; i <= 10; i++ {
		if i == 7 {
			memmap := make(map[int]int)
			memmap[1] = 1
		}
		memmap[i] = 2
		fmt.Println(memmap)
	}

}
