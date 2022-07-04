package main

type trustedtext_s struct {
	author string
	tags   []string
	body   string
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

func return_hash(_trusted_text_element trustedtext_s) string {
	return "testhash"
}
