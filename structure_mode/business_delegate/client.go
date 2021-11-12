package business_delegate

type Client struct {
	businessService *BusinessDelegate
}

func NewClient(service *BusinessDelegate) *Client {
	client := new(Client)
	client.businessService = service
	return client
}

func (c *Client) DoTask() {
	c.businessService.DoTask()
}
