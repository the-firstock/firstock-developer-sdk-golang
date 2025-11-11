// Copyright (c) [2025] [abc]
// SPDX-License-Identifier: MIT
package Firstock

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

type firstock struct{}

var thefirstock = &apifunctions{}

// Call Login function to login to Firstock
func (fs *firstock) Login(reqBody LoginRequest) (loginResponse *LoginResponse, errLogin *ErrorResponseModel) {
	var loginRequest LoginRequest = LoginRequest{
		UserId:     reqBody.UserId,
		Password:   encodePassword(reqBody.Password),
		TOTP:       reqBody.TOTP,
		VendorCode: reqBody.VendorCode,
		APIKey:     reqBody.APIKey,
	}
	loginResponse = &LoginResponse{}

	loginResponseMap, code, _ := thefirstock.LoginFunction(
		loginRequest,
	)

	if code == status_internal_server_error {
		errLogin = internalServerErrorResponse()
		return
	} else if code == status_ok {
		// Extract SUserToken from login response
		jsonData, err := json.Marshal(loginResponseMap)
		if err != nil {
			return nil, internalServerErrorResponse()
		}

		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, loginResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}

		// Write the following to a config.json file. Create the file if it does not exist.
		err = saveJKeyToConfig(LogoutRequest{
			UserId: loginResponse.Data.Actid,
			JKey:   loginResponse.Data.SUserToken,
		})

		if err != nil {
			fmt.Println("Couldn't save key to config file")
		}
		return
	}
	errLogin = failureResponseStructure(loginResponseMap)
	return
}

// Call Logout function to logout from Firstock
func (fs *firstock) Logout(userId string) (logoutResponse *LogoutResponse, errLogout *ErrorResponseModel) {
	var logout LogoutRequest
	logoutResponse = &LogoutResponse{}
	logout.UserId = userId
	logout.JKey = ""

	// Read jKey for userId from config.json
	logout.JKey, errLogout = readJkey(userId)
	if logout.JKey == "" {
		return
	}

	logoutInfo, code, _ := thefirstock.LogoutFunction(logout)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(userId)
	} else if code == status_internal_server_error {
		errLogout = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(logoutInfo)
		if err != nil {
			return nil, internalServerErrorResponse()
		}

		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, logoutResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		removeJKeyFromConfig(logout.UserId)
		return
	}

	errLogout = failureResponseStructure(logoutInfo)
	return
}

// Call UserDetails function to fetch user details from Firstock
func (fs *firstock) UserDetails(userId string) (userDetailsResponse *UserDetailsResponse, errUserDetails *ErrorResponseModel) {
	var userDetailsRequest UserDetailsRequest
	userDetailsResponse = &UserDetailsResponse{}
	userDetailsRequest.UserId = userId

	// Read jKey for userId from config.json
	userDetailsRequest.JKey, errUserDetails = readJkey(userId)
	if userDetailsRequest.JKey == "" {
		return
	}

	userDetails, code, _ := thefirstock.UserDetailsFunction(userDetailsRequest)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(userId)
	} else if code == status_internal_server_error {
		errUserDetails = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(userDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, userDetailsResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}
	errUserDetails = failureResponseStructure(userDetails)
	return
}

func (fs *firstock) PlaceOrder(req PlaceOrderRequest) (placeOrderResponse *PlaceOrderResponse, errPlaceOrder *ErrorResponseModel) {
	placeOrderResponse = &PlaceOrderResponse{}

	// Read jKey for userId from config.json
	jkey, errPlaceOrder := readJkey(req.UserId)
	if jkey == "" {
		return
	}

	reqBody := PlaceOrderRequestBody{
		UserId:          req.UserId,
		JKey:            jkey,
		Exchange:        req.Exchange,
		Retention:       req.Retention,
		Product:         req.Product,
		PriceType:       req.PriceType,
		TradingSymbol:   req.TradingSymbol,
		TransactionType: req.TransactionType,
		Price:           req.Price,
		TriggerPrice:    req.TriggerPrice,
		Quantity:        req.Quantity,
		Remarks:         req.Remarks,
	}

	placeOrderDetails, code, _ := thefirstock.PlaceOrderFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(req.UserId)
	} else if code == status_internal_server_error {
		errPlaceOrder = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(placeOrderDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, placeOrderResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}
	errPlaceOrder = failureResponseStructure(placeOrderDetails)
	return
}

