package main

import (
	"time"
	"crawler/models"
	"strconv"
)

func main() {
	courses := make([]models.Course, 1000*1000)
	for i := 0; i < 1000*1000; i++ {
		index := strconv.Itoa(i)
		courses[i] = models.Course{
			CourseId:       index,
			Title:          "title" + index,
			FirstCategory:  "first" + index,
			SecondCategory: "second" + index,
			PlayCount:      i,
			Author:         "author" + index,
		}
	}

	start := time.Now().Unix()
	models.SaveCourses(courses)
	end := time.Now().Unix()

	println(end - start)
}