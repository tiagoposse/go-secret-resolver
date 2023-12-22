package resolvers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type AwsResolver struct {
	client *secretsmanager.Client
}

func NewAwsResolver(ctx context.Context) (*AwsResolver, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &AwsResolver{
		client: secretsmanager.NewFromConfig(cfg),
	}, nil
}

func (aws *AwsResolver) ResolveSecret(ctx context.Context, secretName string) (string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	}

	output, err := aws.client.GetSecretValue(ctx, input)
	if err != nil {
		return "", err
	}

	return *output.SecretString, nil
}
