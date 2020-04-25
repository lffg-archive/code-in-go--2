package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range files {
			file, err := os.Open(filename)
			if err != nil {
				panic(err)
			}
			countLines(file, counts)
		}
	}

	if len(counts) != 0 {
		fmt.Println("==== Count Result ====")
		for line, count := range counts {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	stats, err := file.Stat()
	if err != nil {
		panic(err)
	}
	if stats.Size() == 0 {
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}
}
