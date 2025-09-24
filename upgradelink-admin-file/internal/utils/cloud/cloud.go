package cloud

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/zeromicro/go-zero/core/logx"

	"upgradelink-admin-file/ent"
	"upgradelink-admin-file/ent/storageprovider"
)

type CloudServiceGroup struct {
	CloudStorage map[string]*s3.S3

	ProviderData map[string]struct {
		Id       uint64
		Folder   string
		Bucket   string
		Endpoint string
		UseCdn   bool
		CdnUrl   string
	}

	DefaultProvider string
}

// NewCloudServiceGroup returns the S3 service client group
func NewCloudServiceGroup(db *ent.Client) *CloudServiceGroup {
	cloudServices := &CloudServiceGroup{}
	cloudServices.CloudStorage = make(map[string]*s3.S3)
	cloudServices.ProviderData = make(map[string]struct {
		Id       uint64
		Folder   string
		Bucket   string
		Endpoint string
		UseCdn   bool
		CdnUrl   string
	})

	data, err := db.StorageProvider.Query().Where(storageprovider.StateEQ(true)).All(context.Background())
	if err != nil {
		logx.Errorw("failed to load provider config from database, make sure database has been initialize and has config data",
			logx.Field("detail", err))
	}

	for _, v := range data {
		sess := session.Must(session.NewSession(
			&aws.Config{
				Region:      aws.String(v.Region),
				Credentials: credentials.NewStaticCredentials(v.SecretID, v.SecretKey, ""),
				Endpoint:    aws.String(v.Endpoint),
				HTTPClient: &http.Client{
					Timeout: 600 * time.Second, // 设置总超时时间为30秒，可根据需要调整
				},
			},
		))
		svc := s3.New(sess)

		cloudServices.CloudStorage[v.Name] = svc
		cloudServices.ProviderData[v.Name] = struct {
			Id       uint64
			Folder   string
			Bucket   string
			Endpoint string
			UseCdn   bool
			CdnUrl   string
		}{Id: v.ID, Folder: v.Folder, Bucket: v.Bucket, Endpoint: v.Endpoint, UseCdn: v.UseCdn, CdnUrl: v.CdnURL}

		if v.IsDefault {
			cloudServices.DefaultProvider = v.Name
		}
	}

	return cloudServices
}
