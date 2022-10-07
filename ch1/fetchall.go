// fetchall fetches all the urls given
// as command line arguments in parallel and
// compute the elapsed time for each one, as
// well as the total time
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

var protocolRegexp, _ = regexp.Compile("^https?://")

func main() {
	start_time := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Starts a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // Receives from channel ch
	}
	fmt.Printf("%.2fs elapsed in total\n", time.Since(start_time).Seconds())
}

func fetch(url string, ch chan<- string) {
	if !protocolRegexp.Match([]byte(url)) {
		url = "http://" + url
	}
	start_time := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	elapsed_secs := time.Since(start_time).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", elapsed_secs, nbytes, url)
}
