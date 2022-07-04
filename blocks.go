package main

type trustedtext_s struct {
	author string
	tags   []string
	body   string
}

func Instantiate(_author string, _tags []string, _body string) trustedtext_s {
	return trustedtext_s{author: _author, tags: _tags, body: _body}
}

func Edit(_existing_element trustedtext_s, _new_author string, _new_body string) trustedtext_s {
	return trustedtext_s{author: _new_author, tags: _existing_element.tags, body: _new_body}
}
