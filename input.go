package main

import (
	"flag"
	"os"
	"path/filepath"

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
