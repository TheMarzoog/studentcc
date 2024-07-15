package student

import "fmt"

type Student struct {
	ID           string
	Name         string
	Email        string
	Phone        string
	Camp         string
	ProjectGrade float64
}

func New(ID, name, email, phone, camp string) *Student {
	return &Student{
		ID:           ID,
		Name:         name,
		Phone:        phone,
		Camp:         camp,
		ProjectGrade: 0.0,
	}
}

func (s *Student) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Email: %s, Phone: %s, Camp: %s, Project Grade: %.2f", s.ID, s.Name, s.Email, s.Phone, s.Camp, s.ProjectGrade)
}
