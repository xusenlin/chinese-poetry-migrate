package tableStruct

import (
	"time"
)

type Author struct {
	ID        uint       `gorm:"primary_key"json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"json:"deleted_at"`
	Name      string     `json:"name"`
	Star      int        `json:"star"`
	Desc      string     `json:"desc"`
	Dynasty   string     `json:"dynasty"`
}
