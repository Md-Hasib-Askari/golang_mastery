package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var user User
var PORT = ":1234"
var DATA = make(map[string]string)

/*
defaultHandler
  - Simple handler that serves as a catch-all for any requests
    that don't match the other handlers. It logs the request
    and responds with a 404 Not Found status and a message thanking the visitor.
*/
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from ", r.Host)
	w.WriteHeader(http.StatusNotFound)

	body := "Thanks for visting\n"
	fmt.Fprintln(w, body)
}

/*
timeHandler
  - Handler that serves the current time. It logs the request,
    gets the current time formatted as RFC1123, and responds with a message containing the time.
*/
func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from ", r.Host)
	t := time.Now().Format(time.RFC1123)

	body := "The current time is: " + t + "\n"
	fmt.Fprintln(w, body)
}

/*
addHandler
  - Handler that allows adding a new user. It logs the request,
    checks that the request method is POST, reads the request body,
*/
func addHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from ", r.Host)

	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Error: Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error: Unable to read request body", http.StatusBadRequest)
		return
	}

	// Parse the JSON data into a User struct
	err = json.Unmarshal(d, &user)
	if err != nil {
		http.Error(w, "Error: Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate the user data
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Error: Missing username or password", http.StatusBadRequest)
		return
	}

	// Store the user data in the DATA map
	DATA[user.Username] = user.Password
	fmt.Fprintln(w, "User added successfully")
	w.WriteHeader(http.StatusCreated)
}

/*
getHandler
  - Handler that allows retrieving a user. It logs the request,
    checks that the request method is GET, reads the request body,
    parses it as JSON, and checks if the user exists in the DATA map.
  - If the user is found, it responds with a message containing the user data;
    otherwise, it responds with a 404 Not Found status and an error message.
*/
func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from ", r.Host)

	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Error: Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the request body and parse it as JSON
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error: Unable to read request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(d, &user)
	if err != nil {
		http.Error(w, "Error: Invalid JSON", http.StatusBadRequest)
		return
	}

	_, ok := DATA[user.Username]
	if !ok {
		http.Error(w, "Error: User not found", http.StatusNotFound)
		return
	}

	if ok && user.Username != "" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n", d)
	} else {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Error: User not found", http.StatusNotFound)
	}

}

/*
deleteHandler
  - Handler that allows deleting a user. It logs the request,
    checks that the request method is DELETE, reads the request body,
    parses it as JSON, and checks if the user exists in the DATA map.
  - If the user is found and the password matches, it deletes the user from the DATA map and responds with a success message;
    otherwise, it responds with an appropriate error message and status code.
*/
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, "from ", r.Host)

	// Only allow DELETE requests
	if r.Method != http.MethodDelete {
		http.Error(w, "Error: Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the request body and parse it as JSON
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error: Unable to read request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(d, &user)
	if err != nil {
		http.Error(w, "Error: Invalid JSON", http.StatusBadRequest)
		return
	}

	// Check if the user exists in the DATA map
	_, ok := DATA[user.Username]
	if !ok {
		http.Error(w, "Error: User not found", http.StatusNotFound)
		return
	}

	// Delete the user from the DATA map
	if ok && user.Username != "" {
		if user.Password == DATA[user.Username] {
			delete(DATA, user.Username)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "User deleted successfully")
		} else {
			http.Error(w, "Error: Incorrect password", http.StatusUnauthorized)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Error: User not found", http.StatusNotFound)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) > 1 {
		PORT = ":" + arguments[1]
	}

	// Create a new ServeMux to handle incoming HTTP requests
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	// Register the handlers for different routes
	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/add", addHandler)
	mux.HandleFunc("/get", getHandler)
	mux.HandleFunc("/delete", deleteHandler)
	mux.HandleFunc("/", defaultHandler)

	fmt.Println("Starting server on port", PORT)

	// Start the server and log any errors that occur
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}
