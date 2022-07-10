package main

import (
	"errors"
)

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
	tags   []string
	head_hash string
	head_hash_tree map[string]bool
}

// Genesis is the function called to create a new trusted text chain. 
// This is always initiated with the same first message. Partly because 
// the first element lacks a 'previous hash' which makes it harder to 
// consider all the text bodies 'validated'. By defering the first real 
// message to actually be the second message, we can ensure that all real
// text data is held in blocks with both sides of the chain. 
func Genesis(_author string, _tags []string, _private_key string) (trustedtext_chain_s, error) {
	original_instruction := tt_body{
		instruction_type: "publish",
		instruction: "This is the origin message of a trusted text chain",
	}

	first_element, err := Instantiate(
		_author,
		original_instruction,
		_private_key,
	)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	

	inital_block_map := make(map[string]trustedtext_s)
	inital_block_map[first_element.hash] = first_element

	inital_head_tree := make(map[string]bool)
	inital_head_tree[first_element.hash] = true

	new_chain := trustedtext_chain_s{
		original_author: _author,
		tt_chain:        inital_block_map,
		tags: _tags,
		head_hash: 		 first_element.hash,
		head_hash_tree:  inital_head_tree,
	}
	return new_chain, nil
}

// Amend is the function called to increment a chain with a new tt block. This is 'stateless' 
// such that it creates a new chain, which is a copy of the previous, but for the inclusion of 
// a new block at the end. 
func Amend(_existing_ttc trustedtext_chain_s, _new_block trustedtext_s) (trustedtext_chain_s, error) {
	if len(_existing_ttc.tt_chain) == 0 {
		return trustedtext_chain_s{}, errors.New("cannot amend an empty chain")
	}
	current_head_hash := Head_hash(_existing_ttc)


	_new_block.head_hash_at_creation = current_head_hash

	_existing_ttc.tt_chain[_new_block.hash] =  _new_block
	return _existing_ttc, nil
}

// Head_hash is a function called to find a core chain identifier. This is the hash of the header block. 
// This header block may be moved over time, and points to the block which contains the current definitive 
// record of the trusted text element.
func Head_hash(_existing_trustedtext trustedtext_chain_s) string {
	return _existing_trustedtext.head_hash
}



// Return_head_block gives back the block object which is currently pointed to by the head hash.
func Return_head_block(_existing_ttc trustedtext_chain_s) (trustedtext_s, error) {
	current_head_hash := Head_hash(_existing_ttc)
	return Return_specified_hash(_existing_ttc, current_head_hash)
}

// Return_specified_hash returns a specific block in the chain
func Return_specified_hash(_existing_ttc trustedtext_chain_s, _specified_hash string) (trustedtext_s, error) {
	hash_found := _existing_ttc.tt_chain[_specified_hash].body != tt_body{}
	if !hash_found {
		return trustedtext_s{}, errors.New("head block not found in chain")
	}
	return _existing_ttc.tt_chain[_specified_hash], nil
}


func Process_incoming_block(_existing_ttc trustedtext_chain_s, _incoming_block trustedtext_s) (trustedtext_chain_s, error) {

	// Validate Block
	block_has_valid_signature, err := Verify_hex_encoded_values(_incoming_block.author, _incoming_block.hash, _incoming_block.hash_signature)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	if !block_has_valid_signature {
		return trustedtext_chain_s{}, errors.New("incoming block has invalid signature")
	}
	// Process block instruction


	// Append block
	
	_existing_ttc, err = Amend(_existing_ttc, _incoming_block)

	if err != nil {
		return trustedtext_chain_s{}, err
	}

	return _existing_ttc, nil
}

// func Dispatch_instruction_processor(_block_body string) func(trustedtext_chain_s) trustedtext_chain_s {

// }