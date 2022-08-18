package database

import (
	"strings"

	"gorm.io/gorm"
)

type Waf struct {
	gorm.Model
	ID                   string `gorm:"primary_key"`
	Tag                  string `gorm:"not null, unique"`
	AllowedControlUsers  string
	AllowedDownloadUsers string
	Data                 []byte `gorm:"type:blob" json:"-"`
}

func (w *Waf) AllowedControlUsersList() []string {
	return strings.Split(w.AllowedControlUsers, ",")
}

func (w *Waf) AllowedDownloadUsersList() []string {
	return strings.Split(w.AllowedDownloadUsers, ",")
}
