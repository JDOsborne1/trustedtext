package main

import "testing"

func Test_Signed_instantiation(t *testing.T) {
	lab_book_1, err := Instantiate("Dexter", []string{"Labs"}, "DeeDee Better not interfere with this one")
	if err != nil {
		t.Log("Erroring on valid instantiation input")
		t.Fail()
	}

	if lab_book_1.hash == "" {
		t.Log("Blocks should be instantiated with a hash")
		t.Fail()
	}
}

func Test_Instantiate_input_validation(t *testing.T) {
	var err error
	_, err = Instantiate("Dexter", []string{}, "DeeDee Better not interfere with this one")
	if err != nil {
		t.Log("Erroring on valid instantiation input")
		t.Fail()
	}

	_, err = Instantiate("Dexter", []string{}, "")
	if err == nil {
		t.Log("Failing to prevent invalid block creation")
		t.Fail()
	}
	_, err = Instantiate("", []string{}, "DeeDee Better not interfere with this one")
	if err == nil {
		t.Log("Failing to prevent invalid block creation")
		t.Fail()
	}
}