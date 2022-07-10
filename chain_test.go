package main

import "testing"



func generate_additonal_test_block(_existing_chain trustedtext_chain_s) trustedtext_s {
	new_block, _ := Instantiate("Dexter", "Intruder alert, DeeDee in the lab. Again!", junk_pri_key)
	return new_block
}


func generate_standard_test_chain(_init_only bool) trustedtext_chain_s {
	test_ttc, _ := Genesis(
		"Dexter",
		[]string{"lab"},
		junk_pri_key,
	)
	if _init_only {
		return test_ttc
	}



	new_block, _ := Instantiate("Dexter", "Intruder alert, DeeDee in the lab", junk_pri_key)
	
	test_ttc, _ = Amend(
		test_ttc,
		new_block,
	)
	return test_ttc
}

const second_standard_message = "18b1bf12c37a2146fc1025f91dced3728960cd70"

const first_standard_message = "4e8a9fbbb44f2756c61681626ac6bcc65e620d31"

// func Test_printer(t *testing.T) {
// 	t.Log("Standard test chain is ", generate_standard_test_chain(false))
// 	t.Fail()
// }

func Test_genesis_validation(t *testing.T) {
	var err error
	_, err = Genesis(
		"",
		[]string{"lab"},
		junk_pri_key,
	)
	if err == nil {
		t.Log("Genesis doesnt reject chains with no author")
		t.Fail()
	}

	_, err = Genesis("Dexter", []string{}, junk_pri_key)

	if err != nil {
		t.Log("Genesis inappropriately rejects chains with no tags")
		t.Fail()
	}
}

func Test_basic_amend(t *testing.T) {
	lab_chain_1  := generate_standard_test_chain(false)


	new_block := generate_additonal_test_block(lab_chain_1)
	_, err := Amend(lab_chain_1, new_block)

	if err != nil {
		t.Log("Amend fails on valid input")
		t.Fail()
	}
}

func Test_amend_functionality(t *testing.T) {
	lab_chain_1 := generate_standard_test_chain(false)
	existing_head_hash := lab_chain_1.head_hash
	existing_chain_length := len(lab_chain_1.tt_chain)

	new_block := generate_additonal_test_block(lab_chain_1)	

	lab_chain_2, _ := Amend(lab_chain_1, new_block)

	if lab_chain_2.head_hash != existing_head_hash {
		t.Log("Amend interferes with head_hash")
		t.Fail()
	}

	if len(lab_chain_2.tt_chain) - existing_chain_length  != 1 {
		t.Log("Amend increments chain length inappropriately")
		t.Fail()
	}
}


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

func Test_return_head_hash_functionality(t *testing.T) {
	lab_chain_1  := generate_standard_test_chain(false)
	head_block, err := Return_head_block(lab_chain_1)
	if err != nil {
		t.Log("Head block doesn't return appropriately")
		t.Fail()
	}
	if head_block.body != lab_chain_1.tt_chain[first_standard_message].body {
		t.Log("Head block doesn't return appropriately")
		t.Fail()
	}

	lab_chain_1, _ = Move_head_hash(lab_chain_1, lab_chain_1.tt_chain[second_standard_message].hash)
	new_head_block, err := Return_head_block(lab_chain_1)
	if err != nil {
		t.Log("Head block doesn't return properly after moving")
		t.Fail()
	}
	if new_head_block.body != lab_chain_1.tt_chain[second_standard_message].body {
		t.Log("Head block doesn't return properly after moving")
		t.Fail()
	}
	

}