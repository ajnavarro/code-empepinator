package evolutionator

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ajnavarro/code-empepinator/sandbox"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"
)

const samplesCount = 5

const name = "multiply"

const paramsCount = 2

func TestEvolutionator(t *testing.T) {
	fmt.Println("STARTING")
	var pairs []*Pair
	executor := sandbox.NewJavascript(name)
	for i := 0; i < samplesCount; i++ {
		var params []float64
		for j := 0; j < paramsCount; j++ {
			params = append(params, rand.Float64()*float64(rand.Int()))
		}

		result, err := executor.Execute(jscode, params...)
		if err != nil {
			panic(err)
		}

		p := &Pair{
			Input:  params,
			Output: result,
		}

		pairs = append(pairs, p)
	}

	fmt.Println("PAIRS GENERATED CORRECTLY")

	err := executor.Parse(jscode)
	assert.NoError(t, err)

	res, err := Optimize(executor.AST, pairs, name)

	assert.NoError(t, err)
	litter.Dump(res)
}
