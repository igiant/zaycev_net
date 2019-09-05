package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func saveCompositions(c compositions, arr []int) {
	resultChan := make(chan string, len(arr))
	out := ""
	left := 0
	for _, num := range arr {
		go saveFile(resultChan, c[num-1])
	}
	fmt.Printf("Скачивается %d %s:\n", len(arr), enditive(len(arr), "файл", "файла", "файлов"))
	for i := 0; i < len(arr); i++ {
		out = <-resultChan
		left = len(arr) - 1 - i
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
		filename = "_" + filename
	}
	addr := createAddr(p.scheme, p.host, c.url, "", 0)
	fileAddr := getFileAddr(addr)
	if fileAddr == "" {
		ch <- "Ошибка при скачивании файла: " + filename
		return
	}
	body := getSiteBody(fileAddr)
	err := ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		ch <- "Ошибка при сохранении файла: " + filename
		return
	}
	ch <- "'" + filename + "' сохранен..."
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
