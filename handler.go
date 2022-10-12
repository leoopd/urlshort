package urlshort

import (
	"net/http"
)

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

// func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

// }
