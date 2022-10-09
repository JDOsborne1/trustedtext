package main

import (
	"encoding/json"
	"errors"
	"file"
	"net/http"
	"trustedtext"

	"golang.org/x/exp/maps"
)

type Peer_detail struct {
	Claimed_name string
	Path         string
}




func Synchronise_with_peers(_peerlist []trustedtext.Peer_detail, _store file.Storage) error {
	if len(_peerlist) == 0 {
		return errors.New("cant validate against empty peerlist")
	}

	var err error

	for _, peer := range _peerlist {
		err = synchronise_with_peer(_store, peer)
		if err != nil {
			return err
		}
	}

	return nil
}

func synchronise_with_peer(_store file.Storage, _peer trustedtext.Peer_detail) error {
	existing_chain, err := _store.Chain.Read_chain()
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

	err = _store.Chain.Write_chain(new_chain)
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
	known_peers := &[]Peer_detail{}
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
