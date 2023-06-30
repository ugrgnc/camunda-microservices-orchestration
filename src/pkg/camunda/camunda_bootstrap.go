package camunda

import (
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"github.com/ugrgnc/camunda-microservices-orchestration/bootstrap"
)

func createCamundaConnection(cfg *bootstrap.CamundaConfig) (zbc.Client, error) {

	credentialsProvider, _ := createCredentialsProvider(cfg)
	client, err := zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:      cfg.ZEEBE_ADDRESS,
		CredentialsProvider: credentialsProvider,
	})
	if err != nil {
		panic(err)
	}

	return client, nil
}

func createCredentialsProvider(cfg *bootstrap.CamundaConfig) (*zbc.OAuthCredentialsProvider, error) {

	cp, err := zbc.NewOAuthCredentialsProvider(&zbc.OAuthProviderConfig{
		ClientID:     cfg.ZEEBE_CLIENT_ID,
		ClientSecret: cfg.ZEEBE_CLIENT_SECRET,
		Audience:     cfg.ZEEBE_TOKEN_AUDIENCE,
	})
	if err != nil {
		panic(err)
	}

	return cp, nil
}
