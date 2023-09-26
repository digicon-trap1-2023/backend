package infrastructure

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"encoding/base64"
	"encoding/json"

	"github.com/digicon-trap1-2023/backend/util"

	_ "image/jpeg"
	_ "image/png"
)

type S3Client struct {
	lambdaUrl      string
	fileBucketName string
}

type S3PutRequest struct {
	File        string `json:"file"`
	FileId      string `json:"file_id"`
	ContentType string `json:"content_type"`
}

func postRequest(url string, req *S3PutRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println(response)
		return fmt.Errorf("failed to post request: %s", response.Status)
	}

	return nil
}

func NewS3Client() (*S3Client, error) {
	return &S3Client{
		lambdaUrl:      util.ReadEnvs("AWS_LAMBDA_URL"),
		fileBucketName: util.ReadEnvs("AWS_S3_BUCKET_NAME"),
	}, nil
}

func (client *S3Client) PutObjectMock(key string, data io.ReadSeeker) error {
	return nil
}

func (client *S3Client) PutObject(key string, fh *multipart.FileHeader) (string, error) {
	file, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	contentType := fh.Header.Get("Content-Type")
	
	parts := strings.Split(contentType, "/")
	extension := ""
	if len(parts) > 1 {
		extension = parts[1]
	}

	req := &S3PutRequest{
		File:        encoded,
		FileId:      key,
		ContentType: contentType,
	}

	if err := postRequest(client.lambdaUrl, req); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s", client.GetObjectUrl(key), extension), nil
}

func (client *S3Client) GetObjectUrl(objectKey string) string {
	return fmt.Sprintf("https://%s.s3.amazonaws.com/image/%s", client.fileBucketName, objectKey)
}
