package main

import (
	//"encoding/json"
	"testing"
)

func generate_standard_test_block() (trustedtext_s, error) {
	dexters_instruction_1 := tt_body{
		Instruction_type: "publish",
		Instruction:      "DeeDee Better not interfere with this one",
	}
	return instantiate(junk_pub_key, dexters_instruction_1, junk_pri_key)
}

// func Test_return_standard_block(t *testing.T) {
// 	test_block, _ := generate_standard_test_block()
// 	text_block, _ := json.Marshal(test_block)
// 	t.Log(string(text_block))
// 	t.Fail()
// }

func Test_Basic_instantiation_works(t *testing.T) {
	_, err := generate_standard_test_block()
	if err != nil {
		t.Log("Erroring on valid instantiation input", "Error:", err)
		t.Fail()
	}
}

func Test_Signed_instantiation(t *testing.T) {
	lab_book_1, _ := generate_standard_test_block()

	if lab_book_1.Hash == "" {
		t.Log("Blocks should be instantiated with a hash")
		t.Fail()
	}
}

func Test_Instantiate_input_validation(t *testing.T) {
	var err error
	dexters_instruction_1 := tt_body{
		Instruction_type: "publish",
		Instruction:      "DeeDee Better not interfere with this one",
	}
	_, err = instantiate(junk_pub_key, dexters_instruction_1, junk_pri_key)
	if err != nil {
		t.Log("Erroring on valid instantiation input", "Error:", err)
		t.Fail()
	}

	_, err = instantiate(junk_pub_key, tt_body{}, junk_pri_key)
	if err == nil {
		t.Log("Failing to prevent invalid block creation")
		t.Fail()
	}

	_, err = instantiate("", dexters_instruction_1, junk_pri_key)
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

	if len(signed_book_1.Hash) == 0 {
		t.Log("signing doesn't generate hash on block")
		t.Fail()
	}

}

func Test_that_all_authors_are_valid_pub_keys(t *testing.T) {
	var err error
	dexters_instruction_1 := tt_body{
		Instruction_type: "publish",
		Instruction:      "DeeDee Better not interfere with this one",
	}
	_, err = instantiate(junk_pub_key, dexters_instruction_1, junk_pri_key)
	if err != nil {
		t.Log("Instantiate block fails on valid pair", "Error:", err)
		t.Fail()
	}

	_, err = instantiate("Junk string", dexters_instruction_1, junk_pri_key)
	if err == nil {
		t.Log("Instantiate fails to block an invalid pair")
		t.Fail()
	}
}

func Test_that_subsequent_hashing_works(t *testing.T) {
	tb, _ := generate_standard_test_block()
	initial_hash := tb.Hash
	new_hash, _ := return_hash(tb)
	if initial_hash != new_hash {
		t.Log("Rehashing the same block produces changing results")
	}
}
