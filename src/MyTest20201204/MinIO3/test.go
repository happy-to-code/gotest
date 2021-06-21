package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
)

var minioClient *minio.Client
var ctx context.Context

func init() {
	var err error
	ctx = context.Background()
	endpoint := "121.224.59.167:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false
	// 初使化 minio client对象。
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	// 创建一个叫mymusic的存储桶。
	bucketName := "zyf-test"

	var bt bytes.Buffer

	file, err := minioClient.GetObject(ctx, bucketName, "30018", minio.GetObjectOptions{})
	if err != nil {
		fmt.Printf("err1:%+v\n", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("err2:%+v\n", err)
				return
			}
		}

		bt.WriteString(string(data))
	}
	fmt.Printf("res:%s\n\n", bt.String())
	// fromMinio, err := GetInfoFromMinio("RSR006_0")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("===%s\n", fromMinio)
}

func GetInfoFromMinio(fileName string) ([]byte, error) {
	// 桶名称
	// bucketName := viper.GetString("OSS.BucketName")
	bucketName := "my-oss-test"
	// 给文件名加后缀
	fileName = fileName + ".json"

	var bt bytes.Buffer
	file, err := minioClient.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		bt.Write(data)
	}
	return bt.Bytes(), nil
}
