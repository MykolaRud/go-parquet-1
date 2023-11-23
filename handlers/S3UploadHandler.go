package handlers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"os"
	"parquet_1/infrastructures"
	"path/filepath"
)

func NewS3UploadHandler() *S3UploadHandler {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(infrastructures.GetS3Credentials()),
	)
	if err != nil {
		panic("Couldn't init S3 config")
	}
	cfg.BaseEndpoint = aws.String("http://127.0.0.1:9000")
	client := s3.NewFromConfig(cfg)

	return &S3UploadHandler{
		client: client,
	}
}

type S3UploadHandler struct {
	client *s3.Client
}

func (handler *S3UploadHandler) Save(ctx context.Context, filename, bucketName string) {

	log.Println("start saving")

	uploadFile, err := os.Open(filename)
	if err != nil {
		log.Println("Failed opening file", filename, err)
	}
	defer uploadFile.Close()

	_, err = handler.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filepath.Base(filename)),
		Body:   uploadFile,
	})
	if err != nil {
		fmt.Println(err)
	}

	log.Println("File has been uploaded")
}

func (handler *S3UploadHandler) UploadFile(filename, bucket string) string {
	//uploadFile, err := os.Open(filename)
	//if err != nil {
	//	log.Println("Failed opening file", filename, err)
	//	return ""
	//}
	//defer uploadFile.Close()
	//
	//uploader := manager.NewUploader(handler.client)
	//result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
	//	Bucket: aws.String(bucket),
	//	Key:    aws.String(strings.Replace(filename, "\\", "-", -1)),
	//	Body:   io.Reader(uploadFile),
	//})
	//
	//log.Println(result.Location)
	//
	//if err != nil {
	//	log.Fatalln("Failed to upload file ", filename, err)
	//}
	//
	//return result.Location

	return ""
}
