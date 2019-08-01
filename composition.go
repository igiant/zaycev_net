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

func (c *compositions) fixLength() {
	for i := range *c {
		if (*c)[i].artist == "" && (*c)[i].song == "" && (*c)[i].url == "" {
			*c = (*c)[:i]
			return
		}
	}
}
