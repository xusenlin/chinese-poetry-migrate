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

type Song struct {
	tableStruct.BasePoetry
	Paragraphs []string
}

func YuanSong(jsonDir string) error {

	jsons, err := ioutil.ReadDir(jsonDir)
	if err != nil {
		return err
	}
	for _, fileInfo := range jsons {
		name := fileInfo.Name()
		if !strings.HasSuffix(name,".json"){
			continue
		}
		err := migrateSong(jsonDir + "/" + name)
		if err != nil {
			return err
		}
	}
	return nil
}

func  migrateSong(path string) error {
	fmt.Println("元曲数据迁移文件:" + path)
	var yuanQu []Song

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &yuanQu)
	if err != nil {
		return err
	}

	bar := pb.StartNew(len(yuanQu))
	for _, s := range yuanQu {
		base := tableStruct.BasePoetry{
			Title:  s.Title,
			Author: s.Author,}

		db.Conn.Create(&tableStruct.YuanQu{
			BasePoetry: base,
			Paragraphs: strings.Join(s.Paragraphs, "||"),
		})
		bar.Increment()
	}
	bar.Finish()

	return nil
}