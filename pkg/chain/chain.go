package trustedtext

import (
	"errors"

	"golang.org/x/exp/maps"
)

type Chain struct {
	Original_author string
	Tt_chain        map[string]Block
	Tags            []string
	Head_hash       string
	Head_hash_tree  map[string]bool
}

type Chain_store interface {
	Write_chain(Chain) error
	Read_chain() (Chain, error)
}

// genesis is the function called to create a new trusted text chain.
// This is always initiated with the same first message. Partly because
// the first element lacks a 'previous hash' which makes it harder to
// consider all the text bodies 'validated'. By defering the first real
// message to actually be the second message, we can ensure that all real
// text data is held in blocks with both sides of the chain.
func genesis(_author string, _tags []string, _private_key string) (Chain, error) {
	original_instruction := body{
		Instruction_type: "publish",
		Instruction:      "This is the origin message of a trusted text chain",
	}

	first_element, err := instantiate(
		_author,
		original_instruction,
		_private_key,
	)
	if err != nil {
		return Chain{}, err
	}

	inital_block_map := make(map[string]Block)
	inital_block_map[first_element.Hash] = first_element

	inital_head_tree := make(map[string]bool)
	inital_head_tree[first_element.Hash] = true

	new_chain := Chain{
		Original_author: _author,
		Tt_chain:        inital_block_map,
		Tags:            _tags,
		Head_hash:       first_element.Hash,
		Head_hash_tree:  inital_head_tree,
	}
	return new_chain, nil
}

// amend is the function called to increment a chain with a new tt block. This is 'stateless'
// such that it creates a new chain, which is a copy of the previous, but for the inclusion of
// a new block at the end.
func amend(_existing_ttc Chain, _new_block Block) (Chain, error) {
	if len(_existing_ttc.Tt_chain) == 0 {
		return Chain{}, errors.New("cannot amend an empty chain")
	}
	current_head_hash := head_hash(_existing_ttc)

	_new_block.Head_hash_at_creation = current_head_hash

	_existing_ttc.Tt_chain[_new_block.Hash] = _new_block
	return _existing_ttc, nil
}

// head_hash is a function called to find a core chain identifier. This is the hash of the header block.
// This header block may be moved over time, and points to the block which contains the current definitive
// record of the trusted text element.
func head_hash(_existing_trustedtext Chain) string {
	return _existing_trustedtext.Head_hash
}

// return_head_block gives back the block object which is currently pointed to by the head hash.
func return_head_block(_existing_ttc Chain) (Block, error) {
	current_head_hash := head_hash(_existing_ttc)
	return Return_specified_hash(_existing_ttc, current_head_hash)
}

// Return_specified_hash returns a specific block in the chain
func Return_specified_hash(_existing_ttc Chain, _specified_hash string) (Block, error) {
	hash_found := _existing_ttc.Tt_chain[_specified_hash].Body != body{}
	if !hash_found {
		return Block{}, errors.New("specified block not found in chain")
	}
	return _existing_ttc.Tt_chain[_specified_hash], nil
}

func Process_incoming_block(_existing_ttc Chain, _incoming_block Block) (Chain, error) {

	// Validate Block
	block_has_valid_signature, err := verify_hex_encoded_values(_incoming_block.Author, _incoming_block.Hash, _incoming_block.Hash_signature)
	if err != nil {
		return Chain{}, err
	}
	if !block_has_valid_signature {
		return Chain{}, errors.New("incoming block has invalid signature")
	}

	// Check if block is already added
	hashes_in_chain := maps.Keys(_existing_ttc.Tt_chain)
	in_chain_map := Util_slice_to_bool_map(hashes_in_chain)
	hash_already_in_chain := in_chain_map[_incoming_block.Hash]
	if hash_already_in_chain {
		return Chain{}, errors.New("incoming block already in chain")
	}

	// Process block instruction
	processor := dispatch_instruction_processor(_incoming_block)

	_existing_ttc, err = processor(_existing_ttc)
	if err != nil {
		return Chain{}, err
	}

	// Append block

	_existing_ttc, err = amend(_existing_ttc, _incoming_block)

	if err != nil {
		return Chain{}, err
	}

	return _existing_ttc, nil
}

func dispatch_instruction_processor(_block Block) func(Chain) (Chain, error) {
	if _block.Body.Instruction_type == "head_change" {
		return func(_input_ttc Chain) (Chain, error) {
			return action_head_move_block(_input_ttc, _block)
		}
	}
	return func(_input_ttc Chain) (Chain, error) {
		return _input_ttc, nil
	}
}

func Process_multiple_blocks(_incoming_chain Chain, _incoming_list_of_blocks []Block) (Chain, error) {
	var err error

	for _, block := range _incoming_list_of_blocks {
		_incoming_chain, err = Process_incoming_block(_incoming_chain, block)
		if err != nil {
			return Chain{}, err
		}
	}

	return _incoming_chain, nil
}

// is_hash_in_chain is a function to determine if a hash is a part of the the trusted text chain
func is_hash_in_chain(_trusted_text_chain Chain, _comparison_hash string) bool {
	all_hashes := maps.Keys(_trusted_text_chain.Tt_chain)
	check_map := Util_slice_to_bool_map(all_hashes)
	return check_map[_comparison_hash]
}
