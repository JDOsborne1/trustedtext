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

func give_block(w http.ResponseWriter, r *http.Request) {
	parsed_q := r.URL.Query()
	requested_hash := parsed_q["block_hash"][0]
	requested_block, _ := Return_specified_hash(test_chain, requested_hash)
	text_block, _ := json.Marshal(requested_block)
	fmt.Fprint(w, string(text_block))
}

func give_known_blocks(w http.ResponseWriter, r *http.Request) {
	output_encoder := json.NewEncoder(w)
	output_encoder.Encode(maps.Keys(test_chain.Tt_chain))
}

func submit_block(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var post_deposit []byte
	var err error
	post_deposit, err = ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	resultant_block := &trustedtext_s{}
	err = json.Unmarshal(post_deposit, resultant_block)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	new_chain, err := Process_incoming_block(test_chain, *resultant_block)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	if err != nil {
		test_chain = new_chain
	}

	w.WriteHeader(http.StatusCreated)
}

func share_peerlist(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	used_config, err := read_config(default_config_path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	peerlist, err := read_peerlist(used_config)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	marshalled_peerlist, err := json.Marshal(peerlist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, string(marshalled_peerlist))
}

func add_peer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	var post_deposit []byte
	var err error
	post_deposit, err = ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	resultant_peer := &peer_detail{}
	err = json.Unmarshal(post_deposit, resultant_peer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	used_config, err := read_config(default_config_path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	existing_peerlist, err := read_peerlist(used_config)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	new_peerlist := append(existing_peerlist, *resultant_peer)

	write_peerlist(new_peerlist, used_config)
}



func peer_check_handler(w http.ResponseWriter, r *http.Request) {

	used_config, _ := read_config(default_config_path)
	peerlist, _ := read_peerlist(used_config)
	err := check_with_peers(peerlist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
}