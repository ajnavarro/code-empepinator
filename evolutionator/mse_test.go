package evolutionator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMSE(t *testing.T) {
	dataset := []float64{0.5, -1, 0, -1, -1, -1}
	res := MSE(dataset)
	fixed := fmt.Sprintf("%.3f", res)
	assert.Equal(t, "0.708", fixed)
}
