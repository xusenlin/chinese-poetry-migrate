package tableStruct

import (
	"chinese-poetry-migrate/db"
)

func AutoMigrate()  {
	db.Conn.AutoMigrate(
		&Authors{},
		&TangShi{},
		&SongCi{},
		&YuanQu{},
		&SiShuWuJing{},
		&YouMengYing{},
		&SongShi{},
	)
}