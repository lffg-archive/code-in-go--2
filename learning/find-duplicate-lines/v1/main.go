package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		if isTerm() && text == ".exit" {
			break
		}
		counts[text]++
	}

	fmt.Println("==== Count Result ====")
	for line, count := range counts {
		fmt.Printf("%d\t%s\n", count, line)
	}
}

func isTerm() bool {
	stats, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	return ((stats.Mode() & os.ModeCharDevice) == os.ModeCharDevice)
}
