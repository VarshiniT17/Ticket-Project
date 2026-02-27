package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
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

// Generate unique 4-digit ticket number per category
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

// Assign admin based on category
func assignAdmin(category string) string {
	for _, admin := range admins {
		if strings.EqualFold(admin.Category, category) {
			return admin.Name
		}
	}
	return "No Admin Found"
}

// Validate category
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

// CREATE TICKET
func createTicketHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method != "POST" {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var input Ticket
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

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

// TRACK TICKET BY ID
func getTicketByIDHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	// Expected URL: /api/ticket/id/{id}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	idStr := parts[4]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	for _, t := range tickets {
		if t.TicketID == id {
			json.NewEncoder(w).Encode(t)
			return
		}
	}

	http.Error(w, "Ticket not found", http.StatusNotFound)
}

//////////////////////////////////////////////////////
// MAIN
//////////////////////////////////////////////////////

func main() {
	rand.Seed(time.Now().UnixNano())

	// Serve static frontend files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// API Routes
	http.HandleFunc("/api/create", createTicketHandler)
	http.HandleFunc("/api/ticket/id/", getTicketByIDHandler)

	fmt.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}