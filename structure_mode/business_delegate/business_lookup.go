package business_delegate

type BusinessLookup struct{}

func (b *BusinessLookup) GetBusinessService(serviceType string) BusinessService {
	if serviceType == "EJB" {
		return new(EJBService)
	} else {
		return new(JMSService)
	}
}
