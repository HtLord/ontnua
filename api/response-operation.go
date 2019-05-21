package api

/*
	New Json Response with default operation is success and return success response data
*/
func NewJRes() map[string]interface{} {
	resData := map[string]interface{}{
		"API-Call":      "accept",
		"Reject-Reason": "nil",
	}
	return resData
}
