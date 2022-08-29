package trustedtext

import (
	"testing"
)

func Test_head_hash_functionality(t *testing.T) {
	lab_chain_1 := generate_standard_test_chain(true)
	if len(lab_chain_1.Head_hash) == 0 {
		t.Log("chain not instantiated with a head hash")
		t.Fail()
	}

	lab_chain_2 := generate_standard_test_chain(false)
	var err error

	_, err = move_head_hash(lab_chain_2, lab_chain_2.Tt_chain[second_standard_message].Hash)
	if err != nil {
		t.Log("Fails to accept a valid hash to change to")
		t.Fail()
	}

	_, err = move_head_hash(lab_chain_2, "randomstring")
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
	lab_chain_1, _ = move_head_hash(lab_chain_1, lab_chain_1.Tt_chain[second_standard_message].Hash)
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
	serialised_instruction, err := serialise_head_change(test_instruction)
	if err != nil {
		t.Log("Fails to serialise valid input", "Error:", err)
		t.Fail()
	}

	deserialised_instruction, err := deserialise_head_change(serialised_instruction)
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

func Test_move_head_block_core_functions(t *testing.T) {

	lab_chain_2 := generate_standard_test_chain(false)
	var err error

	head_move_block, err := Generate_head_move_block(junk_pub_key, second_standard_message, junk_pri_key)
	if err != nil {
		t.Log("Cannot generate a head move block", "Error:", err)
		t.Fail()
	}

	amended_chain, err := action_head_move_block(lab_chain_2, head_move_block)
	if err != nil {
		t.Log("Cannot action the head move block", "Error:", err)
		t.Fail()
	}
	amended_chain, err = amend(amended_chain, head_move_block)

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


func Test_generating_head_move_blocks(t *testing.T) {
	var err error

	_, err = Generate_head_move_block(
		"faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a",
		"newh_hash",
		"366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a",
	)
	if err != nil {
		t.Log("cannot generate head move block")
		t.Fail()
	}


	new_keypair, _ := helper_generate_key_pair()

	_, err = Generate_head_move_block(
		new_keypair.pub_key,
		"024c74fed7eaf14ffbb71fba7b2423d1d868b550",
		new_keypair.pri_key, 
	)

	if err != nil {
		t.Log("cannot generate head move block", err)
		t.Fail()
	}


	_, err =  Generate_head_move_block(
		junk_pub_key,
		second_standard_message,
		junk_pri_key,
	)

	if err != nil {
		t.Log("fails to generate viable head move block")
		t.Fail()
	}
}



func Test_amending_head_hash_using_processor(t *testing.T) {
	var err error
	lab_chain_1 := generate_standard_test_chain(false)

	head_move_block_1, _ := Generate_head_move_block(
		junk_pub_key,
		"newh_hash",
		junk_pri_key,
	)
	
	_, err = Process_incoming_block(
		lab_chain_1,
		head_move_block_1,
	)

	if err == nil {
		t.Log("Doesn't fail when trying to move to a hash which doesn't exist", err)
		t.Fail()
	}



	new_keypair, _ := helper_generate_key_pair()

	head_move_block_2, _ := Generate_head_move_block(
		new_keypair.pub_key,
		"024c74fed7eaf14ffbb71fba7b2423d1d868b550",
		new_keypair.pri_key, 
	)

	_, err = Process_incoming_block(
		lab_chain_1,
		head_move_block_2,
	)

	if err == nil {
		t.Log("fails to reject head changes which aren't by the original author")
		t.Fail()
	}


	head_move_block_3, _ :=  Generate_head_move_block(
		junk_pub_key,
		second_standard_message,
		junk_pri_key,
	)


	new_chain, err := Process_incoming_block(
		lab_chain_1,
		head_move_block_3,
	)

	if err != nil {
		t.Log("fails to process valid head chain block with error", err)
		t.Fail()
	}

	if new_chain.Head_hash != second_standard_message {
		t.Log("head hash change process fails to actually change head hash")
		t.Fail()
	}

	first_head_still_in_tree := new_chain.Head_hash_tree[first_standard_message]

	if !first_head_still_in_tree {
		t.Log("fails to retain record of previous head hashes")
		t.Fail()
	}

	
}