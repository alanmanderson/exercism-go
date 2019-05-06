// Returns the time plus 1 gigasecond
package gigasecond

import "time"

// Return t + 1 gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
