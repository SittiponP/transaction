# Go Client Module Documentation

## Overview

Go Client Module is a library designed to interaction with HTTP server for broadcasting transactions and monitoring their status until finalization.

## Installation

Install the module using 'go get' command:

go get github.com/SittiponP/transaction

## Usage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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

## Configuration

The Go client module requires the baseURL of the server to interact with.
baseURL := "https://your-server-base-url.com" Replace this URL with your URL

## Error Handling

The module perform a basic error handling and return a error messages.