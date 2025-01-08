package mysql

import (
	"time"
)

type File struct {
	Id        int64     `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	Input     string    `xorm:"comment('输入文件路径') VARCHAR(255)"`
	Output    string    `xorm:"comment('输出文件路径') VARCHAR(255)"`
	Success   bool      `xorm:"comment('是否成功') Bool"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func init() {
	GetMysql().Sync2(File{})
}

/*
下载成功后插入
*/

func (f *File) InsertOne() (int64, error) {
	return GetMysql().InsertOne(f)
}
