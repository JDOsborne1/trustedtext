package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"
)

const junk_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const junk_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"

func Test_key_pair(t *testing.T) {
	message := []byte("junk message")
	var err error
	decoded_pri_key, err := hex.DecodeString(junk_pri_key)
	if err != nil {
		t.Log("Cannot decode test primary key", "Error:", err)
		t.Fail()
	}
	decoded_pub_key, err := hex.DecodeString(junk_pub_key)
	if err != nil {
		t.Log("Cannot decode test public key", "Error:", err)
		t.Fail()
	}
	signature := ed25519.Sign(decoded_pri_key, message)
	valid := ed25519.Verify(decoded_pub_key, message, signature)
	if !valid {
		t.Log("Key pair doesn't generate a valid signature")
		t.Fail()
	}

}

func Test_that_hashes_can_be_validated(t *testing.T) {
	test_block, _ := generate_standard_test_block()

	valid_signature, err := verify_hex_encoded_values(test_block.Author, test_block.Hash, test_block.Hash_signature)
	if err != nil {
		t.Log("Cannot verify valid hex encoded values", "Error:", err)
		t.Fail()
	}
	if !valid_signature {
		t.Log("true pairs aren't verifiable")
		t.Fail()
	}
}

func Test_that_pairs_are_properly_checked(t *testing.T) {
	var is_valid bool
	var err error

	is_valid, err = encoded_key_pair_is_valid(junk_pub_key, junk_pri_key)
	if err != nil {
		t.Log("Error in validating legal inputs", "Error:", err)
		t.Fail()
	}
	if !is_valid {
		t.Log("Validation rejects true pairs")
		t.Fail()
	}

	is_valid, err = encoded_key_pair_is_valid("asdasd", junk_pri_key)
	if err != nil {
		t.Log("Error in validating legal, but incorrect inputs", "Error:", err)
		t.Fail()
	}
	if is_valid {
		t.Log("Validation allows false pairs")
		t.Fail()
	}

}

func Test_that_validate_function_works_basic(t *testing.T) {
	tb, _ := generate_standard_test_block()

	valid, err := verify_block_is_valid(tb)
	if err != nil {
		t.Log("Cannot run verify function on standard block", "Error:", err)
		t.Fail()
	}
	if !valid {
		t.Log("Validate function fails valid input")
	}
}
