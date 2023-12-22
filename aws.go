package resolvers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type AwsResolver struct {
	client *secretsmanager.Client
}

func NewAwsResolver() (*AwsResolver, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	return &AwsResolver{
		client: secretsmanager.NewFromConfig(cfg),
	}, nil
}

func (aws *AwsResolver) ResolveSecret(secretName string) (string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	}

	output, err := aws.client.GetSecretValue(context.TODO(), input)
	if err != nil {
		return "", err
	}

	return *output.SecretString, nil
}
