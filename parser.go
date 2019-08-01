package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func parser(addr string) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", addr, nil)
	//if err != nil {
	//	return
	//}
	request.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	resp, _ := client.Do(request)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

//getList возращает список композиций согласно запроса
func getList(query string, min, max int) []string {
	return []string{}
}
