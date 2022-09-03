package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"trustedtext"

	"golang.org/x/exp/maps"
)

// fork_chain_essentials takes an existing trusted text chain, and generates a copy of it.
// This copy of the 'essentials' only takes the elements of the head_hash_tree with it,
// producing the effect of preserving and copying the core progression of the trusted text
// elements, and ignoring any un-promoted blocks
func fork_chain_essentials(_trusted_text_chain trustedtext.Trustedtext_chain_s) trustedtext.Trustedtext_chain_s {
	essential_keys := maps.Keys(_trusted_text_chain.Head_hash_tree)
	_trusted_text_chain.Tt_chain = trustedtext.Util_subset_map(_trusted_text_chain.Tt_chain, essential_keys)
	return _trusted_text_chain
}

// get_head_hashes_missing_from_comp takes a trusted text chain and a comparison list, and returns any missing keys
func get_head_hashes_missing_from_comp(_trusted_text_chain trustedtext.Trustedtext_chain_s, _comparison_list []string) []string {
	all_head_hashes := _trusted_text_chain.Head_hash_tree
	anti_set_map := trustedtext.Util_anti_set_map(all_head_hashes, _comparison_list)
	return maps.Keys(anti_set_map)
}

func Synchronise_with_peers(_peerlist []trustedtext.Peer_detail, _config trustedtext.Config_struct) error {
	if len(_peerlist) == 0 {
		return errors.New("cant validate against empty peerlist")
	}

	var err error

	for _, peer := range _peerlist {
		err = synchronise_with_peer(_config, peer)
		if err != nil {
			return err
		}
	}

	return nil
}

func synchronise_with_peer(_config trustedtext.Config_struct, _peer trustedtext.Peer_detail) error {
	existing_chain, err := trustedtext.Read_chain(_config)
	if err != nil {
		return err
	}

	current_blocks := maps.Keys(existing_chain.Tt_chain)

	peers_missing_blocks, err := check_with_a_peer(_peer, current_blocks)
	if err != nil {
		return err
	}

	returned_blocks, err := retrieve_blocklist_from_peer(peers_missing_blocks, _peer)
	if err != nil {
		return err
	}

	new_chain, err := trustedtext.Process_multiple_blocks(existing_chain, returned_blocks)
	if err != nil {
		return err
	}

	err = trustedtext.Write_chain(new_chain, _config)
	if err != nil {
		return err
	}

	return nil
}

func retrieve_blocklist_from_peer(_blocklist []string, _peer trustedtext.Peer_detail) ([]trustedtext.Trustedtext_s, error) {
	returned_blocklist := []trustedtext.Trustedtext_s{}
	for _, block := range _blocklist {
		retrieved_block, err := retrieve_from_a_peer(_peer, block)
		if err != nil {
			return []trustedtext.Trustedtext_s{}, err
		}
		returned_blocklist = append(returned_blocklist, retrieved_block)
	}
	return returned_blocklist, nil
}

func helper_format_external_block_list(_path string) (map[string]bool, error) {
	resp, err := http.Get(_path + "/all_blocks")
	if err != nil {
		return make(map[string]bool), err
	}
	response_decoder := json.NewDecoder(resp.Body)
	known_blocks_of_peer := &[]string{}
	err = response_decoder.Decode(known_blocks_of_peer)
	if err != nil {
		return make(map[string]bool), err
	}

	// Determine missing elements
	peer_blocks_map := trustedtext.Util_slice_to_bool_map(*known_blocks_of_peer)

	return peer_blocks_map, nil
}

func helper_format_external_peer_names(_path string) (map[string]bool, error) {
	resp, err := http.Get(_path + "/all_peers")
	if err != nil {
		return make(map[string]bool), err
	}
	response_decoder := json.NewDecoder(resp.Body)
	known_peers := &[]trustedtext.Peer_detail{}
	err = response_decoder.Decode(known_peers)
	if err != nil {
		return make(map[string]bool), err
	}

	peer_names := []string{}

	for _, peer := range *known_peers {
		peer_names = append(peer_names, peer.Claimed_name)
	}

	// Determine missing elements
	peer_details := trustedtext.Util_slice_to_bool_map(peer_names)

	return peer_details, nil
}

func check_with_a_peer(_peer trustedtext.Peer_detail, _existing_blocks []string) ([]string, error) {

	// Get and decode known blocks
	peer_blocks_map, err := helper_format_external_block_list("http://" + _peer.Path)
	if err != nil {
		return []string{}, err
	}

	new_keys_of_peer := trustedtext.Util_anti_set_map(peer_blocks_map, _existing_blocks)

	return maps.Keys(new_keys_of_peer), nil
}

func helper_retrieve_and_format_external_block(_path string) (trustedtext.Trustedtext_s, error) {
	resp, err := http.Get(_path)
	if err != nil {
		return trustedtext.Trustedtext_s{}, err
	}
	response_decoder := json.NewDecoder(resp.Body)
	returned_block := &trustedtext.Trustedtext_s{}

	err = response_decoder.Decode(returned_block)
	if err != nil {
		return trustedtext.Trustedtext_s{}, err
	}

	return *returned_block, nil
}


func retrieve_from_a_peer(peer trustedtext.Peer_detail, block_hash string) (trustedtext.Trustedtext_s, error) {
	composed_path := "http://" + peer.Path + "/block" + "/" + block_hash

	return helper_retrieve_and_format_external_block(composed_path)

}
