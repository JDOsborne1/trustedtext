package main

import "testing"


func generate_standard_test_chain(_init_only bool) (trustedtext_chain_s, error) {
	test_ttc, err := Genesis(
		"Dexter",
		[]string{"lab"},
		junk_pri_key,
	)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	if _init_only {
		return test_ttc, nil
	}

	
	existing_block_tags := test_ttc.tt_chain[test_ttc.head_hash].tags

	new_block, _ := Instantiate("Dexter", existing_block_tags, "Intruder alert, DeeDee in the lab", junk_pri_key)
	
	test_ttc, err = Amend(
		test_ttc,
		new_block,
	)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	return test_ttc, nil
}


func Test_standard_generation(t *testing.T) {
	_, err := generate_standard_test_chain(false)

	if err != nil {
		t.Log("Fail on basic generation")
		t.Fail()
	}
}

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
	lab_chain_1, _ := generate_standard_test_chain(false)

	existing_block_tags := lab_chain_1.tt_chain[lab_chain_1.head_hash].tags

	new_block, _ := Instantiate("Dexter", existing_block_tags, "Intruder alert, DeeDee in the lab. Again!", junk_pri_key)
	_, err := Amend(lab_chain_1, new_block)

	if err != nil {
		t.Log("Amend fails on valid input")
		t.Fail()
	}
}

func Test_amend_functionality(t *testing.T) {
	lab_chain_1, _ := generate_standard_test_chain(false)
	existing_head_hash := lab_chain_1.head_hash
	existing_chain_length := len(lab_chain_1.tt_chain)
	existing_block_tags := lab_chain_1.tt_chain[lab_chain_1.head_hash].tags

	new_block, _ := Instantiate("Dexter", existing_block_tags, "Intruder alert, DeeDee in the lab. Again!", junk_pri_key)
	

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
	lab_chain_1, _ := generate_standard_test_chain(true)
	if len(lab_chain_1.head_hash) == 0 {
		t.Log("chain not instantiated with a head hash")
		t.Fail()
	}

	lab_chain_2, _ := generate_standard_test_chain(false)
	var err error

	_, err = Move_head_hash(lab_chain_2, lab_chain_2.tt_chain["59e326b089627118b9a01f7a7666c27d8ec67195"].hash)
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
	lab_chain_1, _ := generate_standard_test_chain(false)
	if !lab_chain_1.head_hash_tree["cc7ec24ab30220f7efdfdf2f2023301f6265d1f4"] {
		t.Log("Genesis block not in head hash tree")
		t.Fail()
	}
	lab_chain_1, _ = Move_head_hash(lab_chain_1, lab_chain_1.tt_chain["59e326b089627118b9a01f7a7666c27d8ec67195"].hash)
	if !lab_chain_1.head_hash_tree["59e326b089627118b9a01f7a7666c27d8ec67195"] {
		t.Log("Subsequent head hashes not added to head hash tree")
		t.Fail()
	}
	
}

func Test_return_head_hash_functionality(t *testing.T) {
	lab_chain_1, _ := generate_standard_test_chain(false)
	head_block, err := Return_head_block(lab_chain_1)
	if err != nil {
		t.Log("Head block doesn't return appropriately")
		t.Fail()
	}
	if head_block.body != lab_chain_1.tt_chain["cc7ec24ab30220f7efdfdf2f2023301f6265d1f4"].body {
		t.Log("Head block doesn't return appropriately")
		t.Fail()
	}

	lab_chain_1, _ = Move_head_hash(lab_chain_1, lab_chain_1.tt_chain["59e326b089627118b9a01f7a7666c27d8ec67195"].hash)
	new_head_block, err := Return_head_block(lab_chain_1)
	if err != nil {
		t.Log("Head block doesn't return properly after moving")
		t.Fail()
	}
	if new_head_block.body != lab_chain_1.tt_chain["59e326b089627118b9a01f7a7666c27d8ec67195"].body {
		t.Log("Head block doesn't return properly after moving")
		t.Fail()
	}
	

}