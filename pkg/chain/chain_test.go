package trustedtext

import (
	"testing"

	"golang.org/x/exp/maps"
)

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
	lab_chain_1 := helper_generate_standard_test_chain(false)

	new_block := helper_generate_additonal_test_block(lab_chain_1)
	_, err := amend(lab_chain_1, new_block)

	if err != nil {
		t.Log("Amend fails on valid input", "Error:", err)
		t.Fail()
	}
}

func Test_amend_functionality(t *testing.T) {
	lab_chain_1 := helper_generate_standard_test_chain(false)
	existing_head_hash := lab_chain_1.Head_hash
	existing_chain_length := len(lab_chain_1.Tt_chain)

	new_block := helper_generate_additonal_test_block(lab_chain_1)

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
	lab_chain_1 := helper_generate_standard_test_chain(false)
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
	lab_chain_1 := helper_generate_standard_test_chain(false)

	new_block := helper_generate_additonal_test_block(lab_chain_1)
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

func Test_that_hash_membership_checks(t *testing.T) {
	lab_chain_2 := helper_generate_standard_test_chain(false)

	hash_1_included_in_chain_2 := is_hash_in_chain(lab_chain_2, "b83030a13322e34fe61ef7dfe6d4750cab4d7429")
	hash_2_included_in_chain_2 := is_hash_in_chain(lab_chain_2, "f655762bf9c727eb04a71072b26e23c13b7d765c")

	if !(hash_1_included_in_chain_2 && hash_2_included_in_chain_2) {
		t.Log("Fails to report correctly on present hashes")
		t.Fail()
	}
	hash_3_included_in_chain_2 := is_hash_in_chain(lab_chain_2, "d519546909952540c4fdaed62481ac6c8cef071e")

	if hash_3_included_in_chain_2 {
		t.Log("Fails to report correctly on missing hashes")
		t.Fail()
	}

	new_block := helper_generate_additonal_test_block(lab_chain_2)
	lab_chain_3, _ := Process_incoming_block(lab_chain_2, new_block)
	hash_1_included_in_chain_3 := is_hash_in_chain(lab_chain_3, "b83030a13322e34fe61ef7dfe6d4750cab4d7429")
	hash_2_included_in_chain_3 := is_hash_in_chain(lab_chain_3, "f655762bf9c727eb04a71072b26e23c13b7d765c")
	hash_3_included_in_chain_3 := is_hash_in_chain(lab_chain_3, "d519546909952540c4fdaed62481ac6c8cef071e")

	if !(hash_1_included_in_chain_3 && hash_2_included_in_chain_3 && hash_3_included_in_chain_3) {
		t.Log("Fails to report correctly on present hashes")
		t.Fail()
	}
}
