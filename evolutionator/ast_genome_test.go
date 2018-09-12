package evolutionator

import (
	"testing"

	"github.com/sanity-io/litter"
)

func TestASTGenome(t *testing.T) {
	g := jsGenome{}

	litter.Dump(g)
}
