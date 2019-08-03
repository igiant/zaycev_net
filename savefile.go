package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func saveCompositions(c compositions, min, max int) {
	resultChan := make(chan string, max-min+1)
	out := ""
	left := 0
	for i := min - 1; i < max; i++ {
		go saveFile(resultChan, c[i])
	}
	fmt.Printf("Скачивается %d %s:\n", max-min+1, enditive(max-min+1, "файл", "файла", "файлов"))
	for i := 0; i < max-min+1; i++ {
		out = <-resultChan
		left = max - min - i
		if left != 0 {
			fmt.Printf("%s (осталось: %d)\n", out, left)
		} else {
			fmt.Println(out)
		}
	}
	fmt.Println("Загрузки завершены!")
}

func saveFile(ch chan string, c composition) {
	filename := fmt.Sprintf("%s – %s.mp3", c.artist, c.song)
	for exists(filename) {
		filename += "_"
	}
	addr := createAddr(p.scheme, p.host, c.url, "", 0)
	fileAddr := getFileAddr(addr)
	if fileAddr == "" {
		ch <- "Ошибка при скачивании файла: " + filename
		return
	}
	body := getSiteBody(fileAddr)
	err := ioutil.WriteFile(filename, body, 0664)
	if err != nil {
		ch <- "Ошибка при сохранении файла: " + filename
		return
	}
	ch <- "'" + filename + "' сохранен..."
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
