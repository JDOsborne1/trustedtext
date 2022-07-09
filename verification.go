package main

import (
	"crypto/ed25519"
	"encoding/hex"
)

func Verify_hex_encoded_values(_author_public_key_hex string, _message_body string, _hash_signature_hex string) (bool, error) {
	signature_of_new_message, err := hex.DecodeString(_hash_signature_hex)
	
	if err != nil {
		return false, err
	}
	
	decoded_original_author, err := hex.DecodeString(_author_public_key_hex)
	if err != nil {
		return false, err
	}
	
	return ed25519.Verify(decoded_original_author, []byte(_message_body), signature_of_new_message), nil

}