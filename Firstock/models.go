// Copyright (c) [2025] [abc]
// SPDX-License-Identifier: MIT
package Firstock

// Models for Login
type LoginRequest struct {
	UserId     string `json:"userId"`
	Password   string `json:"password"`
	TOTP       string `json:"totp"`
	VendorCode string `json:"vendorCode"`
	APIKey     string `json:"apiKey"`
}

type Data struct {
	Actid      string `json:"actid"`
	UserName   string `json:"userName"`
	SUserToken string `json:"susertoken"`
	Email      string `json:"email"`
}

// Model for User Details
type UserDetailsRequest struct {
	UserId string `json:"userId"`
	JKey   string `json:"jkey"`
}

// Model for Logout
type LogoutRequest struct {
	UserId string `json:"userId"`
	JKey   string `json:"jkey"`
}

type LogoutResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Models for Place Order
type PlaceOrderRequest struct {
	UserId          string `json:"userId"`
	Exchange        string `json:"exchange"`
	Retention       string `json:"retention"`
	Product         string `json:"product"`
	PriceType       string `json:"priceType"`
	TradingSymbol   string `json:"tradingSymbol"`
	TransactionType string `json:"transactionType"`
	Price           string `json:"price"`
	TriggerPrice    string `json:"triggerPrice"`
	Quantity        string `json:"quantity"`
	Remarks         string `json:"remarks"`
}

type PlaceOrderRequestBody struct {
	UserId          string `json:"userId"`
	JKey            string `json:"jkey"`
	Exchange        string `json:"exchange"`
	Retention       string `json:"retention"`
	Product         string `json:"product"`
	PriceType       string `json:"priceType"`
	TradingSymbol   string `json:"tradingSymbol"`
	TransactionType string `json:"transactionType"`
	Price           string `json:"price"`
	TriggerPrice    string `json:"triggerPrice"`
	Quantity        string `json:"quantity"`
	Remarks         string `json:"remarks"`
}

// Models for OrderMargin
type OrderMarginRequest struct {
	UserId          string `json:"userId"`
	Exchange        string `json:"exchange"`
	TransactionType string `json:"transactionType"`
	Product         string `json:"product"`
	TradingSymbol   string `json:"tradingSymbol"`
	Quantity        string `json:"quantity"`
	PriceType       string `json:"priceType"`
	Price           string `json:"price"`
}

type OrderMarginRequestBody struct {
	UserId          string `json:"userId"`
	JKey            string `json:"jkey"`
	Exchange        string `json:"exchange"`
	TransactionType string `json:"transactionType"`
	Product         string `json:"product"`
	TradingSymbol   string `json:"tradingSymbol"`
	Quantity        string `json:"quantity"`
	PriceType       string `json:"priceType"`
	Price           string `json:"price"`
}

// Models for Single Order History

type OrderRequest struct {
	UserId      string `json:"userId"`
	OrderNumber string `json:"orderNumber"`
}

type OrderRequestBody struct {
	UserId      string `json:"userId"`
	JKey        string `json:"jkey"`
	OrderNumber string `json:"orderNumber"`
}

// Model for Trade Book, RmsLimit
type BaseRequest struct {
	UserId string `json:"userId"`
	JKey   string `json:"jkey"`
}

// Models for Get Expiry

type GetInfoRequest struct {
	UserId        string `json:"userId"`
	Exchange      string `json:"exchange"`
	TradingSymbol string `json:"tradingSymbol"`
}

type GetInfoRequestBody struct {
	UserId        string `json:"userId"`
	JKey          string `json:"jkey"`
	Exchange      string `json:"exchange"`
	TradingSymbol string `json:"tradingSymbol"`
}

// Model for ModifyOrder
type ModifyOrderRequest struct {
	UserId         string `json:"userId"`
	OrderNumber    string `json:"orderNumber"`
	PriceType      string `json:"priceType"`
	TradingSymbol  string `json:"tradingSymbol"`
	Price          string `json:"price"`
	TriggerPrice   string `json:"triggerPrice"`
	Quantity       string `json:"quantity"`
	Product        string `json:"product"`
	Retention      string `json:"retention"`
	Mkt_protection string `json:"mkt_protection"`
}

