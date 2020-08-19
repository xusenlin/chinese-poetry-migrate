package main

import (
	"chinese-poetry-migrate/db"
	"chinese-poetry-migrate/migrate"
	"chinese-poetry-migrate/tableStruct"
)



func main() {
	db.InitConn()
	tableStruct.AutoMigrate()

	err := migrate.Poetry("./chinese-poetry/json")
	if err != nil {
		panic(err)
	}
	err = migrate.YuanSong("./chinese-poetry/yuanqu")
	if err != nil {
		panic(err)
	}
}