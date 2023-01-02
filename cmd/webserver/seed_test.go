package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"trustedtext"
)

func test_helper_post_block_to_path(_block trustedtext.Block, _path string) (*http.Response, error) {
	marshalled_test_block, err := json.MarshalIndent(_block, "", " ")
	if err != nil {
		return &http.Response{}, err
	}

	test_block_reader := bytes.NewReader(marshalled_test_block)

	submission_request, err := http.NewRequest("POST", _path, test_block_reader)

	if err != nil {
		return &http.Response{}, err
	}

	submission_request.Header.Set("Content-Type", "application/json")

	sending_client := &http.Client{}

	return sending_client.Do(submission_request)
}

func test_helper_post_peer_to_path(_peer Peer_detail, _path string) (*http.Response, error) {
	marshalled_test_block, err := json.MarshalIndent(_peer, "", " ")
	if err != nil {
		return &http.Response{}, err
	}

	test_peer_reader := bytes.NewReader(marshalled_test_block)

	submission_request, err := http.NewRequest("POST", _path, test_peer_reader)

	if err != nil {
		return &http.Response{}, err
	}

	submission_request.Header.Set("Content-Type", "application/json")

	sending_client := &http.Client{}

	return sending_client.Do(submission_request)
}
