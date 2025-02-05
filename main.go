package main

import (
	"MP3Recover/conv"
	"MP3Recover/mysql"
	"MP3Recover/util"
	"os"
)

func main() {
	root := os.Getenv("root")
	if root == "" {
		root = "/home/zen/ugreen/alist/music/audio"
	}
	util.SetLog(root)
	mysql.SetMysql()
	files := util.GetFiles(root)
	for i := range files {
		conv.Mp3(files[i])
	}
}
