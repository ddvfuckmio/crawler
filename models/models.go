package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

const (
	MYSQL = "mysql"
	URL   = "root:root@(127.0.0.1:3306)/crawler?charset=utf8&parseTime=True&loc=Local"
	SALT  = "ddv"
)

type User struct {
	Age  int
	Name string
}

type Course struct {
	CourseId       string    `gorm:"column:courseId"`
	Title          string    `gorm:"column:title"`
	Author         string    `gorm:"column:author"`
	FirstCategory  string    `gorm:"column:firstCategory"`
	SecondCategory string    `gorm:"column:secondCategory"`
	Introduction   string    `gorm:"column:introduction"`
	PlayCount      int       `gorm:"column:playCount"`
	CreatedAt      time.Time `gorm:"column:createdAt"`
}

type Category struct {
	FirstCategory  string
	SecondCategory string
	Url            string
}

func (Course) TableName() string {
	return "courses"
}

func SaveCourses(courses []Course) {
	//log.Print("开始插入...")
	db, _ = gorm.Open(MYSQL, URL)
	for _, course := range courses {
		//log.Println(i)
		course.CreatedAt = time.Now()
		db.Create(course)
	}
	db.Close()
}
