func (client Client) PutObject(ctx context.Context, key string, data []byte) error {
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

func (client Client) GetObjectUrl(objectKey string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", client.fileBucketName, client.region, objectKey)
}
