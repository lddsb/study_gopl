package main

import (
	"os"
	"time"
	"fmt"
	"strings"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	stop := time.Now()
	fmt.Println("for ... i time consuming:", stop.Sub(start))

	s, sep = "", ""
	start = time.Now()
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	stop = time.Now()
	fmt.Println("for ... range time consuming:", stop.Sub(start))

	start = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	stop = time.Now()
	fmt.Println("strings.Join time consuming:", stop.Sub(start))
}
