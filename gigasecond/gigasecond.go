package gigasecond

import "time"

const GIGASECOND = time.Duration(1e9) * time.Second;

func AddGigasecond(t time.Time) time.Time {
	return t.Add(GIGASECOND);
}
