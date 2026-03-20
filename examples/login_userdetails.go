package main

import (
	"fmt"
	"log"

	Firstock "github.com/the-firstock/firstock-developer-sdk-golang/Firstock"
)

func main() {
	// ---------------------------------------------------------------------------
	// Credentials - replace with your actual values
	// ---------------------------------------------------------------------------
	userId := "Your_userId"
	password := "your_Password"
	totp := "your_totp"
	vendorCode := "Your_Vendor_code"
	apiKey := "Your_api_key"

	// ---------------------------------------------------------------------------
	// Step 1: Login
	// ---------------------------------------------------------------------------
	fmt.Println("==> Logging in...")

	loginRequest := Firstock.LoginRequest{
		UserId:     userId,
		Password:   password,
		TOTP:       totp,
		VendorCode: vendorCode,
		APIKey:     apiKey,
	}

	login, err := Firstock.Login(loginRequest)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	fmt.Println("Login successful!")
	fmt.Printf("Login Response: %+v\n", login)

	// ---------------------------------------------------------------------------
	// Step 2: Fetch User Details
	// ---------------------------------------------------------------------------
	fmt.Println("\n==> Fetching user details...")

	userDetails, err := Firstock.UserDetails(userId)
	if err != nil {
		log.Fatalf("Failed to fetch user details: %v", err)
	}

	fmt.Println("User details fetched successfully!")
	fmt.Printf("User Details: %+v\n", userDetails)
}
