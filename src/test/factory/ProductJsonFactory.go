package factory

func GetCreateNewProductJsonRequestWithoutSomeParams() map[string]interface{} {
	return map[string]interface{}{
		"name":  "Test Product",
		"price": 10.99,
	}
}

func GetPageableJsonRequestWithoutSomeParams() map[string]interface{} {
	return map[string]interface{}{
		"Page":     1,
		"PageSize": 15,
	}
}

func GetPageableJsonRequestInvalidParam() map[string]interface{} {
	return map[string]interface{}{
		"Page":     "oi",
		"PageSize": 15,
		"Sort":     "ID desc",
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

func GetCreateProductResponseWithoutSomeParams() map[string]interface{} {
	var statusCode float64 = 400
	return map[string]interface{}{
		"error": map[string]interface{}{
			"message": "Key: 'NewProductRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag\nKey: 'NewProductRequest.Stock' Error:Field validation for 'Stock' failed on the 'required' tag",
		},
		"httpStatus": statusCode,
		"success":    false,
	}
}

func GetUpdateProductResponseWithoutSomeParams() map[string]interface{} {
	var statusCode float64 = 400
	return map[string]interface{}{
		"error": map[string]interface{}{
			"message": "Key: 'UpdateProductRequest.ID' Error:Field validation for 'ID' failed on the 'required' tag\nKey: 'UpdateProductRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag\nKey: 'UpdateProductRequest.Stock' Error:Field validation for 'Stock' failed on the 'required' tag",
		},
		"httpStatus": statusCode,
		"success":    false,
	}
}

func GetCreateProductResponseWithInvalidParam() map[string]interface{} {
	var statusCode float64 = 400
	return map[string]interface{}{
		"error": map[string]interface{}{
			"message": "code=400, message=Unmarshal type error: expected=int, got=string, field=stock, offset=77, internal=json: cannot unmarshal string into Go struct field NewProductRequest.stock of type int",
		},
		"httpStatus": statusCode,
		"success":    false,
	}
}

func GetUpdateProductResponseWithInvalidParam() map[string]interface{} {
	var statusCode float64 = 400
	return map[string]interface{}{
		"error": map[string]interface{}{
			"message": "code=400, message=Unmarshal type error: expected=int, got=string, field=stock, offset=77, internal=json: cannot unmarshal string into Go struct field UpdateProductRequest.stock of type int",
		},
		"httpStatus": statusCode,
		"success":    false,
	}
}

func GetAllProductsResponseWithInvalidParam() map[string]interface{} {
	var statusCode float64 = 400
	return map[string]interface{}{
		"error": map[string]interface{}{
			"message": "Key: 'UpdateProductRequest.ID' Error:Field validation for 'ID' failed on the 'required' tag\nKey: 'UpdateProductRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'UpdateProductRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag\nKey: 'UpdateProductRequest.Price' Error:Field validation for 'Price' failed on the 'required' tag\nKey: 'UpdateProductRequest.Stock' Error:Field validation for 'Stock' failed on the 'required' tag",
		},
		"httpStatus": statusCode,
		"success":    false,
	}
}

func GetAllProductsResponseWithoutSomeParams() map[string]interface{} {
	var statusCode float64 = 400
	return map[string]interface{}{
		"error": map[string]interface{}{
			"message": "Key: 'UpdateProductRequest.ID' Error:Field validation for 'ID' failed on the 'required' tag\nKey: 'UpdateProductRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'UpdateProductRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag\nKey: 'UpdateProductRequest.Price' Error:Field validation for 'Price' failed on the 'required' tag\nKey: 'UpdateProductRequest.Stock' Error:Field validation for 'Stock' failed on the 'required' tag",
		},
		"httpStatus": statusCode,
		"success":    false,
	}
}
