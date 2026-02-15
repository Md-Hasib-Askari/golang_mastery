package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var u1 = User{"alice", "password123"}
var u2 = User{"bob", "securepassword"}
var u3 = User{"", "charliepass"}

const addEndPoint = "/add"
const getEndPoint = "/get"
const deleteEndPoint = "/delete"
const timeEndPoint = "/time"

/*
deleteEndpoint
- This function sends a DELETE request to the server to delete a user.
- It marshals the user data into JSON, creates an HTTP request, and sends it to the server.
- It then reads and prints the response status and body, and returns the response status code.
*/
func deleteEndpoint(user User) int {
	// Marshal the user data into JSON format
	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// Create request body
	u := bytes.NewReader(jsonData)

	// create HTTP request
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080"+deleteEndPoint, u)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// create HTTP client and send the request
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println("Response status: ", resp.Status)
	println("Response body: ", string(data))

	return resp.StatusCode
}

/*
getEndpoint
- This function sends a GET request to the server to retrieve user data.
- It marshals the user data into JSON, creates an HTTP request, and sends it to the server.
- It then reads and prints the response status and body, and returns the response status code.
*/
func getEndpoint(user User) int {
	userMarshall, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	u := bytes.NewReader(userMarshall)

	// create HTTP request
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+getEndPoint, u)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// create HTTP client and send the request
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println("Response status: ", resp.Status)
	println("Response body: ", string(data))

	return resp.StatusCode
}

/*
addEndpoint
- This function sends a POST request to the server to add a new user.
- It marshals the user data into JSON, creates an HTTP request, and sends it to the server.
- It then reads and prints the response status and body, and returns the response status code.
*/
func addEndpoint(user User) int {
	userMarshall, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	u := bytes.NewReader(userMarshall)

	// create HTTP request
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080"+addEndPoint, u)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// create HTTP client and send the request
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println("Response status: ", resp.Status)
	println("Response body: ", string(data))

	return resp.StatusCode
}

func timeEndpoint() int {
	// create HTTP request
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+timeEndPoint, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// create HTTP client and send the request
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println("Response status: ", resp.Status)
	println("Response body: ", string(data))

	return resp.StatusCode
}

func main() {
	if len(os.Args) != 2 {
		println("Usage: go run 2_client.go <endpoint>")
		println("Endpoints: add, get, delete, time")
		return
	}

	deleteEndpoint(u1)
	getEndpoint(u2)
}
