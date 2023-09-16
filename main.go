package main

import (
	"fmt"
	"net/http"

	"github.com/zohaibsoomro/go-load-balancer/models"

)

func main() {
	//create a balancer in golang
	lb := models.NewLoadBalancer(3000)
	lb.Add(
		*models.NewServer("https://www.bing.com"),
		*models.NewServer("https://www.duckduckgo.com"),
		*models.NewServer("https://www.google.com"),
		*models.NewServer("https://www.youtube.com"),
	)
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.ServerProxy(w, r)
	}
	http.HandleFunc("/", handleRedirect)
	http.ListenAndServe(":"+fmt.Sprint(lb.Port), nil)
}
