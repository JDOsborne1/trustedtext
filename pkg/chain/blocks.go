package trustedtext

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
)

// Block is the data struct for a block of trusted text, of which all the fields can
// be validated against one another to ensure validity. The exception being the `Head_hash_at_creation`
// which is not included since it lends itself to races without helping with the trust level.
type Block struct {
	Author                string
	Body                  body
	Head_hash_at_creation string
	Hash                  string
	Hash_signature        string
}

type body struct {
	Instruction_type string
	Instruction      string
}

// This function is called to generate a base instance of the trustedtext block, based
// on its arguments, and then sign it. A block can be valid without tags, as it remains
// globally unique to the original author, however it must have both an author and body.
func instantiate(_author string, _body body, _private_key string) (Block, error) {
	if len(_author) == 0 {
		return Block{}, errors.New("cannot have a missing author")
	}

	if len(_body.Instruction_type) == 0 {
		return Block{}, errors.New("cannot have a missing instruction type")
	}
	if len(_body.Instruction) == 0 {
		return Block{}, errors.New("cannot have an empty instruction")
	}
	valid_signature_pairs, err := encoded_key_pair_is_valid(_author, _private_key)
	if err != nil {
		return Block{}, err
	}
	if !valid_signature_pairs {
		return Block{}, errors.New("author and key combination don't match")
	}

	tt_no_hash := Block{Author: _author, Body: _body}
	tt_with_hash, err := hash_tt(tt_no_hash)
	if err != nil {
		return Block{}, err
	}

	signature, err := sign_tt(tt_with_hash.Hash, _private_key)
	if err != nil {
		return Block{}, err
	}
	tt_with_hash.Hash_signature = signature

	return tt_with_hash, nil
}

// This function wraps the signing process for the trusted text blocks. It will call the
// hashing function, and then return a version of the input with a populated hash element,
// derived from the core content of the text element
func hash_tt(_existing_trustedtext Block) (Block, error) {
	content_hash, err := return_hash(_existing_trustedtext)
	if err != nil {
		return Block{}, err
	}
	_existing_trustedtext.Hash = content_hash
	return _existing_trustedtext, nil
}

// This function wraps the underlying hashing process, reducing it simply to
// block in - string out. This structure should be locked in early, since any
// change to it will almost certainly invalidate all the hashing chains
func return_hash(_trusted_text_element Block) (string, error) {
	elements := _trusted_text_element.Author +
		_trusted_text_element.Body.Instruction_type +
		_trusted_text_element.Body.Instruction

	hasher := sha1.New()

	_, err := hasher.Write([]byte(elements))

	if err != nil {
		return "", err
	}

	bytestring_hash := hasher.Sum(nil)
	return hex.EncodeToString(bytestring_hash), nil
}

// generate_tt_body is the instantiator function for a `tt_body` object which contains the structured form of
// a trustedtext message. The intantiator has paired validation with the instruction processor, so that only the
// kinds of instruction there is a valid processor for can be created.
func generate_tt_body(_instruction_type string, _instruction_body string) (body, error) {
	if _instruction_type != "publish" && _instruction_type != "head_change" {
		return body{}, errors.New("invalid instruction type, cannot generate an instruction for: " + _instruction_type)
	}
	new_instruction := body{
		Instruction_type: _instruction_type,
		Instruction:      _instruction_body,
	}

	return new_instruction, nil
}

// Generate_block is the wrapper for the two 'new' instance generators needed in the creation of
// a new block of trustedtext.
func Generate_block(_instruction_type string, _instruction_body string, _public_key string, _private_key string) (Block, error) {
	new_instruction, err := generate_tt_body(_instruction_type, _instruction_body)
	if err != nil {
		return Block{}, err
	}

	new_block, err := instantiate(_public_key, new_instruction, _private_key)
	if err != nil {
		return Block{}, err
	}

	return new_block, nil

}
