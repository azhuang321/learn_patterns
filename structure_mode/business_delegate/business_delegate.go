package business_delegate

type BusinessDelegate struct {
	lookService     *BusinessLookup
	businessService BusinessService
	serviceType     string
}

func New() *BusinessDelegate {
	delegate := new(BusinessDelegate)
	delegate.lookService = new(BusinessLookup)
	return delegate
}

func (b *BusinessDelegate) SetServiceType(serviceType string) {
	b.serviceType = serviceType
}

func (b *BusinessDelegate) DoTask() {
	b.businessService = b.lookService.GetBusinessService(b.serviceType)
	b.businessService.DoProcessing()
}
