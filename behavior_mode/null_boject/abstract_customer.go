package null_boject

type AbstractCustomer interface {
	IsNil() bool
	GetName() string
}

type RealCustomer struct {
	name string
}

func NewRealCustomer(name string) *RealCustomer {
	rp := new(RealCustomer)
	rp.name = name
	return rp
}

func (r *RealCustomer) GetName() string {
	return r.name
}

func (r *RealCustomer) IsNil() bool {
	return false
}

type NullCustomer struct {
	name string
}

func NewNullCustomer() *NullCustomer {
	rp := new(NullCustomer)
	rp.name = ""
	return rp
}

func (r *NullCustomer) GetName() string {
	return "not available"
}

func (r *NullCustomer) IsNil() bool {
	return true
}
