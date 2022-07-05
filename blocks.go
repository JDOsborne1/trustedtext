package main

import (
	"crypto/sha1"
	"errors"
)

type trustedtext_s struct {
	author string
	tags   []string
	body   string
	previous_hash string
	hash   string
}
// This function is called to generate a base instance of the trustedtext block, based
// on its arguments, and then sign it. A block can be valid without tags, as it remains
// globally unique to the original author, however it must have both an author and body.
func Instantiate(_author string, _tags []string, _body string) (trustedtext_s, error) {
	if len(_author) == 0 {
		return trustedtext_s{}, errors.New("cannot have a missing author")
	}
	if len(_body) == 0 {
		return trustedtext_s{}, errors.New("cannot have an empty body")
	}
	unsigned_tt := trustedtext_s{author: _author, tags: _tags, body: _body}
	return sign_tt(unsigned_tt), nil
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