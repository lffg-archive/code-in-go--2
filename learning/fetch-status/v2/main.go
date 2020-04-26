package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

func main() {
	urls := os.Args[1:]

	regex := regexp.MustCompile(`^https?://`)
	start := time.Now()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Time", "Bytes", "URL", "Error?"})

	for _, url := range urls {
		if !regex.MatchString(url) {
			url = "https://" + url
		}

		bytes, secs, err := getStatus(url)
		if err != nil {
			table.Append([]string{"0", "0", url, err.Error()})
			continue
		}

		fmt.Printf("Finished \"%s\".\n", url)
		table.Append([]string{fmt.Sprintf("%.2fs", secs), strconv.FormatInt(bytes, 10), url, ""})
	}

	fmt.Println()
	table.Render()
	fmt.Printf("%.2fs elapsed.\n", time.Since(start).Seconds())
}

func getStatus(url string) (bytes int64, secs float64, err error) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}

	bytes, err = io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		return 0, 0, err
	}

	return bytes, time.Since(start).Seconds(), nil
}
