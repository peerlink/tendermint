package types

import (
	"sort"
	"time"
)

func Now() time.Time {
	return time.Now().Round(0).UTC()
}

func WeightedMedian(timeToVotingPower map[time.Time]int64) (res time.Time) {
	totalVotingPower := int64(0)
	for _, v := range timeToVotingPower {
		totalVotingPower += v
	}

	median := totalVotingPower / 2

	// To store the keys in slice in sorted order
	var keys []time.Time
	for k := range timeToVotingPower {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i].UnixNano() < keys[j].UnixNano() })

	for _, k := range keys {
		if median <= timeToVotingPower[k] {
			res = k
			break
		}
		median -= timeToVotingPower[k]
	}
	return
}
