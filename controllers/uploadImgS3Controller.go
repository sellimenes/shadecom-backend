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

func UploadImages(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files[]"]

	for _, file := range files {
		// Dosyayı aç
		f, err := file.Open()
		if err != nil {
			log.Printf("failed to open file %q, %v", file.Filename, err)
			continue
		}

		// Dosyayı yükle
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		client := s3.NewFromConfig(cfg)
		uploader := manager.NewUploader(client)

		result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String("shadecom"),
			Key:    aws.String(file.Filename),
			Body:   f,
			ACL:    "public-read",
		})

		if err != nil {
			log.Printf("failed to upload file, %v", err)
			continue
		}

		c.JSON(200, gin.H{
			"result": result,
		})
	}
}
