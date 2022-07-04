package main

type tt_struct struct {
	author string
	tags []string
	body string
}

func Instantiate(_author string, _tags []string, _body string) tt_struct {
	return tt_struct{author: _author, tags: _tags, body: _body}
}

func Edit(_existing_element tt_struct, _new_author string, _new_body string) tt_struct {
	return tt_struct{author: _new_author, tags: _existing_element.tags, body: _new_body}
}