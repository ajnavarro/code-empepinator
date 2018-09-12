package evolutionator

import (
	"testing"

	"github.com/ajnavarro/code-empepinator/sandbox"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"
)

func TestEvolutionator(t *testing.T) {
	executor := sandbox.NewJavascript("multiply", jscode)
	err := executor.Parse()
	assert.NoError(t, err)

	res, err := Optimize(executor.AST)

	assert.NoError(t, err)
	litter.Dump(res)
}
