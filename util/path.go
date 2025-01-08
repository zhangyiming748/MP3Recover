package util

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
)

func GetFiles(root string) []string {
	log.Println("开始查找文件")
	var filePaths []string
	//root := "/path/to/your/directory" // 请将此路径替换为你需要遍历的目录
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			if filepath.Ext(absPath) != ".mp3" {
				file, _ := os.Open(absPath)

				// We only have to pass the file header = first 261 bytes
				head := make([]byte, 261)
				file.Read(head)

				if filetype.IsAudio(head) {
					fmt.Printf("%s is an audio\n", absPath)
					filePaths = append(filePaths, absPath)
				} else {
					fmt.Printf("%s Not an audio\n", absPath)
				}
			}

		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	for _, filePath := range filePaths {
		fmt.Println(filePath)
	}
	return filePaths
}
