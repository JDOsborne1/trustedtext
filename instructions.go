package main

import "encoding/json"

type head_change_instruction struct {
	New_head string
}

func Serialise_head_change(_change_to_serialise head_change_instruction) (string, error) {
	json_change, err := json.Marshal(_change_to_serialise)
	if err != nil {
		return "", err
	}
	return string(json_change), nil
}