package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	regex := regexp.MustCompile(`^https?://`)

	fmt.Println("Time\tBytes\tURL")

	for _, url := range os.Args[1:] {
		if !regex.MatchString(url) {
			url = "https://" + url
		}

		start := time.Now()
		resp := fetch(url)
		bytes, _ := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		secs := time.Since(start).Seconds()

		fmt.Printf("%.2fs\t%d\t%s\n", secs, bytes, url)
	}
}

func fetch(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}
