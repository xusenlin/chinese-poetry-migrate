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

type YMY struct {
	Content string   `json:"content"`
	Comment []string  `json:"comment"`
}

func YouMengYing(jsonDir string) error {

	jsons, err := ioutil.ReadDir(jsonDir)
	if err != nil {
		return err
	}
	for _, fileInfo := range jsons {
		name := fileInfo.Name()
		if !strings.HasSuffix(name, ".json") {
			continue
		}
		err := migrateYouMengYing(jsonDir + "/" + name)
		if err != nil {
			return err
		}
	}
	return nil
}

func migrateYouMengYing(path string) error {
	fmt.Println("幽梦影数据迁移文件:" + path)
	var ymy []YMY

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &ymy)
	if err != nil {
		return err
	}

	bar := pb.StartNew(len(ymy))
	for _, y := range ymy {

		db.Conn.Create(&tableStruct.YouMengYing{
			Content:    y.Content,
			Comment: strings.Join(y.Comment, "||"),
		})
		bar.Increment()
	}
	bar.Finish()

	return nil
}
