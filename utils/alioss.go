package util

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"hmshop/global"
	"log"
	"mime/multipart"
	"strings"
)

var client *oss.Client // 全局变量用来存储OSS客户端实例
var endpoint string
var accessKeyId string
var accessKeySecret string
var bucketName string
var region string

func InitAliOss() {

	endpoint = global.AppConfig.AliOss.Endpoint
	accessKeyId = global.AppConfig.AliOss.AccessKeyId
	accessKeySecret = global.AppConfig.AliOss.AccessKeySecret
	bucketName = global.AppConfig.AliOss.BucketName
	region = global.AppConfig.AliOss.Region
	var err error
	client, err = oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// 输出客户端信息。
	log.Printf("Client: %#v\n", client)

}

func UploadFile(file *multipart.FileHeader) (string, error) {
	// 生成 UUID 作为文件的唯一标识符
	fileUUID := uuid.New().String()

	// 打开文件
	srcFile, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer srcFile.Close()

	// 获取文件的 MIME 类型
	fileExt := file.Filename[strings.LastIndex(file.Filename, "."):] // 获取文件扩展名

	// 获取 Bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", fmt.Errorf("failed to get bucket: %v", err)
	}

	// 上传文件到 OSS，文件名使用 UUID
	objectKey := fmt.Sprintf("%s%s", fileUUID, fileExt)
	err = bucket.PutObject(objectKey, srcFile)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to OSS: %v", err)
	}

	// 文件上传成功后，记录日志。
	ossURL := fmt.Sprintf("https://%s.oss-%s.aliyuncs.com/%s", bucketName, region, objectKey)
	log.Printf("File uploaded successfully to %s/%s", bucketName, objectKey)
	return ossURL, nil
}

// uploadFile 用于将本地文件上传到OSS存储桶。
// 参数：
//
//	bucketName - 存储空间名称。
//	objectName - Object完整路径，完整路径中不包含Bucket名称。
//	localFileName - 本地文件的完整路径。
//	endpoint - Bucket对应的Endpoint。
//
// 如果成功，记录成功日志；否则，返回错误。
