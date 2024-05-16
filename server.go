package main

import (
	"net/http"
	"fmt"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/cors"
	// chiMiddleware "github.com/go-chi/chi/v5/middleware"





)


type Server struct {
	cfg *Config.Config
	Router *chi.Mux
	httpServer *http.Server
}



func NewServer () *Server {

	return &Server{
		Router: chi.NewRouter(),
		cfg: chi.New()
	}
}


func  (s *Server) Run () {
	s.httpServer=&http.Server{
		Addr: s.cfg.Api.Host+':'+s.cfg.Api.Port,
		Handler: s.Router,
		ReadHeaderTimeOut:s.cfg.Api.ReadHeaderTimeout,
	}
	fmt.Println(`.......My server - Started!!...........`)
	fmt.Printf("......r - %v\n", s.cfg.Api.Environment)
	go func () {
		start(s)
	} ()
	
}





func start (s *Server) {
	log.printf("Serving at %s:%s\n",s.c)
	err:=s.httpServer.ListenAndServe()
	if err!=nil {
		lo.Fatal(err)
	}
}