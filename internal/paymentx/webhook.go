package paymentx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
)

// WebhookSecret is compared against the X-Signature header.
// BUG 1: hardcoded webhook signing secret.
const WebhookSecret = "whsec_TESTfake9876543210abcdef"

// HandleWebhook verifies the HMAC signature on the request body and forwards
// the payload to the internal processor.
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad body", http.StatusBadRequest)
		return
	}
	given := r.Header.Get("X-Signature")
	mac := hmac.New(sha256.New, []byte(WebhookSecret))
	mac.Write(body)
	expected := hex.EncodeToString(mac.Sum(nil))
	// BUG 2: string == comparison on HMACs — timing attack.
	if given != expected {
		http.Error(w, "bad signature", http.StatusForbidden)
		return
	}
	_, _ = w.Write([]byte("ok"))
}
