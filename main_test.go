package main

import "testing"

func TestGetRange(t *testing.T) {
	type Test struct {
		value string
		min   int
		max   int
	}
	var arrTests = []Test{
		{"1", 1, 1},
		{"1-5", 1, 5},
		{"0", 1, 1},
		{"7", 7, 7},
		{"5-", 5, -1},
		{"-3", 1, 3},
		{"5-1", 1, 5},
		{"-1-1", 1, 1},
		{"a-b", 1, 1},
	}
	var resMin, resMax int
	for _, test := range arrTests {
		resMin, resMax = getRange(test.value)
		if resMin != test.min || resMax != test.max {
			t.Errorf("Ожидалось (%d, %d), получили (%d, %d) = getRange(%s)",
				test.min, test.max, resMin, resMax, test.value)
		}
	}
}

func TestCreateAddr(t *testing.T) {
	result := createAddr("http", "zaycev.net", "search.html", "one two three", -1)
	if result != "http://zaycev.net/search.html?query_search=one+two+three" {
		t.Errorf("Ожидалось %s, получили %s", result, "http://zaycev.net/search.html?query_search=one+two+three")
	}
	result = createAddr("http", "zaycev.net", "search.html", "один+два+три", 2)
	if result != "http://zaycev.net/search.html?page=2&query_search=%D0%BE%D0%B4%D0%B8%D0%BD%2B%D0%B4%D0%B2%D0%B0%2B%D1%82%D1%80%D0%B8" {
		t.Errorf("Ожидалось %s, получили %s", result,
			"http://zaycev.net/search.html?page=2&query_search=%D0%BE%D0%B4%D0%B8%D0%BD%2B%D0%B4%D0%B2%D0%B0%2B%D1%82%D1%80%D0%B8")
	}
	result = createAddr("https", "zaycev.net", "", "", 0)
	if result != "https://zaycev.net" {
		t.Errorf("Ожидалось %s, получили %s", result, "https://zaycev.net")
	}
}

func TestEnditive(t *testing.T) {
	type testData struct {
		count  int
		form1  string
		form2  string
		form3  string
		result string
	}
	var datas = []testData{
		{1, "1", "2", "3", "1"},
		{5, "1", "2", "3", "3"},
		{11, "1", "2", "3", "3"},
		{2, "1", "2", "3", "2"},
		{1001, "1", "2", "3", "1"},
		{0, "1", "2", "3", "3"},
		{21, "1", "2", "3", "1"},
		{100000054, "1", "2", "3", "2"},
	}
	var result string
	for _, data := range datas {
		result = enditive(data.count, data.form1, data.form2, data.form3)
		if result != data.result {
			t.Errorf("Ожидалось %s, получили %s = enditive(%d, %s, %s, %s)",
				data.result,
				result,
				data.count,
				data.form1,
				data.form2,
				data.form3,
			)
		}
	}
}

func TestReadINI(t *testing.T) {
	result := readINI("zaycev_net.ini")
	if result.scheme == "" || result.host == "" ||
		result.path == "" || result.list == "" ||
		result.artist == "" || result.song == "" ||
		result.duration == "" || result.url == "" ||
		result.data == "" || result.chank == 0 {
		t.Error("Неправильное чтение ini-файла")
	}
}

func TestGetSiteBody(t *testing.T) {
	p := readINI("zaycev_net.ini")
	addr := createAddr(p.scheme, p.host, "", "", 0)
	if getSiteBody(addr) == nil {
		t.Errorf("Проблема с получением содержимого сайта %s", addr)
	}
}
