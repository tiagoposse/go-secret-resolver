package resolvers

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

type GoogleResolver struct {
	client *secretmanager.Client
}

func NewGoogleResolver() (*GoogleResolver, error) {
	client, err := secretmanager.NewClient(context.TODO())
	if err != nil {
		return nil, err
	}

	return &GoogleResolver{
		client: client,
	}, nil
}

func (google *GoogleResolver) ResolveSecret(secretName string) (string, error) {
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", "your-project-id", secretName),
	}

	result, err := google.client.AccessSecretVersion(context.TODO(), req)
	if err != nil {
		return "", err
	}

	return string(result.Payload.Data), nil
}
