package service

import (
	"fmt"
	"wilikidi/gin/model"
)

func InsertAStudent(s model.Student) string {

	fmt.Println(s.Name)

	return ""
}
