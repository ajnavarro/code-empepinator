package astinator

import (
	"fmt"

	"github.com/sanity-io/litter"

	"gopkg.in/bblfsh/client-go.v2"
	"gopkg.in/bblfsh/client-go.v2/tools"
	"gopkg.in/bblfsh/sdk.v1/protocol"
	"gopkg.in/bblfsh/sdk.v1/uast"
	protocol2 "gopkg.in/bblfsh/sdk.v2/protocol"
)

func TextToUAST(jscode string) *uast.Node {
	client, err := bblfsh.NewClient("0.0.0.0:9432")
	if err != nil {
		panic(err)
	}

	res, err := client.NewParseRequest().Language("javascript").Mode(protocol2.Mode_Semantic).Content(jscode).Do()
	if err != nil {
		panic(err)
	}

	// Always check the Response.Status before further processing!
	if res.Status != protocol.Ok {
		panic("Parsing failed")
	}

	return res.UAST
}

func UASTToText(uast *uast.Node) string {
	iter, err := tools.NewIterator(uast, tools.PreOrder)
	if err != nil {
		panic(err)
	}
	defer iter.Dispose()

	var result string

	for node := range iter.Iterate() {

		if node.InternalType == "Alias" {
			litter.Dump(node)
			return ""
		}
		continue

		if node.Token != "" {
			switch node.InternalType {
			case "FunctionGroup":
				result = fmt.Sprintf("%s %s", result, "function")
			case "Identifier":
				result = fmt.Sprintf("%s %s", result, node.Token)
			}
		}
		fmt.Printf("InternalType: %v\n", node.InternalType)
		fmt.Printf("Properties: %v\n", node.Properties)
		//fmt.Printf("Children: %v\n", node.Children)
		fmt.Printf("Token: %v\n", node.Token)
		fmt.Printf("StartPosition: %v\n", node.StartPosition)
		fmt.Printf("EndPosition: %v\n", node.EndPosition)
		fmt.Printf("Roles: %v\n", node.Roles)
		fmt.Println()

	}

	return result
}
