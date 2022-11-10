package main

import (
	"net/http"
)

// util_error_wrapper serves to simplify the code flow a bit by wrapping the error handling flow
func util_error_wrapper(_response_writer http.ResponseWriter, _possible_error error) {
	if _possible_error != nil {
		http.Error(_response_writer, _possible_error.Error(), http.StatusInternalServerError)
	}
}
