package infrastructure

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/digicon-trap1-2023/backend/util"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	_ "image/jpeg"
	_ "image/png"
)

type S3Client struct {
	s3             *s3.Client
	fileBucketName string
	region         string
}

func NewClient() (*S3Client, error) {
	sdkConfig, err := config.LoadDefaultConfig(context.Background())

	if err != nil {
		sdkConfig, err = config.LoadDefaultConfig(
			context.TODO(),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					util.ReadEnvs("AWS_ACCESS_KEY"),
					util.ReadEnvs("AWS_SECRET_KEY"),
					""),
			),
			config.WithRegion(util.ReadEnvs("AWS_REGION")),
		)
		if err != nil {
			return nil, err
		}
	}

	s3 := s3.NewFromConfig(sdkConfig)
	return &S3Client{
		s3:             s3,
		fileBucketName: util.ReadEnvs("AWS_S3_BUCKET_NAME"),
		region:         util.ReadEnvs("AWS_REGION"),
	}, nil
}
func (client S3Client) PutObject(ctx context.Context, key string, data []byte) error {
	_, err := client.s3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(client.fileBucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
		ACL:    types.ObjectCannedACL(*aws.String("public-read")),
	})
	if err != nil {
		return err
	}

	return nil
}

func (client S3Client) GetObjectUrl(objectKey string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", client.fileBucketName, client.region, objectKey)
}
