package cmd

import (
	"fmt"
	"mailvalidator/model"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)


var db *gorm.DB
var err error

func InitDb(){

    dsn := viper.GetString("DB_WRITER_USER")+":@tcp("+viper.GetString("DB_READER_HOST")+":"+strconv.Itoa(viper.GetInt("DB_READER_PORT"))+")/"+ viper.GetString("DB_TABLE_NAME") +"?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println(err)
	}

	var email model.Email
	db.First(&email, 1)

	print(email.Address)

}

func exist(address string) bool {

	var email model.Email
	return db.Where("address= ?", address).First(&email).Error == nil

}

func insert(recipient model.Recipient)  {
	email := model.Email{Id: 99, Address: recipient.Email, Time_create: time.Now().String()}
	result := db.Create(&email)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println(result.RowsAffected)
}