package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var currentId int

var tickets Tickets

// Give us some seed data
func init() {
	CreateTestTicket(Ticket{Name: "Test Ticket"})
}

func ZendeskFindTicket(id int) CommentArray {
	var baseurl = "https://coreos.zendesk.com/api/v2/"

	var username = os.Getenv("ZEN_USER")
	username += "/token"
	var token = os.Getenv("ZEN_TOKEN")
	var ticketid = id

	if username == "" {
		fmt.Println("You must supply the ZEN_USER environment variable")
		os.Exit(1)
	}
	if token == "" {
		fmt.Println("You must supply the ZEN_TOKEN environment variable")
		os.Exit(1)
	}
	ticketurl := baseurl + "tickets/" + strconv.Itoa(ticketid) + "/comments.json"
	//fmt.Printf("Requesting info on ticket url: %s\n", ticketurl)
	req, err := http.NewRequest("GET", ticketurl, nil)
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	t := CommentArray{}
	err = json.Unmarshal([]byte(body), &t)

	return t
	// return empty Ticket if not found
	//	if strings.Contains(t.Name, "InvalidEndpoint") {
	//		return Ticket{}
	//	} else {
	//		return t
	//	}
}

func ListTickets() ListOfTickets {
	var baseurl = "https://coreos.zendesk.com/api/v2/"

	var username = os.Getenv("ZEN_USER")
	username += "/token"
	var token = os.Getenv("ZEN_TOKEN")
	if username == "" {
		fmt.Println("You must supply the ZEN_USER environment variable")
		os.Exit(1)
	}
	if token == "" {
		fmt.Println("You must supply the ZEN_TOKEN environment variable")
		os.Exit(1)
	}
	ticketurl := baseurl + "tickets.json"
	//fmt.Printf("Requesting info on ticket url: %s\n", ticketurl)
	req, err := http.NewRequest("GET", ticketurl, nil)
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	t := ListOfTickets{}
	err = json.Unmarshal([]byte(body), &t)

	return t
}

func CreateTestTicket(t Ticket) Ticket {
	currentId += 1
	t.Id = currentId
	tickets = append(tickets, t)
	return t
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
