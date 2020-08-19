package tableStruct

import (
	"time"
)

type BasePoetry struct {
	ID        uint       `gorm:"primary_key"json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"json:"deleted_at"`
	Title     string     `json:"title"`
	Star      int        `json:"star"`
	Author    string     `json:"author"`
}

type TangShi struct { //唐诗
	BasePoetry
	Paragraphs string `json:"paragraphs"`
}

type SongCi struct { //宋词
	BasePoetry
	Paragraphs string `json:"paragraphs"`
}

type YuanQu struct { //元曲
	BasePoetry
	Paragraphs string `json:"paragraphs"`
}