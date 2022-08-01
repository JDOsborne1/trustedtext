package main

import (
	"testing"

	"golang.org/x/exp/maps"
)

func generate_additonal_test_block(_existing_chain trustedtext_chain_s) trustedtext_s {
	dexters_instruction_2 := tt_body{
		Instruction_type: "publish",
		Instruction:      "Intruder alert, DeeDee in the lab. Again!",
	}
	new_block, _ := instantiate(junk_pub_key, dexters_instruction_2, junk_pri_key)
	return new_block
}

func generate_standard_test_chain(_init_only bool) trustedtext_chain_s {
	dexters_instruction_1 := tt_body{
		Instruction_type: "publish",
		Instruction:      "Intruder alert, DeeDee in the lab",
	}
	test_ttc, _ := genesis(
		junk_pub_key,
		[]string{"lab"},
		junk_pri_key,
	)
	if _init_only {
		return test_ttc
	}

	new_block, _ := instantiate(junk_pub_key, dexters_instruction_1, junk_pri_key)

	test_ttc, _ = amend(
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
	_, err = genesis(
		"",
		[]string{"lab"},
		junk_pri_key,
	)
	if err == nil {
		t.Log("Genesis doesnt reject chains with no author")
		t.Fail()
	}

	_, err = genesis(junk_pub_key, []string{}, junk_pri_key)

	if err != nil {
		t.Log("Genesis inappropriately rejects chains with no tags", "Error:", err)
		t.Fail()
	}
}

func Test_basic_amend(t *testing.T) {
	lab_chain_1 := generate_standard_test_chain(false)

	new_block := generate_additonal_test_block(lab_chain_1)
	_, err := amend(lab_chain_1, new_block)

	if err != nil {
		t.Log("Amend fails on valid input", "Error:", err)
		t.Fail()
	}
}

func Test_amend_functionality(t *testing.T) {
	lab_chain_1 := generate_standard_test_chain(false)
	existing_head_hash := lab_chain_1.Head_hash
	existing_chain_length := len(lab_chain_1.Tt_chain)

	new_block := generate_additonal_test_block(lab_chain_1)

	lab_chain_2, _ := amend(lab_chain_1, new_block)

	if lab_chain_2.Head_hash != existing_head_hash {
		t.Log("Amend interferes with head_hash")
		t.Fail()
	}

	if len(lab_chain_2.Tt_chain)-existing_chain_length != 1 {
		t.Log("Amend increments chain length inappropriately")
		t.Log("Initial chain length is:", existing_chain_length, "new chain length is:", len(lab_chain_2.Tt_chain))
		t.Fail()
	}
}

func Test_return_head_hash_functionality(t *testing.T) {
	lab_chain_1 := generate_standard_test_chain(false)
	head_block, err := return_head_block(lab_chain_1)
	if err != nil {
		t.Log("Head block doesn't return appropriately", "Error:", err)
		t.Fail()
	}
	if head_block.Body != lab_chain_1.Tt_chain[first_standard_message].Body {
		t.Log("Head block doesn't return appropriately")
		t.Fail()
	}

	lab_chain_1, _ = move_head_hash(lab_chain_1, lab_chain_1.Tt_chain[second_standard_message].Hash)
	new_head_block, err := return_head_block(lab_chain_1)
	if err != nil {
		t.Log("Head block doesn't return properly after moving", "Error:", err)
		t.Fail()
	}
	if new_head_block.Body != lab_chain_1.Tt_chain[second_standard_message].Body {
		t.Log("Head block doesn't return properly after moving")
		t.Fail()
	}

}

func Test_distribute_validation(t *testing.T) {
	lab_chain_1 := generate_standard_test_chain(false)

	new_block := generate_additonal_test_block(lab_chain_1)
	existing_hash := maps.Keys(lab_chain_1.Tt_chain)[1]
	existing_block := lab_chain_1.Tt_chain[existing_hash]

	var err error

	_, err = Process_incoming_block(lab_chain_1, existing_block)
	if err == nil {
		t.Log("Validation doesn't catch existing block")
		t.Fail()
	}

	_, err = Process_incoming_block(lab_chain_1, new_block)
	if err != nil {
		t.Log("Validation fails on valid blocks")
		t.Fail()
	}
}
