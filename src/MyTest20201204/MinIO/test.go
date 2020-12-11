package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func main() {
	ctx := context.Background()
	endpoint := "b2904236d6.zicp.vip:8098"
	accessKeyID := "minioadmin"
	secretAccessKey := "Pass@7899"
	useSSL := false

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// 创建一个叫mymusic的存储桶。
	bucketName := "my-oss-test"
	location := "cn-north-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	buckets, _ := minioClient.ListBuckets(ctx)
	for _, bucket := range buckets {
		fmt.Println("bucket:", bucket)
	}

	// 上传一个zip文件。
	objectName := "333.jpeg"
	filePath := "E:\\20.06.16Project\\GoTest\\src\\MyTest20201204\\MinIO\\333.jpeg"
	contentType := "jpeg"

	// minioClient.
	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
	log.Printf("-------------------------")
	log.Printf("===>%+v\n", n)

	// 根据名称获取内容
	// object, err := minioClient.GetObject(ctx,bucketName, "111.jpeg", minio.GetObjectOptions{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// localFile, err := os.Create("./local-file.jpg")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if _, err = io.Copy(localFile, object); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
