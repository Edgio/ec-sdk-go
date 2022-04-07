package mathhelper

import (
	"math"
	"math/rand"
	"time"
)

func CalculateSleepWithJitter(
	min, max time.Duration,
	attemptNum int,
) time.Duration {
	// Calculate the initial sleep period before jitter attemptNum starts at
	// 0 so we add 1
	sleep := math.Pow(2, float64(attemptNum+1)) * float64(min)

	// The final sleep time will be a random number between sleep/2 and sleep
	sleepWithJitter := sleep/2 + RandBetween(0, sleep/2)

	if sleepWithJitter > float64(max) {
		return max
	}

	return time.Duration(sleepWithJitter)
}

func RandBetween(min float64, max float64) float64 {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + rand.Float64()*(max-min)
}
