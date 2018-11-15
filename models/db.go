package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"github.com/pkg/errors"
)

var db *gorm.DB

const (
	MYSQL = "mysql"
	URL   = "root:root@(127.0.0.1:3306)/data?charset=utf8&parseTime=True&loc=Local"
	SALT  = "ddv"
)

type User struct {
	Age  int
	Name string
}

type Course struct {
	CourseId       string `gorm:"column:courseId"`
	Title          string `gorm:"column:title"`
	FirstCategory  string `gorm:"column:firstCategory"`
	SecondCategory string `gorm:"column:secondCategory"`
	Introduction   string `gorm:"column:introduction"`
	PlayCount      int    `gorm:"column:playCount"`
}

func (Course) TableName() string {
	return "courses"
}

func SaveCourses(courses []Course) {
	log.Print("开始插入...")
	db, _ = gorm.Open(MYSQL, URL)
	for i, course := range courses {
		log.Println(i)
		db.Create(course)
	}
	defer db.Close()
}

func main() {
	users:=make([]*User,10)
	users[0] = &User{
		Age:10,
		Name:"ddv",
	}
	log.Println(users[0])
	checkUsers(users)
	log.Println(users[0])
}
func checkUsers(users []*User) {
	users = nil
}
func check(i int) {
	panic(errors.New("1"))
	defer log.Println("check done...")
}

func modify(users []User) {
	for index, v := range users {
		if (v.Age == 6) {
			//users = append(users[:index], users[index+1:]...)
			users[index].Age = 100
			break
		}
	}
}
