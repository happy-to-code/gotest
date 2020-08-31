package main

import (
	"github.com/blok-service/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1048576
	}
	s.ListenAndServe()
}