type ModifyOrderRequestBody struct {
	UserId         string `json:"userId"`
	JKey           string `json:"jkey"`
	OrderNumber    string `json:"orderNumber"`
	PriceType      string `json:"priceType"`
	TradingSymbol  string `json:"tradingSymbol"`
	Price          string `json:"price"`
	TriggerPrice   string `json:"triggerPrice"`
	Quantity       string `json:"quantity"`
	Product        string `json:"product"`
	Retention      string `json:"retention"`
	Mkt_protection string `json:"mkt_protection"`
}

type BrokerageCalculatorRequest struct {
	UserId          string `json:"userId"`
	Exchange        string `json:"exchange"`
	TradingSymbol   string `json:"tradingSymbol"`
	TransactionType string `json:"transactionType"`
	Product         string `json:"Product"`
	Quantity        string `json:"quantity"`
	Price           string `json:"price"`
	StrikePrice     string `json:"strike_price"`
	InstName        string `json:"inst_name"`
	LotSize         string `json:"lot_size"`
}

type BrokerageCalculatorRequestBody struct {
	UserId          string `json:"userId"`
	JKey            string `json:"jkey"`
	Exchange        string `json:"exchange"`
	TradingSymbol   string `json:"tradingSymbol"`
	TransactionType string `json:"transactionType"`
	Product         string `json:"Product"`
	Quantity        string `json:"quantity"`
	Price           string `json:"price"`
	StrikePrice     string `json:"strike_price"`
	InstName        string `json:"inst_name"`
	LotSize         string `json:"lot_size"`
}

type BasketListParam struct {
	Exchange        string `json:"exchange"`
	TransactionType string `json:"transactionType"`
	Product         string `json:"product"`
	TradingSymbol   string `json:"tradingSymbol"`
	Quantity        string `json:"quantity"`
	PriceType       string `json:"priceType"`
	Price           string `json:"price"`
}

type BasketMarginRequest struct {
	UserId           string            `json:"userId"`
	Exchange         string            `json:"exchange"`
	TransactionType  string            `json:"transactionType"`
	Product          string            `json:"product"`
	TradingSymbol    string            `json:"tradingSymbol"`
	Quantity         string            `json:"quantity"`
	PriceType        string            `json:"priceType"`
	Price            string            `json:"price"`
	BasketListParams []BasketListParam `json:"BasketList_Params"`
}

type BasketMarginRequestBody struct {
	UserId           string            `json:"userId"`
	JKey             string            `json:"jKey"`
	Exchange         string            `json:"exchange"`
	TransactionType  string            `json:"transactionType"`
	Product          string            `json:"product"`
	TradingSymbol    string            `json:"tradingSymbol"`
	Quantity         string            `json:"quantity"`
	PriceType        string            `json:"priceType"`
	Price            string            `json:"price"`
	BasketListParams []BasketListParam `json:"BasketList_Params"`
}

type ProductConversionRequest struct {
	UserId          string `json:"userId"`
	TradingSymbol   string `json:"tradingSymbol"`
	Exchange        string `json:"exchange"`
	PreviousProduct string `json:"previousProduct"`
	Product         string `json:"product"`
	Quantity        string `json:"quantity"`
}

type ProductConversionRequestBody struct {
	UserId          string `json:"userId"`
	JKey            string `json:"jkey"`
	TradingSymbol   string `json:"tradingSymbol"`
	Exchange        string `json:"exchange"`
	PreviousProduct string `json:"previousProduct"`
	Product         string `json:"product"`
	Quantity        string `json:"quantity"`
}

type SearchScripsRequest struct {
	UserId string `json:"userId"`
	SText  string `json:"stext"`
}

type SearchScripsBody struct {
	UserId string `json:"userId"`
	JKey   string `json:"jkey"`
	SText  string `json:"stext"`
}

