package server

import (
	"net/http"
	"time"

	"gitlab.com/dc0d/go-workshop/interface_adapters/web/api"
)

// Start .
func Start() {
	router := api.NewRouter()

	s := newServer()
	router.Logger.Fatal(router.StartServer(s))
}

func newServer() *http.Server {
	return &http.Server{
		Addr:              ":8090",
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
}
