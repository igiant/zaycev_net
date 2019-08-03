package main

import "fmt"

type composition struct {
	artist   string
	song     string
	duration string
	url      string
}

func (c composition) String() string {
	return fmt.Sprintf("%s – %s (%s)", c.artist, c.song, c.duration)
}

type compositions []composition

//FileURL - получение ссылки на композицию
type FileURL struct {
	SongURL string `json:"url"`
}

type params struct {
	scheme   string
	host     string
	path     string
	list     string
	artist   string
	song     string
	duration string
	url      string
	data     string
	chank    int
}

// enditive возвращает правильную форму существительного в зависимости от числа num
// form1 - соответствует 1 шт, form2 - соответствует от 2 до 4 шт, form3 - остальные
// например enditive(119, "песня", "песни", "песен") == "песен"
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
