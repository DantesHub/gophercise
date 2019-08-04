package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

type YAMLURLMap struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

type protocoll struct {
	PathToYAML string
}


func main() {
	mux := defaultMux()
	protocoll := protocoll{}

	mapHandler := makeMapHandler(mux)
	yamlBytes := getYamlBytes(protocoll.PathToYAML)
	handler := makeYAMLHandler(yamlBytes, &mapHandler)


	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}




//MapHandler will return a http.HandleFunc that will map paths to there
//respective URLs if none is provided default handler will be used
func MapHandler(pathsToUrls map[string]string, defaultt http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		origURL, ok := pathsToUrls[req.URL.Path]
		if ok {
			http.Redirect(w,req,origURL,301)
		} else {
			defaultt.ServeHTTP(w,req)
		}
	}
}


//defaultMux is a functions that returns a pointer to the default
//server mux which directs user to the index page
func defaultMux() *http.ServeMux{
	mux := defaultMux()
	mux.HandleFunc("/", index)
	return mux
}

//YAMLHandler will take the YAML passed in and then return
//a http.HandlerFunc that will map paths to there URLs,
//if no path is provided default will be called instead
func YAMLHandler(YAML []byte, defaultt http.Handler) (http.HandlerFunc, error) {
	var maps [] YAMLURLMap
	err := yaml.Unmarshal(YAML, &maps)
	if err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, mapVar := range maps {
			if mapVar.Path == r.URL.Path {
				http.Redirect(w, r, mapVar.URL, 301)
				return
			}
		}
		defaultt.ServeHTTP(w,r)
	}, nil
}
//reads the YAML file and returns the file in bytes
func getYamlBytes(pathToFile string) []byte {
	bytes, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return nil
	}
	return bytes
}


func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the index page!")
}

//creates a mapHandler with 3 URL and paths
func makeMapHandler(mux *http.ServeMux) http.HandlerFunc {
	return MapHandler(map[string]string{
		"/handle-yaml":        "http://ghodss.com/2014/the-right-way-to-handle-yaml-in-golang/",
		"/love-golang":        "https://medium.com/@saginadir/why-i-love-golang-90085898b4f7",
		"/clean-architecture": "https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f",
	}, mux)
}

//creates a YAMLHandler with the passed in bytes
func makeYAMLHandler(yamlBytes []byte, defaultHandler *http.HandlerFunc) http.HandlerFunc {
	handler, err := YAMLHandler(yamlBytes, defaultHandler)
	if err != nil {
		panic(err)
	}
	return handler
}

