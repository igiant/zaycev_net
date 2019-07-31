package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	show     bool
	download string
)

func init() {
	flag.BoolVar(&show, "l", false, "Показать список найденных композиций")
	flag.BoolVar(&show, "s", false, "Показать список найденных композиций")
	flag.BoolVar(&show, "list", false, "Показать список найденных композиций")
	flag.BoolVar(&show, "show", false, "Показать список найденных композиций")
	flag.StringVar(&download, "d", "1", "Диапозон загружаемых композиций: '1-3' or '1' or '1-'")
	flag.StringVar(&download, "download", "1", "Диапозон загружаемых композиций: '1-3' or '1' or '1-'")
}

func main() {
	flag.Parse()
	search := strings.Join(flag.Args(), "+")
	fmt.Println(show, download, search)
}
