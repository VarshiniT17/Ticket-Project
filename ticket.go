package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

//////////////////////////////////////////////////////
// STRUCTS
//////////////////////////////////////////////////////

type Ticket struct {
	TicketID     int
	TicketNumber int
	Name         string
	Description  string
	Category     string
	AssignedTo   string
	Status       string
	CreatedAt    string
}

type Admin struct {
	Name     string
	Category string
}

//////////////////////////////////////////////////////
// GLOBAL VARIABLES
//////////////////////////////////////////////////////

var tickets []Ticket
var ticketCounter int
var categoryTicketNumbers = make(map[string]map[int]bool)

var admins = []Admin{
	{"Alice", "IT"},
	{"Bob", "HR"},
	{"Charlie", "Finance"},
}

//////////////////////////////////////////////////////
// LOGIC FUNCTIONS
//////////////////////////////////////////////////////

func generateTicketNumber(category string) int {
	for {
		num := rand.Intn(9000) + 1000

		if categoryTicketNumbers[category] == nil {
			categoryTicketNumbers[category] = make(map[int]bool)
		}

		if !categoryTicketNumbers[category][num] {
			categoryTicketNumbers[category][num] = true
			return num
		}
	}
}

func assignAdmin(category string) string {
	for _, admin := range admins {
		if strings.EqualFold(admin.Category, category) {
			return admin.Name
		}
	}
	return "No Admin Found"
}

func isValidCategory(category string) bool {
	for _, admin := range admins {
		if strings.EqualFold(admin.Category, category) {
			return true
		}
	}
	return false
}

//////////////////////////////////////////////////////
// HTTP HELPERS
//////////////////////////////////////////////////////

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}

//////////////////////////////////////////////////////
// HANDLERS
//////////////////////////////////////////////////////

func createTicketHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method != "POST" {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var input Ticket
	json.NewDecoder(r.Body).Decode(&input)

	if input.Name == "" || input.Description == "" || !isValidCategory(input.Category) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	category := strings.ToUpper(input.Category)

	ticketCounter++

	ticket := Ticket{
		TicketID:     ticketCounter,
		TicketNumber: generateTicketNumber(category),
		Name:         input.Name,
		Description:  input.Description,
		Category:     category,
		AssignedTo:   assignAdmin(category),
		Status:       "Open",
		CreatedAt:    time.Now().Format("02-01-2006 15:04:05"),
	}

	tickets = append(tickets, ticket)

	json.NewEncoder(w).Encode(ticket)
}

func getTicketsHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	json.NewEncoder(w).Encode(tickets)
}

//////////////////////////////////////////////////////
// MAIN
//////////////////////////////////////////////////////

func main() {
	rand.Seed(time.Now().UnixNano())

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// APIs
	http.HandleFunc("/api/create", createTicketHandler)
	http.HandleFunc("/api/tickets", getTicketsHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
