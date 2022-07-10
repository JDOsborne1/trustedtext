package main

import (
	"testing"
)

func generate_standard_test_block() (trustedtext_s, error) {
	return Instantiate(junk_pub_key, "DeeDee Better not interfere with this one", junk_pri_key)
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

	_, err = Instantiate("", "DeeDee Better not interfere with this one", junk_pri_key)
	if err == nil {
		t.Log("Failing to prevent invalid block creation")
		t.Fail()
	}
}

func Test_Signing_adds_hash(t *testing.T) {
	lab_book_1, _ := generate_standard_test_block()

	signed_book_1, err := hash_tt(lab_book_1)

	if err != nil {
		t.Log("signing fails", "Error:", err)
		t.Fail()
	}

	if len(signed_book_1.hash) == 0 {
		t.Log("signing doesn't generate hash on block")
		t.Fail()
	}

}


func Test_that_all_authors_are_valid_pub_keys(t *testing.T) {
	var err error
	_, err = Instantiate(junk_pub_key, "TEST BODY", junk_pri_key)
	if err != nil {
		t.Log("Instantiate block fails on valid pair", "Error:", err)
		t.Fail()
	}

	_, err = Instantiate("Junk string", "TEST BODY", junk_pri_key)
	if err == nil {
		t.Log("Instantiate fails to block an invalid pair")
		t.Fail()
	}
}