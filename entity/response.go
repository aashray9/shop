package entity

func SuccessResponse(result map[string]interface{}) map[string]interface{} {
	result["success"] = true
	return result
}
