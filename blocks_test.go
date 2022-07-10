package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"
)
const junk_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const junk_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"

func generate_standard_test_block() (trustedtext_s, error) {
	return Instantiate(junk_pub_key, "DeeDee Better not interfere with this one", junk_pri_key)
}

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

func Test_Basic_instantiation_works(t *testing.T) {
	_, err := generate_standard_test_block()
	if err != nil {
		t.Log("Erroring on valid instantiation input", "Error:", err)
		t.Fail()
	}
}

func Test_Signed_instantiation(t *testing.T) {
	lab_book_1, _ := generate_standard_test_block()

	if lab_book_1.hash == "" {
		t.Log("Blocks should be instantiated with a hash")
		t.Fail()
	}
}

func Test_Instantiate_input_validation(t *testing.T) {
	var err error

	_, err = Instantiate(junk_pub_key, "DeeDee Better not interfere with this one", junk_pri_key)
	if err != nil {
		t.Log("Erroring on valid instantiation input", "Error:", err)
		t.Fail()
	}

	_, err = Instantiate(junk_pub_key, "", junk_pri_key)
	if err == nil {
		t.Log("Failing to prevent invalid block creation")
		t.Fail()
	}

	_, err = Instantiate("",  "DeeDee Better not interfere with this one", junk_pri_key)
	if err == nil {
		t.Log("Failing to prevent invalid block creation")
		t.Fail()
	}
}

func Test_Signing_adds_hash(t *testing.T) {
	lab_book_1, _ := generate_standard_test_block()

	signed_book_1, err := sign_tt(lab_book_1) 

	if err != nil {
		t.Log("signing fails", "Error:", err)
		t.Fail()
	}

	if len(signed_book_1.hash) == 0 {
		t.Log("signing doesn't generate hash on block")
		t.Fail()
	}


}