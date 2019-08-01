package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strconv"
)

//parse - получение содержимого страницы по адресу
func getSiteBody(addr string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return nil
	}
	request.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() { _ = resp.Body.Close() }()
	if err != nil || resp.StatusCode != 200 {
		return nil
	}
	return body
}

//createAddr - формирует адрес с параметрами get-запроса
func createAddr(scheme, host, path, search string, page int) string {
	query := url.Values{}
	if search != "" {
		query.Set("query_search", search)
	}
	//TODO сделать парсинг с пагинацией
	if page > 1 {
		query.Add("page", strconv.Itoa(page))
	}
	u := &url.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     path,
		RawQuery: query.Encode(),
	}
	return u.String()
}

func getComposition(body []byte, selector string) []composition {
	reader := bytes.NewReader(body)
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil
	}
	result := make(compositions, 100)
	document.Find(selector).Each(func(listIndex int, list *goquery.Selection) {
		list.Find(p.artist).Each(func(itemIndex int, item *goquery.Selection) {
			result[itemIndex].artist = strings.TrimSpace(item.Text())
		})
		list.Find(p.song).Each(func(itemIndex int, item *goquery.Selection) {
			result[itemIndex].song = strings.TrimSpace(item.Text())
		})
		list.Find(p.duration).Each(func(itemIndex int, item *goquery.Selection) {
			result[itemIndex].duration = strings.TrimSpace(item.Text())
		})
		list.Find(p.url).Each(func(itemIndex int, item *goquery.Selection) {
			result[itemIndex].url = item.AttrOr(p.data, "")
		})

	})
	return result
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

func getFileAddr(addr string) string {
	body := getSiteBody(addr)
	fileAddr := FileUrl{}
	err := json.Unmarshal(body, &fileAddr)
	if err != nil {
		return ""
	}
	return fileAddr.Url
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
	ch <- filename + " сохранен,"
}

//getList возращает список композиций согласно запроса
func saveCompositions(c compositions, min, max int) {
	resultChan := make(chan string, max-min+1)
	for i := min - 1; i < max; i++ {
		go saveFile(resultChan, c[i])
	}
	for i := 0; i < max-min+1; i++ {
		fmt.Println(<-resultChan)
	}
	fmt.Println("Загрузки завершены!")
}
