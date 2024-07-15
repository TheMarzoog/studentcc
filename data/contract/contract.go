package contract

import (
	"encoding/json"
	"fmt"
	"studentcc/model/student"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type DataContract struct {
	contractapi.Contract
}

func (dc *DataContract) GetStudent(ctx contractapi.TransactionContextInterface, ID string) (*student.Student, error) {
	studentJSON, err := ctx.GetStub().GetState(ID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if studentJSON == nil {
		return nil, fmt.Errorf("student %s does not exist", ID)
	}
	var student student.Student
	err = json.Unmarshal(studentJSON, &student)

	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (dc *DataContract) GetAllStudents(ctx contractapi.TransactionContextInterface) ([]*student.Student, error) {
	studentsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	defer studentsIterator.Close()
	var students []*student.Student
	for studentsIterator.HasNext() {
		result, err := studentsIterator.Next()
		if err != nil {
			return nil, err
		}
		var student student.Student
		err = json.Unmarshal(result.Value, &student)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}
	return students, nil
}

func (dc *DataContract) PutStudent(ctx contractapi.TransactionContextInterface, student *student.Student) error {
	studentJSON, err := json.Marshal(student)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(student.ID, studentJSON)
}
