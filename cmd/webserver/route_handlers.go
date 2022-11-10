package main

import (
	"encoding/json"
	"file"
	"fmt"
	"io"
	"net/http"
	"trustedtext"

	"golang.org/x/exp/maps"
)

// give_block writes out a specified block from a storage layer
func give_block(w http.ResponseWriter, r *http.Request, _store file.Storage, _block_hash string) {
	existing_chain, err := _store.Chain.Read_chain()
	util_error_wrapper(w, err)

	requested_block, err := trustedtext.Return_specified_hash(existing_chain, _block_hash)
	util_error_wrapper(w, err)
	text_block, err := json.Marshal(requested_block)
	util_error_wrapper(w, err)

	fmt.Fprint(w, string(text_block))
}

// Gives the head block as it would be after processing the content as markdown. This is to make it viable 
// to show a webpage using just trustedtext blocks
func give_head_block_md_processed(w http.ResponseWriter, r *http.Request, _store file.Storage) {

	requested_block, err := give_head_block_raw(w, r, _store)
	util_error_wrapper(w, err)

	text_block, err := process_md_block(requested_block.Body.Instruction)
	util_error_wrapper(w, err)

	fmt.Fprint(w, text_block)

}

// This function gives the head block unprocessed, to allow comparison on the JSON level
func give_head_block_unprocessed(w http.ResponseWriter, r *http.Request, _store file.Storage) {

	requested_block, err := give_head_block_raw(w, r, _store)
	util_error_wrapper(w, err)

	text_block, err := json.Marshal(requested_block)
	util_error_wrapper(w, err)

	fmt.Fprint(w, string(text_block))

}


// The baseline function to give the head block, unprocessed, so that it can have multiple downstream versions
func give_head_block_raw(w http.ResponseWriter, r *http.Request, _store file.Storage) (trustedtext.Trustedtext_s, error) {

	existing_chain, err := _store.Chain.Read_chain()
	util_error_wrapper(w, err)

	head_hash := existing_chain.Head_hash

	return trustedtext.Return_specified_hash(existing_chain, head_hash)
}

// Handler to accept incoming blocks in the form of a Post request, and action it within the provided store
func submit_block(w http.ResponseWriter, r *http.Request, _store file.Storage) {
	var post_deposit []byte
	var err error
	post_deposit, err = io.ReadAll(r.Body)
	util_error_wrapper(w, err)

	resultant_block := &trustedtext.Trustedtext_s{}
	err = json.Unmarshal(post_deposit, resultant_block)
	util_error_wrapper(w, err)


	existing_chain, err := _store.Chain.Read_chain()
	util_error_wrapper(w, err)

	new_chain, err := trustedtext.Process_incoming_block(existing_chain, *resultant_block)
	util_error_wrapper(w, err)

	if err == nil {
		err := _store.Chain.Write_chain(new_chain)
		if err != nil {
			util_error_wrapper(w, err)
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// Listing handler, which shares the set of all known blocks in the chain.
func give_known_blocks(w http.ResponseWriter, r *http.Request, _store file.Storage) {

	existing_chain, err := _store.Chain.Read_chain()
	util_error_wrapper(w, err)

	output_encoder := json.NewEncoder(w)
	err = output_encoder.Encode(maps.Keys(existing_chain.Tt_chain))
	util_error_wrapper(w, err)
}

// Peer management route which shares the peerlist of the server
func share_peerlist(w http.ResponseWriter, r *http.Request, _store file.Storage) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}


	peerlist, err := _store.Peerlist.Read_peerlist()
	util_error_wrapper(w, err)

	marshalled_peerlist, err := json.Marshal(peerlist)
	util_error_wrapper(w, err)

	fmt.Fprint(w, string(marshalled_peerlist))
}

// Peer management request which allows you to add a new peer to a given server
func add_peer(w http.ResponseWriter, r *http.Request, _store file.Storage) {
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


	existing_peerlist, err := _store.Peerlist.Read_peerlist()
	util_error_wrapper(w, err)

	new_peerlist := append(existing_peerlist, *resultant_peer)

	_store.Peerlist.Write_peerlist(new_peerlist)
	w.WriteHeader(http.StatusCreated)
}

// Calling route which triggers the Synchronisation process
func peer_check(w http.ResponseWriter, r *http.Request, _store file.Storage) {

	peerlist, err := _store.Peerlist.Read_peerlist()
	util_error_wrapper(w, err)

	err = Synchronise_with_peers(peerlist, _store)
	util_error_wrapper(w, err)
	w.WriteHeader(http.StatusAccepted)
}
