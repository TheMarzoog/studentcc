package memory

import (
	"fmt"
	"studentcc/model/student"
)

type DataMemory map[string]*student.Student

func New() *DataMemory {
	sm := make(DataMemory)
	return &sm
}

func (sm *DataMemory) GetStudent(_ any, ID string) (*student.Student, error) {
	student, existing := (*sm)[ID]
	if !existing {
		return nil, fmt.Errorf("no student with ID %s", ID)
	}
	return student, nil
}

func (sm *DataMemory) GetAllStudents(_ any) ([]*student.Student, error) {
	var students []*student.Student
	for _, s := range *sm {
		students = append(students, s)
	}
	if len(students) == 0 {
		return nil, fmt.Errorf("no students in state memory")
	}
	return students, nil
}

func (sm *DataMemory) PutStudent(_ any, student *student.Student) error {
	if _, err := sm.GetStudent(nil, student.ID); err == nil {
		return err
	}
	(*sm)[student.ID] = student
	fmt.Printf("Student %v added successfully", &student)
	return nil
}
