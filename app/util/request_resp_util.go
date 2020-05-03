package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	code    int
	status  string
	message string
}

func ParseRequest(r *http.Request) ([]byte, error) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}
	return reqBody, nil
}

func GetDefaultFailureResponse(code int, status string, message string) *Response {
	if (code == 0) {
		code = 404
	}
	if (IsEmpty(status)) {
		status = "failed"
	}
	if (IsEmpty(message)) {
		status = "Internal Server Error"
	}
	return &Response{
		code:    code,
		status:  status,
		message: message,
	}
}

func WriteResponseIntoOutputStream(w http.ResponseWriter,response interface{}) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "Something went wrong!")
	}
}
