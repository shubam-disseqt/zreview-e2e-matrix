package legacy

import (
	"fmt"
	"strconv"
	"strings"
)

// Record is a parsed CSV record: id,amount,note.
type Record struct {
	ID     int
	Amount int
	Note   string
}

// ParseRecord parses "id,amount,note", returning an error on malformed input.
func ParseRecord(line string) (Record, error) {
	parts := strings.SplitN(line, ",", 3)
	if len(parts) < 3 {
		return Record{}, fmt.Errorf("expected 3 fields, got %d", len(parts))
	}
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return Record{}, fmt.Errorf("invalid id %q: %w", parts[0], err)
	}
	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		return Record{}, fmt.Errorf("invalid amount %q: %w", parts[1], err)
	}
	return Record{ID: id, Amount: amount, Note: parts[2]}, nil
}
