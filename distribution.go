package main

import "golang.org/x/exp/maps"

// util_make_boolean_map_from_slice is a function which takes a slice of strings, and turns it into 
// a convenient shape for easy checking
func util_make_boolean_map_from_slice(input_slice_of_keys []string) map[string]bool {
	mapped_keys_to_keep := make(map[string]bool)

	for _, key := range input_slice_of_keys {
		mapped_keys_to_keep[key] = true
	}	

	return mapped_keys_to_keep
}

// util_subset_map is a function to return a subset of the input map, based on a slice of keys to keep. 
func util_subset_map(_original_map map[string]trustedtext_s, _keys_to_keep []string) map[string]trustedtext_s {
	mapped_keys_to_keep := util_make_boolean_map_from_slice(_keys_to_keep)

	new_map := make(map[string]trustedtext_s)
	for key , value := range _original_map {
		if mapped_keys_to_keep[key] {
			new_map[key] = value
		}
	} 

	return new_map
}

// Fork_chain_essentials takes an existing trusted text chain, and generates a copy of it. 
// This copy of the 'essentials' only takes the elements of the head_hash_tree with it, 
// producing the effect of preserving and copying the core progression of the trusted text 
// elements, and ignoring any un-promoted blocks
func Fork_chain_essentials(_trusted_text_chain trustedtext_chain_s) trustedtext_chain_s {
	essential_keys := maps.Keys(_trusted_text_chain.head_hash_tree)
	_trusted_text_chain.tt_chain = util_subset_map(_trusted_text_chain.tt_chain, essential_keys)
	return _trusted_text_chain
}