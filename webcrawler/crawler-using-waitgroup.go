package main

import (
	"flag"
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
)

var visited = make(map[string]bool)

func main() {
	var wg sync.WaitGroup
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Please specify start page and depth(example http://google.com 2)")
		os.Exit(1)
	}
	url := args[0]
	depth, _ := strconv.Atoi(args[1])
	wg.Add(1)
	retrieve(url, depth, &wg)
	wg.Wait()
}
func retrieve(uri string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 {
		fmt.Sprintf("<- Done with %v, depth 0.\n", uri)
		return
	}
	fmt.Println("fetching ", uri)
	resp, err := http.Get(uri)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)
	for _, link := range links {
		absolute := parseUrl(link, uri)
		if !visited[absolute] {
			wg.Add(1)
			go retrieve(absolute, depth-1, wg)
		}
	}
	return
}

func parseUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return " "
	}
	baseurl, err := url.Parse(base)
	if err != nil {
		return " "
	}
	uri = baseurl.ResolveReference(uri)
	return uri.String()
}
