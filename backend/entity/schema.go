package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string

	Ambulances []Ambulance `gorm:"foreignKey:RecorderID"`
	// 1 user ใช้ได้หลาย ambulance
	Ambulances1 []Ambulance `gorm:"foreignKey:OfficerID"`
	// 1 user มีได้หลาย notify
	Notifys []Notify `gorm:"foreignKey:OfficerID"`
}

type Ambulance struct {
	gorm.Model
	Registration    string
	Recording       time.Time
	RecorderID      *uint
	Recorder        User `gorm:"references:id"`
	AmbulanceTypeID *uint
	AmbulanceType   AmbulanceType `gorm:"references:id"`
	StatusID        *uint
	Status          Status `gorm:"references:id"`
	BrandID         *uint
	Brand           Brand `gorm:"references:id"`

	// OfficerID ทำหน้าที่เป็น FK
	OfficerID *uint
	// เป็นข้อมูล user เมื่อ join ตาราง
	Officer   User
	Register  []Register  `gorm:"foreignKey:AmbulanceID"`
	CheckList []CheckList `gorm:"foreignKey:AmbulanceID"`
}
type AmbulanceType struct {
	gorm.Model
	TypeName   string
	Ambulances []Ambulance `gorm:"foreignKey:AmbulanceTypeID"`
}
type Status struct {
	gorm.Model
	StatusName string
	Ambulances []Ambulance `gorm:"foreignKey:StatusID"`
}
type Brand struct {
	gorm.Model
	BrandName  string
	Ambulances []Ambulance `gorm:"foreignKey:BrandID"`
}

type Car_path struct {
	gorm.Model

	Path_titel string `gorm:"uniqueIndex"`

	CheckList []CheckList `gorm:"foreignKey:Car_pathID"`
}
type Path_status struct {
	gorm.Model
	Status    string      `gorm:"uniqueIndex"`
	CheckList []CheckList `gorm:"foreignKey:Path_statusID"`
}

type CheckList struct {
	gorm.Model
	Checked_time time.Time

	Car_pathID *uint
	Car_path   Car_path

	Path_statusID *uint
	Path_status   Path_status

	AmbulanceID *uint
	Ambulance   Ambulance
}
type Notify struct {
	gorm.Model
	Address    string `gorm:"uniqueIndex"`
	NotifyTime time.Time
	// OfficerID ทำหน้าที่เป็น FK
	OfficerID *uint
	// เป็นข้อมูล user เมื่อ join ตาราง
	Officer  User
	Register []Register `gorm:"foreignKey:NotifyID"`
}

type AssessmentSheet struct {
	gorm.Model
	Value      string `gorm:"uniqueIndex"`
	Assesstime time.Time

	Register []Register `gorm:"foreignKey:AssessmentSheetID"`
}

type Register struct {
	gorm.Model
	RegisterTime time.Time

	// AssessmentSheet ทำหน้าที่เป็น FK
	AssessmentSheetID *uint
	AssessmentSheet   AssessmentSheet

	// Notify ทำหน้าที่เป็น FK
	NotifyID *uint
	Notify   Notify

	// Ambulance ทำหน้าที่เป็น FK
	AmbulanceID *uint
	Ambulance   Ambulance
}
