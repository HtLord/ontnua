package model

type DocExist struct {
	IsExist bool
}

func DocExistWrapper() DocExist {
	return DocExist{IsExist: true}
}

func DocNotExistWrapper() DocExist {
	return DocExist{IsExist: false}
}

type ErrorResponse struct {
	ErrorMsg string
}

func ErrorResponseWrapper(err error) ErrorResponse {
	return ErrorResponse{ErrorMsg: err.Error()}
}
