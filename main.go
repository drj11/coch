package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var count [256]int
	buf := make([]byte, 1)
	in := bufio.NewReader(os.Stdin)
	for {
		n, _ := in.Read(buf)
		if n == 0 {
			break
		}
		count[buf[0]] += 1
	}

	for i, c := range count {
		fmt.Printf("%d %d\n", i, c)
	}
}
