// /pkg/repository/s3_repository.go
package repository

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Repository interface {
	ListFiles(bucket string) ([]types.Object, error)
}

type s3Repository struct {
	s3Client *s3.Client
}

func NewS3Repository(s3Client *s3.Client) S3Repository {
	return &s3Repository{s3Client: s3Client}
}

func (repo *s3Repository) ListFiles(bucket string) ([]types.Object, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	}

	var files []types.Object
	paginator := s3.NewListObjectsV2Paginator(repo.s3Client, input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Printf("failed to get page, %v", err)
			return nil, err
		}

		files = append(files, page.Contents...)
	}

	return files, nil
}
