package common

import (
	"context"
	"errors"
	"fmt"
	"sagooiot/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func init() {
	path := g.Cfg().MustGet(context.Background(), "system.upload.path").String()
	service.RegisterFileSystem(&LocalSysFile{
		UploadPath: path,
	})
}

type LocalSysFile struct {
	UploadPath string
}

func (l *LocalSysFile) Save(ctx context.Context, file *ghttp.UploadFile, path string) error {
	_, err := file.Save(l.UploadPath+path, false)
	return err
}

func (l *LocalSysFile) Remove(ctx context.Context, path string) error {
	return gfile.Remove(l.UploadPath + path)
}

func (l *LocalSysFile) Download(ctx context.Context, response *ghttp.Response, path string) error {
	response.ServeFileDownload(l.UploadPath + path)
	return nil
}

type MinioSystemFile struct {
	bucketName string
	client     *minio.Client
}

func NewMinioSystemFile() *MinioSystemFile {
	bucketName := "sagooiot"
	endpoint := "localhost:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	return &MinioSystemFile{
		bucketName: bucketName,
		client:     minioClient,
	}
}

func (m *MinioSystemFile) Save(ctx context.Context, file *ghttp.UploadFile, path string) error {
	fileRader, err := file.Open()
	if err != nil {
		return err
	}
	_, err = m.client.PutObject(ctx, m.bucketName, path+"/"+file.Filename, fileRader, file.Size, minio.PutObjectOptions{})
	return err
}

func (m *MinioSystemFile) Remove(ctx context.Context, path string) error {
	return m.client.RemoveObject(ctx, m.bucketName, path, minio.RemoveObjectOptions{})
}

func (m *MinioSystemFile) Download(ctx context.Context, response *ghttp.Response, path string) error {
	_, err := m.client.StatObject(ctx, m.bucketName, path, minio.StatObjectOptions{
		Checksum: true,
	})
	if err != nil {
		return errors.New("文件不存在")
	}
	object, err := m.client.GetObject(ctx, m.bucketName, path, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	info, err := object.Stat()
	if err != nil {
		return err
	}
	g.Log().Debugf(ctx, "info: %+v", info)
	response.Header().Set("Content-Type", "application/force-download")
	response.Header().Set("Accept-Ranges", "bytes")
	response.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename=%s`, info.Key))
	response.ServeContent(info.Key, info.LastModified, object)
	return nil
}
