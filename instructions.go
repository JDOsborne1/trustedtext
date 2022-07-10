package main

import (
	"encoding/json"
	"errors"
)

type head_change_instruction struct {
	New_head string
}

func Serialise_head_change(_change_to_serialise head_change_instruction) (string, error) {
	json_change, err := json.Marshal(_change_to_serialise)
	if err != nil {
		return "", err
	}
	return string(json_change), nil
}

func Generate_head_move_block(_author string, _new_head_hash string, _private_key string) (trustedtext_s, error) {
	change_instruction := head_change_instruction{New_head: _new_head_hash}

	serialised_change, err := Serialise_head_change(change_instruction)

	if err != nil {
		return trustedtext_s{}, err
	}
	
	new_element, err := Instantiate(
		_author,
		serialised_change,
		_private_key,
	)
	if err != nil {
		return trustedtext_s{}, err
	}
	return new_element, nil
}

func Amend_with_head_move_block(_existing_ttc trustedtext_chain_s, _author string, _new_head_hash string, _private_key string) (trustedtext_chain_s, error) {

	new_element, err := Generate_head_move_block(
		_author,
		_new_head_hash,
		_private_key,
	)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	
	head_change_by_original_author, err := Verify_hex_encoded_values(_existing_ttc.original_author, new_element.body, new_element.hash_signature)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	
	if !head_change_by_original_author {
		return trustedtext_chain_s{}, errors.New("head change block is not signed by original author")
	}
	
	_existing_ttc, err = Amend(_existing_ttc, new_element)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	
	_existing_ttc, err = Move_head_hash(_existing_ttc, _new_head_hash)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	

	return _existing_ttc, nil

}

// Move_head_hash is the function which executes the change of the head hash. At present this only validates 
// that the suggested hash is actually in the chain
func Move_head_hash(_existing_ttc trustedtext_chain_s, _new_head_hash string) (trustedtext_chain_s, error) {
	hash_found := _existing_ttc.tt_chain[_new_head_hash].body != ""
	if !hash_found {
		return trustedtext_chain_s{}, errors.New("suggested new hash not in chain")
	}
	_existing_ttc.head_hash = _new_head_hash
	_existing_ttc.head_hash_tree[_new_head_hash] = true
	return _existing_ttc, nil
}