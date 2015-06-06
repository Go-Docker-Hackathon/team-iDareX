package youtube

import(
	"os/exec"
	"fmt"
	"strings"
)

func YoutubeDl(url string) (fileName string, err error) {
	format := selectVideoFormat(url)
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
	output, err := cmd.Output();
	if( err != nil ){
		fmt.Println("get url formats error:", url)
	}
	fmt.Println("output contnent is:", string(output))
	lines := strings.Split(string(output), "\n")
	lastLine := lines[len(lines)-2]
	
	lastLineFields := strings.Fields(lastLine)
	return lastLineFields[0]
	
}