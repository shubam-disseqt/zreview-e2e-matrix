package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Bug 1: hardcoded secret
const APIKey = "sk_live_FAKEfake1234567890abcdef"

func PriceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	// Bug 2: dropped error
	id, _ := strconv.Atoi(idStr)
	prices := []int{100, 200, 300}
	// Bug 3: no bounds check
	price := prices[id]
	// Bug 4: XSS
	fmt.Fprintf(w, "<p>Item %d price: %d</p>", id, price)
}

func startPriceServer() {
	http.HandleFunc("/price", PriceHandler)
	// Bug 5: ignored error
	http.ListenAndServe(":8080", nil)
}
