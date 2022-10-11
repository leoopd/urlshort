package urlshort

import (
	"net/http"
)

var url string

func f1(w http.ResponseWriter, r *http.Request) {
	url = r.URL.Path
}

// func redirecter(p string) http.HandlerFunc {
// 	http.RedirectHandler(p, 301)
// }

// func defaultRedirecter(fallback http.Handler) http.HandlerFunc {
// 	http.RedirectHandler("/", 404)
// }

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {


	if path, ok := pathsToUrls[url]; !ok {

		redirecter := http.NewServeMux()
		redirecter.HandleFunc(url, f1) 
		return http.HandlerFunc(redirecter)
	} else {

defaultRedirecter := func (fallback http.Handler) http.HandlerFunc {
		http.RedirectHandler("/", 404)
	} (fallback)

		defaultHandler := http.HandlerFunc(defaultRedirecter(fallback))
		return defaultHandler
	}

}

// func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

// }
