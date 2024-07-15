package main

import (
	"fmt"
	// CHAINCODE MODE
	// DEV MOVE
	"studentcc/data/memory"
	"studentcc/model/student"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 1. variable declaration and initialization
// 1.1 types and pointers
// 2. if  else
// 3. for loops
// 4. functions
// 5. structs
// 6. methods
//
/*


 */

type State struct {
	DataContract *memory.DataMemory // DEV MODE
	// contract.DataContract // CHAINCODE MODE
}

func (s *State) Register(ctx contractapi.TransactionContextInterface, ID, name, email, phone, camp string) error {
	newStudent := student.New(ID, name, email, phone, camp)
	err := s.DataContract.PutStudent(ctx, newStudent)
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *State) GetStudent(ctx contractapi.TransactionContextInterface, ID string) (*student.Student, error) {
	return s.DataContract.GetStudent(ctx, ID)
}

func (s *State) GetAllStudent(ctx contractapi.TransactionContextInterface) ([]*student.Student, error) {
	return s.DataContract.GetAllStudents(ctx)
}

func (s *State) UpdateEmail(ctx contractapi.TransactionContextInterface, ID string, email string) error {
	oldInfo, err := s.GetStudent(ctx, ID)
	if err != nil {
		return err
	}
	oldInfo.Email = email
	return s.DataContract.PutStudent(ctx, oldInfo)
}

func (s *State) UpdatePhone(ctx contractapi.TransactionContextInterface, ID string, phone string) error {
	oldInfo, err := s.GetStudent(ctx, ID)
	if err != nil {
		return err
	}
	oldInfo.Phone = phone
	return s.DataContract.PutStudent(ctx, oldInfo)
}

func (s *State) UpdateCamp(ctx contractapi.TransactionContextInterface, ID string, camp string) error {
	oldInfo, err := s.GetStudent(ctx, ID)
	if err != nil {
		return err
	}
	oldInfo.Camp = camp
	return s.DataContract.PutStudent(ctx, oldInfo)
}

func (s *State) Grade(ctx contractapi.TransactionContextInterface, ID string, projectGrade float64) error {
	if projectGrade < 0 || projectGrade > 100 {
		return fmt.Errorf("The grade should be between 0 and 100")
	}
	oldInfo, err := s.GetStudent(ctx, ID)
	if err != nil {
		return err
	}
	oldInfo.ProjectGrade = projectGrade
	return s.DataContract.PutStudent(ctx, oldInfo)
}

// func mainChaincode() {

// 	chaincode, err := contractapi.NewChaincode(new(State))

// 	if err != nil {
// 		fmt.Printf("Error creating student chaincode: %v", err)
// 		return
// 	}
// 	if err := chaincode.Start(); err != nil {
// 		fmt.Printf("Error starting student chaincode: %v", err)
// 	}
// }

func main() {

	// mainChaincode()

	state := State{
		DataContract: memory.New(),
	}

	err := state.Register(nil, "201702068", "Marzoog AlGhazwi", "marzoog@marzoog.co", "0569910469", "Blockchain")

	if err != nil {
		fmt.Println(err)
	}

	err = state.Register(nil, "201702068", "Marzoog AlGhazwi", "marzoog@marzoog.co", "0569910469", "Blockchain")

	if err != nil {
		fmt.Println(err)
	}

	err = state.Register(nil, "201910291", "Ayyub", "ayyub@psu.edu.sa", "0591829182", "Blockchain")

	if err != nil {
		fmt.Println(err)
	}

	st1, _ := state.GetStudent(nil, "201702068")

	fmt.Printf("The student is: %v\n", st1)

	allStudents, _ := state.GetAllStudent(nil)

	for _, s := range allStudents {
		fmt.Println(s)
	}

	fmt.Println("Updating Marzoog Email...")

	err = state.UpdateEmail(nil, "201702068", "malghazwi@psu.edu.sa")

	allStudents, _ = state.GetAllStudent(nil)

	for _, s := range allStudents {
		fmt.Println(s)
	}

	fmt.Println("Updating Marzoog Phone...")

	err = state.UpdatePhone(nil, "201702068", "0569999112")

	allStudents, _ = state.GetAllStudent(nil)

	for _, s := range allStudents {
		fmt.Println(s)
	}

	fmt.Println("Grading Marzoog...")

	err = state.Grade(nil, "201", 90)

	if err != nil {
		fmt.Println(err)
	}

	allStudents, _ = state.GetAllStudent(nil)

	for _, s := range allStudents {
		fmt.Println(s)
	}
}
