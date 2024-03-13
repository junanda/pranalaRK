package jawaban

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataStore struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   string `json:"age"`
}

func Conenct() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", "root", "password", "127.0.0.1", "3306", "test_pranala")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect Database...", err.Error())
		return nil
	}

	log.Println("Connect Database Success....")
	return db
}

func CreateData(data interface{}, db *gorm.DB) {
	err := db.Create(data).Error
	if err != nil {
		log.Println("Failed Store Data:", err.Error())
	}

	log.Println("Success Store Data")
}

func ReadData(email string, db *gorm.DB) {
	var dataRead DataStore

	err := db.Where("email == ?", email).First(&dataRead).Error
	if err != nil {
		log.Println("Failed Read data: ", err.Error())
	}

	fmt.Println(dataRead)
}
