package main

import (
	"context"
	"fmt"
	"log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {

	endpoint := "minio-dp-p01"

	fmt.Println("Enter access key: ")
	var accessKey string
	fmt.Scan(&accessKey)

	fmt.Println("Enter access key: ")
	var secretKey string
	fmt.Scan(&secretKey)

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	bucketNames := []string{}

	for _, bucketName := range bucketNames {
		exists, err := client.BucketExists(context.Background(), bucketName)
		if err != nil {
			log.Printf("Unable to ping bucket %s: %v\n", bucketName, err)
		}
		if exists {
			fmt.Printf("Bucket %s exists", bucketName)
		} else {
			fmt.Printf("Bucket %s does not exist", bucketName)
		}
	}

}
