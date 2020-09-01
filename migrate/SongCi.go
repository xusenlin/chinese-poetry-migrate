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




func SongCi(jsonDir string) error {

	jsons, err := ioutil.ReadDir(jsonDir)
	if err != nil {
		return err
	}
	for _, fileInfo := range jsons {
		name := fileInfo.Name()
		if !strings.HasSuffix(name, ".json") {
			continue
		}
		if strings.HasPrefix(name, "author.") {
			err := migrateSongAuthor(jsonDir + "/" + name)
			if err != nil {
				return err
			}
		}
		if strings.HasPrefix(name, "ci.") {
			err := migrateSongCi(jsonDir + "/" + name)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func migrateSongAuthor(path string) error {

	type Author struct {
		Description     string
		Name   string
	}
	fmt.Println("作者数据迁移文件:"+path)
	var authors []Author

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &authors)
	if err != nil {
		return err
	}

	bar := pb.StartNew(len(authors))

	for _, author := range authors {
		db.Conn.Create(&tableStruct.Authors{
			Name:strings.Trim(author.Name," "),
			Dynasty:DynastySong,
			Desc:author.Description,
		})
		bar.Increment()
	}
	bar.Finish()

	return nil
}

func migrateSongCi(path string) error {
	type SongCi struct {
		Author     string
		Rhythmic   string
		Paragraphs []string
	}
	fmt.Println("宋词数据迁移文件:" + path)
	var sc []SongCi

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &sc)
	if err != nil {
		return err
	}

	bar := pb.StartNew(len(sc))
	for _, s := range sc {
		base := tableStruct.BasePoetry{
			Title:  s.Rhythmic,
			Author: s.Author,}

		db.Conn.Create(&tableStruct.SongCi{
			BasePoetry: base,
			Paragraphs: strings.Join(s.Paragraphs, "||"),
		})
		bar.Increment()
	}
	bar.Finish()

	return nil
}
