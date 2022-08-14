package trustedtext

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
)

func verify_hex_encoded_values(_author_public_key_hex string, _hash_of_message_body string, _hash_signature_hex string) (bool, error) {
	signature_of_new_message, err := hex.DecodeString(_hash_signature_hex)

	if err != nil {
		return false, err
	}

	decoded_original_author, err := hex.DecodeString(_author_public_key_hex)
	if err != nil {
		return false, err
	}

	decoded_message_hash, err := hex.DecodeString(_hash_of_message_body)
	if err != nil {
		return false, err
	}

	return ed25519.Verify(decoded_original_author, decoded_message_hash, signature_of_new_message), nil

}

func sign_tt(_hash_of_message_body string, _private_key string) (string, error) {
	decoded_key, err := hex.DecodeString(_private_key)
	if err != nil {
		return "", err
	}
	decoded_hash, err := hex.DecodeString(_hash_of_message_body)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(ed25519.Sign(decoded_key, decoded_hash)), nil

}

func encoded_key_pair_is_valid(_encoded_public_key string, _encoded_private_key string) (bool, error) {
	decoded_pri_key, err := hex.DecodeString(_encoded_private_key)
	if err != nil {
		return false, err
	}

	publicKey := make([]byte, 32)
	copy(publicKey, decoded_pri_key[32:])

	encoded_regenerated_pub_key := hex.EncodeToString(publicKey)

	keys_match := encoded_regenerated_pub_key == _encoded_public_key

	return keys_match, nil
}

func verify_block_is_valid(_input_block Trustedtext_s) (bool, error) {
	rehash_of_body, err := return_hash(_input_block)
	if err != nil {
		return false, err
	}
	if rehash_of_body != _input_block.Hash {
		return false, errors.New("body content doesn't match body hash")
	}
	signature_is_valid, err := verify_hex_encoded_values(_input_block.Author, _input_block.Hash, _input_block.Hash_signature)
	if err != nil {
		return false, err
	}
	if !signature_is_valid {
		return false, errors.New("hash signature not verified")
	}
	return true, nil
}
