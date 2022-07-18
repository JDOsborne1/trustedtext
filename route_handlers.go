package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/exp/maps"
)

func give_block(w http.ResponseWriter, r *http.Request) {
	parsed_q := r.URL.Query()
	requested_hash := parsed_q["block_hash"][0]
	requested_block, _ := Return_specified_hash(test_chain, requested_hash)
	text_block, _ := json.Marshal(requested_block)
	fmt.Fprint(w, string(text_block))
}

func give_known_blocks(w http.ResponseWriter, r *http.Request) {
	output_encoder := json.NewEncoder(w)
	output_encoder.Encode(maps.Keys(test_chain.tt_chain))
}

func extract_submitted_block(r *http.Request) (trustedtext_s, error) {
	var post_deposit []byte
	var err error
	post_deposit, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return trustedtext_s{}, err
	}
	resultant_block := &trustedtext_s{}
	err = json.Unmarshal(post_deposit, resultant_block)
	if err != nil {
		return trustedtext_s{}, err
	}
	return *resultant_block, nil
}