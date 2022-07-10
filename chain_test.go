package main

import "testing"



func generate_additonal_test_block(_existing_chain trustedtext_chain_s) trustedtext_s {
	dexters_instruction_2 := tt_body{
		instruction_type: "publish",
		instruction: "Intruder alert, DeeDee in the lab. Again!",
	}
	new_block, _ := Instantiate(junk_pub_key, dexters_instruction_2, junk_pri_key)
	return new_block
}


func generate_standard_test_chain(_init_only bool) trustedtext_chain_s {
	dexters_instruction_1 := tt_body{
		instruction_type: "publish",
		instruction: "Intruder alert, DeeDee in the lab",
	}
	test_ttc, _ := Genesis(
		junk_pub_key,
		[]string{"lab"},
		junk_pri_key,
	)
	if _init_only {
		return test_ttc
	}



	new_block, _ := Instantiate(junk_pub_key, dexters_instruction_1, junk_pri_key)
	
	test_ttc, _ = Amend(
		test_ttc,
		new_block,
	)
	return test_ttc
}

const first_standard_message = "b83030a13322e34fe61ef7dfe6d4750cab4d7429"
const second_standard_message = "f655762bf9c727eb04a71072b26e23c13b7d765c"


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

	_, err = Genesis(junk_pub_key, []string{}, junk_pri_key)

	if err != nil {
		t.Log("Genesis inappropriately rejects chains with no tags", "Error:", err)
		t.Fail()
	}
}

func Test_basic_amend(t *testing.T) {
	lab_chain_1  := generate_standard_test_chain(false)


	new_block := generate_additonal_test_block(lab_chain_1)
	_, err := Amend(lab_chain_1, new_block)

	if err != nil {
		t.Log("Amend fails on valid input", "Error:", err)
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
		t.Log( "Initial chain length is:", existing_chain_length, "new chain length is:", len(lab_chain_2.tt_chain))
		t.Fail()
	}
}


func Test_return_head_hash_functionality(t *testing.T) {
	lab_chain_1  := generate_standard_test_chain(false)
	head_block, err := Return_head_block(lab_chain_1)
	if err != nil {
		t.Log("Head block doesn't return appropriately", "Error:", err)
		t.Fail()
	}
	if head_block.body != lab_chain_1.tt_chain[first_standard_message].body {
		t.Log("Head block doesn't return appropriately")
		t.Fail()
	}

	lab_chain_1, _ = Move_head_hash(lab_chain_1, lab_chain_1.tt_chain[second_standard_message].hash)
	new_head_block, err := Return_head_block(lab_chain_1)
	if err != nil {
		t.Log("Head block doesn't return properly after moving", "Error:", err)
		t.Fail()
	}
	if new_head_block.body != lab_chain_1.tt_chain[second_standard_message].body {
		t.Log("Head block doesn't return properly after moving")
		t.Fail()
	}
	

}