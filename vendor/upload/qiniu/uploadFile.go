package qiniu

import (
	"encoding/json"
	"fmt"
	"github.com/qiniu/api/io"
	"log"
	"strings"
	"time"
)

func uploadFile(videoUrl, uploadToken string) string {
	var ret io.PutRet
	var returnStr ReturnJson
	extra := &io.PutExtra{}
	fileName := getFileName(videoUrl)
	key := time.Now().Format("20060102150405") + fileName
	fmt.Println(&ret, uploadToken, key, videoUrl, extra)
	err := io.PutFile(nil, &ret, uploadToken, key, videoUrl, extra)
	if err != nil {
		log.Print("io.PutFile failed: ", err)
		returnStr.Success = "0"
		returnStr.Data = "io.PutFile failed!"
		return getJson(returnStr)
	}
	returnStr.Success = "1"
	returnStr.Data = ReturnSuccess{ret.Hash, ret.Key}
	return getJson(returnStr)
}

func getFileName(videoUrl string) string {
	callback := func(c rune) bool { return c == '/' }
	UrlArr := strings.FieldsFunc(videoUrl, callback)
	return UrlArr[len(UrlArr)-1]
}

func getJson(returnJson ReturnJson) string {
	jsonStr, err := json.Marshal(returnJson)
	if err != nil {

	}
	return string(jsonStr)
}
