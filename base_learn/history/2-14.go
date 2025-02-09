package main

import (
	"base/model"
	"fmt"
)

type school struct {
	name     string
	addr     string
	students []*model.Student
}

func (s *school) getName() string {
	return s.name
}

func main() {
	s := &school{
		name: "school",
		addr: "幸福小区1号",
		students: []*model.Student{
			{
				Name: "小明",
				Age:  10,
			},
			{
				Name: "小李",
				Age:  12,
			},
		},
	}
	name := s.getName()
	fmt.Println("name = ", name)

	for _, student := range s.students {
		//studentName := model.GetName(student)
		studentName := student.GetName()
		student.SetAge(17)
		fmt.Println("studentName = ", studentName)
		fmt.Println("age = ", student.Age)
	}
}
