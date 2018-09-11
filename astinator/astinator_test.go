package astinator

import (
	"testing"

	"github.com/sanity-io/litter"
)

var jscode = `
function multiply(p1,p2) {
    var a = p1;
    var b = p2;
    var result = p1 *p2;

    return result;
}
`

func TestAST(t *testing.T) {
	uast := TextToUAST(jscode)

	litter.Dump(uast)
}

func TestText(t *testing.T) {
	uast := TextToUAST(jscode)

	res := UASTToText(uast)

	litter.Dump(res)

}
