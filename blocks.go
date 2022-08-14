package trustedtext

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
)

type Trustedtext_s struct {
	Author                string
	Body                  tt_body
	Head_hash_at_creation string
	Hash                  string
	Hash_signature        string
}

type tt_body struct {
	Instruction_type string
	Instruction      string
}

// This function is called to generate a base instance of the trustedtext block, based
// on its arguments, and then sign it. A block can be valid without tags, as it remains
// globally unique to the original author, however it must have both an author and body.
func instantiate(_author string, _body tt_body, _private_key string) (Trustedtext_s, error) {
	if len(_author) == 0 {
		return Trustedtext_s{}, errors.New("cannot have a missing author")
	}

	if len(_body.Instruction_type) == 0 {
		return Trustedtext_s{}, errors.New("cannot have a missing instruction type")
	}
	if len(_body.Instruction) == 0 {
		return Trustedtext_s{}, errors.New("cannot have an empty instruction")
	}
	valid_signature_pairs, err := encoded_key_pair_is_valid(_author, _private_key)
	if err != nil {
		return Trustedtext_s{}, err
	}
	if !valid_signature_pairs {
		return Trustedtext_s{}, errors.New("author and key combination don't match")
	}

	tt_no_hash := Trustedtext_s{Author: _author, Body: _body}
	tt_with_hash, err := hash_tt(tt_no_hash)
	if err != nil {
		return Trustedtext_s{}, err
	}

	signature, err := sign_tt(tt_with_hash.Hash, _private_key)
	if err != nil {
		return Trustedtext_s{}, err
	}
	tt_with_hash.Hash_signature = signature

	return tt_with_hash, nil
}

// This function wraps the signing process for the trusted text blocks. It will call the
// hashing function, and then return a version of the input with a populated hash element,
// derived from the core content of the text element
func hash_tt(_existing_trustedtext Trustedtext_s) (Trustedtext_s, error) {
	content_hash, err := return_hash(_existing_trustedtext)
	if err != nil {
		return Trustedtext_s{}, err
	}
	_existing_trustedtext.Hash = content_hash
	return _existing_trustedtext, nil
}

// This function wraps the underlying hashing process, reducing it simply to
// block in - string out. This structure should be locked in early, since any
// change to it will almost certainly invalidate all the hashing chains
func return_hash(_trusted_text_element Trustedtext_s) (string, error) {
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

func generate_tt_body(_instruction_type string, _instruction_body string) (tt_body, error) {
	if _instruction_type != "publish" && _instruction_type != "head_change" {
		return tt_body{}, errors.New("invalid instruction type, cannot generate an instruction for: " + _instruction_type)
	}
	new_instruction := tt_body{
		Instruction_type: _instruction_type,
		Instruction:      _instruction_body,
	}

	return new_instruction, nil
}

func Generate_block(_instruction_type string, _instruction_body string, _public_key string, _private_key string) (Trustedtext_s, error) {
	new_instruction, err := generate_tt_body(_instruction_type, _instruction_body)
	if err != nil {
		return Trustedtext_s{}, err
	}

	new_block, err := instantiate(_public_key, new_instruction, _private_key)
	if err != nil {
		return Trustedtext_s{}, err
	}

	return new_block, nil

}
