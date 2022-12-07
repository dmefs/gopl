package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	foundin := make(map[string][]string)

	files := os.Args[1:]
	for _, arg := range files {
		if f, err := os.Open(arg); err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		} else {
			countLines(f, counts, foundin)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%v\t%s\n", n, foundin[line], line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, foundin map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		for _, s := range foundin[line] {
			if s == f.Name() {
				return
			}
		}
		foundin[line] = append(foundin[line], f.Name())
	}
}
