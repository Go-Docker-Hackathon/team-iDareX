package qiniu

import (
	"github.com/qiniu/api/rs"
)

import . "github.com/qiniu/api/conf"

func UploadQiniu(videoUrl string) (key string, err error) {
	
	ACCESS_KEY = "A3opmNKS8XPKzt5ks5C2um__tYL2E6dZu81Xjzim"
	SECRET_KEY = "SCaNo1SxnFnjyxOEJm08OBkCIhA0R7HOHOmIMAZU"
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
