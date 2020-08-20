package migrate

import (
	"chinese-poetry-migrate/db"
	"chinese-poetry-migrate/tableStruct"
	"encoding/json"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io/ioutil"
	"strings"
)

type SSWJ struct {
	Chapter    string     `json:"chapter"`
	Paragraphs []string     `json:"paragraphs"`
}

func SiShuWuJing(jsonDir string) error {

	jsons, err := ioutil.ReadDir(jsonDir)
	if err != nil {
		return err
	}
	for _, fileInfo := range jsons {
		name := fileInfo.Name()
		if !strings.HasSuffix(name,".json"){
			continue
		}
		err := migrateSiShuWuJing(jsonDir + "/" + name)
		if err != nil {
			return err
		}
	}
	return nil
}

func migrateSiShuWuJing(path string) error {
	fmt.Println("四书五经数据迁移文件:" + path)
	var sishuwujing []SSWJ

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &sishuwujing)
	if err != nil {
		return err
	}

	bar := pb.StartNew(len(sishuwujing))
	for _, s := range sishuwujing {

		db.Conn.Create(&tableStruct.SiShuWuJing{
			Chapter: s.Chapter,
			Paragraphs: strings.Join(s.Paragraphs, "||"),
		})
		bar.Increment()
	}
	bar.Finish()

	return nil
}