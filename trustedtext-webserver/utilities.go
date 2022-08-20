package main

import (
	"fmt"
	"net/http"
)

func util_error_wrapper(_response_writer http.ResponseWriter, _possible_error error) {
	if _possible_error != nil {
		fmt.Fprint(_response_writer,  _possible_error)
		_response_writer.WriteHeader(http.StatusInternalServerError)
	}
}