type OptionChainRequest struct {
	UserId      string `json:"userId"`
	Exchange    string `json:"exchange"`
	Symbol      string `json:"symbol"`
	Expiry      string `json:"expiry"`
	Count       string `json:"count"`
	StrikePrice string `json:"strikePrice"`
}

type OptionChainRequestBody struct {
	UserId      string `json:"userId"`
	JKey        string `json:"jkey"`
	Exchange    string `json:"exchange"`
	Symbol      string `json:"symbol"`
	Expiry      string `json:"expiry"`
	Count       string `json:"count"`
	StrikePrice string `json:"strikePrice"`
}

type MultiQuoteData struct {
	Exchange      string `json:"exchange"`
	TradingSymbol string `json:"tradingSymbol"`
}

type GetMultiQuotesRequest struct {
	UserId string           `json:"userId"`
	Data   []MultiQuoteData `json:"data"`
}

type GetMultiQuotesRequestBody struct {
	UserId string           `json:"userId"`
	JKey   string           `json:"jkey"`
	Data   []MultiQuoteData `json:"data"`
}

type TimePriceSeriesIntervalRequest struct {
	UserId        string `json:"userId"`
	Exchange      string `json:"exchange"`
	Interval      string `json:"interval"`
	TradingSymbol string `json:"tradingSymbol"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
}

type TimePriceSeriesIntervalRequestBody struct {
	UserId        string `json:"userId"`
	JKey          string `json:"jkey"`
	Exchange      string `json:"exchange"`
	Interval      string `json:"interval"`
	TradingSymbol string `json:"tradingSymbol"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
}

// Error Response model
type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponseModel struct {
	Code   string      `json:"code"`
	Error  ErrorDetail `json:"error"`
	Name   string      `json:"name"`
	Status string      `json:"status"`
}

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
	Error   any    `json:"error"`
}

//--------------------------------------------------Response Structure-----------------------------------------------

