package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var visited = make(map[string]int)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Please specify start page and depth(example:~#go run crawl-channel.go http://golang.org 2)")
		os.Exit(1)
	}
	url := args[0]
	depth, _ := strconv.Atoi(args[1])

	ch := make(chan string)
	urlRegexp, _ := regexp.Compile(`http.*\"`)
	go func() {
		crawler(url, depth, urlRegexp, ch)
		close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
}

func crawler(url string, depth int, urlRegexp *regexp.Regexp, ch chan string) {
	if depth <= 0 {
		return
	}
	urls := retrieve_urls(url, urlRegexp)
	for _, url := range urls {
		ch <- fmt.Sprintf("%v", url)
		crawler(url, depth-1, urlRegexp, ch)
	}
}
func retrieve_urls(uri string, urlRegexp *regexp.Regexp) []string {
	urls := []string{}
	resp, err := http.Get(uri)
	if err != nil {
		return urls
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	lines := strings.Split(string(body), "\n")
	for _, v := range lines {
		matched, _ := regexp.MatchString("<a href=\"http", v)
		if matched {
			matchedurl := string(urlRegexp.Find([]byte(v)))
			matchedurl = fixUrl(matchedurl)
			visited[matchedurl]++
			if visited[matchedurl] > 1 {
				continue
			} else {
				urls = append(urls, matchedurl)
			}
		}
	}
	return urls
}

func fixUrl(url string) string {
	url = strings.Split(url, " ")[0]
	url = strings.Split(url, "#")[0]
	url = strings.Split(url, "\"")[0]
	return url
}
