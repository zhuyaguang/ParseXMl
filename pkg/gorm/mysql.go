package gorm

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"patentExtr/pkg"
)
import "gorm.io/driver/mysql"


func ConnectMysql() *gorm.DB {
	// Configure the database connection (always check errors)
	// mysql 10.101.32.33 用户名：root  密码：123456
	dsn := "root:123456@tcp(10.101.32.33:3306)/itech4u?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil{
		log.Fatal(err)
	}
	return db
}

func Create(p pkg.Patent, db *gorm.DB)  {
	result :=db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "publication_no"}},
		DoNothing: true,
	}).Create(&p)
	fmt.Println("RowsAffected==",result.RowsAffected)
}

func Search(db *gorm.DB)  {
	p :=pkg.Patent{}
	db.First(p, 1)
	fmt.Println(p)
}

