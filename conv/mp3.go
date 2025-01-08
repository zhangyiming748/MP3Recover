package conv

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ffmpeg -i "$flac_file" -c:a libmp3lame -filter:a "volume=1.35" -map_chapters -1 "$mp3_file"
func Mp3(input string) {
	output := strings.Replace(input, filepath.Ext(input), ".mp3", 1)
	cmd := exec.Command("ffmpeg", "-i", input, "-c:a", "libmp3lame", "-ac", "1", "-map_chapters", "-1", output)
	log.Printf("将要执行的命令:%s\n", cmd.String())
	console, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("命令执行错误,命令原文:%v\t错误原文:%v\n", cmd.String(), err)
	} else {
		log.Printf("命令:%s执行成功,控制台输出:%s\n", cmd.String(), string(console))
		if err := os.Remove(input); err != nil {
			log.Printf("删除原始文件:%v失败:%v\n", input, err)
		} else {
			log.Printf("删除原始文件:%v\n", input)
		}
	}
}
