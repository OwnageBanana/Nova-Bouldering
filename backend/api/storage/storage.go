package storage

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageService struct {
	Client      *s3.Client
	Bucket      string
	Endpoint    string
	TokenValue  string
}

func Init() (*StorageService, error) {
	// bucket := os.Getenv("NOVA_R2_BUCKET")
	// if bucket == "" {
	// 	return nil, fmt.Errorf("NOVA_R2_BUCKET environment variable is not set")
	// }

	// accountID := os.Getenv("NOVA_R2_ACCOUNT_ID")
	// if accountID == "" {
	// 	return nil, fmt.Errorf("NOVA_R2_ACCOUNT_ID environment variable is not set")
	// }

	// accessKeyID := os.Getenv("NOVA_R2_ACCESS_KEY_ID")
	// if accessKeyID == "" {
	// 	return nil, fmt.Errorf("NOVA_R2_ACCESS_KEY_ID environment variable is not set")
	// }

	// secretAccessKey := os.Getenv("NOVA_R2_SECRET_ACCESS_KEY")
	// if secretAccessKey == "" {
	// 	return nil, fmt.Errorf("NOVA_R2_SECRET_ACCESS_KEY environment variable is not set")
	// }
	endpoint := os.Getenv("NOVA_R2_ENDPOINT")
	if endpoint == "" {
		return nil, fmt.Errorf("NOVA_R2_ENDPOINT environment variable is not set")
	}
	tokenValue := os.Getenv("NOVA_R2_TOKEN_VALUE") // ig this is the user token instead of account token
	if tokenValue == "" {
		return nil, fmt.Errorf("NOVA_R2_TOKEN_VALUE environment variable is not set")
	}
	accessKeyID := os.Getenv("NOVA_R2_ACCESS_ID")
	if accessKeyID == "" {
		return nil, fmt.Errorf("NOVA_R2_ACCESS_ID environment variable is not set")
	}
	secretAccessKey := os.Getenv("NOVA_R2_SECRET_ACCESS_KEY")
	if secretAccessKey == "" {
		return nil, fmt.Errorf("NOVA_R2_SECRET_ACCESS_KEY environment variable is not set")
	}
	bucket := "nova-bouldering";
	// endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID)

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})

	return &StorageService{
		Client:     client,
		Bucket:     bucket,
		Endpoint:   endpoint,
		TokenValue: tokenValue,
	}, nil
}

func (s *StorageService) Retrieve(ctx context.Context, key string) (io.ReadCloser, error) {
	output, err := s.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve object %s: %w", key, err)
	}
	return output.Body, nil
}

func (s *StorageService) Write(ctx context.Context, key string, body io.Reader, contentType string) error {
	_, err := s.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.Bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return fmt.Errorf("failed to write object %s: %w", key, err)
	}
	return nil
}

func (s *StorageService) Delete(ctx context.Context, key string) error {
	_, err := s.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object %s: %w", key, err)
	}
	return nil
}
