package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWeightedMedian(t *testing.T) {
	m := make(map[time.Time]int64)

	t1 := Now()
	t2 := t1.Add(5 * time.Second)
	t3 := t1.Add(10 * time.Second)

	// total voting power is 100 in all test cases. Voting power of faulty processes is 33.

	m[t1] = 33 // faulty processes
	m[t2] = 40 // correct processes
	m[t3] = 27 // correct processes

	median := WeightedMedian(m)
	assert.Equal(t, t2, median)

	m[t1] = 40 // correct processes
	m[t2] = 27 // correct processes
	m[t3] = 33 // faulty processes

	median = WeightedMedian(m)
	assert.Equal(t, t2, median)

	t4 := t1.Add(15 * time.Second)
	t5 := t1.Add(60 * time.Second)

	m[t1] = 10 // correct processes
	m[t2] = 20 // correct processes
	m[t3] = 23 // faulty processes
	m[t4] = 20 // correct processes
	m[t5] = 10 // faulty processes

	median = WeightedMedian(m)
	assert.Equal(t, t3, median) // median always returns value between values of correct processes
}
