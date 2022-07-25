package main

import "golang.org/x/exp/maps"



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

