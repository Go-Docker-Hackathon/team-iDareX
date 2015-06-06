package youtube

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type VideoQuality struct {
	FormatId   string
	Extension  string
	Resolution string
	Note       string
}

func YoutubeDl(workRequest WorkRequest) (fileName string, err error) {
//	format := selectVideoFormat(url)
	format := workRequest.FormatId
	url := workRequest.Url
	fmt.Println("format:", format)

	fmt.Println("youtube-dl", "-o", "./steamerDataDir/%(title)s-%(format)-%(id)s.%(ext)s", "-f", format, url)
	cmd := exec.Command("youtube-dl", "-o", "./steamerDataDir/%(title)s-%(format)s-%(id)s.%(ext)s", "-f", format, url)
	output, err := cmd.Output()

	if err != nil {
		return "", err
	}
	fmt.Println("download content: ", string(output))

	cmd = exec.Command("youtube-dl", "--get-filename", "-o", "./steamerDataDir/%(title)s-%(format)s-%(id)s.%(ext)s", "-f", format, url)
	output, err = cmd.Output()
	if err != nil {
		return "", err
	}

	fileName = strings.TrimSpace(string(output))
	fmt.Println("download fileName:", fileName)

	return fileName, nil
}

func selectVideoFormat(url string) string {
	fmt.Println("exec command youtube-dl -F", url)
	cmd := exec.Command("youtube-dl", "-F", url)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("get url formats error:", url)
	}
	fmt.Println("output contnent is:", string(output))
	lines := strings.Split(string(output), "\n")
	lastLine := lines[len(lines)-2]

	lastLineFields := strings.Fields(lastLine)
	return lastLineFields[0]

}

func GetVideoQuality(url string) string {
	videoQualities := make([]VideoQuality, 0)
	//获取视频信息
	cmd := exec.Command("youtube-dl", "-F", url)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("获取视频信息失败")
	}
	//解析字符串信息
	outputArr := strings.Split(string(output), "\n")
	for i := 0; i < len(outputArr); i++ {

		conArr := strings.Fields(outputArr[i])

		if len(conArr) > 4 && conArr[3] == "DASH" && conArr[4] == "video" {

			conStr3 := ""
			for i := 3; i < len(conArr[3:])+3; i++ {
				conStr3 += conArr[i]
			}

			videoQualities = append(videoQualities, VideoQuality{
				FormatId:   conArr[0],
				Extension:  conArr[1],
				Resolution: conArr[2],
				Note:       conStr3,
			})
		}
	}
	str, _ := json.Marshal(videoQualities)
	return string(str)
}
