package resolvers

import (
	"errors"
	"fmt"
	"os"
)

type ResolverField struct {
	File   *string `yaml:"file"`
	Value  *string `yaml:"value"`
	Aws    *string `yaml:"aws"`
	Google *string `yaml:"google"`
	Azure  *string `yaml:"azure"`
}

type Resolver struct {
	aws    *AwsResolver
	az     *AzureResolver
	google *GoogleResolver
}

func NewResolver() *Resolver {
	return &Resolver{}
}

// resolve resolves the secret value based on the given name.
func (r *Resolver) Resolve(field *ResolverField) error {
	var err error

	if field.Value != nil {
		return nil
	}

	if field.File != nil {
		val, err := getFileValue(*field.File)
		if err != nil {
			return fmt.Errorf("reading file: %s", *field.File)
		}

		field.Value = &val
		return nil
	}

	// Check for AWS secret
	if field.Aws != nil {
		if r.aws == nil {
			if r.aws, err = NewAwsResolver(); err != nil {
				return fmt.Errorf("creating aws resolver: %w", err)
			}
		}

		val, err := r.aws.ResolveSecret(*field.Aws)
		if err != nil {
			return fmt.Errorf("resolving aws secret: %s", *field.Aws)
		}
		field.Value = &val
		return nil
	}

	// Check for Google secret
	if field.Google != nil {
		if r.google == nil {
			if r.google, err = NewGoogleResolver(); err != nil {
				return fmt.Errorf("creating google resolver: %w", err)
			}
		}

		val, err := r.google.ResolveSecret(*field.Google)
		if err != nil {
			return fmt.Errorf("resolving google secret: %s", *field.Google)
		}
		field.Value = &val
		return nil
	}

	// Check for Azure secret
	if field.Azure != nil {
		if r.az == nil {
			if r.az, err = NewAzureResolver(); err != nil {
				return fmt.Errorf("creating azure resolver: %w", err)
			}
		}

		val, err := r.az.ResolveSecret(*field.Azure)
		if err != nil {
			return fmt.Errorf("resolving azure secret: %s", *field.Azure)
		}

		field.Value = &val
		return nil
	}

	return errors.New("field cannot be resolved")
}

// func getEnvVar(key string) (string, error) {
// 	val, exists := os.LookupEnv(key)
// 	if !exists {
// 		return "", fmt.Errorf("env variable %s not found", key)
// 	}
// 	return val, nil
// }

func getFileValue(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
