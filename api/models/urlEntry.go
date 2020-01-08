package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type UrlEntry struct {
	gorm.Model
	UUID string
	OriginalUrl string
	CreatedAt time.Time
}