// --------------------------------------------------Login Response---------------------------------------------------
type LoginResponse struct {
	Status  string `json:"status"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

//--------------------------------------------------User Details Response--------------------------------------------

type UserDetailsResponse struct {
	Status string          `json:"status"`
	Data   UserDetailsData `json:"data"`
}

type UserDetailsData struct {
	Actid       string   `json:"actid"`
	Email       string   `json:"email"`
	Exchange    []string `json:"exchange"`
	Orarr       []string `json:"orarr"`
	RequestTime string   `json:"requestTime"`
	Uprev       string   `json:"uprev"`
	UserName    string   `json:"userName"`
}

//--------------------------------------------------Place Order Response--------------------------------------------

type PlaceOrderResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    PlaceOrderData `json:"data"`
}

type PlaceOrderData struct {
	OrderNumber string `json:"orderNumber"`
	RequestTime string `json:"requestTime"`
}

//--------------------------------------------------Order Margin Response--------------------------------------------

type OrderMarginResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    OrderMarginData `json:"data"`
}

type OrderMarginData struct {
	AvailableMargin  string `json:"availableMargin"`
	Cash             string `json:"cash"`
	MarginOnNewOrder string `json:"marginOnNewOrder"`
	Remarks          string `json:"remarks"`
	RequestTime      string `json:"requestTime"`
}

//--------------------------------------------------Single Order History Response--------------------------------------------

type SingleOrderHistoryModel struct {
	AveragePrice     string `json:"averagePrice"`
	Exchange         string `json:"exchange"`
	ExchangeOrderNum string `json:"exchangeOrderNum"`
	ExchangeTime     string `json:"exchangeTime"`
	FillShares       string `json:"fillShares"`
	OrderNumber      string `json:"orderNumber"`
	OrderTime        string `json:"orderTime"`
	Price            string `json:"price"`
	PriceType        string `json:"priceType"`
	Product          string `json:"product"`
	Quantity         string `json:"quantity"`
	RejectReason     string `json:"rejectReason"`
	Remarks          string `json:"remarks"`
	ReportType       string `json:"reportType"`
	Retention        string `json:"retention"`
	Status           string `json:"status"`
	TickSize         string `json:"tickSize"`
	Token            string `json:"token"`
	TradingSymbol    string `json:"tradingSymbol"`
	TransactionType  string `json:"transactionType"`
	UserID           string `json:"userId"`
}

type SingleOrderHistoryResponse struct {
	Status  string                    `json:"status"`
	Message string                    `json:"message"`
	Data    []SingleOrderHistoryModel `json:"data"`
}

//--------------------------------------------------Cancel Order Response--------------------------------------------

type CancelOrderResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    CancelOrderResponseData `json:"data"`
}

type CancelOrderResponseData struct {
	OrderNumber string `json:"orderNumber"`
	RejReason   string `json:"rejreason"`
	RequestTime string `json:"requestTime"`
}

//--------------------------------------------------Modify Order Response--------------------------------------------

type ModifyOrderResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    ModifyOrderResponseData `json:"data"`
}

type ModifyOrderResponseData struct {
	OrderNumber string `json:"orderNumber"`
	RejReason   string `json:"rejreason"`
	RequestTime string `json:"requestTime"`
}

//--------------------------------------------------Trade Book Response--------------------------------------------

type TradeBookResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Data    []TradeBookEntry `json:"data"`
}

type TradeBookEntry struct {
	Exchange           string `json:"exchange"`
	ExchangeUpdateTime string `json:"exchangeUpdateTime"`
	ExchOrdID          string `json:"exchordid"`
	FillID             string `json:"fillId"`
	FillPrice          string `json:"fillPrice"`
	FillQuantity       string `json:"fillQuantity"`
	FillTime           string `json:"fillTime"`
	FillShares         string `json:"fillshares"`
	LotSize            string `json:"lotSize"`
	OrderNumber        string `json:"orderNumber"`
	OrderTime          string `json:"orderTime"`
	PriceFactor        string `json:"priceFactor"`
	PricePrecision     string `json:"pricePrecision"`
	PriceType          string `json:"priceType"`
	Product            string `json:"product"`
	Quantity           string `json:"quantity"`
	Retention          string `json:"retention"`
	TickSize           string `json:"tickSize"`
	Token              string `json:"token"`
	TradingSymbol      string `json:"tradingSymbol"`
	TransactionType    string `json:"transactionType"`
	UserID             string `json:"userId"`
}

//--------------------------------------------------RMS Limit Response--------------------------------------------

type RmsLimitResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    RmsLimitData `json:"data"`
}

type RmsLimitData struct {
	BrkCollAmt  string `json:"brkcollamt"`
	Cash        string `json:"cash"`
	Collateral  string `json:"collateral"`
	Expo        string `json:"expo"`
	MarginUsed  string `json:"marginused"`
	PayIn       string `json:"payin"`
	PeakMar     string `json:"peak_mar"`
	Premium     string `json:"premium"`
	RequestTime string `json:"requestTime"`
	Span        string `json:"span"`
}

//--------------------------------------------------Position Book Response--------------------------------------------

type PositionBookResponse struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Data    []PositionBookData `json:"data"`
}

type PositionBookData struct {
	RealizedPNL         string `json:"RealizedPNL"`
	DayBuyAmount        string `json:"dayBuyAmount"`
	DayBuyAveragePrice  string `json:"dayBuyAveragePrice"`
	DayBuyQuantity      string `json:"dayBuyQuantity"`
	DaySellAveragePrice string `json:"daySellAveragePrice"`
	DaySellQuantity     string `json:"daySellQuantity"`
	Exchange            string `json:"exchange"`
	LotSize             string `json:"lotSize"`
	NetAveragePrice     string `json:"netAveragePrice"`
	NetQuantity         string `json:"netQuantity"`
	NetUploadPrice      string `json:"netUploadPrice"`
	Product             string `json:"product"`
	TickSize            string `json:"tickSize"`
	Token               string `json:"token"`
	TotalMTM            string `json:"totalMTM"`
	TotalPNL            string `json:"totalPNL"`
	TradingSymbol       string `json:"tradingSymbol"`
	UploadPrice         string `json:"uploadPrice"`
	UserID              string `json:"userId"`
}

//--------------------------------------------------Holdings Response--------------------------------------------

type HoldingsResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    []HoldingEntry `json:"data"`
}

type HoldingEntry struct {
	Exchange      string `json:"exchange"`
	TradingSymbol string `json:"tradingSymbol"`
}

// -------------------------------------------------------Holdings Details Response--------------------------------------

type HoldingsDetailsResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    []Holding `json:"data"`
}

type Holding struct {
	ExchangeTradingSymbol []ExchangeTradingSymbolDetails `json:"exchangeTradingSymbol"`
	SellAmount            string                         `json:"sellAmount"`
	HoldQuantity          string                         `json:"holdQuantity"`
	UploadPrice           string                         `json:"uploadPrice"`
	BTSTQuantity          string                         `json:"BTSTQuantity"`
	UsedQuantity          string                         `json:"usedQuantity"`
	TradeQuantity         string                         `json:"tradeQuantity"`
}

type ExchangeTradingSymbolDetails struct {
	Exchange       string `json:"exchange"`
	Token          string `json:"token"`
	TradingSymbol  string `json:"tradingSymbol"`
	PricePrecision string `json:"pricePrecision"`
	TickSize       string `json:"tickSize"`
	LotSize        string `json:"lotSize"`
}

// -------------------------------------------------------Order Book Response--------------------------------------
type OrderBookResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Data    []OrderBookEntry `json:"data"`
}

type OrderBookEntry struct {
	AveragePrice    string `json:"averagePrice"`
	Exchange        string `json:"exchange"`
	FillShares      string `json:"fillShares"`
	LotSize         string `json:"lotSize"`
	OrderNumber     string `json:"orderNumber"`
	OrderTime       string `json:"orderTime"`
	Price           string `json:"price"`
	PriceType       string `json:"priceType"`
	Product         string `json:"product"`
	Quantity        string `json:"quantity"`
	RejectReason    string `json:"rejectReason"`
	Remarks         string `json:"remarks"`
	Retention       string `json:"retention"`
	Status          string `json:"status"`
	TickSize        string `json:"tickSize"`
	Token           string `json:"token"`
	TradingSymbol   string `json:"tradingSymbol"`
	TransactionType string `json:"transactionType"`
	UserID          string `json:"userId"`
}

// ------------------------------------------------Get Expiry Response----------------------------------------
type GetExpiryResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Data    ExpiryData `json:"data"`
}

type ExpiryData struct {
	ExpiryDates []string `json:"expiryDates"`
}

// ------------------------------------------------Brokerage Calculator Response----------------------------------------

type BrokerageCalculatorResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    BrokerageCalculatorData `json:"data"`
}

type BrokerageCalculatorData struct {
	Brokerage      string `json:"brokerage"`
	ExchangeCharge string `json:"exchange_charge"`
	GST            string `json:"gst"`
	Remarks        string `json:"remarks"`
	SebiCharge     string `json:"sebi_charge"`
	StampDuty      string `json:"stamp_duty"`
}

// ------------------------------------------------Basket Margin Response----------------------------------------
type BasketMarginResponse struct {
	Status  string                   `json:"status"`
	Message string                   `json:"message"`
	Data    BasketMarginResponseData `json:"data"`
}

type BasketMarginResponseData struct {
	BasketMargin     []string `json:"BasketMargin"`
	MarginOnNewOrder string   `json:"MarginOnNewOrder"`
	PreviousMargin   string   `json:"PreviousMargin"`
	Remarks          string   `json:"Remarks"`
	TradedMargin     string   `json:"TradedMargin"`
	RequestTime      string   `json:"requestTime"`
}

// -----------------------------------------------------------------------Get Security Info Response----------------------------------------
type GetSecurityInfoResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Data    SecurityInfoData `json:"data"`
}

type SecurityInfoData struct {
	Exchange       string `json:"exchange"`
	InstrumentName string `json:"instrumentName"`
	LotSize        string `json:"lotSize"`
	Mult           string `json:"mult"`
	PriceFactor    string `json:"priceFactor"`
	PricePrecision string `json:"pricePrecision"`
	Segment        string `json:"segment"`
	SymbolName     string `json:"symbolName"`
	TickSize       string `json:"tickSize"`
	Token          string `json:"token"`
	TradingSymbol  string `json:"tradingSymbol"`
}

// --------------------------------------------------------- Product conversion response -----------------------------------------
type ProductConversionResponse struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Data    ProductConversionData `json:"data"`
}

type ProductConversionData struct {
	Status      string `json:"Status"`
	RequestTime string `json:"requestTime"`
}

// --------------------------------------------------------- Get Quote Response -----------------------------------------
type GetQuoteResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    QuoteData `json:"data"`
}

type QuoteData struct {
	VWAP              string `json:"VWAP"`
	BestBuyOrder1     string `json:"bestBuyOrder1"`
	BestBuyOrder2     string `json:"bestBuyOrder2"`
	BestBuyOrder3     string `json:"bestBuyOrder3"`
	BestBuyOrder4     string `json:"bestBuyOrder4"`
	BestBuyOrder5     string `json:"bestBuyOrder5"`
	BestBuyPrice1     string `json:"bestBuyPrice1"`
	BestBuyPrice2     string `json:"bestBuyPrice2"`
	BestBuyPrice3     string `json:"bestBuyPrice3"`
	BestBuyPrice4     string `json:"bestBuyPrice4"`
	BestBuyPrice5     string `json:"bestBuyPrice5"`
	BestBuyQuantity1  string `json:"bestBuyQuantity1"`
	BestBuyQuantity2  string `json:"bestBuyQuantity2"`
	BestBuyQuantity3  string `json:"bestBuyQuantity3"`
	BestBuyQuantity4  string `json:"bestBuyQuantity4"`
	BestBuyQuantity5  string `json:"bestBuyQuantity5"`
	BestSellOrder1    string `json:"bestSellOrder1"`
	BestSellOrder2    string `json:"bestSellOrder2"`
	BestSellOrder3    string `json:"bestSellOrder3"`
	BestSellOrder4    string `json:"bestSellOrder4"`
	BestSellOrder5    string `json:"bestSellOrder5"`
	BestSellPrice1    string `json:"bestSellPrice1"`
	BestSellPrice2    string `json:"bestSellPrice2"`
	BestSellPrice3    string `json:"bestSellPrice3"`
	BestSellPrice4    string `json:"bestSellPrice4"`
	BestSellPrice5    string `json:"bestSellPrice5"`
	BestSellQuantity1 string `json:"bestSellQuantity1"`
	BestSellQuantity2 string `json:"bestSellQuantity2"`
	BestSellQuantity3 string `json:"bestSellQuantity3"`
	BestSellQuantity4 string `json:"bestSellQuantity4"`
	BestSellQuantity5 string `json:"bestSellQuantity5"`
	CompanyName       string `json:"companyName"`
	DayClosePrice     string `json:"dayClosePrice"`
	DayHighPrice      string `json:"dayHighPrice"`
	DayLowPrice       string `json:"dayLowPrice"`
	DayOpenPrice      string `json:"dayOpenPrice"`
	Exchange          string `json:"exchange"`
	ISIN              string `json:"isin"`
	LastTradedPrice   string `json:"lastTradedPrice"`
	LotSize           string `json:"lotSize"`
	Multipler         string `json:"multipler"`
	OpenInterest      string `json:"openInterest"`
	PriceFactor       string `json:"priceFactor"`
	PricePrecision    string `json:"pricePrecision"`
	RequestTime       string `json:"requestTime"`
	Segment           string `json:"segment"`
	SymbolName        string `json:"symbolName"`
	TickSize          string `json:"tickSize"`
	Token             string `json:"token"`
	TotalBuyQuantity  string `json:"totalBuyQuantity"`
	TotalSellQuantity string `json:"totalSellQuantity"`
	TradingSymbol     string `json:"tradingSymbol"`
}

// ------------------------------------------------------------ Get Quote Ltp Response ----------------------------------------------------
type GetQuoteLtpResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    QuoteLtpData `json:"data"`
}

type QuoteLtpData struct {
	CompanyName     string `json:"companyName"`
	Exchange        string `json:"exchange"`
	LastTradedPrice string `json:"lastTradedPrice"`
	RequestTime     string `json:"requestTime"`
	Token           string `json:"token"`
}

// ----------------------------------------------------------- Get Multi Quotes Response ----------------------------------------------------
type GetMultiQuotesResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    []MultiQuoteEntry `json:"data"`
}

type MultiQuoteEntry struct {
	// Common fields
	CompanyName   string `json:"companyName"`
	Exchange      string `json:"exchange"`
	Identifier    string `json:"identifier"`
	RequestTime   string `json:"requestTime"`
	TradingSymbol string `json:"tradingSymbol"`
	Status        string `json:"status,omitempty"` // only present in error case
	Error         string `json:"error,omitempty"`  // only present in error case

	// Optional quote fields (present only in success case)
	VWAP              string `json:"VWAP,omitempty"`
	BestBuyOrder1     string `json:"bestBuyOrder1,omitempty"`
	BestBuyOrder2     string `json:"bestBuyOrder2,omitempty"`
	BestBuyOrder3     string `json:"bestBuyOrder3,omitempty"`
	BestBuyOrder4     string `json:"bestBuyOrder4,omitempty"`
	BestBuyOrder5     string `json:"bestBuyOrder5,omitempty"`
	BestBuyPrice1     string `json:"bestBuyPrice1,omitempty"`
	BestBuyPrice2     string `json:"bestBuyPrice2,omitempty"`
	BestBuyPrice3     string `json:"bestBuyPrice3,omitempty"`
	BestBuyPrice4     string `json:"bestBuyPrice4,omitempty"`
	BestBuyPrice5     string `json:"bestBuyPrice5,omitempty"`
	BestBuyQuantity1  string `json:"bestBuyQuantity1,omitempty"`
	BestBuyQuantity2  string `json:"bestBuyQuantity2,omitempty"`
	BestBuyQuantity3  string `json:"bestBuyQuantity3,omitempty"`
	BestBuyQuantity4  string `json:"bestBuyQuantity4,omitempty"`
	BestBuyQuantity5  string `json:"bestBuyQuantity5,omitempty"`
	BestSellOrder1    string `json:"bestSellOrder1,omitempty"`
	BestSellOrder2    string `json:"bestSellOrder2,omitempty"`
	BestSellOrder3    string `json:"bestSellOrder3,omitempty"`
	BestSellOrder4    string `json:"bestSellOrder4,omitempty"`
	BestSellOrder5    string `json:"bestSellOrder5,omitempty"`
	BestSellPrice1    string `json:"bestSellPrice1,omitempty"`
	BestSellPrice2    string `json:"bestSellPrice2,omitempty"`
	BestSellPrice3    string `json:"bestSellPrice3,omitempty"`
	BestSellPrice4    string `json:"bestSellPrice4,omitempty"`
	BestSellPrice5    string `json:"bestSellPrice5,omitempty"`
	BestSellQuantity1 string `json:"bestSellQuantity1,omitempty"`
	BestSellQuantity2 string `json:"bestSellQuantity2,omitempty"`
	BestSellQuantity3 string `json:"bestSellQuantity3,omitempty"`
	BestSellQuantity4 string `json:"bestSellQuantity4,omitempty"`
	BestSellQuantity5 string `json:"bestSellQuantity5,omitempty"`
	DayClosePrice     string `json:"dayClosePrice,omitempty"`
	DayHighPrice      string `json:"dayHighPrice,omitempty"`
	DayLowPrice       string `json:"dayLowPrice,omitempty"`
	DayOpenPrice      string `json:"dayOpenPrice,omitempty"`
	InstrumentName    string `json:"instrumentName,omitempty"`
	LastTradedPrice   string `json:"lastTradedPrice,omitempty"`
	LotSize           string `json:"lotSize,omitempty"`
	Multipler         string `json:"multipler,omitempty"`
	OpenInterest      string `json:"openInterest,omitempty"`
	PriceFactor       string `json:"priceFactor,omitempty"`
	PricePrecision    string `json:"pricePrecision,omitempty"`
	Segment           string `json:"segment,omitempty"`
	SymbolName        string `json:"symbolName,omitempty"`
	TickSize          string `json:"tickSize,omitempty"`
	Token             string `json:"token,omitempty"`
	TotalBuyQuantity  string `json:"totalBuyQuantity,omitempty"`
	TotalSellQuantity string `json:"totalSellQuantity,omitempty"`
}

// -------------------------------------------------------------- Index List Response ----------------------------------------------------
type IndexListResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    []IndexEntry `json:"data"`
}

type IndexEntry struct {
	Exchange      string `json:"exchange"`
	Token         string `json:"token"`
	TradingSymbol string `json:"tradingSymbol"`
	Symbol        string `json:"symbol"`
	IndexName     string `json:"idxname"`
}

// ------------------------------------------------------ Get Multi Quotes Ltp Response---------------------------------------------------
type GetMultiQuotesLtpResponse struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Data    []MultiQuoteLtpEntry `json:"data"`
}

type MultiQuoteLtpEntry struct {
	CompanyName     string `json:"companyName"`
	Exchange        string `json:"exchange"`
	Identifier      string `json:"identifier"`
	LastTradedPrice string `json:"lastTradedPrice"`
	RequestTime     string `json:"requestTime"`
	Token           string `json:"token,omitempty"`
	TradingSymbol   string `json:"tradingSymbol"`
	Status          string `json:"status,omitempty"` // "error" if failed
	Error           string `json:"error,omitempty"`  // error message if any
}

// ------------------------------------------------------ Search Scrips Response---------------------------------------------------

type SearchScripsResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    []SearchScripItem `json:"data"`
}

type SearchScripItem struct {
	Token              string `json:"token"`
	Exchange           string `json:"exchange"`
	CompanyName        string `json:"companyName"`
	RepresentationName string `json:"representationName"`
	InstrumentName     string `json:"instrumentName"`
	TradingSymbol      string `json:"tradingSymbol"`
}

//--------------------------------------------------------- Option Chain Response---------------------------------------------------

type OptionChainResponse struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Data    []OptionChainEntry `json:"data"`
}

type OptionChainEntry struct {
	Exchange        string `json:"exchange"`
	LastTradedPrice string `json:"lastTradedPrice"`
	LotSize         string `json:"lotSize"`
	OptionType      string `json:"optionType"`
	ParentToken     string `json:"parentToken"`
	PricePrecision  string `json:"pricePrecision"`
	StrikePrice     string `json:"strikePrice"`
	TickSize        string `json:"tickSize"`
	Token           string `json:"token"`
	TradingSymbol   string `json:"tradingSymbol"`
}

//-------------------------------------------------------- Time Price Series Regular Interval Response---------------------------------------------------

type TimePriceSeriesRegularIntervalResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    []TimePriceSeriesEntry `json:"data"`
}

type TimePriceSeriesEntry struct {
	Time      string  `json:"time"`
	EpochTime int64   `json:"epochTime"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    int     `json:"volume"`
	OI        int     `json:"oi"`
}

//-------------------------------------------------------- Time Price Series Day Interval Response---------------------------------------------------

type TimePriceSeriesDayIntervalResponse struct {
	Status  string                   `json:"status"`
	Message string                   `json:"message"`
	Data    []TimePriceSeriesDayData `json:"data"`
}

type TimePriceSeriesDayData struct {
	Time      string  `json:"time"`
	EpochTime int64   `json:"epochTime"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    int     `json:"volume"`
	OI        int     `json:"oi"`
}
