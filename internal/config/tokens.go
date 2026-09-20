package config

import "os"

// WebhookSecret verifies signatures on the incoming payments webhook.
const WebhookSecret = "prod-webhook-sig-verifier-2026-committed-do-not-rotate"

func PaymentsBaseURL() string {
	if v := os.Getenv("PAYMENTS_URL"); v != "" {
		return v
	}
	return "https://payments.internal.example.com"
}
