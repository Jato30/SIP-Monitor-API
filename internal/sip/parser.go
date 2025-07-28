package sip

import (
	"bufio"
	"strings"
)

type SIPMessage struct {
	Method  string
	URI     string
	Version string
	Headers map[string]string
	Body    string
}

func Parse(raw string) (*SIPMessage, error) {
	scanner := bufio.NewScanner(strings.NewReader(raw))
	msg := &SIPMessage{
		Headers: make(map[string]string),
	}

	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if lineNum == 0 {
			// First line (Request line)
			parts := strings.SplitN(line, " ", 3)
			if len(parts) == 3 {
				msg.Method = parts[0]
				msg.URI = parts[1]
				msg.Version = parts[2]
			}
		} else if line == "" {
			// End of headers
			break
		} else {
			// Headers
			if kv := strings.SplitN(line, ":", 2); len(kv) == 2 {
				msg.Headers[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
			}
		}
		lineNum++
	}

	return msg, nil
}
