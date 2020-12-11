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
	endpoint := "b2904236d6.zicp.vip:8098"
	accessKeyID := "minioadmin"
	secretAccessKey := "Pass@7899"
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
	bucketName := "my-oss-test"

	var bt bytes.Buffer

	file, err := minioClient.GetObject(ctx, bucketName, "fundtet00KEeko8RfViV5_01.json", minio.GetObjectOptions{})
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

}