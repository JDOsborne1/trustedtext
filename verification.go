package main

import (
	"crypto/ed25519"
	"encoding/hex"
)

func Verify_hex_encoded_values(_author_public_key_hex string, _hash_of_message_body string, _hash_signature_hex string) (bool, error) {
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