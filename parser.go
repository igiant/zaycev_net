package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

//parse - получение содержимого страницы по адресу
func parse(addr string) []byte {
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
	query.Set("query_search", search)
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

//getList возращает список композиций согласно запроса
func getList(query string, min, max int) []string {
	return []string{}
}
