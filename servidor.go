package main

import "net/http"

type Country struct {
	Name     string
	Language string
}

//var countries []*Country = []*Country{}
var countries = []*Country{
	{Name: "si", Language: "si"},
	{Name: "no", Language: "no"},
}

func New(addr string) *http.Server {
	initRoutes()

	return &http.Server{
		Addr: addr,
	}
}
