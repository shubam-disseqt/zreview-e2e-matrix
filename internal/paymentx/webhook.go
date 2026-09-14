package paymentx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"os"
)

// webhookSecret is loaded once at package init from STRIPE_WEBHOOK_SECRET.
var webhookSecret = os.Getenv("STRIPE_WEBHOOK_SECRET")

// HandleWebhook verifies the HMAC signature on the request body.
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if webhookSecret == "" {
		http.Error(w, "server misconfigured", http.StatusInternalServerError)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad body", http.StatusBadRequest)
		return
	}
	given, err := hex.DecodeString(r.Header.Get("X-Signature"))
	if err != nil {
		http.Error(w, "bad signature", http.StatusForbidden)
		return
	}
	mac := hmac.New(sha256.New, []byte(webhookSecret))
	mac.Write(body)
	if !hmac.Equal(given, mac.Sum(nil)) {
		http.Error(w, "bad signature", http.StatusForbidden)
		return
	}
	_, _ = w.Write([]byte("ok"))
}
