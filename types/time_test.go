package types

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTimeWindow(t *testing.T) {
	const (
		tolerance = 33
		times     = 1000
	)

	for i := int64(0); i < times; i++ {
		c1 := int64(100000 + i)
		w1 := TimeWindow(c1, tolerance)
		for j := int64(0); j < times/2; j++ {
			c2 := c1 - times/4 + j
			w2 := TimeWindow(c2, tolerance)
			if math.Abs(float64(w1-w2)) < float64(tolerance) {
				require.Equal(t, w1, w2)
			} else {
				require.NotEqual(t, w1, w2)
			}
		}
	}
}
