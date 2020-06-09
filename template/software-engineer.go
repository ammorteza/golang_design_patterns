package main

import (
	"fmt"
	"math/rand"
)

type SoftwareEngineer struct {

}

func (se SoftwareEngineer)PrepareResume(name string) StandardResume{
	fmt.Println(name , " is a software engineer and he/she is preparing his/her resume.")
	return StandardResume{
		Name: name,
	}
}

func (se SoftwareEngineer)FindingCompany() (Company, error){
	fmt.Println("The software engineer is finding company for applying.")
	return Company{
		ID: rand.Intn(100000),
	}, nil
}

func (se SoftwareEngineer)Apply(company Company, resume StandardResume) error{
	fmt.Println(resume.Name , " is applying to ", company.ID)
	return nil
}

func (se SoftwareEngineer)DoingInterview(company Company) error{
	fmt.Println("The software engineer is interviewing with ", company.ID)
	return nil
}

func (se SoftwareEngineer)MakingContract(company Company){
	fmt.Println("The software engineer doing contract with ", company.ID)
}