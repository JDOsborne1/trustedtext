package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"
)

const junk_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const junk_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"

func Test_key_pair(t *testing.T) {
	message :=  []byte("junk message")
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
	decoded_pub_key, _ := hex.DecodeString(test_block.author)
	decoded_hash_sig, _ := hex.DecodeString(test_block.hash_signature)
	decoded_hash, _ := hex.DecodeString(test_block.hash)
	valid_signature := ed25519.Verify(decoded_pub_key, decoded_hash, decoded_hash_sig)
	if !valid_signature {
		t.Log("true pairs aren't verifiable")
		t.Fail()
	}
}
