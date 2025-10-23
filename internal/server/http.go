package server

import (
	"fmt"
	"net/http"
	"os"
)

type RegisterFn func(router *http.ServeMux)

func RunServer(register RegisterFn) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	RunServerOnAddr(addr, register)
}

func RunServerOnAddr(addr string, register RegisterFn) {
	router := http.NewServeMux()
	register(router)

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	server.ListenAndServe()
}
