package utils

import (
	"automic/global"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
)

func OssUpload(file io.Reader, objectName string, objectSize int64, suffix string) (string, error) {
	ctx := context.Background()
	fmt.Printf(global.OssSetting.Endpoint)
	minioClient, err := minio.New(global.OssSetting.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(global.OssSetting.AccessKeyID, global.OssSetting.SecretAccessKey, ""),
	})
	if err != nil {
		global.Logger.Errorf(ctx, "Create oss client  err: %v", err)
	}

	uploadInfo, err := minioClient.PutObject(ctx, global.OssSetting.BucketName, objectName, file, objectSize, minio.PutObjectOptions{ContentType: fmt.Sprintf("%s", suffix)})
	if err != nil {
		global.Logger.Errorf(ctx, "Upload file err: %v", err)
	}

	return uploadInfo.VersionID, err

}

func OssDownload(bucketName string, name string, version_id string) string {
	ctx := context.Background()
	minioClient, err := minio.New(global.OssSetting.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(global.OssSetting.AccessKeyID, global.OssSetting.SecretAccessKey, ""),
	})
	if err != nil {
		global.Logger.Errorf(ctx, "Create oss client  err: %v", err)
	}

	object, err := minioClient.GetObject(ctx, bucketName, name, minio.GetObjectOptions{VersionID: version_id})

	if err != nil {
		global.Logger.Errorf(ctx, "Upload file err: %v", err)
	}

	data, err := io.ReadAll(object)

	return string(data)
}
