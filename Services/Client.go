package Services

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/medic-basic/s3-test/Config"
)

var StorageImpl Storage

type Storage struct {
	s3Client       *s3.Client
	feedBucketName string
	authBucketName string
	feedPrefix     string
	authPrefix     string
}

func InitStorage() {
	s3Client := s3.NewFromConfig(Config.S3Cfg)
	StorageImpl = Storage{
		s3Client:       s3Client,
		feedBucketName: "medic-test-feed-images",
		authBucketName: "medic-test-auth-images",
		feedPrefix:     "#Patient:",
		authPrefix:     "#MedicAuth:",
	}
}
