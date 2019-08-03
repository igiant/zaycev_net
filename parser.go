package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"

	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
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

// getComposition - получение списка композиций.
func getComposition(body []byte, selector string) []composition {
	reader := bytes.NewReader(body)
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil
	}
	result := make(compositions, 0)
	document.Find(selector).Each(func(listIndex int, list *goquery.Selection) {
		list.Find(p.artist).Each(func(itemIndex int, item *goquery.Selection) {
			song := composition{artist: strings.TrimSpace(item.Text())}
			result = append(result, song)
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

func getFileAddr(addr string) string {
	body := getSiteBody(addr)
	fileAddr := FileURL{}
	err := json.Unmarshal(body, &fileAddr)
	if err != nil {
		return ""
	}
	return fileAddr.SongURL
}

// enditive возвращает правильную форму существительного в зависимости от числа num
// form1 - соответствует 1 шт, form2 - соответствует от 2 до 4 шт, form3 - остальные
// например Enditive(119, "яблоко", "яблока", "яблок") вернет "яблок"
func enditive(num int, form1, form2, form3 string) string {
	mod := num % 100
	if mod >= 11 && mod <= 20 {
		return form3
	}
	mod = num % 10
	if mod == 0 || mod >= 5 {
		return form3
	}
	if mod == 1 {
		return form1
	}
	return form2
}
