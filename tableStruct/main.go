package tableStruct

import (
	"chinese-poetry-migrate/db"
)

func AutoMigrate()  {
	db.Conn.AutoMigrate(
		&Author{},
		&TangShi{},
		&SongCi{},
		&YuanQu{},
	)
}