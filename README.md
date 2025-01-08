# MP3Recover
把所有音频文件恢复成原始MP3格式
# 核心代码
```bash
 ffmpeg -i "$flac_file" -c:a libmp3lame -filter:a "volume=1.35" -map_chapters -1 "$mp3_file"
```
