package api

import (
	"github.com/fertech02/Wasa-repository/service/database"
)

func (p *Photo) PhotoFromDatabase(photo database.Photo) {
	p.Pid = photo.Pid
	p.Uid = photo.Uid
	p.File = photo.File
	p.Date = photo.Date
}

type Photo struct {
	Pid  string `json:"pid"`
	Uid  string `json:"uid"`
	File []byte `json:"file"`
	Date string ` json:"date"`
}

func (p *Photo) PhotoToDatabase() database.Photo {
	return database.Photo{
		Pid:  p.Pid,
		Uid:  p.Uid,
		File: p.File,
		Date: p.Date,
	}
}
