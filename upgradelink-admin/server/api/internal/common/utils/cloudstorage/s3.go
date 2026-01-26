package cloudstorage

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/zeromicro/go-zero/core/logx"
)

// S3Config 定义 S3 配置结构体
type S3Config struct {
	Bucket   string
	SecretID string
	SecretKey string
	Endpoint string
	Region   string
}

// S3Service S3 上传服务结构体
type S3Service struct {
	config *S3Config
	client *s3.S3
}

// NewS3Service 创建并初始化 S3 服务实例
func NewS3Service(config *S3Config) *S3Service {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region:      aws.String(config.Region),
			Credentials: credentials.NewStaticCredentials(config.SecretID, config.SecretKey, ""),
			Endpoint:    aws.String(config.Endpoint),
			HTTPClient: &http.Client{
				Timeout: 600 * time.Second, // 设置总超时时间为10分钟，可根据需要调整
			},
		},
	))

	return &S3Service{
		config: config,
		client: s3.New(sess),
	}
}

// UploadFile 上传文件到 S3 存储桶
// @param ctx 上下文
// @param file 文件对象
// @param key 存储在 S3 中的键
// @return 文件的访问 URL
// @return 错误信息
func (s *S3Service) UploadFile(ctx context.Context, file multipart.File, key string) (string, error) {
	// multipart.File 接口实现了 io.ReadSeeker，所以可以直接传递给 uploadReadSeeker 方法
	return s.uploadReadSeeker(ctx, file, key)
}

// UploadReader 上传 io.Reader 到 S3 存储桶
// @param ctx 上下文
// @param reader 数据读取器
// @param key 存储在 S3 中的键
// @return 文件的访问 URL
// @return 错误信息
func (s *S3Service) UploadReader(ctx context.Context, reader io.Reader, key string) (string, error) {
	// AWS SDK 要求 Body 实现 io.ReadSeeker，我们需要将 io.Reader 转换为 io.ReadSeeker
	// 检查 reader 是否已经实现了 io.ReadSeeker
	if rs, ok := reader.(io.ReadSeeker); ok {
		// 如果已经是 io.ReadSeeker，直接使用
		return s.uploadReadSeeker(ctx, rs, key)
	}
	
	// 否则，我们需要将 io.Reader 转换为 io.ReadSeeker
	// 这里我们可以将 reader 的内容读取到一个缓冲区，然后使用 bytes.Reader 作为 ReadSeeker
	buf, err := io.ReadAll(reader)
	if err != nil {
		logx.Errorw("failed to read from reader", logx.Field("detail", err), logx.Field("key", key))
		return "", errors.New("failed to read from reader")
	}
	
	return s.uploadReadSeeker(ctx, bytes.NewReader(buf), key)
}

// UploadReadSeeker 上传 io.ReadSeeker 到 S3 存储桶
// @param ctx 上下文
// @param reader 数据读取器（实现了 io.ReadSeeker）
// @param key 存储在 S3 中的键
// @return 文件的访问 URL
// @return 错误信息
func (s *S3Service) UploadReadSeeker(ctx context.Context, reader io.ReadSeeker, key string) (string, error) {
	return s.uploadReadSeeker(ctx, reader, key)
}

// uploadReadSeeker 内部方法：上传 io.ReadSeeker 到 S3 存储桶
func (s *S3Service) uploadReadSeeker(ctx context.Context, reader io.ReadSeeker, key string) (string, error) {
	_, err := s.client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(key),
		Body:   reader,
	})
	if err != nil {
		logx.Errorw("failed to upload object", logx.Field("detail", err), logx.Field("key", key))
		var aerr awserr.Error
		if errors.As(err, &aerr) && aerr.Code() == request.CanceledErrorCode {
			return "", errors.New("upload canceled due to timeout")
		}
		return "", errors.New("failed to upload object")
	}

	// 构建文件访问 URL
	url := s.buildFileURL(key)
	return url, nil
}

// buildFileURL 构建文件访问 URL
func (s *S3Service) buildFileURL(key string) string {
	return "https://" + s.config.Bucket + "." + s.config.Endpoint + key
}
