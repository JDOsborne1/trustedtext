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