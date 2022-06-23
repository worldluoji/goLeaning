package v1

import "example-sdk-go/sdk"

const (
	defaultEndpoint = "127.0.0.1:8083"
	serviceName     = "ladon"
)

type Client struct {
	sdk.Client
}

func NewClient(config *sdk.Config, credential *sdk.Credential) (client *Client, err error) {
	client = &Client{}
	if config == nil {
		config = sdk.NewConfig().WithEndpoint(defaultEndpoint)
	}

	client.Init(serviceName).WithCredential(credential).WithConfig(config)
	return
}

func NewClientWithSecret(secretID, secretKey string) (client *Client, err error) {
	client = &Client{}
	config := sdk.NewConfig().WithEndpoint(defaultEndpoint)
	client.Init(serviceName).WithSecret(secretID, secretKey).WithConfig(config)
	return
}
