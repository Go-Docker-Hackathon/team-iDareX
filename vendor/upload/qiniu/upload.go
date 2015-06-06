package qiniu

import (
	"fmt"
	"github.com/qiniu/api/rs"
)

func UploadQiniu(videoUrl, bucketName string) string {
	uploadToken := getUploadToken(bucketName)
	fmt.Println(uploadToken)
	return uploadFile(videoUrl, uploadToken)
}

func getUploadToken(bucketName string) string {
	putPolicy := rs.PutPolicy{
		Scope: bucketName,
	}
	return putPolicy.Token(nil)
}
