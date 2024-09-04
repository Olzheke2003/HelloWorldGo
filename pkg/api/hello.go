package api

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello, world! (Get)\n"))

	case http.MethodPost:
		w.Write([]byte("Hello, world! (Post)\n"))
	}

}
