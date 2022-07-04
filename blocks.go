package main

import (
	"crypto/sha1"
)

type trustedtext_s struct {
	author string
	tags   []string
	body   string
	previous_hash string
	hash   string
}

func Instantiate(_author string, _tags []string, _body string) trustedtext_s {
	unsigned_tt := trustedtext_s{author: _author, tags: _tags, body: _body}
	return sign_tt(unsigned_tt)
}

func sign_tt(_existing_trustedtext trustedtext_s) trustedtext_s {
	content_hash := return_hash(_existing_trustedtext)
	_existing_trustedtext.hash = content_hash
	return _existing_trustedtext
}

func collapse_tags(_list_of_tags []string) string {
	tag_list := ""
	for _, tag := range _list_of_tags {
		tag_list = tag_list + tag + ","
	}
	return tag_list
}

func return_hash(_trusted_text_element trustedtext_s) string {
	elements := _trusted_text_element.author +
		_trusted_text_element.body +
		_trusted_text_element.previous_hash +
		collapse_tags(_trusted_text_element.tags)

	hasher := sha1.New()

	hasher.Write([]byte(elements))

	bytestring_hash := hasher.Sum(nil)
	return string(bytestring_hash)
}