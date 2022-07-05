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
	signed_tt, err := sign_tt(unsigned_tt)
	if err != nil {
		return trustedtext_s{}, err
	}
	return signed_tt, nil
}

// This function wraps the signing process for the trusted text blocks. It will call the 
// hashing function, and then return a version of the input with a populated hash element,
// derived from the core content of the text element
func sign_tt(_existing_trustedtext trustedtext_s) (trustedtext_s, error) {
	content_hash, err := return_hash(_existing_trustedtext)
	if err != nil {
		return trustedtext_s{}, err
	}
	_existing_trustedtext.hash = content_hash
	return _existing_trustedtext, nil
}

// Fairly straightforward function to collapse the tags on a block in order to hash them.
// This is a dependency for the hashing wrapper, so must be kept static in order to preserve
// hashing continuity
func collapse_tags(_list_of_tags []string) string {
	tag_list := ""
	for _, tag := range _list_of_tags {
		tag_list = tag_list + tag + ","
	}
	return tag_list
}

// This function wraps the underlying hashing process, reducing it simply to
// block in - string out. This structure should be locked in early, since any 
// change to it will almost certainly invalidate all the hashing chains
func return_hash(_trusted_text_element trustedtext_s) (string, error) {
	elements := _trusted_text_element.author +
		_trusted_text_element.body +
		_trusted_text_element.previous_hash +
		collapse_tags(_trusted_text_element.tags)

	hasher := sha1.New()

	_, err := hasher.Write([]byte(elements))

	if err != nil {
		return "", err
	}

	bytestring_hash := hasher.Sum(nil)
	return string(bytestring_hash), nil
}