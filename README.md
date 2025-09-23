# The Firstock Developer API Golang client - 

To communicate with the Firstock Developer API using Golang, you can use the official Golang client library provided by Firstock.
Licensed under the MIT License.


[Version - 1.0.0]


## Documentation

* Golang client documentation

## Installing the client


```bash
go get firstock
```

## API usage

```golang

// Login
loginRequest := Firstock.LoginRequest{
		UserId:     userId,
		Password:   password,
		TOTP:       totp,
		VendorCode: vendorCode,
		APIKey:     apiKey,
	}
login, err := Firstock.Login(loginRequest)
fmt.Println("Error:", err)
fmt.Println("Result:", login)

// Logout
logout, err := Firstock.Logout(userId)
fmt.Println("Error:", err)
fmt.Println("Result:", logout)

// UserDetails
userDetails, err := Firstock.UserDetails(userId)
fmt.Println("Error:", err)
fmt.Println("Result:", userDetails)

// Place Order
placeOrderRequest := Firstock.PlaceOrderRequest{
		UserId:          userId,
		Exchange:        exchange,
		Retention:       retention,
		Product:         product,
		PriceType:       priceType,
		TradingSymbol:   tradingSymbol,
		TransactionType: transactionType,
		Price:           price,
		TriggerPrice:    triggerPrice,
		Quantity:        quantity,
		Remarks:         remarks,
	}
placeOrder, err := Firstock.PlaceOrder(placeOrderRequest)
fmt.Println("Error:", err)
fmt.Println("Result:", placeOrder)

// Order Margin
orderMarginRequest := Firstock.OrderMarginRequest{
		UserId:          userId,
		Exchange:        exchange,
		TransactionType: transactionType,
		Product:         product,
		TradingSymbol:   tradingSymbol,
		Quantity:        quantity,
		PriceType:       priceType,
		Price:           price,
	}
orderMargin, err := Firstock.OrderMargin(orderMarginRequest)
fmt.Println("Error:", err)
fmt.Println("Result:", orderMargin)

// Order Book
orderBookDetails, err := Firstock.OrderBook(userId)
fmt.Println("Error:", err)
fmt.Println("Result:", orderBookDetails)

// Cancel Order
cancel_order := Firstock.OrderRequest{
		UserId:      userId,
		OrderNumber: order_number,
	}
cancelOrder, err := Firstock.CancelOrder(cancel_order)
fmt.Println("Error:", err)
fmt.Println("Result:", cancelOrder)

// Modify Order
modify_order := Firstock.ModifyOrderRequest{
	UserId:         userId,
	OrderNumber:    "25060500005017",
	PriceType:      "MKT",
	TradingSymbol:  "SAWACA",
	Price:          "",
	TriggerPrice:   "",
	Quantity:       "2",
	Product:        "C",
	Retention:      "DAY",
	Mkt_protection: "0.5",
}
modifyOrder := Firstock.ModifyOrder(modify_order)
fmt.Printf("Modify Order:\n%v\n", modifyOrder)

// Single Order History
single_order_history := Firstock.OrderRequest{
		UserId:      userId,
		OrderNumber: order_number,
	}
singleOrderHistory, err := Firstock.SingleOrderHistory(single_order_history)
fmt.Println("Error:", err)
fmt.Println("Result:", singleOrderHistory)

// Trade Book
tradeBook, err := Firstock.TradeBook(userId)
fmt.Println("Error:", err)
fmt.Println("Result:", tradeBook)

// Position Book
positionBookDetails, err := Firstock.PositionBook(userId)
fmt.Println("Error:", err)
fmt.Println("Result:", positionBookDetails)

// Product Conversion
productConversionRequest := Firstock.ProductConversionRequest{
		UserId:          userId,
		TradingSymbol:   "AVANCE",
		Exchange:        "BSE",
		PreviousProduct: "I", // B = Buy, S = Sell
		Product:         "C", // C = Delivery, I = Intraday, M = Margin Intraday (MIS)
		Quantity:        "1", // As string
		 MessageFlag:     "1",
	}

productConversion, err := Firstock.ProductConversion(productConversionRequest)
fmt.Println("Error:", err)
fmt.Println("Result:", productConversion)

// Holdings
holdingsDetails, err := Firstock.Holdings(userId)
fmt.Println("Error:", err)
fmt.Println("Result:", holdingsDetails)

// Holdings Details
holdingsDetails, err := Firstock.HoldingsDetails(userId)
fmt.Println("Error:", err)
fmt.Println("Result:", holdingsDetails)

// Limit
rmsLimitDetails, err := Firstock.RMSLmit(userId)
fmt.Println("Error:", err)
fmt.Println("Result:", rmsLimitDetails)

// Basket Margin
basketMarginRequest := Firstock.BasketMarginRequest{
		UserId:          userId,
		Exchange:        "NSE",
		TransactionType: "B",           // B = Buy, S = Sell
		Product:         "C",           // C = Delivery, I = Intraday, M = Margin Intraday (MIS)
		TradingSymbol:   "RELIANCE-EQ", // Ensure it's the correct symbol
		Quantity:        "1",           // As string
		PriceType:       "MKT",         // Example: "LMT" for Limit, "MKT" for Market
		Price:           "0",           // As string
		BasketListParams: []Firstock.BasketListParam{
			{
				Exchange:        "NSE",
				TransactionType: "B",
				Product:         "C",
				TradingSymbol:   "IDEA-EQ",
				Quantity:        "1",
				PriceType:       "MKT",
				Price:           "0",
			},
		},
	}

basketMargin, err := Firstock.BasketMargin(basketMarginRequest)
fmt.Println("Error:", err)
fmt.Println("Result:", basketMargin)

// Brokerage Calculator 
brokerageCalculatorRequest := Firstock.BrokerageCalculatorRequest{
		UserId:          userId,
		Exchange:        "NSE",
		TradingSymbol:   "SAWACA",
		TransactionType: "B",
		Product:         "C",
		Quantity:        "1",
		Price:           "0.50",
		StrikePrice:     "0.00",
		InstName:        "EQ",
		LotSize:         "1",
	}
brokerageCalculator, err := Firstock.BrokerageCalculator(brokerageCalculatorRequest)
fmt.Println("Error:", err)
fmt.Println("Result:", brokerageCalculator)

// Get Security Info
getSecurityInfo := Firstock.GetInfoRequest{
		UserId:        userId,
		Exchange:      "NSE",
		TradingSymbol: "NIFTY",
	}

getSecurityInfoDetails, err := Firstock.GetSecurityInfo(getSecurityInfo)
fmt.Println("Error:", err)
fmt.Println("Result:", getSecurityInfoDetails)

// Get Quote
getQuoteReq := Firstock.GetInfoRequest{
		UserId:        userId,
		Exchange:      "NSE",
		TradingSymbol: "NIFTY",
	}

	getQuoteDetails, err := Firstock.GetQuote(getQuoteReq)
	fmt.Println("Error:", err)
	fmt.Println("Result:", getQuoteDetails)

// Get Quote LTP
getQuoteLtpReq := Firstock.GetInfoRequest{
		UserId:        userId,
		Exchange:      "NSE",
		TradingSymbol: "NIFTY",
	}

getQuoteDetails, err := Firstock.GetQuoteLtp(getQuoteLtpReq)
fmt.Println("Error:", err)
fmt.Println("Result:", getQuoteDetails)

// Get Multi Quotes
	getMultiQuotesReq := Firstock.GetMultiQuotesRequest{
		UserId: userId, // replace with actual value
		Data: []Firstock.MultiQuoteData{
			{
				Exchange:      "NSE",
				TradingSymbol: "Nifty 50", // Ensure this matches the broker’s expected format
			},
			{
				Exchange:      "NFO",
				TradingSymbol: "NIFTY03APR25C23500",
			},
		},
	}

getMultiQuotes, err := Firstock.GetMultiQuotes(getMultiQuotesReq)
fmt.Println("Error:", err)
fmt.Println("Result:", getMultiQuotes)

// Get Multi Quote LTP
getMultiQuotesLtpReq := Firstock.GetMultiQuotesRequest{
		UserId: userId, // replace with actual value
		Data: []Firstock.MultiQuoteData{
			{
				Exchange:      "NSE",
	TradingSymbol: "Nifty 50", // Ensure this matches the broker’s expected format
			},
			{
				Exchange:      "NFO",
				TradingSymbol: "NIFTY03APR25C23500",
			},
		},
	}
getMultiQuotesLtp, err := Firstock.GetMultiQuotesLtp(getMultiQuotesLtpReq)
fmt.Println("Error:", err)
fmt.Println("Result:", getMultiQuotesLtp)

// Index List
indexList, err := Firstock.IndexList(userId)
fmt.Println("Error:", err)
fmt.Println("Result:", indexList)

// Get Expiry
getExpiryReq := Firstock.GetInfoRequest{
		UserId:        userId,
		Exchange:      "NSE",
		TradingSymbol: "NIFTY",
	}
getExpiryDetails, err := Firstock.GetExpiry(getExpiryReq)
fmt.Println("Error:", err)
fmt.Println("Result:", getExpiryDetails)

// Option Chain
optionChainRequest := Firstock.OptionChainRequest{
		UserId:      userId,
		Exchange:    "NFO",
		Symbol:      "NIFTY",
		Expiry:      "12JUN25", // Format must match broker format
		Count:       "5",       // Number of strikes above/below
		StrikePrice: "23150",   // ATM strike price
	}
optionChain, err := Firstock.OptionChain(optionChainRequest)
fmt.Println("Error:", err)
fmt.Println("Result:", optionChain)

// Search Scrips
searchScripsRequest := Firstock.SearchScripsRequest{
		UserId: userId,
		SText:  "RELIANCE",
	}
searchScrips, err := Firstock.SearchScrips(searchScripsRequest)
fmt.Println("Error:", err)
fmt.Println("Result:", searchScrips)

// Time Price Series Regular Interval
timePriceSeriesRegularIntervalRequest := Firstock.TimePriceSeriesIntervalRequest{
		UserId:        userId,
		Exchange:      "NSE",
		TradingSymbol: "NIFTY",
		Interval:      "1mi", // 5 minutes interval
		StartTime:     "09:15:00 23-04-2025",
		EndTime:       "15:29:00 23-04-2025",
	}
timePriceSeriesRegularInterval, err := Firstock.TimePriceSeriesRegularInterval(timePriceSeriesRegularIntervalRequest)
fmt.Println("Error:", err)
fmt.Println("Result:", timePriceSeriesRegularInterval)

// Time Price Series Day Interval
timePriceSeriesDayIntervalRequest := Firstock.TimePriceSeriesIntervalRequest{
		UserId:        userId,
		Exchange:      "NSE",
		TradingSymbol: "NIFTY",
		Interval:      "1d", // 5 minutes interval
		StartTime:     "09:15:00 20-04-2025",
		EndTime:       "15:29:00 23-04-2025",
	}
timePriceSeriesDayInterval, err := Firstock.TimePriceSeriesDayInterval(timePriceSeriesDayIntervalRequest)
fmt.Println("Error:", err)
fmt.Println("Result:", timePriceSeriesDayInterval)


Refer to the [Firstock Connect Documentation](https://firstock.in/api/docs/)  for the complete list of supported methods.

## Changelog

Check release notes.
