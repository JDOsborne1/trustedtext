package main

import "testing"

func Test_fork_keeps_core_tree(t *testing.T) {
	lab_chain_2 := generate_standard_test_chain(false)

	lab_chain_2, _ = Move_head_hash(lab_chain_2, lab_chain_2.tt_chain[second_standard_message].Hash)

	lab_chain_2_forked := Fork_chain_essentials(lab_chain_2)

	if len(lab_chain_2_forked.tt_chain) != 2 {
		t.Log("forked chain doesn't carry both head entries")
		t.Fail()
	}
	deedees_instruction := tt_body{
		instruction_type: "publish",
		instruction:      "Haha, look at all the fun dials to turn",
	}
	new_block, _ := Instantiate("DeeDee", deedees_instruction, junk_pri_key)
	lab_chain_3, _ := Amend(lab_chain_2, new_block)

	lab_chain_3_forked := Fork_chain_essentials(lab_chain_3)

	if len(lab_chain_3_forked.tt_chain) != 2 {
		t.Log("Forking doesn't drop unheaded blocks")
		t.Fail()
	}

}
