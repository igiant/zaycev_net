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
