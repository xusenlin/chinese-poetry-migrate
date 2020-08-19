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
type Poet struct {
	tableStruct.BasePoetry
	Paragraphs []string
}


func Poetry(jsonDir string) error {

	jsons, err := ioutil.ReadDir(jsonDir)
	if err != nil {
		return err
	}
	for _, fileInfo := range jsons {
		name := fileInfo.Name()
		if !strings.HasSuffix(name,".json"){
			continue
		}
		if strings.Contains(name, "authors") {
			err := migrateAuthor(jsonDir + "/" + name)
			if err != nil {
				return err
			}
		}
		if strings.Contains(name, "poet") {
			err := migratePoetry(jsonDir + "/" + name)
			if err != nil {
				return err
			}
		}
	}
	return nil
}





func migrateAuthor(path string) error {

	fmt.Println("作者数据迁移文件:"+path)
	var authors []tableStruct.Author

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

		if strings.Contains(path, "tang") {
			author.Dynasty = DynastyTang
		}
		if strings.Contains(path, "song") {
			author.Dynasty = DynastySong
		}
		author.Name = strings.Trim(author.Name," ")
		db.Conn.Create(&author)
		bar.Increment()
	}
	bar.Finish()

	return nil
}


func  migratePoetry(path string) error {

	fmt.Println("诗词数据迁移文件:"+path)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var poetry []Poet
	err = json.Unmarshal(b, &poetry)
	if err != nil {
		return err
	}
	bar := pb.StartNew(len(poetry))
	for _, p := range poetry{
		base := tableStruct.BasePoetry{
			Title:  p.Title,
			Author: p.Author,}


		if strings.Contains(path, "tang") {
			db.Conn.Create(&tableStruct.TangShi{
				BasePoetry: base,
				Paragraphs: strings.Join(p.Paragraphs, "||"),
			})
		}
		if strings.Contains(path, "song") {
			db.Conn.Create(&tableStruct.SongCi{
				BasePoetry: base,
				Paragraphs: strings.Join(p.Paragraphs, "||"),
			})
		}
		bar.Increment()
	}
	bar.Finish()

	return nil
}

