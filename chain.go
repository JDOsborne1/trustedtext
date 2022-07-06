package main

import "errors"

type Trustedtext_chain_i interface {
	Genesis(_author string, _tags []string) Trustedtext_chain_i
	Amend(_existing_ttc Trustedtext_chain_i, _author string, _body string) Trustedtext_chain_i
	Most_recent_hash(_existing_ttc Trustedtext_chain_i) string
	Head_hash(_existing_ttc Trustedtext_chain_i) string
	Move_head_hash(_existing_ttc Trustedtext_chain_i, _new_head_hash string) Trustedtext_chain_i
	Return_head_block(_existing_ttc Trustedtext_chain_i) trustedtext_s
}

type trustedtext_chain_s struct {
	original_author string
	tt_chain        map[string]trustedtext_s
	head_hash string
}

// Genesis is the function called to create a new trusted text chain. 
// This is always initiated with the same first message. Partly because 
// the first element lacks a 'previous hash' which makes it harder to 
// consider all the text bodies 'validated'. By defering the first real 
// message to actually be the second message, we can ensure that all real
// text data is held in blocks with both sides of the chain. 
func Genesis(_author string, _tags []string) (trustedtext_chain_s, error) {
	first_element, err := Instantiate(
		_author,
		_tags,
		"This is the origin message of a trusted text chain",
	)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	inital_map := make(map[string]trustedtext_s)
	inital_map[first_element.hash] = first_element
	new_chain := trustedtext_chain_s{
		original_author: _author,
		tt_chain:        inital_map,
		head_hash: 		 first_element.hash,
	}
	return new_chain, nil
}

// Amend is the function called to increment a chain with a new tt block. This is 'stateless' 
// such that it creates a new chain, which is a copy of the previous, but for the inclusion of 
// a new block at the end. 
func Amend(_existing_ttc trustedtext_chain_s, _author string, _body string) (trustedtext_chain_s, error) {
	if len(_existing_ttc.tt_chain) == 0 {
		return trustedtext_chain_s{}, errors.New("cannot amend an empty chain")
	}
	current_head_hash := Head_hash(_existing_ttc)
	new_element, err := Instantiate(
		_author,
		_existing_ttc.tt_chain[current_head_hash].tags,
		_body,
	)
	if err != nil {
		return trustedtext_chain_s{}, err
	}

	new_element.head_hash_at_creation = current_head_hash

	_existing_ttc.tt_chain[new_element.hash] =  new_element
	return _existing_ttc, nil
}

// Head_hash is a function called to find a core chain identifier. This is the hash of the header block. 
// This header block may be moved over time, and points to the block which contains the current definitive 
// record of the trusted text element.
func Head_hash(_existing_trustedtext trustedtext_chain_s) string {
	return _existing_trustedtext.head_hash
}

// get_valid_hashes is a function which returns an unsorted map of the block hashes present in the chain.
// When this map is called with an index of a non-existent hash, it will return the boolean default (false).
func get_valid_hashes(_existing_trustedtext trustedtext_chain_s) map[string]bool {
	valid_hashes :=  make(map[string]bool)

	for _, block := range _existing_trustedtext.tt_chain {
		valid_hashes[block.hash] = true
	}
	return valid_hashes
}

// Move_head_hash is the function which executes the change of the head hash. At present this only validates 
// that the suggested hash is actually in the chain
func Move_head_hash(_existing_ttc trustedtext_chain_s, _new_head_hash string) (trustedtext_chain_s, error) {
	valid_hashes := get_valid_hashes(_existing_ttc)
	if !valid_hashes[_new_head_hash] {
		return trustedtext_chain_s{}, errors.New("suggested new hash not in chain")
	}
	_existing_ttc.head_hash = _new_head_hash
	return _existing_ttc, nil
}

// Return_head_block gives back the block object which is currently pointed to by the head hash.
func Return_head_block(_existing_ttc trustedtext_chain_s) (trustedtext_s, error) {
	current_head_hash := Head_hash(_existing_ttc)
	for _, block := range _existing_ttc.tt_chain {
		if block.hash == current_head_hash {
			return block, nil
		}
	}
	return trustedtext_s{}, errors.New("head block not found in chain")
}