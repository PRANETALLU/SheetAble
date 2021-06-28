package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Composer struct {
	Name      string    `gorm:"primary_key" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Composer) Prepare() {
	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Composer) SaveComposer(db *gorm.DB) (*Composer, error) {
	err := db.Debug().Model(&Sheet{}).Create(&c).Error
	if err != nil {
		return &Composer{}, err
	}
	return c, nil
}

func (c *Composer) GetAllComposer(db *gorm.DB) (*[]Composer, error) {
	/*
		This method will return max 20 composer, to find more or specific one you need to specify it.
		Currently it sorts it by the newest updates
	*/
	var err error
	composers := []Composer{}

	err = db.Debug().Order("updated_at desc").Limit(20).Find(&composers).Error

	if err != nil {
		return &[]Composer{}, err
	}
	return &composers, err
}