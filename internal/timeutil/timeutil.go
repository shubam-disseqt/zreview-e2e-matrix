package timeutil

import "time"

// FormatDurationCompact renders a duration like "1h23m" or "45s".
// It omits zero units and never returns an empty string.
func FormatDurationCompact(d time.Duration) string {
	if d < time.Second {
		return "<1s"
	}
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	if hours > 0 && minutes > 0 {
		return formatHM(hours, minutes)
	}
	if hours > 0 {
		return formatH(hours)
	}
	if minutes > 0 {
		return formatMS(minutes, seconds)
	}
	return formatS(seconds)
}

func formatH(h int) string     { return itoa(h) + "h" }
func formatHM(h, m int) string { return itoa(h) + "h" + itoa(m) + "m" }
func formatMS(m, s int) string { return itoa(m) + "m" + itoa(s) + "s" }
func formatS(s int) string     { return itoa(s) + "s" }

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [11]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}
