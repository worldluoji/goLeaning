package main

import (
	"fmt"

	"github.com/ory/ladon"

	"example-sdk-go/sdk"
	testv1 "example-sdk-go/service/test/v1"
)

func main() {
	client, _ := testv1.NewClientWithSecret("XhbY3aCrfjdYcP1OFJRu9xcno8JzSbUIvGE2", "bfJRvlFwsoW9L30DlG87BBW0arJamSeK")

	req := testv1.NewAuthzRequest()
	req.Resource = sdk.String("resources:users")
	req.Action = sdk.String("delete")
	req.Subject = sdk.String("users:admin")
	ctx := ladon.Context(map[string]interface{}{"username": "luoji"})
	req.Context = &ctx

	resp, err := client.Authz(req)
	if err != nil {
		fmt.Println("err1", err)
		return
	}
	fmt.Printf("get response body: `%s`\n", resp.String())
	fmt.Printf("allowed: %v\n", resp.Allowed)
}
