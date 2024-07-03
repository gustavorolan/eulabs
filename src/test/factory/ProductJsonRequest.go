package factory

func GetCreateNewProductJsonRequest() map[string]interface{} {
	return map[string]interface{}{
		"name":  "Test Product",
		"price": 10.99,
	}
}
