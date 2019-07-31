package main

import (
	"flag"
	"fmt"
	"strconv"
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
	flag.StringVar(&download, "download", "-1", "Диапозон загружаемых композиций: '1-3' or '1' or '1-'")
}

//getRange возвращает диапозон значений 'min', 'max' из строки 's', либо 1 значение, если в 's' не содержится диапозон
func getRange(s string) (min, max int) {
	var (
		arrRange = []string{}
		err      error
	)
	if strings.Contains(s, "-") {
		arrRange = strings.Split(s, "-")
		min, err = strconv.Atoi(arrRange[0])
		if err != nil {
			min = 1
		}
		max, err = strconv.Atoi(arrRange[1])
		if err != nil {
			max = 1
		}
		if min < 1 {
			min = 1
		}
		if max < 1 {
			max = min
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

func main() {
	var min, max int
	flag.Parse()
	search := strings.Join(flag.Args(), "+")
	if download != "-1" {
		min, max = getRange(download)
	}
	fmt.Println(show, download, search, min, max)
}
