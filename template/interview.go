package main

type StandardResume struct {
	Name 		string
}

type Company struct {
	ID 			int
}

type iinterview interface {
	PrepareResume(string) StandardResume
	FindingCompany() (Company, error)
	Apply(company Company, resume StandardResume) error
	DoingInterview(company Company) error
	MakingContract(company Company)
}

type JobSeeker struct {
	iinterview
}

func (js JobSeeker)ApplyAndDoExam(name string) error{
	resume := js.iinterview.PrepareResume(name)
	company, err := js.FindingCompany()
	if err != nil{
		return err
	}
	
	if err := js.Apply(company, resume); err != nil{
		return err
	}
	
	if err := js.DoingInterview(company); err != nil{
		return err
	}
	
	js.MakingContract(company)
	return nil
}