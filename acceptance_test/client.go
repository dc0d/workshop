package acceptance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type client struct {
	clientID string
}

func newClient(clientID string) *client {
	return &client{clientID: clientID}
}

func (c *client) deposit(amount int, t time.Time) {
	var command command
	command.Command = "DEPOSIT"
	command.Data.ClientID = c.clientID
	command.Data.Amount = amount
	command.Data.Time = t

	sendTransactionCommand(command)
}

func (c *client) withdraw(amount int, t time.Time) {
	var command command
	command.Command = "WITHDRAW"
	command.Data.ClientID = c.clientID
	command.Data.Amount = amount
	command.Data.Time = t

	sendTransactionCommand(command)
}

func (c *client) bankStatement() string {
	path := fmt.Sprintf("/api/bank/%v/statement", c.clientID)
	body, status, err := get(path)
	if err != nil {
		panic(err)
	}
	if status != http.StatusOK {
		panic(fmt.Sprintf("Expected status code %v but got %v", http.StatusText(http.StatusOK), http.StatusText(status)))
	}
	return body
}

func sendTransactionCommand(command command) {
	path := "/api/bank/transactions"

	_, status, err := post(path, command)
	if err != nil {
		panic(err)
	}

	expectedStatus := http.StatusOK
	if status != expectedStatus {
		panic(fmt.Sprintf("Expected status code %v but got %v", http.StatusText(expectedStatus), http.StatusText(status)))
	}
}

func post(path string, payload interface{}) (string, int, error) {
	js, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewBuffer(js)
	resp, err := http.Post(fmt.Sprintf("%v%v", serverAddr, path), "application/json", reader)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	return string(body), resp.StatusCode, nil
}

func get(path string) (string, int, error) {
	resp, err := http.Get(fmt.Sprintf("%v%v", serverAddr, path))
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	return string(body), resp.StatusCode, nil
}

type command struct {
	Command string `json:"command"`
	Data    struct {
		ClientID string    `json:"client_id"`
		Amount   int       `json:"amount"`
		Time     time.Time `json:"time"`
	} `json:"data"`
}

var (
	serverAddr = "http://localhost:8090"
)
