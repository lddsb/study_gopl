package main

import (
	"os"
	"fmt"
	"bufio"
	"time"
)

func main() {
	start := time.Now()
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, make(map[string]int))
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			printFileName(f, counts)
			f.Close()
		}
	}
	fmt.Println("time cosuming: ", time.Now().Sub(start))
}

func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func printFileName(f *os.File, counts map[string]int)  {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, f.Name())
		}
	}
}