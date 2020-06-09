package main

import (
	"fmt"
	"math/rand"
)

type ElectronicEngineer struct {

}

func (ee ElectronicEngineer)PrepareResume(name string) StandardResume{
	fmt.Println(name , " is a electronic engineer and he/she is preparing his/her resume.")
	return StandardResume{
		Name: name,
	}
}

func (ee ElectronicEngineer)FindingCompany() (Company, error){
	fmt.Println("The electronic engineer is finding company for applying.")
	return Company{
		ID: rand.Intn(100000),
	}, nil
}

func (ee ElectronicEngineer)Apply(company Company, resume StandardResume) error{
	fmt.Println(resume.Name , " is applying to ", company.ID)
	return nil
}

func (ee ElectronicEngineer)DoingInterview(company Company) error{
	fmt.Println("The electronic engineer is interviewing with ", company.ID)
	return nil
}

func (ee ElectronicEngineer)MakingContract(company Company){
	fmt.Println("The electronic engineer doing contract with ", company.ID)
}