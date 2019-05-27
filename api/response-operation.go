package api

/*
	New Json Response with default operation is success and return success response data
*/
func AcceptRes() map[string]interface{} {
	resData := map[string]interface{}{
		"API-Call":      "accept",
		"Reject-Reason": "nil",
	}
	return resData
}

func AcceptResWith(opts map[string]interface{}) map[string]interface{} {
	resData := map[string]interface{}{
		"API-Call":      "accept",
		"Reject-Reason": "nil",
	}
	for k, v := range opts{
		resData[k] = v
	}
	return resData
}

func RejectRes(err error) map[string]interface{} {
	res := AcceptRes()
	res["API-Call"] = "reject"
	res["Reject-Reason"] = err.Error()
	return res
}

func RejectResWith(tag string, err error) map[string]interface{} {
	res := AcceptRes()
	res["API-Call"] = "reject"
	res["Reject-Tag"] = tag
	res["Reject-Reason"] = err.Error()
	return res
}