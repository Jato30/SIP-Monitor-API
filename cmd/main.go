package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Jato30/SIP-Monitor-API/internal/api"
	"github.com/Jato30/SIP-Monitor-API/internal/sip"
)

func main() {
	metrics := api.NewMetrics()

	// Simulate parsing SIP messages every second (replace with real parser later)
	go func() {
		mockMessages := []string{
			"INVITE sip:bob@domain.com SIP/2.0\r\n",
			"BYE sip:bob@domain.com SIP/2.0\r\n",
			"REGISTER sip:server.com SIP/2.0\r\n",
		}
		for {
			for _, raw := range mockMessages {
				msg, err := sip.Parse(raw)
				if err == nil && msg.Method != "" {
					metrics.Increment(msg.Method)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	http.HandleFunc("/health", api.HealthHandler)
	http.HandleFunc("/metrics", metrics.MetricsHandler)

	fmt.Println("Starting HTTP server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
