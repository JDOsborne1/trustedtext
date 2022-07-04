package main

import "fmt"

type tt_interface interface {
	Instantiate(_author string, _keys []string, _body string) tt_interface
	Edit(_existing_element tt_interface, _new_body string )
}

type tt_struct struct {
	author string
	keys []string
	body string
}

func Instantiate(_author string, _keys []string, _body string) tt_struct {
	return tt_struct{author: _author, keys: _keys, body: _body}
}

func main() {
	test_tt := Instantiate(
		"Johnny Bravo",
		[]string{"diary", "test"},
		"This is the best and worst day ever, and I'm glad I can record it in a trusted way",
	)

	fmt.Println(test_tt)
}