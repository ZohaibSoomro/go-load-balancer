package models

import (
	"net/http"
	"net/http/httputil"
	"net/url"

)

type Server struct {
	Address string
	proxy   *httputil.ReverseProxy
}

func NewServer(address string) *Server {
	url, _ := url.Parse(address)
	return &Server{
		Address: address,
		proxy:   httputil.NewSingleHostReverseProxy(url),
	}
}

func (s *Server) IsAlive() bool {
	return true
}

func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}
