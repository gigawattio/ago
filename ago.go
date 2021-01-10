package ago

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	tokenizer = regexp.MustCompile(`([0-9]+|[0-9]*\.[0-9]+)([^0-9.]+)`)
)

// Time takes an input string of the form ".. ## <time-unit> ago" and returns
// the closest possible time value.
//
// When no match is found, nil will be returned.
func Time(s string, relativeTo time.Time) *time.Time {
	s = strings.ToLower(strings.Trim(s, " \r\n\t"))
	s = strings.ReplaceAll(s, "ago", "")

	tokens := tokenizer.FindAllStringSubmatch(s, -1)
	if len(tokens) == 0 {
		return nil
	}

	t := relativeTo

	for _, token := range tokens {
		amount := token[len(token)-2]
		unit := token[len(token)-1]
		//fmt.Fprintf(os.Stderr, "unit=%q amt=%q\n", unit, amount)
		d := duration(amount, unit)
		if d == nil {
			return nil
		}
		t = t.Add(*d * -1)
	}

	return &t
}

// duration turns an amount and time unit into a *time.Duration.
//
// When the input cannot be converted into a duration, nil will be returned.
func duration(amount string, unit string) *time.Duration {
	amount = trimWhitespace(amount)
	unit = trimWhitespace(unit)

	base, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "err parsing %q: %s\n", amount, err)
		return nil
	}

	switch unit {
	case "s", "sec", "secs", "second", "seconds":
		base *= 1
		break
	case "m", "min", "mins", "minute", "minutes":
		base *= 60 // 60 seconds.
		break
	case "h", "hr", "hrs", "hour", "hours":
		base *= 3600 // 60 minutes.
		break
	case "d", "day", "days":
		base *= 86400 // 24 hours.
		break
	case "w", "wk", "wks", "week", "weeks":
		base *= 604800 // 7 days.
		break
	case "mo", "mos", "month", "months":
		base *= 2592000 // 30 days.
		break
	case "y", "yr", "yrs", "year", "years":
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

// trimWhitespace trims all whitespace characters from either end of the input
// string.
func trimWhitespace(s string) string {
	s = strings.ToLower(strings.Trim(s, " \r\n\t"))
	return s
}
