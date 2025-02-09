package main

import "fmt"

type Skill interface {
	GetName() string
}

type BasketballSkill struct {
}

func (a *BasketballSkill) GetName() string {
	return "篮球"
}

type Student struct {
	Name string
	Age  int
	Skill
}

type Manager interface {
	// Join 学生加入学校
	Join(s *Student)
	// FindStudentByName 通过名字查询学生
	FindStudentByName(name string) *Student
}

type School interface {
	Manager
	// GetName 获取学校名称
	GetName() string
	// GetAddr 获取学校地址
	GetAddr() string
}

type MiddleSchool struct {
	Students []*Student
	Name     string
	Addr     string
}

func (m *MiddleSchool) FindStudentByName(name string) *Student {
	for _, s := range m.Students {
		if s.Name == name {
			return s
		}
	}
	return nil
}

func (m *MiddleSchool) GetName() string {
	return m.Name
}

func (m *MiddleSchool) GetAddr() string {
	return m.Addr
}

func (m *MiddleSchool) Join(s *Student) {
	m.Students = append(m.Students, s)
}

func main() {
	basketball := &BasketballSkill{}
	nike := &Student{
		Name:  "nike",
		Age:   17,
		Skill: basketball,
	}

	var school School
	school = &MiddleSchool{
		Name: "school",
		Addr: "幸福小区1号",
	}
	fmt.Println("name  ", school.GetName())

	school.Join(nike)
	s := school.FindStudentByName("nike")

	fmt.Println("s.Expertise.GetName() = ", s.Skill.GetName())
}
