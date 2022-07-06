package main

import "testing"

func generate_standard_test_chain(_init_only bool) (trustedtext_chain_s, error) {
	test_ttc, err := Genesis(
		"Dexter",
		[]string{"lab"},
	)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	if _init_only {
		return test_ttc, nil
	}

	test_ttc, err = Amend(
		test_ttc,
		"Dexter",
		"DeeDee Better not interfere with this one",
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
	)
	if err == nil {
		t.Log("Genesis doesnt reject chains with no author")
		t.Fail()
	}

	_, err = Genesis("Dexter", []string{})

	if err != nil {
		t.Log("Genesis inappropriately rejects chains with no tags")
		t.Fail()
	}
}

func Test_basic_amend(t *testing.T) {
	lab_chain_1, _ := generate_standard_test_chain(false)

	_, err := Amend(lab_chain_1, "Dexter", "Intruder alert, DeeDee in the lab")

	if err != nil {
		t.Log("Amend fails on valid input")
		t.Fail()
	}
}

func Test_amend_functionality(t *testing.T) {
	lab_chain_1, _ := generate_standard_test_chain(false)
	existing_head_hash := lab_chain_1.head_hash
	existing_chain_length := len(lab_chain_1.tt_chain)

	lab_chain_2, _ := Amend(lab_chain_1, "Dexter", "Intruder alert, DeeDee in the lab")

	if lab_chain_2.head_hash != existing_head_hash {
		t.Log("Amend interferes with head_hash")
		t.Fail()
	}

	if len(lab_chain_2.tt_chain) - existing_chain_length  != 1 {
		t.Log("Amend increments chain length inappropriately")
		t.Fail()
	}
}

func Test_recent_hash_functionality(t *testing.T) {
	lab_chain_1, _ := generate_standard_test_chain(false)
	first_last_hash := Most_recent_hash(lab_chain_1)

	lab_chain_2, _ := Amend(lab_chain_1, "Dexter", "Intruder alert, DeeDee in the lab")

	next_last_hash := Most_recent_hash(lab_chain_2)

	if first_last_hash == next_last_hash {
		t.Log("Last Hashes not moving along with amendments")
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

	_, err = Move_head_hash(lab_chain_2, lab_chain_2.tt_chain[1].hash)
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