package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	router := newRouter()
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	server := newServer(router)

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}

}

func newRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func newServer(r *gin.Engine) *http.Server {
	return &http.Server{
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
