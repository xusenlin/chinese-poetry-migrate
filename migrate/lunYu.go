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

type LY struct {
	Chapter    string     `json:"chapter"`
	Paragraphs []string     `json:"paragraphs"`
}

func LunYu(jsonDir string) error {

	jsons, err := ioutil.ReadDir(jsonDir)
	if err != nil {
		return err
	}
	for _, fileInfo := range jsons {
		name := fileInfo.Name()
		if !strings.HasSuffix(name,".json"){
			continue
		}
		err := migrateLunYu(jsonDir + "/" + name)
		if err != nil {
			return err
		}
	}
	return nil
}

func migrateLunYu(path string) error {
	fmt.Println("论语数据迁移文件:" + path)
	var lunYu []LY

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &lunYu)
	if err != nil {
		return err
	}

	bar := pb.StartNew(len(lunYu))
	for _, s := range lunYu {

		db.Conn.Create(&tableStruct.LunYu{
			Chapter: s.Chapter,
			Paragraphs: strings.Join(s.Paragraphs, "||"),
		})
		bar.Increment()
	}
	bar.Finish()

	return nil
}