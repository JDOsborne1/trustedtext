package main

import "fmt"

type Ttc_interface interface {
	Genesis(_author string, _tags []string) Ttc_interface
	Amend(_existing_ttc Ttc_interface, _author string, _body string) Ttc_interface
}

type ttc_struct struct {
	original_author string
	tt_chain []tt_struct
}

func Genesis(_author string, _tags []string) ttc_struct {
	first_element := Instantiate(
			_author,
			_tags,
			"This is the origin message of a trusted text chain",
		)
	
	new_chain :=  ttc_struct{
		original_author: _author,
		tt_chain: []tt_struct{first_element},
	}
	return new_chain
}

func Amend(_existing_ttc ttc_struct, _author string, _body string) ttc_struct {
	new_element := Instantiate(
		_author,
		_existing_ttc.tt_chain[0].tags,
		_body,
	)
	new_chain := append(_existing_ttc.tt_chain, new_element)
	_existing_ttc.tt_chain = new_chain
	return _existing_ttc
}

func main() {
	test_ttc := Genesis(
		"Johnny Bravo",
		[]string{"diary", "test"},
	)
	fmt.Println(test_ttc)
	test_ttc = Amend(
		test_ttc,
		"Johnny Bravo",
		"This is the best and worst day ever, and I'm glad I can record it in a trusted way",
	)
	fmt.Println(test_ttc)

	test_ttc = Amend(
		test_ttc, 
		"Johnny Bravo", 
		"Actually, now I've recorded it, it's simply the best day",
	)
	fmt.Println(test_ttc)

}