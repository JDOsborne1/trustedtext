package trustedtext

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/exp/maps"
)

type Peer_detail struct {
	Claimed_name string
	Path         string
}

// fork_chain_essentials takes an existing trusted text chain, and generates a copy of it.
// This copy of the 'essentials' only takes the elements of the head_hash_tree with it,
// producing the effect of preserving and copying the core progression of the trusted text
// elements, and ignoring any un-promoted blocks
func fork_chain_essentials(_trusted_text_chain Trustedtext_chain_s) Trustedtext_chain_s {
	essential_keys := maps.Keys(_trusted_text_chain.Head_hash_tree)
	_trusted_text_chain.Tt_chain = util_subset_map(_trusted_text_chain.Tt_chain, essential_keys)
	return _trusted_text_chain
}

// get_head_hashes_missing_from_comp takes a trusted text chain and a comparison list, and returns any missing keys
func get_head_hashes_missing_from_comp(_trusted_text_chain Trustedtext_chain_s, _comparison_list []string) []string {
	all_head_hashes := _trusted_text_chain.Head_hash_tree
	anti_set_map := util_anti_set_map(all_head_hashes, _comparison_list)
	return maps.Keys(anti_set_map)
}

func Synchronise_with_peers(_peerlist []Peer_detail, _config config_struct) error {
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

func synchronise_with_peer(_config config_struct, _peer Peer_detail) error {
	existing_chain, err := Read_chain(_config)
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

	new_chain, err := process_multiple_blocks(existing_chain, returned_blocks)
	if err != nil {
		return err
	}

	err = Write_chain(new_chain, _config)
	if err != nil {
		return err
	}

	return nil
}

func retrieve_blocklist_from_peer(_blocklist []string, _peer Peer_detail) ([]Trustedtext_s, error) {
	returned_blocklist := []Trustedtext_s{}
	for _, block := range _blocklist {
		retrieved_block, err := retrieve_from_a_peer(_peer, block)
		if err != nil {
			return []Trustedtext_s{}, err
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
	peer_blocks_map := util_make_boolean_map_from_slice(*known_blocks_of_peer)

	return peer_blocks_map, nil
}

func check_with_a_peer(_peer Peer_detail, _existing_blocks []string) ([]string, error) {

	// Get and decode known blocks
	peer_blocks_map, err := helper_format_external_block_list("http://" + _peer.Path)
	if err != nil {
		return []string{}, err
	}

	new_keys_of_peer := util_anti_set_map(peer_blocks_map, _existing_blocks)

	return maps.Keys(new_keys_of_peer), nil
}

func helper_retrieve_and_format_external_block(_path string) (Trustedtext_s, error) {
	resp, err := http.Get(_path)
	if err != nil {
		return Trustedtext_s{}, err
	}
	response_decoder := json.NewDecoder(resp.Body)
	returned_block := &Trustedtext_s{}

	err = response_decoder.Decode(returned_block)
	if err != nil {
		return Trustedtext_s{}, err
	}

	return *returned_block, nil
}

func retrieve_from_a_peer(peer Peer_detail, block_hash string) (Trustedtext_s, error) {
	composed_path := "http://" + peer.Path + "/block" + "/" + block_hash

	return helper_retrieve_and_format_external_block(composed_path)

}
