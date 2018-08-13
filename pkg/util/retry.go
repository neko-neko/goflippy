package util

import (
	"math/rand"
	"time"
)

// seed given randmizer
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// Maximum retry wait seconds
const maximumWaitSeconds = 32

// Operator is retry function
type Operator func() error

// Retry function
func Retry(retry int, f Operator) error {
	var err error

	for i := 1; i <= retry; i++ {
		err = f()
		if err == nil {
			return nil
		}

		base := (1 << uint(i) * 1000) + rnd.Int63n(1000)
		interval := time.Duration(Min(base, maximumWaitSeconds*1000)) * time.Millisecond
		time.Sleep(interval)
	}

	return err
}
