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

	all := false

	for {
		n, _ := in.Read(buf)
		if n == 0 {
			break
		}
		count[buf[0]] += 1
	}

	for i, c := range count {
		if c > 0 || all {
			fmt.Printf("%d %d\n", i, c)
		}
	}
}
