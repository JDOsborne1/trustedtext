package main

import (
	"testing"
)

func Test_fork_keeps_core_tree(t *testing.T) {
	lab_chain_2 := generate_standard_test_chain(false)

	lab_chain_2, _ = Move_head_hash(lab_chain_2, lab_chain_2.tt_chain[second_standard_message].Hash)

	lab_chain_2_forked := Fork_chain_essentials(lab_chain_2)

	if len(lab_chain_2_forked.tt_chain) != 2 {
		t.Log("forked chain doesn't carry both head entries")
		t.Fail()
	}
	deedees_instruction := tt_body{
		Instruction_type: "publish",
		Instruction:      "Haha, look at all the fun dials to turn",
	}
	new_block, _ := Instantiate("DeeDee", deedees_instruction, junk_pri_key)
	lab_chain_3, _ := Amend(lab_chain_2, new_block)

	lab_chain_3_forked := Fork_chain_essentials(lab_chain_3)

	if len(lab_chain_3_forked.tt_chain) != 2 {
		t.Log("Forking doesn't drop unheaded blocks")
		t.Fail()
	}

}

func Test_that_hash_membership_checks(t *testing.T) {
	lab_chain_2 := generate_standard_test_chain(false)

	hash_1_included_in_chain_2 := Is_hash_in_chain(lab_chain_2, "b83030a13322e34fe61ef7dfe6d4750cab4d7429")
	hash_2_included_in_chain_2 := Is_hash_in_chain(lab_chain_2, "f655762bf9c727eb04a71072b26e23c13b7d765c")
	
	if !(hash_1_included_in_chain_2 && hash_2_included_in_chain_2) {
		t.Log("Fails to report correctly on present hashes")
		t.Fail()
	}
	hash_3_included_in_chain_2 := Is_hash_in_chain(lab_chain_2, "d519546909952540c4fdaed62481ac6c8cef071e")

	if hash_3_included_in_chain_2 {
		t.Log("Fails to report correctly on missing hashes")
		t.Fail()
	}
	
	new_block := generate_additonal_test_block(lab_chain_2)
	lab_chain_3, _ := Process_incoming_block(lab_chain_2, new_block)
	hash_1_included_in_chain_3 := Is_hash_in_chain(lab_chain_3, "b83030a13322e34fe61ef7dfe6d4750cab4d7429")
	hash_2_included_in_chain_3 := Is_hash_in_chain(lab_chain_3, "f655762bf9c727eb04a71072b26e23c13b7d765c")
	hash_3_included_in_chain_3 := Is_hash_in_chain(lab_chain_3, "d519546909952540c4fdaed62481ac6c8cef071e")

	if !(hash_1_included_in_chain_3 && hash_2_included_in_chain_3 && hash_3_included_in_chain_3) {
		t.Log("Fails to report correctly on present hashes")
		t.Fail()
	}
}