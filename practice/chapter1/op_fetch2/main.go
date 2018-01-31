package main

import (
	"os"
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

func main()  {
	for _, url := range os.Args[1:] {
		if ok := strings.HasPrefix(url, "http://"); !ok {
			url = "http://"+url;
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
