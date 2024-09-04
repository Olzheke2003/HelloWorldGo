package api

import "net/http"

type api struct {
	addr   string
	router *http.ServeMux
}

func New(addr string, r *http.ServeMux) *api {
	return &api{
		addr:   addr,
		router: r,
	}
}
func (api *api) FillEndPoints() {
	api.router.HandleFunc("/api/hello", hello)
	api.router.HandleFunc("/api/goodbye", goodbye)
}

func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.addr, api.router)
}
