package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// gorm.io/driver/mysql
// gorm.io/gorm
// user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local

type Account struct {
	ID       int64     `gorm:"column:id;primary_key"`
	UserID   string    `gorm:"column:user_id"`
	Password string    `gorm:"column:password"`
	Nickname string    `gorm:"column:nickname"`
	Ct       time.Time `gorm:"column:created_at"`
	Ut       time.Time `gorm:"column:updated_at"`
}

func (a Account) TableName() string {
	table := "account"
	return table
}

func main() {
	db := connDB()
	var account Account
	//if err := db.Find(&accounts).Error; err != nil {
	//	fmt.Println(err)
	//	return
	//}
	if err := db.Where("id = ?",
		2).First(&account).Error; err != nil {
		fmt.Println(err)
		return
	}
	//db.Find()
	//db.Delete()
	//db.Update()

	//db.Where().Offset().Limit().Find()
	fmt.Println(account)
}

func connDB() *gorm.DB {
	mysqlDB, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/cms_account?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	mysqlDB = mysqlDB.Debug()
	return mysqlDB
}
