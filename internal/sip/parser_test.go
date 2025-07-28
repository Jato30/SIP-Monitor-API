package sip

import (
	"testing"
)

func TestParseBasicSIP(t *testing.T) {
	raw := `INVITE sip:bob@domain.com SIP/2.0
Via: SIP/2.0/UDP pc33.domain.com
From: "Alice" <sip:alice@domain.com>;tag=1928301774
To: <sip:bob@domain.com>
Call-ID: a84b4c76e66710@pc33.domain.com
CSeq: 314159 INVITE

`

	msg, err := Parse(raw)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if msg.Method != "INVITE" {
		t.Errorf("Expected INVITE, got %s", msg.Method)
	}
	if msg.Headers["Call-ID"] != "a84b4c76e66710@pc33.domain.com" {
		t.Errorf("Missing Call-ID")
	}
}
