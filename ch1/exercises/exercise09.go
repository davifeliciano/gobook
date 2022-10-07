package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	protocolRegexp, _ := regexp.Compile("^https?://")
	for _, url := range os.Args[1:] {
		if !protocolRegexp.Match([]byte(url)) {
			fmt.Printf("%s: missing protocol in url. prepending http://\n", os.Args[0])
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		body, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: reading %s: %v\n", os.Args[0], url, err)
			os.Exit(1)
		}
		fmt.Printf("Returned status: %s\n", resp.Status)
		fmt.Printf("Body:\n\n%v", body)
	}
}
