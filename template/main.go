package main

import "fmt"

func main(){
	jobSeeker := JobSeeker{}
	jobSeeker.iinterview = ElectronicEngineer{}
	if err := jobSeeker.ApplyAndDoExam("Ali Amzajerdi"); err != nil{
		fmt.Println("Ali could not find a job until now!")
	}

	jobSeeker.iinterview = SoftwareEngineer{}
	if err := jobSeeker.ApplyAndDoExam("Morteza Amzajerdi"); err != nil{
		fmt.Println("Morteza could not find a job until now!")
	}
}