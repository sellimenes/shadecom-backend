package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func UploadImages(c *gin.Context) {
    form, err := c.MultipartForm()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get multipart form"})
        return
    }

    files := form.File["files[]"]

    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load AWS config"})
        return
    }

    client := s3.NewFromConfig(cfg)
    uploader := manager.NewUploader(client)

    for _, file := range files {
        f, err := file.Open()
        if err != nil {
            log.Printf("failed to open file %q, %v", file.Filename, err)
            continue
        }

        _, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
            Bucket: aws.String("shadecom"),
            Key:    aws.String(file.Filename),
            Body:   f,
            ACL:    "public-read",
        })

        if err != nil {
            log.Printf("failed to upload file, %v", err)
            continue
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Files uploaded successfully",
    })
}
