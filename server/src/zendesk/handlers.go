package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func TicketsIndex(w http.ResponseWriter, r *http.Request) {
	t := ListTickets()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	/*for _, ticket := range t.Tickets {
		if err := json.NewEncoder(w).Encode(ticket); err != nil {
			panic(err)
		}
	}*/

	json := json.NewEncoder(w).Encode(t.Tickets)
	if json != nil {
		log.Fatal(json)
	}

	return
}

func TicketShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var ticketId int
	var err error
	if ticketId, err = strconv.Atoi(vars["ticketId"]); err != nil {
		panic(err)
	}
	ticket := ZendeskFindTicket(ticketId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//for i, _ := range ticket.Comments {
	if err := json.NewEncoder(w).Encode(ticket.Comments); err != nil {
		panic(err)
	}
	//	}
	return

	// If we didn't find it, 404
	//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//	w.WriteHeader(http.StatusNotFound)
	//	if err := json.NewEncoder(w).Encode(ticket); err != nil {
	//		panic(err)
	//	}

}

// vim: ts=4 sw=4 expandtab :
