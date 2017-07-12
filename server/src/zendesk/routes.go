package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TicketsIndex",
		"GET",
		"/api/tickets",
		TicketsIndex,
	},
	Route{
		"TicketShow",
		"GET",
		"/api/tickets/{ticketId}",
		TicketShow,
	},
	//Route{
	//	"TodoCreate",
	//	"POST",
	//	"/todos",
	//	TodoCreate,
	//	},
}
