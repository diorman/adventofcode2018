package main

import "testing"

func TestGetGuardIDMultipliedByMinute(t *testing.T) {
	logs := []string{
		"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:05] falls asleep",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 00:55] wakes up",
		"[1518-11-01 23:58] Guard #99 begins shift",
		"[1518-11-02 00:40] falls asleep",
		"[1518-11-02 00:50] wakes up",
		"[1518-11-03 00:05] Guard #10 begins shift",
		"[1518-11-03 00:24] falls asleep",
		"[1518-11-03 00:29] wakes up",
		"[1518-11-04 00:02] Guard #99 begins shift",
		"[1518-11-04 00:36] falls asleep",
		"[1518-11-04 00:46] wakes up",
		"[1518-11-05 00:03] Guard #99 begins shift",
		"[1518-11-05 00:45] falls asleep",
		"[1518-11-05 00:55] wakes up",
	}

	tt := []struct {
		id       string
		strategy strategy
		result   int
	}{
		{"strategy 1", strategy1, 2400},
		{"strategy 2", strategy2, 4455},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			if result := getGuardIDMultipliedByMinute(logs, tc.strategy); result != tc.result {
				t.Errorf("got: %v, want: %v", result, tc.result)
			}
		})
	}
}
