package main

import (
	"net/http"
)

func util_error_wrapper(_response_writer http.ResponseWriter, _possible_error error) {
	if _possible_error != nil {
		http.Error(_response_writer, _possible_error.Error(), http.StatusInternalServerError)
	}
}
