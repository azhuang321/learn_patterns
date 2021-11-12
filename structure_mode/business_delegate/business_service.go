package business_delegate

import "fmt"

type BusinessService interface {
	DoProcessing()
}

type EJBService struct{}

func (s *EJBService) DoProcessing() {
	fmt.Println("processing task by invoking EJB service")
}

type JMSService struct{}

func (s *JMSService) DoProcessing() {
	fmt.Println("processing task by invoking JMS service")
}
