package ago

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var intervalExpr = regexp.MustCompile(`^(?:second|minute|hour|day|week|month|year)s?$`)

// Time takes an input string of the form ".. ## <time-unit> ago" and returns
// the closest possible time value.
//
// When no match is found, nil will be returned.
func Time(s string, relativeTo time.Time) *time.Time {
	tokens := strings.Split(s, " ")
	if len(tokens) < 2 {
		return nil
	}

	prevToken := tokens[0]

	for _, token := range tokens[1:] {
		if intervalExpr.MatchString(token) {
			if d := Duration(prevToken, token); d != nil {
				ts := relativeTo.Add(time.Duration(*d * -1))
				return &ts
			}
		}
	}

	return nil
}

// Duration turns an amount and time unit into a *time.Duration.
//
// When the input cannot be converted into a duration, nil will be returned.
func Duration(amount string, unit string) *time.Duration {
	unit = strings.ToLower(unit)

	base, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return nil
	}

	switch unit {
	case "seconds", "second":
		base *= 1
		break
	case "minutes", "minute":
		base *= 60
		break
	case "hours", "hour":
		base *= 3600
		break
	case "days", "day":
		base *= 86400 // 24 hours.
		break
	case "weeks", "week":
		base *= 604800 // 7 days.
		break
	case "months", "month":
		base *= 2592000 // 30 days.
		break
	case "years", "sear":
		base *= 31536000 // 365 days.
		break
	default:
		return nil
	}

	s := fmt.Sprintf("%.0fs", base)

	d, err := time.ParseDuration(s)
	if err != nil {
		return nil
	}
	return &d
}
