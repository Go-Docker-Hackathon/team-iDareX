package youtube

import(
	"fmt"
	"io/ioutil"
	"os/exec"
)

func YoutubeDl(url string){
	cmd := exec.Command("youtube-dl", "--all-formats", url)
	_, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}
	
	if err := cmd.Start(); err != nil {
		panic(err.Error())
	}
	
	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}
}