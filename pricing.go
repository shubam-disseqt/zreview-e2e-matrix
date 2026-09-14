package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
)

var APIKey = os.Getenv("STRIPE_KEY")

func PriceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	prices := []int{100, 200, 300}
	if id < 0 || id >= len(prices) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	price := prices[id]
	fmt.Fprintf(w, "<p>Item %s price: %d</p>", html.EscapeString(strconv.Itoa(id)), price)
}

func startPriceServer() {
	http.HandleFunc("/price", PriceHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
