package main

import (
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"strconv"
	"strings"
)

var (
	show     bool
	download string
	p        params
)

func init() {
	flag.BoolVar(&show, "l", false, "Показать список композиций")
	flag.BoolVar(&show, "s", false, "Показать список композиций")
	flag.BoolVar(&show, "list", false, "Показать список композиций")
	flag.BoolVar(&show, "show", false, "Показать список композиций")
	flag.StringVar(&download, "d", "-1", "Диапазон загружаемых композиций")
	flag.StringVar(&download, "download", "-1", "Диапазон загружаемых композиций")
	p = readINI("zaycev_net.ini")
}

func printHelp() {
	fmt.Println("\n Параметры запуска:\n")
	fmt.Println(" ––––––––––––––––––\n")

	fmt.Println("  '-l (-s, --show, --list) <параметры поиска>' - Показать список найденных композиций\n")
	fmt.Println("  '-d=<от>-<до> (--download=<от>-<до>) <параметры поиска>' - Загрузить песни в указанном диапазоне")

}

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

func trimSongs(c compositions) compositions {
	for i := range c {
		if c[i].artist == "" && c[i].song == "" && c[i].url == "" {
			return c[:i]
		}
	}
	return c
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

func readINI(filename string) params {
	result := params{}
	cfg, err := ini.Load(filename)
	if err != nil {
		return result
	}
	result.scheme = cfg.Section("url").Key("scheme").String()
	result.host = cfg.Section("url").Key("host").String()
	result.path = cfg.Section("url").Key("path").String()
	result.list = cfg.Section("selectors").Key("list").String()
	result.artist = cfg.Section("selectors").Key("artist").String()
	result.song = cfg.Section("selectors").Key("song").String()
	result.duration = cfg.Section("selectors").Key("duration").String()
	result.url = cfg.Section("selectors").Key("url").String()
	result.data = cfg.Section("selectors").Key("data").String()
	result.chank, err = cfg.Section("output").Key("chank").Int()
	if err != nil {
		result.chank = 20
	}
	return result
}

func main() {
	if len(os.Args) == 1 {
		printHelp()
		return
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
	songs = trimSongs(songs)
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
