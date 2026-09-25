package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/currency"
	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/pricing"
)

// priceHandler answers /price?base=12.34&pct=10 with the discounted amount.
func priceHandler(w http.ResponseWriter, r *http.Request) {
	pct, err := strconv.ParseInt(r.URL.Query().Get("pct"), 10, 64)
	if err != nil {
		http.Error(w, "pct must be an integer", http.StatusBadRequest)
		return
	}
	q, err := pricing.Compute(r.URL.Query().Get("base"), pct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, currency.Format(q.FinalCents))
}

// totalHandler sums repeated ?amount= params, e.g. /total?amount=1.00&amount=2.50.
func totalHandler(w http.ResponseWriter, r *http.Request) {
	var sum int64
	for _, a := range r.URL.Query()["amount"] {
		c, err := currency.ToCents(a)
		if err != nil {
			http.Error(w, "bad amount: "+a, http.StatusBadRequest)
			return
		}
		sum += c
	}
	fmt.Fprintln(w, currency.Format(sum))
}
