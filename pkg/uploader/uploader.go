package uploader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	_ "embed"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

//go:embed gcp-credentials.json
var credentials []byte

const (
	CACHE_CONTROL = "Cache-Control:private, max-age=0, no-transform" // disables bucket caching
	CONTENT_TYPE  = "image/svg+xml"
)

type GoogleBucketUploader struct {
	bucket     string
	objectPath string
}

func NewGoogleBucketUploader(bucket, objectPath string) *GoogleBucketUploader {
	return &GoogleBucketUploader{
		bucket:     bucket,
		objectPath: objectPath,
	}
}

// UploadSVG uploads an SVG by using google storage package.
func (u *GoogleBucketUploader) UploadSVG(pair, svg string) error {
	err := fileUpload(u.bucket, fmt.Sprintf("%v/%v/chart.svg", u.objectPath, pair), []byte(svg))
	if err != nil {
		return err
	}

	// TODO: add flag to disable uploading
	// os.WriteFile(fmt.Sprintf("%v.html", pair), []byte(svg), 0644)

	return nil
}

func fileUpload(bucket, object string, data []byte) error {
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(credentials))
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	bucketWriter := client.Bucket(bucket).Object(object).NewWriter(ctx)
	defer bucketWriter.Close()
	bucketWriter.ChunkSize = 0 // note retries are not supported for chunk size 0.
	bucketWriter.CacheControl = CACHE_CONTROL
	bucketWriter.ContentType = CONTENT_TYPE

	buf := bytes.NewBuffer(data)
	if _, err = io.Copy(bucketWriter, buf); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}

	return nil
}
