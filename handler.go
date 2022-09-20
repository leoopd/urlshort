package urlshort

import (
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request, longUrl string) {
	http.Redirect(w, r, longUrl, http.StatusSeeOther)
}

func lookupUrl(pathsToUrls map[string]string, w http.ResponseWriter, r *http.Request) string {

	longUrl, ok := pathsToUrls[r.URL.Path]

	if ok {
		redirect(w, r, longUrl)
	} else {
		redirect(w, r, "/")
	}

	return "nil"
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	longUrl, ok := pathsToUrls[r.URL.Path]
	mux := http.NewServeMux()

	

}

// func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

// }
