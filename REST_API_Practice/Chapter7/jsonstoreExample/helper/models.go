package models

import (
	"github.com/jinzhu/gorm"
)

type Shipment struct {
	gorm.Model
	Packages []Package
	Data     string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Package struct {
	gorm.Model
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

func (Shipment) TableName() string {
	return "Shipment"
}

func (Package) TableName() string {
	return "Package"
}

func InitDB() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open("postgres", "postgres://root:secret@localhost/root?sslmode=disable")
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Shipment{}, &Package{})
	return db, nil
}
