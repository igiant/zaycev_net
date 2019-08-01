package main

import (
	"bytes"
	"io/ioutil"
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
	result := make(compositions, 50)
	document.Find(selector).Each(func(listIndex int, list *goquery.Selection) {
		list.Find("div > div > div.musicset-track__artist > a").Each(func(itemIndex int, item *goquery.Selection) {
			result[itemIndex].artist = strings.TrimSpace(item.Text())
		})
		list.Find("div > div > div.musicset-track__track-name > a").Each(func(itemIndex int, item *goquery.Selection) {
			result[itemIndex].song = strings.TrimSpace(item.Text())
		})
		list.Find("div > div > div.musicset-track__duration").Each(func(itemIndex int, item *goquery.Selection) {
			result[itemIndex].duration = strings.TrimSpace(item.Text())
		})
		list.Find("div.musicset-track.clearfix").Each(func(itemIndex int, item *goquery.Selection) {
			result[itemIndex].url = item.AttrOr("data-url", "")
		})

	})
	return result
}

//getList возращает список композиций согласно запроса
func saveCompositions(c compositions, min, max int) {

}
