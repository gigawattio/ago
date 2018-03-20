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
			input:    "1 hour",
			expected: timeP(now.Add(-1 * time.Hour)),
		},
		{
			input:    "60 hours ago",
			expected: timeP(now.Add(-60 * time.Hour)),
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
			unit:     "seconds",
			expected: durationP(60 * time.Second),
		},
		{
			amount:   "60",
			unit:     "second",
			expected: durationP(60 * time.Second),
		},
		{
			amount:   "1",
			unit:     "hour",
			expected: durationP(1 * time.Hour),
		},
		{
			amount:   "60",
			unit:     "hours",
			expected: durationP(60 * time.Hour),
		},
	}

	for i, testCase := range testCases {
		actual := Duration(testCase.amount, testCase.unit)
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
