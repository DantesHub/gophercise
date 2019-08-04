package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type storyHandler struct {
	storyChapters map[string]storyChapter
	template  *template.Template
}


type storyChapter struct {
	Title string
	Story []string
	Options []storyOption
}

type storyOption struct {
	Text string
	Arc  string
}

func (sh storyHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	var path string
	if req.URL.Path == "/" {
		path = "intro"
	} else {
		path = strings.TrimLeft(req.URL.Path, "/")
	}
	storyChapter := sh.storyChapters[path]
	sh.template.Execute(resp, storyChapter)
}

func main() {
	f,err := os.Open("gopher.json")
	if err != nil {
		panic(err)
	}

	buffer := new(bytes.Buffer)
	_ , err = buffer.ReadFrom(f)
	if err != nil {
		panic(err)
	}

	var storyChapters map[string]storyChapter
	err = json.Unmarshal(buffer.Bytes(), &storyChapters)
	if err != nil {
		panic(err)
	}

	temp, err := template.ParseFiles("storytemplate.html")
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", storyHandler{storyChapters,temp})

}
