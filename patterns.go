package textpatterns

// InitiationShare returns a's share of observed conversation starts.
// Invalid or empty totals return zero.
func InitiationShare(a, b int) float64 {
	if a < 0 || b < 0 || a+b == 0 {
		return 0
	}
	return float64(a) / float64(a+b)
}
