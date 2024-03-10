package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Replace this baseURL to actual URL
const baseURL = "https://mock-node-wgqbnxruha-as.a.run.app"

type TransactionPayload struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type TranscationResponse struct {
	TxHash string `json:"tx_hash"`
}

type TranscationStatusResponse struct {
	TxStatus string `json:"tx_status"`
}

// Client module for broadcasting and monitoring transactions
type ClientModule struct {
	baseURL string
}

// Create a new client modlue
func NewClientModule() *ClientModule {
	return &ClientModule{baseURL: baseURL}
}

// Broadcasting Function
func (c *ClientModule) BroadcastTransaction(payload TransactionPayload) (string, error) {
	url := c.baseURL + "/broadcast"
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to broadcast transaction. status code %d", resp.StatusCode)
	}

	var response TranscationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.TxHash, nil
}

// Monitor TransactionStatus
func (c *ClientModule) MonitorTransactionStatus(txHash string) (string, error) {
	url := c.baseURL + "/check/" + txHash

	for {
		resp, err := http.Get(url)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("failed to check transaction status. Status Code: %d", resp.StatusCode)
		}

		var response TranscationStatusResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return "", err
		}

		if response.TxStatus == "CONFIRMED" || response.TxStatus == "FAILED" || response.TxStatus == "DNE" {
			return response.TxStatus, nil
		}

		// Stop before checking again
		time.Sleep(5 * time.Second)
	}
}

func main() {
	// Create new Client instance
	client := NewClientModule()

	// Example payload broadcasting transaction
	payload := TransactionPayload{
		Symbol:    "ETH",
		Price:     4500,
		Timestamp: uint64(time.Now().Unix()),
	}

	// Broadcasting the Transaction
	txHash, err := client.BroadcastTransaction(payload)
	if err != nil {
		fmt.Println("Error broadcasting transaction:", err)
		return
	}

	fmt.Println("Transaction broadcasting. TxHash:", txHash)

	// Monitor Transaction Status
	status, err := client.MonitorTransactionStatus(txHash)
	if err != nil {
		fmt.Println("Error monitoring transaction status:", err)
		return
	}

	fmt.Println("Transaction status:", status)
}
