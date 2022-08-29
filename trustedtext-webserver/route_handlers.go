package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"trustedtext"

	"golang.org/x/exp/maps"
)

func give_block(w http.ResponseWriter, r *http.Request, _block_hash string) {
	config, err := trustedtext.Read_config(default_config_path)
	util_error_wrapper(w, err)
	existing_chain, err := trustedtext.Read_chain(config)
	util_error_wrapper(w, err)

	requested_block, err := trustedtext.Return_specified_hash(existing_chain, _block_hash)
	util_error_wrapper(w, err)
	text_block, err := json.Marshal(requested_block)
	util_error_wrapper(w, err)

	fmt.Fprint(w, string(text_block))
}

func give_head_block_md_processed(w http.ResponseWriter, r *http.Request) {

	requested_block, err := give_head_block_raw(w, r)
	util_error_wrapper(w, err)

	text_block, err := process_md_block(requested_block.Body.Instruction)
	util_error_wrapper(w, err)

	fmt.Fprint(w, text_block)

}

func give_head_block_unprocessed(w http.ResponseWriter, r *http.Request) {

	requested_block, err := give_head_block_raw(w, r)
	util_error_wrapper(w, err)

	text_block, err := json.Marshal(requested_block)
	util_error_wrapper(w, err)

	fmt.Fprint(w, string(text_block))

}



func give_head_block_raw(w http.ResponseWriter, r *http.Request) (trustedtext.Trustedtext_s, error) {
	config, err := trustedtext.Read_config(default_config_path)
	util_error_wrapper(w, err)

	existing_chain, err := trustedtext.Read_chain(config)
	util_error_wrapper(w, err)

	head_hash := existing_chain.Head_hash

	return trustedtext.Return_specified_hash(existing_chain, head_hash)
}

func submit_block(w http.ResponseWriter, r *http.Request) {
	var post_deposit []byte
	var err error
	post_deposit, err = io.ReadAll(r.Body)
	util_error_wrapper(w, err)

	resultant_block := &trustedtext.Trustedtext_s{}
	err = json.Unmarshal(post_deposit, resultant_block)
	util_error_wrapper(w, err)

	config, err := trustedtext.Read_config(default_config_path)
	util_error_wrapper(w, err)
	existing_chain, err := trustedtext.Read_chain(config)
	util_error_wrapper(w, err)

	new_chain, err := trustedtext.Process_incoming_block(existing_chain, *resultant_block)
	util_error_wrapper(w, err)

	if err == nil {
		err := trustedtext.Write_chain(new_chain, config)
		if err != nil {
			util_error_wrapper(w, err)
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func give_known_blocks(w http.ResponseWriter, r *http.Request) {
	config, err := trustedtext.Read_config(default_config_path)
	util_error_wrapper(w, err)

	existing_chain, err := trustedtext.Read_chain(config)
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
	used_config, err := trustedtext.Read_config(default_config_path)
	util_error_wrapper(w, err)

	peerlist, err := trustedtext.Read_peerlist(used_config)
	util_error_wrapper(w, err)

	marshalled_peerlist, err := json.Marshal(peerlist)
	util_error_wrapper(w, err)

	fmt.Fprint(w, string(marshalled_peerlist))
}

func add_peer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	var post_deposit []byte
	var err error
	post_deposit, err = io.ReadAll(r.Body)
	util_error_wrapper(w, err)

	resultant_peer := &trustedtext.Peer_detail{}
	err = json.Unmarshal(post_deposit, resultant_peer)
	util_error_wrapper(w, err)

	used_config, err := trustedtext.Read_config(default_config_path)
	util_error_wrapper(w, err)

	existing_peerlist, err := trustedtext.Read_peerlist(used_config)
	util_error_wrapper(w, err)

	new_peerlist := append(existing_peerlist, *resultant_peer)

	trustedtext.Write_peerlist(new_peerlist, used_config)
	w.WriteHeader(http.StatusCreated)
}

func peer_check(w http.ResponseWriter, r *http.Request) {

	used_config, err := trustedtext.Read_config(default_config_path)
	util_error_wrapper(w, err)

	peerlist, err := trustedtext.Read_peerlist(used_config)
	util_error_wrapper(w, err)

	err = Synchronise_with_peers(peerlist, used_config)
	util_error_wrapper(w, err)
	w.WriteHeader(http.StatusAccepted)
}
