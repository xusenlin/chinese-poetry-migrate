package main

import (
	"chinese-poetry-migrate/db"
	"chinese-poetry-migrate/migrate"
	"chinese-poetry-migrate/tableStruct"
)



func main() {
	db.InitConn()
	tableStruct.AutoMigrate()

	//err := migrate.Poetry("./chinese-poetry/json")
	//if err != nil {
	//	panic(err)
	//}
	//err = migrate.YuanQu("./chinese-poetry/yuanqu")
	//if err != nil {
	//	panic(err)
	//}
	//err = migrate.LunYu("./chinese-poetry/lunyu")
	//if err != nil {
	//	panic(err)
	//}
	//err = migrate.SiShuWuJing("./chinese-poetry/sishuwujing")
	//if err != nil {
	//	panic(err)
	//}
	err := migrate.YouMengYing("./chinese-poetry/youmengying")
	if err != nil {
		panic(err)
	}

}