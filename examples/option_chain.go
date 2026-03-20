package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	Firstock "github.com/the-firstock/firstock-developer-sdk-golang/Firstock"
)

func main() {
	// ---------------------------------------------------------------------------
	// Read userId from config.json (saved during login)
	// ---------------------------------------------------------------------------
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Failed to read config.json: %v", err)
	}

	var config map[string]interface{}
	if err := json.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("Failed to parse config.json: %v", err)
	}

	// userId is the first key in config.json
	var userId string
	for key := range config {
		userId = key
		break
	}
	fmt.Println("Using userId:", userId)

	// ---------------------------------------------------------------------------
	// Step 1: Get Expiry
	// ---------------------------------------------------------------------------
	fmt.Println("\n==> Fetching expiry dates...")

	getExpiryResponse, errGetExpiry := Firstock.GetExpiry(Firstock.GetInfoRequest{
		UserId:        userId,
		Exchange:      "NSE",
		TradingSymbol: "NIFTY",
	})
	if errGetExpiry != nil {
		log.Fatalf("Failed to fetch expiry dates: %+v", errGetExpiry)
	}

	fmt.Printf("Expiry Response: %+v\n", getExpiryResponse)

	// Extract the first (nearest) expiry date from the response
	rawExpiry := getExpiryResponse.Data.ExpiryDates[0]

	// Convert expiry from DDMMMYYYY to DDMMMYY format (e.g. 24MAR2026 -> 24MAR26)
	expiry := rawExpiry[:len(rawExpiry)-4] + rawExpiry[len(rawExpiry)-2:]
	fmt.Println("Using expiry:", expiry)

	// ---------------------------------------------------------------------------
	// Step 2: Fetch Option Chain
	// ---------------------------------------------------------------------------
	fmt.Println("\n==> Fetching Option Chain...")

	optionChainResponse, errOptionChain := Firstock.OptionChain(Firstock.OptionChainRequest{
		UserId:      userId,
		Exchange:    "NFO",
		Symbol:      "NIFTY",
		StrikePrice: "23150", // Replace with current NIFTY market price
		Expiry:      expiry,
		Count:       "5",
	})
	if errOptionChain != nil {
		log.Fatalf("Failed to fetch option chain: %+v", errOptionChain)
	}

	fmt.Println("Option Chain fetched successfully!")
	fmt.Printf("Option Chain: %+v\n", optionChainResponse)
}