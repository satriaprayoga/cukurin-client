package models

import "time"

type Advert struct {
	AdvertiseID int `json:"advertise_id" gorm:"PRIMARY_KEY"`
	AddAdvert
}

type AddAdvert struct {
	Title         string    `json:"title" gorm:"type:varchar(100)"`
	Desc          string    `json:"desc" gorm:"type:varchar(255)"`
	Status        string    `json:"status" gorm:"type:varchar(2)"`
	SlideDuration int       `json:"slide_duration" gorm:"type:integer"`
	StartDate     time.Time `json:"start_date" gorm:"type:timestamp(0) without time zone"`
	EndDate       time.Time `json:"end_date" gorm:"type:timestamp(0) without time zone"`
	FileID        int       `json:"file_id" gorm:"primary_key;auto_increment:true"`
}

type ListAdvert struct {
	AdvertiseID int `json:"advertise_id"`
	AddAdvert
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	FileType string `json:"file_type"`
}
