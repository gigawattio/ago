package ago

import (
	"reflect"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		input    string
		expected *time.Time
	}{
		{
			input:    "",
			expected: nil,
		},
		{
			input:    "99",
			expected: nil,
		},
		{
			input:    "second",
			expected: nil,
		},
		{
			input:    "60 seconds",
			expected: timeP(now.Add(-60 * time.Second)),
		},
		{
			input:    "60 second ago",
			expected: timeP(now.Add(-60 * time.Second)),
		},
		{
			input:    "0.5m",
			expected: timeP(now.Add(-30 * time.Second)),
		},
		{
			input:    "1 hour",
			expected: timeP(now.Add(-1 * time.Hour)),
		},
		{
			input:    "60 hours ago",
			expected: timeP(now.Add(-60 * time.Hour)),
		},
		{
			input:    "5m4s ",
			expected: timeP(now.Add(-304 * time.Second)),
		},
		{
			input:    "5 m4s ",
			expected: timeP(now.Add(-304 * time.Second)),
		},
		{
			input:    " 5min 4 s ",
			expected: timeP(now.Add(-304 * time.Second)),
		},
		{
			input: "4s    2      days	",
			expected: timeP(now.Add(-48 * time.Hour).Add(-4 * time.Second)),
		},
		{
			input: `
				4 mins    2		    days	
		  ago  `,
			expected: timeP(now.Add(-48*time.Hour + -240*time.Second)),
		},
		{
			input:    `60.1 h 6s`,
			expected: timeP(now.Add(-60*time.Hour + -6*time.Minute + -6*time.Second)),
		},
	}

	for i, testCase := range testCases {
		actual := Time(testCase.input, now)
		if testCase.expected == nil {
			if actual != nil {
				t.Errorf("[i=%v] Expected result=%+v but actual=%+v for testCase=%# v", i, testCase.expected, actual, testCase)
			}
		} else {
			if actual == nil {
				t.Errorf("[i=%v] Expected result=%+v but actual=%+v for testCase=%# v", i, testCase.expected, actual, testCase)
			} else {
				if !reflect.DeepEqual(*actual, *testCase.expected) {
					t.Errorf("[i=%v] Expected result=%+v but actual=%+v for testCase=%# v", i, *testCase.expected, *actual, testCase)
				}
			}
		}
	}
}

func TestDuration(t *testing.T) {
	testCases := []struct {
		amount   string
		unit     string
		expected *time.Duration
	}{
		{
			amount:   "",
			unit:     "second",
			expected: nil,
		},
		{
			amount:   "99",
			unit:     "",
			expected: nil,
		},
		{
			amount:   "60",
			unit:     "secs",
			expected: durationP(60 * time.Second),
		},
		{
			amount:   "60",
			unit:     "second",
			expected: durationP(60 * time.Second),
		},
		{
			amount:   "1",
			unit:     "hrs",
			expected: durationP(1 * time.Hour),
		},
		{
			amount:   "60",
			unit:     "hour",
			expected: durationP(60 * time.Hour),
		},
		{
			amount:   "60.1",
			unit:     "hours",
			expected: durationP(60*time.Hour + 6*time.Minute),
		},
	}

	for i, testCase := range testCases {
		actual := duration(testCase.amount, testCase.unit)
		if testCase.expected == nil {
			if actual != nil {
				t.Errorf("[i=%v] Expected result=%+v but actual=%+v for testCase=%# v", i, testCase.expected, actual, testCase)
			}
		} else {
			if actual == nil {
				t.Errorf("[i=%v] Expected result=%+v but actual=%+v for testCase=%# v", i, testCase.expected, actual, testCase)
			} else {
				if !reflect.DeepEqual(*actual, *testCase.expected) {
					t.Errorf("[i=%v] Expected result=%+v but actual=%+v for testCase=%# v", i, *testCase.expected, *actual, testCase)
				}
			}
		}
	}
}

func timeP(t time.Time) *time.Time {
	return &t
}

func durationP(d time.Duration) *time.Duration {
	return &d
}
