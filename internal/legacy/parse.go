package legacy

import (
	"strconv"
	"strings"
)

// Record is a parsed CSV record: id,amount,note.
type Record struct {
	ID     int
	Amount int
	Note   string
}

// ParseRecord parses "id,amount,note".
// BUG: strconv.Atoi errors ignored — a non-numeric field silently becomes 0.
// Downstream logic treats 0 as valid, hiding malformed input.
func ParseRecord(line string) Record {
	parts := strings.SplitN(line, ",", 3)
	if len(parts) < 3 {
		return Record{}
	}
	id, _ := strconv.Atoi(parts[0])
	amount, _ := strconv.Atoi(parts[1])
	return Record{
		ID:     id,
		Amount: amount,
		Note:   parts[2],
	}
}
