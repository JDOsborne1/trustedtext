package main

import "testing"

func Test_head_hash_functionality(t *testing.T) {
	lab_chain_1  := generate_standard_test_chain(true)
	if len(lab_chain_1.head_hash) == 0 {
		t.Log("chain not instantiated with a head hash")
		t.Fail()
	}

	lab_chain_2  := generate_standard_test_chain(false)
	var err error

	_, err = Move_head_hash(lab_chain_2, lab_chain_2.tt_chain[second_standard_message].hash)
	if err != nil {
		t.Log("Fails to accept a valid hash to change to")
		t.Fail()
	}

	_, err = Move_head_hash(lab_chain_2, "randomstring")
	if err == nil {
		t.Log("Fails to reject an invalid hash")
		t.Fail()
	}

}

func Test_head_hash_history(t *testing.T) {
	lab_chain_1 := generate_standard_test_chain(false)
	if !lab_chain_1.head_hash_tree[first_standard_message] {
		t.Log("Genesis block not in head hash tree")
		t.Fail()
	}
	lab_chain_1, _ = Move_head_hash(lab_chain_1, lab_chain_1.tt_chain[second_standard_message].hash)
	if !lab_chain_1.head_hash_tree[second_standard_message] {
		t.Log("Subsequent head hashes not added to head hash tree")
		t.Fail()
	}
	
}