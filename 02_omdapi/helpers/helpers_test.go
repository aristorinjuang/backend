package helpers

import (
	"testing"
	"time"
)

var samples = []struct {
	unformatted string
	formatted   string
}{
	{"120 min", "120m"},
	{"240 min", "240m"},
	{"76 min", "76m"},
	{"0 min", "0m"},
	{"1200 min", "1200m"},
	{"99999 min", "99999m"},
	{"99 min ", "99m"},
	{"99 minute", "99m"},
	{"99 minutes", "99m"},
	{"99 m", "99m"},
	{"99", "0m"},
	{"1 mad", "0m"},
	{"1 hour", "0m"},
	{"1 h", "0m"},
	{"1m", "1m"},
}

func TestGetSeconds(t *testing.T) {
	for _, s := range samples {
		unformatted := GetSeconds(s.unformatted)
		formatted, _ := time.ParseDuration(s.formatted)

		if *unformatted != formatted {
			t.Errorf(`got "%s" from "%s", expected "%s"`, unformatted, s.unformatted, formatted)
		}
	}
}

func BenchmarkGetSeconds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range samples {
			GetSeconds(s.unformatted)
		}
	}
}
