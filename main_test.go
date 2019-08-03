package main

import "testing"

func TestGetRange(t *testing.T) {
	type Test struct {
		value string
		result []int
	}
	var arrTests = []Test{
		{"1", []int{1}},
		{"1-5", []int{1, 2, 3, 4, 5}},
		{"0", []int{1}},
		{"1,5,3-7", []int{1, 3, 4, 5, 6, 7}},
		{"5-", []int{1}},
		{"-3", []int{1}},
		{"5-1", []int{1, 2, 3, 4, 5}},
		{"-1-1", []int{1}},
		{"a-b", []int{1}},
	}
	var result []int
	for _, test := range arrTests {
		result = getRange(test.value)
		if !compareIntSlice(result, test.result) {
			t.Errorf("Ожидалось %v, получили %v = getRange(%s)", test.result, result, test.value)
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
		forms [3]string
		result string
	}
	var datas = []testData{
		{1, [3]string{"1", "2", "3"}, "1"},
		{5, [3]string{"1", "2", "3"}, "3"},
		{11, [3]string{"1", "2", "3"}, "3"},
		{2, [3]string{"1", "2", "3"}, "2"},
		{1001, [3]string{"1", "2", "3"}, "1"},
		{0, [3]string{"1", "2", "3"}, "3"},
		{21, [3]string{"1", "2", "3"}, "1"},
		{100000054, [3]string{"1", "2", "3"}, "2"},
	}
	var result string
	for _, data := range datas {
		result = enditive(data.count, data.forms[0], data.forms[1], data.forms[2])
		if result != data.result {
			t.Errorf("Ожидалось %s, получили %s = enditive(%d, %s, %s, %s)",
				data.result,
				result,
				data.count,
				data.forms[0],
				data.forms[1],
				data.forms[2],
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

func compareIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSliceIsSorted(t *testing.T) {
	type data struct {
		value []int
		result bool
	}
	var datas = []data{
		{[]int{1, 3, 5}, true},
		{[]int{1, 1, 1}, true},
		{[]int{1, 3, 2}, false},
		{[]int{3, 2, 1}, false},
		{[]int{2, 2, 3}, true},
	}
	for _, item := range datas {
		result := sliceIsSorted(item.value) 
		if result != item.result {
			t.Errorf("Ожидалось %v, получили %v при %v", item.result, result, item.value)
		}
	}
}