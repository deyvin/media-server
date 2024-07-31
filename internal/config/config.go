package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

type Config struct {
	DBConfig *DBConfig
	AWS      *AWSConfig
}

type DBConfig struct {
	DBUrl string
}

type AWSConfig struct {
	S3Client *s3.Client
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func LoadDBConfig() (*DBConfig, error) {
	loadEnv()

	config := &DBConfig{
		DBUrl: os.Getenv("DATABASE_URL"),
	}

	return config, nil
}

func LoadAWSConfig() (*AWSConfig, error) {
	loadEnv()

	awsRegion := os.Getenv("AWS_REGION")
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			awsAccessKey,
			awsSecretKey,
			"",
		)),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	return &AWSConfig{
		S3Client: s3Client,
	}, nil
}
