package main

import (
	"fmt"
	"os"
)

func main() {
	var count [256]int
	buf := make([]byte, 1)
	for {
		n, _ := os.Stdin.Read(buf)
		if n == 0 {
			break
		}
		count[buf[0]] += 1
	}

	for i, c := range count {
		fmt.Printf("%d %d\n", i, c)
	}
}
