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
	config, err := read_config(default_config_path)
	util_error_wrapper(w, err)
	existing_chain, err := read_chain(config)
	util_error_wrapper(w, err)

	requested_block, err := return_specified_hash(existing_chain, _block_hash)
	util_error_wrapper(w, err)
	text_block, err := json.Marshal(requested_block)
	util_error_wrapper(w, err)

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

	config, err := read_config(default_config_path)
	util_error_wrapper(w, err)
	existing_chain, err := read_chain(config)
	util_error_wrapper(w, err)


	new_chain, err := process_incoming_block(existing_chain, *resultant_block)
	util_error_wrapper(w, err)

	if err != nil {
		write_chain(new_chain, config)
	}

	w.WriteHeader(http.StatusCreated)
}

func give_known_blocks(w http.ResponseWriter, r *http.Request) {
	config, err := read_config(default_config_path)
	util_error_wrapper(w, err)

	existing_chain, err := read_chain(config)
	util_error_wrapper(w, err)
	
	output_encoder := json.NewEncoder(w)
	err = output_encoder.Encode(maps.Keys(existing_chain.Tt_chain))
	util_error_wrapper(w, err)
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
