package qiniu

import (
	"github.com/qiniu/api/rs"
)

import . "github.com/qiniu/api/conf"

const DOMAIN = "http://7xjj3q.com1.z0.glb.clouddn.com/"

func UploadQiniu(videoUrl string) (key string, err error) {

	ACCESS_KEY = "WoXW_P0kS1TVQ_eXVcs-gRfgxoYUGHFs-wQ9OPKw"
	SECRET_KEY = "X5pHXl9qB079ZX8ppeHTslUkZfqTLm3qhUdac63z"
	BUCKET_NAME := "youtubetmp"

	uploadToken := getUploadToken(BUCKET_NAME)
	key, err = uploadFile(videoUrl, uploadToken)

	return
}

func getUploadToken(bucketName string) string {
	putPolicy := rs.PutPolicy{
		Scope: bucketName,
	}
	return putPolicy.Token(nil)
}
