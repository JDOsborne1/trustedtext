package trustedtext

import (
	"encoding/json"
	"errors"
)

type head_change_instruction struct {
	New_head string
}

// serialise_head_change takes an instruction for a changed head hash, and serialises it into a JSON string.
// This string is then used as the body of an instruction block.
func serialise_head_change(_change_to_serialise head_change_instruction) (string, error) {
	json_change, err := json.Marshal(_change_to_serialise)
	if err != nil {
		return "", err
	}
	return string(json_change), nil
}

// deserialise_head_change inverts the serialisation of an instruction string, and turns it back into an instruction object.
func deserialise_head_change(instruction_body_to_deserialise string) (head_change_instruction, error) {
	instruction := &head_change_instruction{}
	err := json.Unmarshal([]byte(instruction_body_to_deserialise), instruction)
	if err != nil {
		return head_change_instruction{}, err
	}
	return *instruction, nil
}

// Generate_head_move_block creates a new block, which contains only an instruction, signed by the private key of the author.
func Generate_head_move_block(_author string, _new_head_hash string, _private_key string) (Trustedtext_s, error) {
	change_instruction := head_change_instruction{New_head: _new_head_hash}

	serialised_change, err := serialise_head_change(change_instruction)

	instruction_body := tt_body{
		Instruction_type: "head_change",
		Instruction:      serialised_change,
	}

	if err != nil {
		return Trustedtext_s{}, err
	}

	new_element, err := instantiate(
		_author,
		instruction_body,
		_private_key,
	)
	if err != nil {
		return Trustedtext_s{}, err
	}
	return new_element, nil
}

// action_head_move_block first validates that the instruction was from the original author.
// It then moves the head hash.
func action_head_move_block(_existing_ttc Trustedtext_chain_s, _head_move_block Trustedtext_s) (Trustedtext_chain_s, error) {
	head_change_by_original_author, err := verify_hex_encoded_values(_existing_ttc.Original_author, _head_move_block.Hash, _head_move_block.Hash_signature)
	if err != nil {
		return Trustedtext_chain_s{}, err
	}

	if !head_change_by_original_author {
		return Trustedtext_chain_s{}, errors.New("head change block is not signed by original author")
	}

	head_change_value, err := deserialise_head_change(_head_move_block.Body.Instruction)
	if err != nil {
		return Trustedtext_chain_s{}, err
	}

	_existing_ttc, err = move_head_hash(_existing_ttc, head_change_value.New_head)
	if err != nil {
		return Trustedtext_chain_s{}, err
	}

	return _existing_ttc, nil
}

// move_head_hash is the function which executes the change of the head hash. At present this only validates
// that the suggested hash is actually in the chain
func move_head_hash(_existing_ttc Trustedtext_chain_s, _new_head_hash string) (Trustedtext_chain_s, error) {
	hash_found := _existing_ttc.Tt_chain[_new_head_hash].Body != tt_body{}
	if !hash_found {
		return Trustedtext_chain_s{}, errors.New("suggested new hash not in chain")
	}
	_existing_ttc.Head_hash = _new_head_hash
	_existing_ttc.Head_hash_tree[_new_head_hash] = true
	return _existing_ttc, nil
}