func (fs *firstock) OrderMargin(req OrderMarginRequest) (orderMarginResponse *OrderMarginResponse, errOrderMargin *ErrorResponseModel) {
	orderMarginResponse = &OrderMarginResponse{}

	// Read jKey for userId from config.json
	jkey, errOrderMargin := readJkey(req.UserId)
	if jkey == "" {
		return
	}

	reqBody := OrderMarginRequestBody{
		UserId:          req.UserId,
		JKey:            jkey,
		Exchange:        req.Exchange,
		TransactionType: req.TransactionType,
		Product:         req.Product,
		TradingSymbol:   req.TradingSymbol,
		Quantity:        req.Quantity,
		PriceType:       req.PriceType,
		Price:           req.Price,
	}

	orderMarginDetails, code, _ := thefirstock.OrderMarginFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(req.UserId)
	} else if code == status_internal_server_error {
		errOrderMargin = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(orderMarginDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, orderMarginResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errOrderMargin = failureResponseStructure(orderMarginDetails)
	return
}

func (fs *firstock) SingleOrderHistory(req OrderRequest) (singleOrderHistoryResponse *SingleOrderHistoryResponse, errSingleOrderHistory *ErrorResponseModel) {
	singleOrderHistoryResponse = &SingleOrderHistoryResponse{}

	// Read jKey for userId from config.json
	jkey, errSingleOrderHistory := readJkey(req.UserId)
	if jkey == "" {
		return
	}

	reqBody := OrderRequestBody{
		UserId:      req.UserId,
		JKey:        jkey,
		OrderNumber: req.OrderNumber,
	}

	singleOrderHistoryDetails, code, _ := thefirstock.SingleOrderHistoryFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(req.UserId)
	} else if code == status_internal_server_error {
		errSingleOrderHistory = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(singleOrderHistoryDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, singleOrderHistoryResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errSingleOrderHistory = failureResponseStructure(singleOrderHistoryDetails)
	return
}

func (fs *firstock) CancelOrder(req OrderRequest) (cancelOrderResponse *CancelOrderResponse, errCancelOrder *ErrorResponseModel) {
	cancelOrderResponse = &CancelOrderResponse{}

	// Read jKey for userId from config.json
	jkey, errCancelOrder := readJkey(req.UserId)
	if jkey == "" {
		return
	}

	reqBody := OrderRequestBody{
		UserId:      req.UserId,
		JKey:        jkey,
		OrderNumber: req.OrderNumber,
	}

	cancelOrderDetails, code, _ := thefirstock.CancelOrderFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(req.UserId)
	} else if code == status_internal_server_error {
		errCancelOrder = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(cancelOrderDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, cancelOrderResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errCancelOrder = failureResponseStructure(cancelOrderDetails)
	return
}

func (fs *firstock) ModifyOrder(req ModifyOrderRequest) (modifyOrderResponse *ModifyOrderResponse, errModifyOrder *ErrorResponseModel) {
	modifyOrderResponse = &ModifyOrderResponse{}

	// Read jKey for userId from config.json
	jkey, errModifyOrder := readJkey(req.UserId)
	if jkey == "" {
		return
	}

	reqBody := ModifyOrderRequestBody{
		UserId:         req.UserId,
		JKey:           jkey,
		OrderNumber:    req.OrderNumber,
		PriceType:      req.PriceType,
		TradingSymbol:  req.TradingSymbol,
		Price:          req.Price,
		TriggerPrice:   req.TriggerPrice,
		Quantity:       req.Quantity,
		Product:        req.Product,
		Retention:      req.Retention,
		Mkt_protection: req.Mkt_protection,
	}

	modifyOrderDetails, code, _ := thefirstock.ModifyOrderFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(req.UserId)
	} else if code == status_internal_server_error {
		errModifyOrder = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(modifyOrderDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, modifyOrderResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errModifyOrder = failureResponseStructure(modifyOrderDetails)
	return
}

func (fs *firstock) TradeBook(userId string) (tradeBookResponse *TradeBookResponse, errTradeBook *ErrorResponseModel) {
	tradeBookResponse = &TradeBookResponse{}

	// Read jKey for userId from config.json
	jkey, errTradeBook := readJkey(userId)
	if jkey == "" {
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	tradeBookDetails, code, _ := thefirstock.TradeBookFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(userId)
	} else if code == status_internal_server_error {
		errTradeBook = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(tradeBookDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, tradeBookResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errTradeBook = failureResponseStructure(tradeBookDetails)
	return
}

func (fs *firstock) RMSLmit(userId string) (rmsLmitResponse *RmsLimitResponse, errRmsLimit *ErrorResponseModel) {
	rmsLmitResponse = &RmsLimitResponse{}
	// Read jKey for userId from config.json
	jkey, errRmsLimit := readJkey(userId)
	if jkey == "" {
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	rmsLimitDetails, code, _ := thefirstock.RmsLimitFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(userId)
	} else if code == status_internal_server_error {
		errRmsLimit = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(rmsLimitDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, rmsLmitResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errRmsLimit = failureResponseStructure(rmsLimitDetails)
	return
}

func (fs *firstock) PositionBook(userId string) (positionBookResponse *PositionBookResponse, errPositionBook *ErrorResponseModel) {

	positionBookResponse = &PositionBookResponse{}
	// Read jKey for userId from config.json
	jkey, errPositionBook := readJkey(userId)
	if jkey == "" {
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	positionBookDetails, code, _ := thefirstock.PositionBookFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(userId)
	} else if code == status_internal_server_error {
		errPositionBook = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(positionBookDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, positionBookResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errPositionBook = failureResponseStructure(positionBookDetails)
	return
}

func (fs *firstock) Holdings(userId string) (holdingsResponse *HoldingsResponse, errHoldings *ErrorResponseModel) {
	holdingsResponse = &HoldingsResponse{}
	// Read jKey for userId from config.json
	jkey, errHoldings := readJkey(userId)
	if jkey == "" {
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	holdingsDetails, code, _ := thefirstock.HoldingsFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(userId)
	} else if code == status_internal_server_error {
		errHoldings = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(holdingsDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, holdingsResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errHoldings = failureResponseStructure(holdingsDetails)
	return
}

func (fs *firstock) HoldingsDetails(userId string) (holdingsResponse *HoldingsDetailsResponse, errHoldings *ErrorResponseModel) {
	holdingsResponse = &HoldingsDetailsResponse{}
	// Read jKey for userId from config.json
	jkey, errHoldings := readJkey(userId)
	if jkey == "" {
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	holdingsDetails, code, _ := thefirstock.HoldingsDetailsFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(userId)
	} else if code == status_internal_server_error {
		errHoldings = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(holdingsDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, holdingsResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errHoldings = failureResponseStructure(holdingsDetails)
	return
}

func (fs *firstock) OrderBook(userId string) (orderBookResponse *OrderBookResponse, errOrderBook *ErrorResponseModel) {
	orderBookResponse = &OrderBookResponse{}
	// Read jKey for userId from config.json
	jkey, errOrderBook := readJkey(userId)
	if jkey == "" {
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	orderBookDetails, code, _ := thefirstock.OrderBookFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(userId)
	} else if code == status_internal_server_error {
		errOrderBook = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(orderBookDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, orderBookResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errOrderBook = failureResponseStructure(orderBookDetails)
	return
}

func (fs *firstock) GetExpiry(getExpiryRequest GetInfoRequest) (getExpiryResponse *GetExpiryResponse, errGetExpiry *ErrorResponseModel) {
	getExpiryResponse = &GetExpiryResponse{}
	// Read jKey for userId from config.json
	jkey, errGetExpiry := readJkey(getExpiryRequest.UserId)
	if jkey == "" {
		return
	}

	reqBody := GetInfoRequestBody{
		UserId:        getExpiryRequest.UserId,
		JKey:          jkey,
		Exchange:      getExpiryRequest.Exchange,
		TradingSymbol: getExpiryRequest.TradingSymbol,
	}

	getExpiryDetails, code, _ := thefirstock.GetExpiryFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(getExpiryRequest.UserId)
	} else if code == status_internal_server_error {
		errGetExpiry = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(getExpiryDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, getExpiryResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errGetExpiry = failureResponseStructure(getExpiryDetails)
	return
}

func (fs *firstock) BrokerageCalculator(brokerageCalculatorRequest BrokerageCalculatorRequest) (brokerageCalculatorResponse *BrokerageCalculatorResponse, errBrockerageCalculator *ErrorResponseModel) {
	brokerageCalculatorResponse = &BrokerageCalculatorResponse{}
	// Read jKey for userId from config.json
	jkey, errBrockerageCalculator := readJkey(brokerageCalculatorRequest.UserId)
	if jkey == "" {
		return
	}

	reqBody := BrokerageCalculatorRequestBody{
		UserId:          brokerageCalculatorRequest.UserId,
		JKey:            jkey,
		Exchange:        brokerageCalculatorRequest.Exchange,
		TradingSymbol:   brokerageCalculatorRequest.TradingSymbol,
		TransactionType: brokerageCalculatorRequest.TransactionType,
		Product:         brokerageCalculatorRequest.Product,
		Quantity:        brokerageCalculatorRequest.Quantity,
		Price:           brokerageCalculatorRequest.Price,
		StrikePrice:     brokerageCalculatorRequest.StrikePrice,
		InstName:        brokerageCalculatorRequest.InstName,
		LotSize:         brokerageCalculatorRequest.LotSize,
	}

	brockerageCalculatorDetails, code, _ := thefirstock.BrokerageCalculatorFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(brokerageCalculatorRequest.UserId)
	} else if code == status_internal_server_error {
		errBrockerageCalculator = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(brockerageCalculatorDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, brokerageCalculatorResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errBrockerageCalculator = failureResponseStructure(brockerageCalculatorDetails)
	return
}

func (fs *firstock) BasketMargin(basketMarginRequest BasketMarginRequest) (basketMarginResponse *BasketMarginResponse, errbasketMargin *ErrorResponseModel) {
	basketMarginResponse = &BasketMarginResponse{}
	// Read jKey for userId from config.json
	jkey, errbasketMargin := readJkey(basketMarginRequest.UserId)
	if jkey == "" {
		return
	}

	reqBody := BasketMarginRequestBody{
		UserId:           basketMarginRequest.UserId,
		JKey:             jkey,
		Exchange:         basketMarginRequest.Exchange,
		TradingSymbol:    basketMarginRequest.TradingSymbol,
		TransactionType:  basketMarginRequest.TransactionType,
		Product:          basketMarginRequest.Product,
		Quantity:         basketMarginRequest.Quantity,
		Price:            basketMarginRequest.Price,
		PriceType:        basketMarginRequest.PriceType,
		BasketListParams: basketMarginRequest.BasketListParams,
	}

	basketMarginDetails, code, _ := thefirstock.BasketMarginFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(basketMarginRequest.UserId)
	} else if code == status_internal_server_error {
		errbasketMargin = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(basketMarginDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, basketMarginResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errbasketMargin = failureResponseStructure(basketMarginDetails)
	return
}

func (fs *firstock) GetSecurityInfo(getSecurityInfoRequest GetInfoRequest) (getSecurityInfoResponse *GetSecurityInfoResponse, errGetSecurityInfo *ErrorResponseModel) {
	getSecurityInfoResponse = &GetSecurityInfoResponse{}
	// Read jKey for userId from config.json
	jkey, errGetSecurityInfo := readJkey(getSecurityInfoRequest.UserId)
	if jkey == "" {
		return
	}

	reqBody := GetInfoRequestBody{
		UserId:        getSecurityInfoRequest.UserId,
		JKey:          jkey,
		Exchange:      getSecurityInfoRequest.Exchange,
		TradingSymbol: getSecurityInfoRequest.TradingSymbol,
	}

	getSecurityInfoDetails, code, _ := thefirstock.GetSecurityInfoFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(getSecurityInfoRequest.UserId)
	} else if code == status_internal_server_error {
		errGetSecurityInfo = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(getSecurityInfoDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, getSecurityInfoResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errGetSecurityInfo = failureResponseStructure(getSecurityInfoDetails)
	return
}

func (fs *firstock) ProductConversion(productConversionRequest ProductConversionRequest) (productConversionResponse *ProductConversionResponse, errproductConversion *ErrorResponseModel) {
	productConversionResponse = &ProductConversionResponse{}
	// Read jKey for userId from config.json
	jkey, errproductConversion := readJkey(productConversionRequest.UserId)
	if jkey == "" {
		return
	}
	msgFlag := strings.TrimSpace(productConversionRequest.MessageFlag)
	if msgFlag != "" {
		switch msgFlag {
		case "1", "2", "3", "4":
		default:
			errproductConversion = failureResponseStructure(map[string]interface{}{
				"status": "failed",
				"code":   "400",
				"error": map[string]interface{}{
					"message": product_conversion_error_message,
				},
			})
			return
		}
	}

	reqBody := ProductConversionRequestBody{
		UserId:          productConversionRequest.UserId,
		JKey:            jkey,
		TradingSymbol:   productConversionRequest.TradingSymbol,
		Exchange:        productConversionRequest.Exchange,
		PreviousProduct: productConversionRequest.PreviousProduct,
		Product:         productConversionRequest.Product,
		Quantity:        productConversionRequest.Quantity,
	}

	productConversionDetails, code, _ := thefirstock.ProductConversionFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(productConversionRequest.UserId)
	} else if code == status_internal_server_error {
		errproductConversion = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(productConversionDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, productConversionResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errproductConversion = failureResponseStructure(productConversionDetails)
	return
}

// Market Connect
func (fs *firstock) GetQuote(getQuoteRequest GetInfoRequest) (getQuoteResponse *GetQuoteResponse, errGetQuote *ErrorResponseModel) {
	getQuoteResponse = &GetQuoteResponse{}
	// Read jKey for userId from config.json
	jkey, errGetQuote := readJkey(getQuoteRequest.UserId)
	if jkey == "" {
		return
	}

	reqBody := GetInfoRequestBody{
		UserId:        getQuoteRequest.UserId,
		JKey:          jkey,
		Exchange:      getQuoteRequest.Exchange,
		TradingSymbol: getQuoteRequest.TradingSymbol,
	}

	getQuoteDetails, code, _ := thefirstock.GetQuoteFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(getQuoteRequest.UserId)
	} else if code == status_internal_server_error {
		errGetQuote = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(getQuoteDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, getQuoteResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errGetQuote = failureResponseStructure(getQuoteDetails)
	return

}

func (fs *firstock) GetQuoteLtp(getQuoteLtpRequest GetInfoRequest) (getQuoteLtpResponse *GetQuoteLtpResponse, errGetQuoteLtp *ErrorResponseModel) {
	getQuoteLtpResponse = &GetQuoteLtpResponse{}
	// Read jKey for userId from config.json
	jkey, errGetQuoteLtp := readJkey(getQuoteLtpRequest.UserId)
	if jkey == "" {
		return
	}

	reqBody := GetInfoRequestBody{
		UserId:        getQuoteLtpRequest.UserId,
		JKey:          jkey,
		Exchange:      getQuoteLtpRequest.Exchange,
		TradingSymbol: getQuoteLtpRequest.TradingSymbol,
	}

	getQuoteLtpDetails, code, _ := thefirstock.GetQuoteLtpFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(getQuoteLtpRequest.UserId)
	} else if code == status_internal_server_error {
		errGetQuoteLtp = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(getQuoteLtpDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, getQuoteLtpResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errGetQuoteLtp = failureResponseStructure(getQuoteLtpDetails)
	return
}

func (fs *firstock) GetMultiQuotes(getMultiQuotesRequest GetMultiQuotesRequest) (getMultiQuotesResponse *GetMultiQuotesResponse, errGetMultiQuotes *ErrorResponseModel) {
	getMultiQuotesResponse = &GetMultiQuotesResponse{}
	// Read jKey for userId from config.json
	jkey, errGetMultiQuotes := readJkey(getMultiQuotesRequest.UserId)
	if jkey == "" {
		return
	}

	reqBody := GetMultiQuotesRequestBody{
		UserId: getMultiQuotesRequest.UserId,
		JKey:   jkey,
		Data:   getMultiQuotesRequest.Data,
	}

	getMultiQuotesDetails, code, _ := thefirstock.GetMultiQuotesFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(getMultiQuotesRequest.UserId)
	} else if code == status_internal_server_error {
		errGetMultiQuotes = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(getMultiQuotesDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, getMultiQuotesResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errGetMultiQuotes = failureResponseStructure(getMultiQuotesDetails)
	return

}

func (fs *firstock) GetMultiQuotesLtp(getMultiQuotesLtpRequest GetMultiQuotesRequest) (getMultiQuotesLtpResponse *GetMultiQuotesLtpResponse, errGetMultiQuotesLtp *ErrorResponseModel) {
	getMultiQuotesLtpResponse = &GetMultiQuotesLtpResponse{}
	// Read jKey for userId from config.json
	jkey, errGetMultiQuotesLtp := readJkey(getMultiQuotesLtpRequest.UserId)
	if jkey == "" {
		return
	}

	reqBody := GetMultiQuotesRequestBody{
		UserId: getMultiQuotesLtpRequest.UserId,
		JKey:   jkey,
		Data:   getMultiQuotesLtpRequest.Data,
	}

	getMultiQuotesLtpDetails, code, _ := thefirstock.GetMultiQuotesLtpFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(getMultiQuotesLtpRequest.UserId)
	} else if code == status_internal_server_error {
		errGetMultiQuotesLtp = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(getMultiQuotesLtpDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, getMultiQuotesLtpResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errGetMultiQuotesLtp = failureResponseStructure(getMultiQuotesLtpDetails)
	return

}

func (fs *firstock) IndexList(userId string) (indexListResponse *IndexListResponse, errIndexList *ErrorResponseModel) {
	indexListResponse = &IndexListResponse{}
	// Read jKey for userId from config.json
	jkey, errIndexList := readJkey(userId)
	if jkey == "" {
		return
	}

	reqBody := BaseRequest{
		UserId: userId,
		JKey:   jkey,
	}

	indexListDetails, code, _ := thefirstock.IndexListFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(userId)
	} else if code == status_internal_server_error {
		errIndexList = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(indexListDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, indexListResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errIndexList = failureResponseStructure(indexListDetails)
	return

}

func (fs *firstock) SearchScrips(searchScripsRequest SearchScripsRequest) (searchScripsResponse *SearchScripsResponse, errSearchScrips *ErrorResponseModel) {
	searchScripsResponse = &SearchScripsResponse{}
	// Read jKey for userId from config.json
	jkey, errSearchScrips := readJkey(searchScripsRequest.UserId)
	if jkey == "" {
		return
	}

	reqBody := SearchScripsBody{
		UserId: searchScripsRequest.UserId,
		JKey:   jkey,
		SText:  searchScripsRequest.SText,
	}

	searchScripsDetails, code, _ := thefirstock.SearchScripsFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(searchScripsRequest.UserId)
	} else if code == status_internal_server_error {
		errSearchScrips = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(searchScripsDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, searchScripsResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errSearchScrips = failureResponseStructure(searchScripsDetails)
	return
}

func (fs *firstock) OptionChain(optionChainRequest OptionChainRequest) (optionChainResponse *OptionChainResponse, errOptionChain *ErrorResponseModel) {
	optionChainResponse = &OptionChainResponse{}
	// Read jKey for userId from config.json
	jkey, errOptionChain := readJkey(optionChainRequest.UserId)
	if jkey == "" {
		return
	}

	reqBody := OptionChainRequestBody{
		UserId:      optionChainRequest.UserId,
		JKey:        jkey,
		Exchange:    optionChainRequest.Exchange,
		Symbol:      optionChainRequest.Symbol,
		Expiry:      optionChainRequest.Expiry,
		Count:       optionChainRequest.Count,
		StrikePrice: optionChainRequest.StrikePrice,
	}

	optionChainDetails, code, _ := thefirstock.OptionChainFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(optionChainRequest.UserId)
	} else if code == status_internal_server_error {
		errOptionChain = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(optionChainDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, optionChainResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errOptionChain = failureResponseStructure(optionChainDetails)
	return

}

func (fs *firstock) TimePriceSeriesRegularInterval(req TimePriceSeriesIntervalRequest) (timePriceSeriesRegularIntervalResponse *TimePriceSeriesRegularIntervalResponse, errTimePriceSeriesRegularInterval *ErrorResponseModel) {
	timePriceSeriesRegularIntervalResponse = &TimePriceSeriesRegularIntervalResponse{}
	// Read jKey for userId from config.json
	jkey, errTimePriceSeriesRegularInterval := readJkey(req.UserId)
	if jkey == "" {
		return
	}

	reqBody := TimePriceSeriesIntervalRequestBody{
		UserId:        req.UserId,
		JKey:          jkey,
		Exchange:      req.Exchange,
		Interval:      req.Interval,
		TradingSymbol: req.TradingSymbol,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
	}

	timePriceSeriesDetails, code, _ := thefirstock.TimePriceSeriesRegularIntervalFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(req.UserId)
	} else if code == status_internal_server_error {
		errTimePriceSeriesRegularInterval = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(timePriceSeriesDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, timePriceSeriesRegularIntervalResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errTimePriceSeriesRegularInterval = failureResponseStructure(timePriceSeriesDetails)
	return
}

func (fs *firstock) TimePriceSeriesDayInterval(req TimePriceSeriesIntervalRequest) (timePriceSeriesDayIntervalResponse *TimePriceSeriesDayIntervalResponse, errTimePriceSeriesDayInterval *ErrorResponseModel) {
	timePriceSeriesDayIntervalResponse = &TimePriceSeriesDayIntervalResponse{}
	// Read jKey for userId from config.json
	jkey, errTimePriceSeriesDayInterval := readJkey(req.UserId)
	if jkey == "" {
		return
	}

	reqBody := TimePriceSeriesIntervalRequestBody{
		UserId:        req.UserId,
		JKey:          jkey,
		Exchange:      req.Exchange,
		Interval:      req.Interval,
		TradingSymbol: req.TradingSymbol,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
	}

	timePriceSeriesDetails, code, _ := thefirstock.TimePriceSeriesDayIntervalFunction(reqBody)
	if check_if_unauthorized(code) {
		removeJKeyFromConfig(req.UserId)
	} else if code == status_internal_server_error {
		errTimePriceSeriesDayInterval = internalServerErrorResponse()
		return
	} else if code == status_ok {
		jsonData, err := json.Marshal(timePriceSeriesDetails)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		// Unmarshal JSON to struct
		err = json.Unmarshal(jsonData, timePriceSeriesDayIntervalResponse)
		if err != nil {
			return nil, internalServerErrorResponse()
		}
		return
	}

	errTimePriceSeriesDayInterval = failureResponseStructure(timePriceSeriesDetails)
	return

}

type FirstockAPI interface {
	Login(reqBody LoginRequest) (loginResponse *LoginResponse, errRes *ErrorResponseModel)
	Logout(userId string) (logoutResponse *LogoutResponse, errRes *ErrorResponseModel)
	UserDetails(userId string) (userDetailsResponse *UserDetailsResponse, errRes *ErrorResponseModel)
	PlaceOrder(req PlaceOrderRequest) (placeOrderResponse *PlaceOrderResponse, errRes *ErrorResponseModel)
	OrderMargin(req OrderMarginRequest) (orderMarginResponse *OrderMarginResponse, errRes *ErrorResponseModel)
	SingleOrderHistory(req OrderRequest) (singleOrderHistoryResponse *SingleOrderHistoryResponse, errRes *ErrorResponseModel)
	CancelOrder(req OrderRequest) (cancelOrderResponse *CancelOrderResponse, errRes *ErrorResponseModel)
	ModifyOrder(req ModifyOrderRequest) (modifyOrderResponse *ModifyOrderResponse, errRes *ErrorResponseModel)
	TradeBook(userId string) (tradeBookResponse *TradeBookResponse, errRes *ErrorResponseModel)
	RMSLmit(userId string) (rmsLmitResponse *RmsLimitResponse, errRes *ErrorResponseModel)
	PositionBook(userId string) (positionBookResponse *PositionBookResponse, errRes *ErrorResponseModel)
	Holdings(userId string) (holdingsResponse *HoldingsResponse, errRes *ErrorResponseModel)
	HoldingsDetails(userId string) (holdingsResponse *HoldingsDetailsResponse, errRes *ErrorResponseModel)
	OrderBook(userId string) (orderBookResponse *OrderBookResponse, errRes *ErrorResponseModel)
	GetExpiry(getExpiryRequest GetInfoRequest) (getExpiryResponse *GetExpiryResponse, errRes *ErrorResponseModel)
	BrokerageCalculator(brokerageCalculatorRequest BrokerageCalculatorRequest) (brokerageCalculatorResponse *BrokerageCalculatorResponse, errRes *ErrorResponseModel)
	BasketMargin(basketMarginRequest BasketMarginRequest) (basketMarginResponse *BasketMarginResponse, errRes *ErrorResponseModel)
	GetSecurityInfo(getSecurityInfoRequest GetInfoRequest) (getSecurityInfoResponse *GetSecurityInfoResponse, errRes *ErrorResponseModel)
	ProductConversion(productConversionRequest ProductConversionRequest) (productConversionResponse *ProductConversionResponse, errRes *ErrorResponseModel)
	GetQuote(getQuoteRequest GetInfoRequest) (getQuoteResponse *GetQuoteResponse, errRes *ErrorResponseModel)
	GetQuoteLtp(getQuoteLtpRequest GetInfoRequest) (getQuoteLtpResponse *GetQuoteLtpResponse, errRes *ErrorResponseModel)
	GetMultiQuotes(getMultiQuotesRequest GetMultiQuotesRequest) (getMultiQuotesResponse *GetMultiQuotesResponse, errRes *ErrorResponseModel)
	GetMultiQuotesLtp(getMultiQuotesRequest GetMultiQuotesRequest) (getMultiQuotesLtpResponse *GetMultiQuotesLtpResponse, errRes *ErrorResponseModel)
	IndexList(userId string) (indexListResponse *IndexListResponse, errRes *ErrorResponseModel)
	SearchScrips(searchScripsRequest SearchScripsRequest) (searchScripsResponse *SearchScripsResponse, errRes *ErrorResponseModel)
	OptionChain(optionChainRequest OptionChainRequest) (optionChainResponse *OptionChainResponse, errRes *ErrorResponseModel)
	TimePriceSeriesRegularInterval(req TimePriceSeriesIntervalRequest) (timePriceSeriesRegularIntervalResponse *TimePriceSeriesRegularIntervalResponse, errRes *ErrorResponseModel)
	TimePriceSeriesDayInterval(req TimePriceSeriesIntervalRequest) (timePriceSeriesDayIntervalResponse *TimePriceSeriesDayIntervalResponse, errRes *ErrorResponseModel)
	InitializeWebSockets(userId string, model WebSocketModel) (conn *websocket.Conn, errRes *ErrorResponseModel)
	CloseWebSocket(conn *websocket.Conn) (err *ErrorResponseModel)
	Subscribe(conn *websocket.Conn, data []string) (err *ErrorResponseModel)
	Unsubscribe(conn *websocket.Conn, data []string) (err *ErrorResponseModel)
}

// internal instance, not exported
var firstockAPI FirstockAPI = &firstock{}

// -------------------------------- WebSocket --------------------------------------------------
func InitializeWebSockets(userId string, model WebSocketModel) (conn *websocket.Conn, errRes *ErrorResponseModel) {
	return firstockAPI.InitializeWebSockets(userId, model)
}
func CloseWebSocket(conn *websocket.Conn) (err *ErrorResponseModel) {
	return firstockAPI.CloseWebSocket(conn)
}
func Subscribe(conn *websocket.Conn, data []string) (err *ErrorResponseModel) {
	return firstockAPI.Subscribe(conn, data)
}
func Unsubscribe(conn *websocket.Conn, data []string) (err *ErrorResponseModel) {
	return firstockAPI.Unsubscribe(conn, data)
}
func Login(reqBody LoginRequest) (loginResponse *LoginResponse, errRes *ErrorResponseModel) {
	return firstockAPI.Login(reqBody)
}

func Logout(userId string) (logoutResponse *LogoutResponse, errRes *ErrorResponseModel) {
	return firstockAPI.Logout(userId)
}

func UserDetails(userId string) (userDetailsResponse *UserDetailsResponse, errRes *ErrorResponseModel) {
	return firstockAPI.UserDetails(userId)
}

func PlaceOrder(req PlaceOrderRequest) (placeOrderResponse *PlaceOrderResponse, errRes *ErrorResponseModel) {
	return firstockAPI.PlaceOrder(req)
}

func OrderMargin(req OrderMarginRequest) (orderMarginResponse *OrderMarginResponse, errRes *ErrorResponseModel) {
	return firstockAPI.OrderMargin(req)
}

func SingleOrderHistory(req OrderRequest) (singleOrderHistoryResponse *SingleOrderHistoryResponse, errRes *ErrorResponseModel) {
	return firstockAPI.SingleOrderHistory(req)
}

func CancelOrder(req OrderRequest) (cancelOrderResponse *CancelOrderResponse, errRes *ErrorResponseModel) {
	return firstockAPI.CancelOrder(req)
}

func ModifyOrder(req ModifyOrderRequest) (modifyOrderResponse *ModifyOrderResponse, errRes *ErrorResponseModel) {
	return firstockAPI.ModifyOrder(req)
}

func TradeBook(userId string) (tradeBookResponse *TradeBookResponse, errRes *ErrorResponseModel) {
	return firstockAPI.TradeBook(userId)
}

func RMSLmit(userId string) (rmsLmitResponse *RmsLimitResponse, errRes *ErrorResponseModel) {
	return firstockAPI.RMSLmit(userId)
}

func PositionBook(userId string) (positionBookResponse *PositionBookResponse, errRes *ErrorResponseModel) {
	return firstockAPI.PositionBook(userId)
}

func Holdings(userId string) (holdingsResponse *HoldingsResponse, errRes *ErrorResponseModel) {
	return firstockAPI.Holdings(userId)
}

func HoldingsDetails(userId string) (holdingsResponse *HoldingsDetailsResponse, errRes *ErrorResponseModel) {
	return firstockAPI.HoldingsDetails(userId)
}

func OrderBook(userId string) (orderBookResponse *OrderBookResponse, errRes *ErrorResponseModel) {
	return firstockAPI.OrderBook(userId)
}

func GetExpiry(getExpiryRequest GetInfoRequest) (getExpiryResponse *GetExpiryResponse, errRes *ErrorResponseModel) {
	return firstockAPI.GetExpiry(getExpiryRequest)
}

func BrokerageCalculator(brokerageCalculatorRequest BrokerageCalculatorRequest) (brokerageCalculatorResponse *BrokerageCalculatorResponse, errRes *ErrorResponseModel) {
	return firstockAPI.BrokerageCalculator(brokerageCalculatorRequest)
}

func BasketMargin(basketMarginRequest BasketMarginRequest) (basketMarginResponse *BasketMarginResponse, errRes *ErrorResponseModel) {
	return firstockAPI.BasketMargin(basketMarginRequest)
}

func GetSecurityInfo(getSecurityInfoRequest GetInfoRequest) (getSecurityInfoResponse *GetSecurityInfoResponse, errRes *ErrorResponseModel) {
	return firstockAPI.GetSecurityInfo(getSecurityInfoRequest)
}

func ProductConversion(productConversionRequest ProductConversionRequest) (productConversionResponse *ProductConversionResponse, errRes *ErrorResponseModel) {
	return firstockAPI.ProductConversion(productConversionRequest)
}

func GetQuote(getQuoteRequest GetInfoRequest) (getQuoteResponse *GetQuoteResponse, errRes *ErrorResponseModel) {
	return firstockAPI.GetQuote(getQuoteRequest)
}

func GetQuoteLtp(getQuoteLtpRequest GetInfoRequest) (getQuoteLtpResponse *GetQuoteLtpResponse, errRes *ErrorResponseModel) {
	return firstockAPI.GetQuoteLtp(getQuoteLtpRequest)
}

func GetMultiQuotes(getMultiQuotesRequest GetMultiQuotesRequest) (getMultiQuotesResponse *GetMultiQuotesResponse, errRes *ErrorResponseModel) {
	return firstockAPI.GetMultiQuotes(getMultiQuotesRequest)
}

func GetMultiQuotesLtp(getMultiQuotesRequest GetMultiQuotesRequest) (getMultiQuotesLtpResponse *GetMultiQuotesLtpResponse, errRes *ErrorResponseModel) {
	return firstockAPI.GetMultiQuotesLtp(getMultiQuotesRequest)
}

func IndexList(userId string) (indexListResponse *IndexListResponse, errRes *ErrorResponseModel) {
	return firstockAPI.IndexList(userId)
}

func SearchScrips(searchScripsRequest SearchScripsRequest) (searchScripsResponse *SearchScripsResponse, errRes *ErrorResponseModel) {
	return firstockAPI.SearchScrips(searchScripsRequest)
}

func OptionChain(optionChainRequest OptionChainRequest) (optionChainResponse *OptionChainResponse, errRes *ErrorResponseModel) {
	return firstockAPI.OptionChain(optionChainRequest)
}

func TimePriceSeriesRegularInterval(req TimePriceSeriesIntervalRequest) (timePriceSeriesRegularIntervalResponse *TimePriceSeriesRegularIntervalResponse, errRes *ErrorResponseModel) {
	return firstockAPI.TimePriceSeriesRegularInterval(req)
}

func TimePriceSeriesDayInterval(req TimePriceSeriesIntervalRequest) (timePriceSeriesDayIntervalResponse *TimePriceSeriesDayIntervalResponse, errRes *ErrorResponseModel) {
	return firstockAPI.TimePriceSeriesDayInterval(req)
}
