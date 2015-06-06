package qiniu

import (
	"github.com/qiniu/api/io"
	"strings"
	"time"
	
	"fmt"
)

func uploadFile(videoUrl, uploadToken string) (key string, err error){
	var ret io.PutRet
	extra := &io.PutExtra{}
	fileName := getFileName(videoUrl)
	key = time.Now().Format("20060102150405") + "-" + fileName
	err = io.PutFile(nil, &ret, uploadToken, key, videoUrl, extra)
	
	return
}

func getFileName(videoUrl string) string {
	fmt.Println(videoUrl)
	callback := func(c rune) bool { return c == '/' }
	UrlArr := strings.FieldsFunc(videoUrl, callback)
	return UrlArr[len(UrlArr)-1]
}