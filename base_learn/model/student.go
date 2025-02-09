package model

type Student struct {
	Name string
	Age  int
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

func SetAge(s Student, age int) {
	s.Age = age
}

//func GetName(s *Student) string {
//	return s.Name
//}
