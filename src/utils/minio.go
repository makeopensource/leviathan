package utils

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client *minio.Client
	Bucket string
}

func NewMinioClient(endpoint, accessKey, secretKey, bucket string, useSSL bool) (*MinioClient, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return &MinioClient{Client: client, Bucket: bucket}, nil
}

func (m *MinioClient) CreateBucket(ctx context.Context) error {
	exists, err := m.Client.BucketExists(ctx, m.Bucket)
	if err != nil {
		return err
	}
	if !exists {
		return m.Client.MakeBucket(ctx, m.Bucket, minio.MakeBucketOptions{})
	}
	return nil
}

func (m *MinioClient) DeleteBucket(ctx context.Context) error {
	return m.Client.RemoveBucket(ctx, m.Bucket)
}

func (m *MinioClient) UploadFile(ctx context.Context, objectName, filePath string, contentType string) error {
	_, err := m.Client.FPutObject(ctx, m.Bucket, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	return err
}

func (m *MinioClient) ListBuckets(ctx context.Context) ([]minio.BucketInfo, error) {
	return m.Client.ListBuckets(ctx)
}
