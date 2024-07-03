package factory

func GetCreateNewProductJsonRequestWithoutSomeParams() map[string]interface{} {
	return map[string]interface{}{
		"name":  "Test Product",
		"price": 10.99,
	}
}

func GetCreateNewProductJsonRequestInvalidParam() map[string]interface{} {
	return map[string]interface{}{
		"name":        "Test Product",
		"price":       10.99,
		"description": "description",
		"stock":       "oi",
	}
}

func GetProductResponseWithoutSomeParams() map[string]interface{} {
	var statusCode float64 = 400
	return map[string]interface{}{
		"error": map[string]interface{}{
			"message": "Key: 'NewProductRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag\nKey: 'NewProductRequest.Stock' Error:Field validation for 'Stock' failed on the 'required' tag",
		},
		"http_status": statusCode,
		"success":     false,
	}
}

func GetProductResponseWithInvalidParam() map[string]interface{} {
	var statusCode float64 = 500
	return map[string]interface{}{
		"error": map[string]interface{}{
			"message": "code=400, message=Unmarshal type error: expected=int, got=string, field=stock, offset=77, internal=json: cannot unmarshal string into Go struct field NewProductRequest.stock of type int",
		},
		"http_status": statusCode,
		"success":     false,
	}
}
