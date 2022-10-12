package urlshort

import (
	"fmt"
	"net/http"

	yaml "github.com/go-yaml/yaml"
)

type yamlStruct struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	mapHandler := func(w http.ResponseWriter, r *http.Request) {
		if path, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, path, http.StatusTemporaryRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}

	}
	return mapHandler
}

func parseYAML(yamlSlice []byte) yamlStruct {
	var yamlParsed yamlStruct
	
	if err := yaml.Unmarshal(yamlSlice, &yamlParsed); err != nil {
		log.Fatal(err)
	}

	return yamlStruct
}

func makeMapFromParsedYAML(yaml yamlStruct) map[string]string {
yamlMap := make(map[string]string)
for i := 0; i < len()
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	yamlStruct := parseYAML(yml)
	yamlMap := makemakeMapFromParsedYAML(yamlStruct)

	
}
