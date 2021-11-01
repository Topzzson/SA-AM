package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db
}
func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-64-3.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")
	}
	// Migrate the schema
	database.AutoMigrate(
		&Ambulance{},
		&User{},
		&CheckList{},
		&Car_path{},
		&Path_status{},
		&AmbulanceType{}, &Status{}, &Brand{}, &Notify{},
		&AssessmentSheet{},
		&Register{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	// ข้อมูล user
	db.Model(&User{}).Create(&User{
		Name:     "pantita",
		Email:    "pantita@gmail.com",
		Password: string(password),
	})

	var pantita User

	db.Raw("SELECT * FROM users WHERE email = ?", "pantita@gmail.com").Scan(&pantita)

	db.Model(&User{}).Create(&User{
		Name:     "Rattatammanoon",
		Email:    "rattatammanoontop@gmail.com",
		Password: string(password),
	})
	db.Model(&User{}).Create(&User{
		Name:     "Name",
		Email:    "name@example.com",
		Password: string(password),
	})
	var rattatammanoon User
	var name User
	db.Raw("SELECT * FROM users WHERE email = ?", "rattatammanoontop@gmail.com").Scan(&rattatammanoon)
	db.Raw("SELECT * FROM users WHERE email = ?", "name@example.com").Scan(&name)
	db.Model(&User{}).Create(&User{
		Name:     "Chanon",
		Email:    "chanon@gmail.com",
		Password: string(password),
	})

	db.Model(&User{}).Create(&User{
		Name:     "Apple",
		Email:    "Apple@gmail.com",
		Password: string(password),
	})

	var chanon User
	var apple User
	db.Raw("SELECT * FROM users WHERE email = ?", "chanon@gmail.com").Scan(&chanon)
	db.Raw("SELECT * FROM users WHERE email = ?", "Apple@gmail.com").Scan(&apple)
	// --- Type Data
	basicLifeSupport := AmbulanceType{
		TypeName: "รถพยาบาลปฏิบัติการพื้นฐาน",
	}
	db.Model(&AmbulanceType{}).Create(&basicLifeSupport)

	advancedLifeSupport := AmbulanceType{
		TypeName: "รถพยาบาลปฏิบัติการฉุกเฉิน",
	}
	db.Model(&AmbulanceType{}).Create(&advancedLifeSupport)

	//----Brand Data
	brand1 := Brand{
		BrandName: "Honda Stepwgn Spada",
	}
	db.Model(&Brand{}).Create(&brand1)

	brand2 := Brand{
		BrandName: "Toyota Hilux Vigo",
	}
	db.Model(&Brand{}).Create(&brand2)

	brand3 := Brand{
		BrandName: "Nilsson Volvo",
	}
	db.Model(&Brand{}).Create(&brand3)

	//----Status Data
	status1 := Status{
		StatusName: "พร้อมใช้งาน",
	}
	db.Model(&Status{}).Create(&status1)

	status2 := Status{
		StatusName: "ไม่พร้อมใช้งาน",
	}
	db.Model(&Status{}).Create(&status2)

	// ambulance 1
	db.Model(&Ambulance{}).Create(&Ambulance{
		AmbulanceType: basicLifeSupport,
		Registration:  "กข1234",
		Recorder:      name,
		Brand:         brand1,
		Recording:     time.Now(),
		Status:        status1,
	})
	// ambulance 2
	db.Model(&Ambulance{}).Create(&Ambulance{
		AmbulanceType: advancedLifeSupport,
		Registration:  "งม12345",
		Recorder:      pantita,
		Brand:         brand2,
		Recording:     time.Now(),
		Status:        status2,
	})
	// ambulance 3
	db.Model(&Ambulance{}).Create(&Ambulance{
		AmbulanceType: advancedLifeSupport,
		Registration:  "พส8888",
		Recorder:      pantita,
		Brand:         brand3,
		Recording:     time.Now(),
		Status:        status1,
	})

	pathstatus1 := Path_status{
		Status: "Normal",
	}
	db.Model(&Path_status{}).Create(&pathstatus1)

	pathstatus2 := Path_status{
		Status: "Defective/Fix",
	}
	db.Model(&Path_status{}).Create(&pathstatus2)

	pathstatus3 := Path_status{
		Status: "Order",
	}
	db.Model(&Path_status{}).Create(&pathstatus3)

	path1 := Car_path{
		Path_titel: "oil fuel",
	}
	db.Model(&Car_path{}).Create(&path1)
	path2 := Car_path{
		Path_titel: "oil fuel1",
	}
	db.Model(&Car_path{}).Create(&path2)
	path3 := Car_path{
		Path_titel: "oil fuel2",
	}
	db.Model(&Car_path{}).Create(&path3)

	var car1 Ambulance
	db.Raw("SELECT * FROM ambulance WHERE Registration = ?", "พส8888").Scan(&car1)
	db.Model(&CheckList{}).Create(&CheckList{
		Checked_time: time.Now(),
		Ambulance:    car1,
		Car_path:     path1,
		Path_status:  pathstatus3,
	})

	// --- Notify Data
	Notify1 := Notify{
		Address: "bangkok",
		Officer: chanon,
	}
	db.Model(&Notify{}).Create(&Notify1)

	Notify2 := Notify{
		Address: "korat",
		Officer: chanon,
	}
	db.Model(&Notify{}).Create(&Notify2)

	Notify3 := Notify{
		Address: "burirum",
		Officer: apple,
	}
	db.Model(&Notify{}).Create(&Notify3)

	// AssessmentSheet Data
	AssessmentSheet1 := AssessmentSheet{
		Value: "assess001",
	}
	db.Model(&AssessmentSheet{}).Create(&AssessmentSheet1)

	AssessmentSheet2 := AssessmentSheet{
		Value: "assess002",
	}
	db.Model(&AssessmentSheet{}).Create(&AssessmentSheet2)

	AssessmentSheet3 := AssessmentSheet{
		Value: "assess003",
	}
	db.Model(&AssessmentSheet{}).Create(&AssessmentSheet3)

	// Ambulance Data
	Ambulance1 := Ambulance{
		Registration: "ambulance01",
		Officer:      chanon,
	}
	db.Model(&Ambulance{}).Create(&Ambulance1)

	Ambulance2 := Ambulance{
		Registration: "ambulance02",
		Officer:      chanon,
	}
	db.Model(&Ambulance{}).Create(&Ambulance2)

	Ambulance3 := Ambulance{
		Registration: "ambulance03",
		Officer:      apple,
	}
	db.Model(&Ambulance{}).Create(&Ambulance3)

	// watch 1
	db.Model(&Register{}).Create(&Register{
		Ambulance:       Ambulance1,
		Notify:          Notify1,
		RegisterTime:    time.Now(),
		AssessmentSheet: AssessmentSheet1,
	})
	// watch 2
	db.Model(&Register{}).Create(&Register{
		Ambulance:       Ambulance2,
		Notify:          Notify2,
		RegisterTime:    time.Now(),
		AssessmentSheet: AssessmentSheet2,
	})
	// watch 3
	db.Model(&Register{}).Create(&Register{
		Ambulance:       Ambulance3,
		Notify:          Notify3,
		RegisterTime:    time.Now(),
		AssessmentSheet: AssessmentSheet3,
	})

}
