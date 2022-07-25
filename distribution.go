package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/exp/maps"
)

// Fork_chain_essentials takes an existing trusted text chain, and generates a copy of it.
// This copy of the 'essentials' only takes the elements of the head_hash_tree with it,
// producing the effect of preserving and copying the core progression of the trusted text
// elements, and ignoring any un-promoted blocks
func Fork_chain_essentials(_trusted_text_chain trustedtext_chain_s) trustedtext_chain_s {
	essential_keys := maps.Keys(_trusted_text_chain.head_hash_tree)
	_trusted_text_chain.tt_chain = util_subset_map(_trusted_text_chain.tt_chain, essential_keys)
	return _trusted_text_chain
}

// Get_head_hashes_missing_from_comp takes a trusted text chain and a comparison list, and returns any missing keys
 func Get_head_hashes_missing_from_comp(_trusted_text_chain trustedtext_chain_s, _comparison_list []string) []string {
	all_head_hashes := _trusted_text_chain.head_hash_tree
	anti_set_map := util_anti_set_map(all_head_hashes, _comparison_list)
	return maps.Keys(anti_set_map)
 }

// Is_hash_in_chain is a function to determine if a hash is a part of the the trusted text chain
func Is_hash_in_chain(_trusted_text_chain trustedtext_chain_s, _comparison_hash string) bool {
	all_hashes := maps.Keys(_trusted_text_chain.tt_chain)
	check_map := util_make_boolean_map_from_slice(all_hashes)
	return check_map[_comparison_hash]
}


func check_with_peers(peerlist []peer_detail) error{
	if len(peerlist) == 0 {
		return errors.New("cant validate against empty peerlist")
	}
	missing_blocks := []string{}
	for _, peer := range peerlist {
		peers_missing_blocks, err := check_with_a_peer(peer)
		if err != nil {
			return err
		}
		for _, block := range peers_missing_blocks {
			retrieve_from_a_peer(peer, block)
		}
		missing_blocks = append(missing_blocks, peers_missing_blocks...)
	}

	if len(missing_blocks) != 0 {
		fmt.Println("Missing items on the web")
	}
	return nil
}

func check_with_a_peer(peer peer_detail) ([]string, error) {
	
	// Get and decode known blocks
	resp, err := http.Get("http://" + peer.Path + "/known_blocks")
	if err != nil {
		return []string{}, err
	}
	response_decoder := json.NewDecoder(resp.Body)
	known_blocks_of_peer := &[]string{}
	response_decoder.Decode(known_blocks_of_peer)

	// Determine missing elements
	peer_blocks_map := util_make_boolean_map_from_slice(*known_blocks_of_peer)
	my_blocks := maps.Keys(test_chain.tt_chain)

	new_keys_of_peer := util_anti_set_map(peer_blocks_map, my_blocks)

	return maps.Keys(new_keys_of_peer), nil
}

func retrieve_from_a_peer(peer peer_detail, block_hash string) error {
	resp, err := http.Get("http://" + peer.Path + "/block" + "?block_hash=" + block_hash)
	if err != nil {
		return err
	}
	response_decoder := json.NewDecoder(resp.Body) 
	returned_block := &trustedtext_s{}

	err = response_decoder.Decode(returned_block)
	if err != nil {
		return err
	}

	test_chain, err = Process_incoming_block(test_chain, *returned_block)
	if err != nil {
		return nil
	}

	return nil

}