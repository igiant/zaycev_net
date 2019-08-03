package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func showList(c compositions) {
	var s string
	for i, song := range c {
		if (i+1) != len(c) && (i+1)%p.chank == 0 {
			fmt.Printf("%d. %s", i+1, song)
			_, _ = fmt.Scanln(&s)
			continue
		}
		fmt.Printf("%d. %s\n", i+1, song)
	}
}

func main() {
	flag.Parse()
	if len(os.Args) == 1 {
		show = true
	}
	search := strings.Join(flag.Args(), "+")
	var songsList []int
	if download != "-1" {
		songsList = getRange(download)
	}
	path := p.path
	if search == "" {
		path = ""
	}
	addr := createAddr(p.scheme, p.host, path, search, 1)
	body := getSiteBody(addr)
	songs := getComposition(body, p.list)
	if show {
		showList(songs)
	}
	if download != "-1" {
		saveCompositions(songs, songsList)
	}
}
