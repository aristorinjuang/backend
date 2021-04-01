package helpers

import (
	"regexp"
	"time"
)

// GetSeconds converts the runtime to the second duration.
func GetSeconds(runtime string) *time.Duration {
	re := regexp.MustCompile(` min.*| m`)
	trimmedSeconds := re.ReplaceAllString(runtime, "m")
	seconds, _ := time.ParseDuration(trimmedSeconds)

	return &seconds
}
