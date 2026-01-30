// Copyright (c) [2025] [Firstock]
// SPDX-License-Identifier: MIT
package Firstock

const j_key = "jkey"
const status_ok = "200"
const status_internal_server_error = "500"

// URLs
const base_url = "https://api.firstock.in/V1"

const login_url = base_url + "/login"
const logout_url = base_url + "/logout"
const user_details_url = base_url + "/userDetails"
const place_order_url = base_url + "/placeOrder"
const order_margin_url = base_url + "/orderMargin"
const single_order_history_url = base_url + "/singleOrderHistory"
const cancel_order_url = base_url + "/cancelOrder"
const modify_order_url = base_url + "/modifyOrder"
const trade_book_url = base_url + "/tradeBook"
const rms_limit_url = base_url + "/limit"
const position_book_url = base_url + "/positionBook"
const holdings_url = base_url + "/holdings"
const holdings_details_url = base_url + "/holdingsDetails"
const order_book_url = base_url + "/orderBook"
const get_expiry_url = base_url + "/getExpiry"
const brokerage_calculator_url = base_url + "/brokerageCalculator"
const basket_margin_url = base_url + "/basketMargin"
const get_security_info = base_url + "/securityInfo"
const product_conversion_url = base_url + "/productConversion"
const get_quote_url = base_url + "/getQuote"
const get_quote_ltp_url = base_url + "/getQuote/ltp"
const get_multi_quotes_url = base_url + "/getMultiQuotes"
const get_multi_quotes_ltp_url = base_url + "/getMultiQuotes/ltp"
const index_list_url = base_url + "/indexList"
const search_scrips_url = base_url + "/searchScrips"
const option_chain_url = base_url + "/optionChain"
const time_price_series_url = base_url + "/timePriceSeries"

const scheme = "wss"
const host = "socket.firstock.in"
const path = "/V2/ws"
const srcVal = "developer-api"
const maxWebsocketConnectionRetries = 20
const timeInterval = 1
const accept_encoding = "gzip, deflate, br, zstd"
const accept_language = "en-US,en;q=0.9"
const no_cache = "no-cache"
const origin = "https://firstock.in"

// Error messages
const please_login_to_firstock = "Please login to Firstock"
const internal_server_error = "Internal Server Error"
const product_conversion_error_message = "message_flag can only have values `1`, `2`, `3` or `4`"
