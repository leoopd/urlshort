package urlshort

import (
	"net/http"
)

var url string

func f1 (w http.ResponseWriter, r *http.Request) {
	url = r.Response.Body
}

func myHandler ()

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {


}

// func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

// }
