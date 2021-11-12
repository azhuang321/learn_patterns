package main

import "go-patterns/structure_mode/business_delegate"

func main() {
	businessDelegate := business_delegate.New()
	businessDelegate.SetServiceType("EJB")

	client := business_delegate.NewClient(businessDelegate)
	client.DoTask()

	businessDelegate.SetServiceType("JMS")
	client.DoTask()
}
