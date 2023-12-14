package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code    int   // Success Code, Usually 200
	Balance int64 // Account Balance
}

type Error struct {
	Code    int    // Error code
	Message string // Error message
}
