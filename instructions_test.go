package main

import "testing"

func Test_head_hash_functionality(t *testing.T) {
	lab_chain_1 := generate_standard_test_chain(true)
	if len(lab_chain_1.Head_hash) == 0 {
		t.Log("chain not instantiated with a head hash")
		t.Fail()
	}

	lab_chain_2 := generate_standard_test_chain(false)
	var err error

	_, err = Move_head_hash(lab_chain_2, lab_chain_2.Tt_chain[second_standard_message].Hash)
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
	if !lab_chain_1.Head_hash_tree[first_standard_message] {
		t.Log("Genesis block not in head hash tree")
		t.Fail()
	}
	lab_chain_1, _ = Move_head_hash(lab_chain_1, lab_chain_1.Tt_chain[second_standard_message].Hash)
	if !lab_chain_1.Head_hash_tree[second_standard_message] {
		t.Log("Subsequent head hashes not added to head hash tree")
		t.Fail()
	}

}

func Test_serialise_deserialise(t *testing.T) {
	test_instruction := head_change_instruction{
		New_head: "blah",
	}
	var err error
	serialised_instruction, err := Serialise_head_change(test_instruction)
	if err != nil {
		t.Log("Fails to serialise valid input", "Error:", err)
		t.Fail()
	}

	deserialised_instruction, err := Deserialise_head_change(serialised_instruction)
	if err != nil {
		t.Log("Failed to deserialise valid input", "Error:", err)
		t.Fail()
	}

	identical_after_serialisation_loop := deserialised_instruction == test_instruction

	if !identical_after_serialisation_loop {
		t.Log("Serialisation loop doesn't produce equivalent values", "Error:", err)
		t.Fail()
	}

}

func Test_head_move_block(t *testing.T) {

	lab_chain_2 := generate_standard_test_chain(false)
	var err error

	head_move_block, err := Generate_head_move_block(junk_pub_key, second_standard_message, junk_pri_key)
	if err != nil {
		t.Log("Cannot generate a head move block", "Error:", err)
		t.Fail()
	}

	amended_chain, err := Action_head_move_block(lab_chain_2, head_move_block)
	if err != nil {
		t.Log("Cannot action the head move block", "Error:", err)
		t.Fail()
	}
	amended_chain, err = Amend(amended_chain, head_move_block)

	if err != nil {
		t.Log("Cannot generate an amended chain", "Error:", err)
		t.Fail()
	}

	if amended_chain.Head_hash != second_standard_message {
		t.Log("Head hash doesn't change after instruction block")
		t.Fail()
	}

	second_message_in_head_tree := amended_chain.Head_hash_tree[second_standard_message]

	if !second_message_in_head_tree {
		t.Log("New head hash not included in head tree", "Error:", err)
		t.Fail()
	}

}
