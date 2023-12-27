package controllers

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context){
	file, err := c.FormFile("file")

	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO());

	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	f, openErr := file.Open()

	if openErr != nil {
		log.Printf("failed to open file %q, %v", file.Filename, openErr)
		return
	}

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("shadecom"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL: 	"public-read",
	})

	if uploadErr != nil {
		log.Printf("failed to upload file, %v", uploadErr)
		return
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}