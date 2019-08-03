package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)



//getRange возвращает диапозон значений 'min', 'max' из строки 's', либо 1 значение, если в 's' не содержится диапозон
func getRange(s string) (min, max int) {
	var err error
	if strings.Contains(s, "-") {
		arrRange := strings.Split(s, "-")
		min, err = strconv.Atoi(arrRange[0])
		if err != nil {
			min = 1
		}
		if min < 1 {
			min = 1
		}
		if arrRange[1] == "" {
			max = -1
		} else {
			max, err = strconv.Atoi(arrRange[1])
			if err != nil {
				max = 1
			}
			if max < 1 {
				max = min
			}
		}
		if max < min && max != -1 {
			min, max = max, min
		}
	} else {
		min, err = strconv.Atoi(s)
		if err != nil {
			min = 1
		}
		if min < 1 {
			min = 1
		}
		max = min
	}
	return min, max
}

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
	if len(os.Args) == 1 {
		show = true
	}
	var min, max int
	flag.Parse()
	search := strings.Join(flag.Args(), "+")
	if download != "-1" {
		min, max = getRange(download)
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
		if max == -1 {
			max = len(songs)
		}
		saveCompositions(songs, min, max)
	}
}
