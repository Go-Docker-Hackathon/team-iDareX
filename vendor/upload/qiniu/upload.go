package qiniu

import (
	"github.com/qiniu/api/rs"
	"github.com/Go-Docker-Hackathon/team-iDareX/vendor/db/mongo"
	"labix.org/v2/mgo/bson"
)

import . "github.com/qiniu/api/conf"

func UploadQiniu(videoUrl string) (key string, err error) {
	C := mongo.Connect()
	C.Update(bson.M{"fetchurl": videoUrl}, bson.M{"$set": bson.M{"status": 3}}) // start upload
	
	ACCESS_KEY = "A3opmNKS8XPKzt5ks5C2um__tYL2E6dZu81Xjzim"
	SECRET_KEY = "SCaNo1SxnFnjyxOEJm08OBkCIhA0R7HOHOmIMAZU"
	BUCKET_NAME := "youtubetmp"
	
	uploadToken := getUploadToken(BUCKET_NAME)
	key, err = uploadFile(videoUrl, uploadToken)
	
	C.Update(bson.M{"fetchurl": videoUrl}, bson.M{"$set": bson.M{"status": 4}}) // upload finish
	
	downloadurl := "http://7xjhxh.com1.z0.glb.clouddn.com" + key
	C.Update(bson.M{"fetchurl": videoUrl}, bson.M{"$set": bson.M{"downloadurl": downloadurl}}) // set download key
	return
}

func getUploadToken(bucketName string) string {
	putPolicy := rs.PutPolicy{
		Scope: bucketName,
	}
	return putPolicy.Token(nil)
}
