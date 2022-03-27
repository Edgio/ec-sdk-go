// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package mathhelper

import (
	"sort"
	"testing"
	"time"
)

func TestCalculateSleepWithJitter(t *testing.T) {
	maxTestExpected := 1 * time.Second
	notMaxTestExpected := 60 * time.Second

	cases := []struct {
		name        string
		min         time.Duration
		max         time.Duration
		attemptNum  int
		expected    *time.Duration
		notExpected *time.Duration
	}{
		{
			name:       "Testing Max Invocation",
			min:        20 * time.Second, // min more than max to force max
			max:        1 * time.Second,
			attemptNum: 3,
			expected:   &maxTestExpected,
		},
		{
			name:        "Testing Not Max Invocation",
			min:         2 * time.Second,
			max:         60 * time.Second,
			attemptNum:  3,
			notExpected: &notMaxTestExpected,
		},
	}

	for _, c := range cases {
		actual := CalculateSleepWithJitter(c.min, c.max, c.attemptNum)

		if c.expected != nil {
			if *c.expected != actual {
				t.Fatalf(
					"'%s': Expected %d but got %d", c.name, c.expected, actual)
			}
		} else if c.notExpected != nil {
			if *c.notExpected == actual {
				t.Fatalf(
					"'%s': Did not Expect %d and got %d",
					c.name, *c.notExpected, actual)
			}
		} else {
			t.Fatalf("Both expected and notExpected test values were nil")
		}
	}
}

func TestRandBetween(t *testing.T) {
	cases := []struct {
		name       string
		min        float64
		max        float64
		iterations int
	}{
		{
			name:       "Testing Normal Invocation",
			min:        5,
			max:        20,
			iterations: 10,
		},
		// TODO Figure out a method of testing randomness...
	}

	for _, c := range cases {
		actuals := []float64{}
		for i := 0; i < c.iterations; i++ {
			actual := RandBetween(c.min, c.max)
			actuals = append(actuals, actual)
		}

		sort.Float64s(actuals)

		for i, actual := range actuals {
			if i+1 == len(actuals) {
				break
			}

			if actual == actuals[i+1] {
				t.Fatalf(
					"'%s': Found duplicate in randomness check %v",
					c.name, actual)
			}
		}
	}
}
