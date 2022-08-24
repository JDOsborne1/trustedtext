package main

import (
	"net/http"
	"testing"
)

func Test_local_env_exists(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/all_blocks")

	if err != nil {
		t.Log("fails to get response to all blocks")
		t.Fail()
	}

	if resp.StatusCode != 200 {
		t.Log("response fails, with status: ", resp.StatusCode)
		t.Fail()
	}
}

