package resolvers

import (
	"context"
	"errors"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

type AzureResolver struct {
	client *azsecrets.Client
}

func NewAzureResolver() (*AzureResolver, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	endpoint := os.Getenv("AZURE_VAULT_ENDPOINT")
	if endpoint == "" {
		return nil, errors.New("azure vault endpoint not set")
	}

	client, err := azsecrets.NewClient(endpoint, cred, nil)
	if err != nil {
		return nil, err
	}

	return &AzureResolver{
		client: client,
	}, nil
}

func (az *AzureResolver) ResolveSecret(secretName string) (string, error) {
	secret, err := az.client.GetSecret(context.Background(), secretName, "", nil)
	if err != nil {
		return "", err
	}

	return *secret.Value, nil
}
