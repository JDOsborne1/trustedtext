package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/exp/maps"
)

type peer_detail struct {
	Claimed_name string
	Path         string
}

func give_block(w http.ResponseWriter, r *http.Request, _block_hash string) {
	requested_block, _ := return_specified_hash(test_chain, _block_hash)
	text_block, _ := json.Marshal(requested_block)
	fmt.Fprint(w, string(text_block))
}

func submit_block(w http.ResponseWriter, r *http.Request) {
	var post_deposit []byte
	var err error
	post_deposit, err = ioutil.ReadAll(r.Body)
	util_error_wrapper(w, err)

	resultant_block := &trustedtext_s{}
	err = json.Unmarshal(post_deposit, resultant_block)
	util_error_wrapper(w, err)

	new_chain, err := process_incoming_block(test_chain, *resultant_block)
	util_error_wrapper(w, err)

	if err != nil {
		test_chain = new_chain
	}

	w.WriteHeader(http.StatusCreated)
}

func give_known_blocks(w http.ResponseWriter, r *http.Request) {
	output_encoder := json.NewEncoder(w)
	output_encoder.Encode(maps.Keys(test_chain.Tt_chain))
}

func share_peerlist(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	used_config, err := read_config(default_config_path)
	util_error_wrapper(w, err)

	peerlist, err := read_peerlist(used_config)
	util_error_wrapper(w, err)

	marshalled_peerlist, err := json.Marshal(peerlist)
	util_error_wrapper(w, err)

	fmt.Fprint(w, string(marshalled_peerlist))
}

func add_peer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	var post_deposit []byte
	var err error
	post_deposit, err = ioutil.ReadAll(r.Body)
	util_error_wrapper(w, err)

	resultant_peer := &peer_detail{}
	err = json.Unmarshal(post_deposit, resultant_peer)
	util_error_wrapper(w, err)

	used_config, err := read_config(default_config_path)
	util_error_wrapper(w, err)

	existing_peerlist, err := read_peerlist(used_config)
	util_error_wrapper(w, err)

	new_peerlist := append(existing_peerlist, *resultant_peer)

	write_peerlist(new_peerlist, used_config)
	w.WriteHeader(http.StatusCreated)
}

func peer_check_handler(w http.ResponseWriter, r *http.Request) {

	used_config, err := read_config(default_config_path)
	util_error_wrapper(w, err)

	peerlist, err := read_peerlist(used_config)
	util_error_wrapper(w, err)

	err = synchronise_with_peers(peerlist, used_config)
	util_error_wrapper(w, err)
	w.WriteHeader(http.StatusAccepted)
}
