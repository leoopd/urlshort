package urlshort

import (
	"errors"
	"net/http"

	yaml "github.com/go-yaml/yaml"
)

type yamlStruct struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// Looks for the path in the given map and returns a HandlerFunc
// that redirects to the specified url.
// Returns the fallback handler if the path is not defined.
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

// Parses the given yaml byte slice into a struct with the fields
// Path and Url
func parseYAML(yamlSlice []byte) ([]yamlStruct, error) {
	var err error
	var yamlParsed []yamlStruct

	if err := yaml.Unmarshal(yamlSlice, &yamlParsed); err != nil {
		errors.New("YAML could not be parsed")
	}

	return yamlParsed, err
}

// Converts the parsed yaml into a map with the path as key and
// url as the corresponding value
func makeMapFromParsedYAML(yaml []yamlStruct) (map[string]string, error) {
	var err error
	yamlMap := make(map[string]string)
	for i := 0; i < len(yaml); i++ {
		yamlMap[yaml[i].Path] = yaml[i].Url
	}

	if len(yamlMap) == 0 {
		err = errors.New("Map is empty")
	}
	return yamlMap, err
}

// Redirects the user if the given path is in the yaml and has a corresponding
// url. Returns a fallback if it doesn't.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var err error
	yamlStruct, err := parseYAML(yml)
	if err != nil {

		yamlHandler := func(w http.ResponseWriter, r *http.Request) {
			fallback.ServeHTTP(w, r)
		}
		return yamlHandler, err
	}
	yamlMap, err := makeMapFromParsedYAML(yamlStruct)
	if err != nil {

		yamlHandler := func(w http.ResponseWriter, r *http.Request) {
			fallback.ServeHTTP(w, r)
		}
		return yamlHandler, err
	}
	yamlHandler := MapHandler(yamlMap, fallback)
	return yamlHandler, err
}
