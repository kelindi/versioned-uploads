package main

import "time"

type User struct {
	Username string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
}


type DataVersion struct {
	DataID string `gorm:"type:varchar(255);not null"`
	Version int `gorm:"type:int;not null"`
	CID      string `gorm:"type:varchar(255);not null"`
	Date    time.Time `gorm:"type:timestamp;not null"`
	UploadedBy User `gorm:"foreignkey:Username;association_foreignkey:Username"`
}

type Dataset struct {
	DataID  string `gorm:"type:varchar(255);not null"`
	CurrentVersionNumber int `gorm:"type:int;not null"`
	CurrentVersion DataVersion `gorm:"foreignkey:SetID;association_foreignkey:SetID"`
	Owners []User `gorm:"many2many:dataset_owners;"`

}
