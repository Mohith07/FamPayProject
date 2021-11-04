package model

import (
	"time"

	"github.com/labstack/gommon/log"

	"FamPayProject/clients"
)

type BaseModel struct {
	ID        uint64    `gorm:"primary_key;AUTO_INCREMENT" json:"-"`
	CreatedAt time.Time `sql:"DEFAULT:now();not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `sql:"DEFAULT:now();not null" json:"-"`
}

type Video struct {
	BaseModel
	Name        string `gorm:"name"`
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
	URL         string `gorm:"url"`
	PublishedAt string `gorm:"published_at"`
}

func (Video) TableName() string {
	return "video"
}

type VideoSearchConfig struct {
	BaseModel
	Query      string `gorm:"query"`
	MaxResults int64  `gorm:"max_results"`
	OrderBy    string `gorm:"order_by"`
	SearchType       string `gorm:"search_type"`
	IsActive   bool   `gorm:"is_active"`
}

type APIConfig struct {
	BaseModel
	APIKey   string `gorm:"api_key"`
	IsActive bool   `gorm:"is_active"`
}

type YoutubeCronJobRun struct {
	BaseModel
	PublishedAfter string `gorm:"published_after"`
	APIKey         string `gorm:"api_key"`
}

func Init() {
	db := clients.GetDB()
	if err := db.AutoMigrate(&Video{}, &YoutubeCronJobRun{}, &APIConfig{}, &VideoSearchConfig{}); err != nil {
		log.Errorf("error automigrating - " + err.Error())
	}
}
