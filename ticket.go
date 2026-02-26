package main

// Importing required standard libraries
import (
	"bufio"     // Used for buffered input (to read full lines including spaces)
	"fmt"       // Used for formatted I/O (printing to console)
	"math/rand" // Used to generate random numbers
	"os"        // Used to access OS-level features like stdin
	"strings"   // Used for string manipulation (TrimSpace, ToUpper, EqualFold)
	"time"      // Used for date and time operations
)

//////////////////////////////////////////////////////
//                STRUCT DEFINITIONS
//////////////////////////////////////////////////////

// Ticket struct represents a single ticket in the system.
// It groups all related ticket information together.
type Ticket struct {
	TicketID     int    // Auto-increment unique ID (starts from 1)
	TicketNumber int    // Random 4-digit number unique within category
	Name         string // Ticket name entered by user
	Description  string // Ticket description entered by user
	Category     string // Department (IT, HR, Finance)
	AssignedTo   string // Admin responsible for this category
	Status       string // Current ticket status (default: Open)
	CreatedAt    string // System-generated timestamp
}

// Admin struct represents an admin responsible for a category.
type Admin struct {
	Name     string // Admin name
	Category string // Category handled by admin
}

//////////////////////////////////////////////////////
//              GLOBAL VARIABLES
//////////////////////////////////////////////////////

// Slice to store all created tickets in memory.
// A slice is a dynamic array in Go.
var tickets []Ticket

// Counter to auto-generate Ticket IDs sequentially.
// Default value of int in Go is 0.
var ticketCounter int

// Nested map to ensure uniqueness of ticket numbers within each category.
// Structure:
// Category -> (TicketNumber -> true)
var categoryTicketNumbers = make(map[string]map[int]bool)

// Predefined admins mapped to categories.
var admins = []Admin{
	{"Alice", "IT"},
	{"Bob", "HR"},
	{"Charlie", "Finance"},
}

//////////////////////////////////////////////////////
//          FUNCTION: Generate Ticket Number
//////////////////////////////////////////////////////

// generateTicketNumber generates a random 4-digit number
// and ensures it is unique within the given category.
func generateTicketNumber(category string) int {

	for {
		// Generate random number between 1000–9999
		num := rand.Intn(9000) + 1000

		// If category map does not exist yet, initialize it.
		if categoryTicketNumbers[category] == nil {
			categoryTicketNumbers[category] = make(map[int]bool)
		}

		// Check if number already used in this category.
		// If not used, mark it as used and return it.
		if !categoryTicketNumbers[category][num] {
			categoryTicketNumbers[category][num] = true
			return num
		}

		// If number exists, loop continues and generates new number.
	}
}

//////////////////////////////////////////////////////
//          FUNCTION: Assign Admin
//////////////////////////////////////////////////////

// assignAdmin assigns an admin based on category.
// It performs case-insensitive comparison.
func assignAdmin(category string) string {

	for _, admin := range admins {

		// EqualFold compares strings ignoring case
		if strings.EqualFold(admin.Category, category) {
			return admin.Name
		}
	}

	// If no admin found for category
	return "No Admin Found"
}

//////////////////////////////////////////////////////
//          FUNCTION: Validate Category
//////////////////////////////////////////////////////

// isValidCategory checks whether entered category exists.
func isValidCategory(category string) bool {

	for _, admin := range admins {

		if strings.EqualFold(admin.Category, category) {
			return true
		}
	}

	return false
}

//////////////////////////////////////////////////////
//          FUNCTION: Create Ticket
//////////////////////////////////////////////////////

// createTicket handles complete ticket creation process.
func createTicket(reader *bufio.Reader) {

	// Read Ticket Name
	fmt.Print("Enter Ticket Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Remove newline and extra spaces

	// Validate Name
	if name == "" {
		fmt.Println("Ticket Name cannot be empty!")
		return
	}

	// Read Description
	fmt.Print("Enter Description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	// Validate Description
	if description == "" {
		fmt.Println("Description cannot be empty!")
		return
	}

	// Read Category
	fmt.Print("Enter Category (IT/HR/Finance): ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	// Validate Category
	if !isValidCategory(category) {
		fmt.Println("Invalid Category! Ticket not created.")
		return
	}

	// Convert category to uppercase for consistent storage
	category = strings.ToUpper(category)

	// Increment global counter to generate unique Ticket ID
	ticketCounter++

	// Get current system date & time formatted nicely
	currentTime := time.Now().Format("02-01-2006 15:04:05")

	// Create ticket struct instance
	ticket := Ticket{
		TicketID:     ticketCounter,
		TicketNumber: generateTicketNumber(category),
		Name:         name,
		Description:  description,
		Category:     category,
		AssignedTo:   assignAdmin(category),
		Status:       "Open", // Default status
		CreatedAt:    currentTime,
	}

	// Append ticket to slice (dynamic array)
	tickets = append(tickets, ticket)

	// Display success message
	fmt.Println("\n✅ Ticket Created Successfully!")
	fmt.Println("Ticket ID:", ticket.TicketID)
	fmt.Println("Ticket Number:", ticket.TicketNumber)
	fmt.Println("Assigned To:", ticket.AssignedTo)
	fmt.Println("Created At:", ticket.CreatedAt)
}

//////////////////////////////////////////////////////
//          FUNCTION: View Tickets
//////////////////////////////////////////////////////

// viewTickets prints all tickets stored in memory.
func viewTickets() {

	if len(tickets) == 0 {
		fmt.Println("No tickets found.")
		return
	}

	for _, t := range tickets {

		fmt.Println("\n-------------------------------")
		fmt.Println("Ticket ID     :", t.TicketID)
		fmt.Println("Ticket Number :", t.TicketNumber)
		fmt.Println("Name          :", t.Name)
		fmt.Println("Description   :", t.Description)
		fmt.Println("Category      :", t.Category)
		fmt.Println("Assigned To   :", t.AssignedTo)
		fmt.Println("Status        :", t.Status)
		fmt.Println("Created At    :", t.CreatedAt)
	}
}

//////////////////////////////////////////////////////
//                  MAIN FUNCTION
//////////////////////////////////////////////////////

// main is the entry point of the program.
func main() {

	// Seed random generator using current time.
	// Without this, random numbers repeat every run.
	rand.Seed(time.Now().UnixNano())

	// Create a buffered reader to read full-line input.
	reader := bufio.NewReader(os.Stdin)

	// Infinite loop to show menu repeatedly.
	for {

		fmt.Println("\n===== Ticket Management System =====")
		fmt.Println("1. Create Ticket")
		fmt.Println("2. View Tickets")
		fmt.Println("3. Exit")
		fmt.Print("Choose option: ")

		// Read user choice
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Switch case handles menu options
		switch input {

		case "1":
			createTicket(reader)

		case "2":
			viewTickets()

		case "3":
			fmt.Println("Exiting system...")
			return // Terminates program

		default:
			fmt.Println("Invalid choice! Please select 1, 2, or 3.")
		}
	}
}
