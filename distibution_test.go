package main

import "testing"

func Test_fork_keeps_core_tree(t *testing.T) {
	lab_chain_2, _ := generate_standard_test_chain(false)

	lab_chain_2, _ = Move_head_hash(lab_chain_2, lab_chain_2.tt_chain["03d797a80fb52073fbf599047c862c5e7890a960"].hash)

	lab_chain_2_forked := Fork_chain_essentials(lab_chain_2)

	if len(lab_chain_2_forked.tt_chain) != 2 {
		t.Log("forked chain doesn't carry both head entries")
		t.Fail()
	}

	lab_chain_3, _ := Amend(lab_chain_2, "DeeDee", "Haha, look at all the fun dials to turn")

	lab_chain_3_forked := Fork_chain_essentials(lab_chain_3)

	if len(lab_chain_3_forked.tt_chain) != 2 {
		t.Log("Forking doesn't drop unheaded blocks")
		t.Fail()
	}


}