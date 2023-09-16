package models

import (
	"fmt"
	"net/http"

)

type LoadBalancer struct {
	Port            int
	Servers         []Server
	RoundRobinCount int
}

func NewLoadBalancer(port int) *LoadBalancer {
	return &LoadBalancer{
		Port: port,
	}
}

func (lb *LoadBalancer) Add(servers ...Server) {
	lb.Servers = append(lb.Servers, servers...)
}

func (lb *LoadBalancer) GetNextAvailableServer() *Server {
	server := lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	for !server.IsAlive() {
		lb.RoundRobinCount++
		server = lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	}
	lb.RoundRobinCount++
	return &server
}

func (lb *LoadBalancer) ServerProxy(w http.ResponseWriter, r *http.Request) {
	server := lb.GetNextAvailableServer()
	fmt.Println("Redirecting request to server address:", server.Address)
	server.Serve(w, r)
}
