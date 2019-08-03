package main

import (
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/go-ini/ini"
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
	p = readINI(filepath.Join(filepath.Dir(os.Args[0]), "zaycev_net.ini"))
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

//getRange возвращает слайс упорядоченных номеров входящих в перечисляемый диапазон
func getRange(s string) []int {
	result := getSlice(s)
	if len(result) == 1 {
		return result
	}
	if !sliceIsSorted(result) {
		sort.Slice(result, func(a, b int) bool { return result[a] < result[b] })
	}
	return deduplInSortedSlace(result)
}

func getSlice(s string) []int {
	result := make([]int, 0)
	var err error
	if !strings.ContainsAny(s, "-,") {
		number, err := strconv.Atoi(s)
		if err != nil {
			result = append(result, 1)
			return result
		}
		if number < 1 {
			number = 1
		}
		result = append(result, number)
		return result
	}
	var nRange []int
	if strings.Contains(s, ",") {
		nRange, err = list(s)
		if err != nil {
			result = append(result, 1)
			return result
		}
		result = append(result, nRange...)
		return result
	}
	if strings.Contains(s, "-") {
		nRange, err = band(s)
		if err != nil {
			result = append(result, 1)
			return result
		}
		result = append(result, nRange...)
	}
	return result
}

func list(s string) ([]int, error) {
	result := make([]int, 0)
	for _, sets := range strings.Split(s, ",") {
		if !strings.Contains(sets, "-") {
			number, err := strconv.Atoi(sets)
			if err != nil {
				return result, err
			}
			if number < 1 {
				number = 1
			}
			result = append(result, number)
			continue
		}
		nRange, err := band(sets)
		if err != nil {
			return result, err
		}
		result = append(result, nRange...)
	}
	return result, nil
}

func band(s string) ([]int, error) {
	result := make([]int, 0)
	set := strings.Split(s, "-")
	from, err := strconv.Atoi(set[0])
	if err != nil {
		return result, err
	}
	to, err := strconv.Atoi(set[1])
	if err != nil {
		return result, err
	}
	for k := min(from, to); k <= max(from, to); k++ {
		result = append(result, k)
	}
	return result, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sliceIsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func deduplInSortedSlace(arr []int) []int {
	result := make([]int, 0, len(arr))
	result = append(result, arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			result = append(result, arr[i])
		}
	}
	return result
}
