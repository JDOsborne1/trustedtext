package main

import "fmt"

type TT_element interface {

}

type tt_element struct {
	author string
	keys []string
	body string
}

func Instantiate_tt(_author string, _keys []string, _body string) tt_element {
	return tt_element{author: _author, keys: _keys, body: _body}
}

func main() {
	test_tt := Instantiate_tt(
		"Johnny Bravo",
		[]string{"diary", "test"},
		"This is the best and worst day ever, and I'm glad I can record it in a trusted way",
	)

	fmt.Println(test_tt)
}