// Copyright (c) [2025] [Firstock]
// SPDX-License-Identifier: MIT
package Firstock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type apifunctions struct{}

func (fs *apifunctions) LoginFunction(
	reqBody LoginRequest) (loginResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(login_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &loginResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) LogoutFunction(
	reqBody LogoutRequest) (logoutResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(logout_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &logoutResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) UserDetailsFunction(reqBody UserDetailsRequest) (userDetailsResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(user_details_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &userDetailsResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return

}

func (fs *apifunctions) PlaceOrderFunction(req PlaceOrderRequestBody) (placeOrderResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(place_order_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &placeOrderResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) OrderMarginFunction(req OrderMarginRequestBody) (orderMarginResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(order_margin_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &orderMarginResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) SingleOrderHistoryFunction(req OrderRequestBody) (singleOrderHistoryResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(single_order_history_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &singleOrderHistoryResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) CancelOrderFunction(req OrderRequestBody) (cancelOrderResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(cancel_order_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &cancelOrderResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) ModifyOrderFunction(req ModifyOrderRequestBody) (modifyOrderResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(modify_order_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &modifyOrderResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) TradeBookFunction(req BaseRequest) (tradeBookResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(trade_book_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &tradeBookResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) RmsLimitFunction(req BaseRequest) (rmsLimitResponse map[string]interface{}, statusCode string, err error) {
	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(rms_limit_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &rmsLimitResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) PositionBookFunction(req BaseRequest) (positionBookResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(position_book_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &positionBookResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) HoldingsFunction(req BaseRequest) (holdingsResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(holdings_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &holdingsResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) HoldingsDetailsFunction(req BaseRequest) (holdingsResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(holdings_details_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &holdingsResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) OrderBookFunction(req BaseRequest) (orderBookResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(order_book_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &orderBookResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) GetExpiryFunction(req GetInfoRequestBody) (getExpiryResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(get_expiry_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &getExpiryResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) BrokerageCalculatorFunction(req BrokerageCalculatorRequestBody) (brokerageCalculatorResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(brokerage_calculator_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &brokerageCalculatorResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) BasketMarginFunction(req BasketMarginRequestBody) (basketMarginResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(basket_margin_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &basketMarginResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) GetSecurityInfoFunction(req GetInfoRequestBody) (getSecurityInfoResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(get_security_info, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &getSecurityInfoResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) ProductConversionFunction(req ProductConversionRequestBody) (productConversionResponse map[string]interface{}, statusCode string, err error) {
	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(product_conversion_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &productConversionResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

// ---------------------------------------Connect---------------------------------
func (fs *apifunctions) GetQuoteFunction(req GetInfoRequestBody) (getQuoteResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(get_quote_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &getQuoteResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) GetQuoteLtpFunction(req GetInfoRequestBody) (getQuoteLtpResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(get_quote_ltp_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &getQuoteLtpResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) GetMultiQuotesFunction(req GetMultiQuotesRequestBody) (getMultiQuotesResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(get_multi_quotes_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &getMultiQuotesResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) GetMultiQuotesLtpFunction(req GetMultiQuotesRequestBody) (getMultiQuotesLtpResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}
	resp, err := http.Post(get_multi_quotes_ltp_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &getMultiQuotesLtpResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) IndexListFunction(req BaseRequest) (indexListResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(index_list_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &indexListResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) SearchScripsFunction(req SearchScripsBody) (searchScripsResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(search_scrips_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &searchScripsResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) OptionChainFunction(req OptionChainRequestBody) (optionChainResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(option_chain_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &optionChainResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) TimePriceSeriesRegularIntervalFunction(req TimePriceSeriesIntervalRequestBody) (timePriceSeriesRegularIntervalResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(time_price_series_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &timePriceSeriesRegularIntervalResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}

func (fs *apifunctions) TimePriceSeriesDayIntervalFunction(req TimePriceSeriesIntervalRequestBody) (timePriceSeriesDayIntervalResponse map[string]interface{}, statusCode string, err error) {

	jsonPayload, err := json.Marshal(req)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(time_price_series_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, "500", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "500", fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &timePriceSeriesDayIntervalResponse); err != nil {
		return nil, "500", fmt.Errorf("failed to parse response: %w", err)
	}

	statusCode = strconv.Itoa(resp.StatusCode)
	return
}
