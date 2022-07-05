package main

type Trustedtext_chain_i interface {
	Genesis(_author string, _tags []string) Trustedtext_chain_i
	Amend(_existing_ttc Trustedtext_chain_i, _author string, _body string) Trustedtext_chain_i
	Most_recent_hash(_existing_ttc Trustedtext_chain_i) string
	Head_hash(_existing_ttc Trustedtext_chain_i) string
}

type trustedtext_chain_s struct {
	original_author string
	tt_chain        []trustedtext_s
	head_hash string
}

func Genesis(_author string, _tags []string) trustedtext_chain_s {
	first_element := Instantiate(
		_author,
		_tags,
		"This is the origin message of a trusted text chain",
	)

	new_chain := trustedtext_chain_s{
		original_author: _author,
		tt_chain:        []trustedtext_s{first_element},
		head_hash: 		 first_element.hash,
	}
	return new_chain
}

func Amend(_existing_ttc trustedtext_chain_s, _author string, _body string) trustedtext_chain_s {
	new_element := Instantiate(
		_author,
		_existing_ttc.tt_chain[0].tags,
		_body,
	)

	new_element.previous_hash = Most_recent_hash(_existing_ttc)

	_existing_ttc.tt_chain =  append(_existing_ttc.tt_chain, new_element)
	return _existing_ttc
}

func Most_recent_hash(_existing_ttc trustedtext_chain_s) string {
	chain_length := len(_existing_ttc.tt_chain)
	last_element := _existing_ttc.tt_chain[chain_length-1]
	return last_element.hash
}

func Head_hash(_existing_trustedtext trustedtext_chain_s) string {
	return _existing_trustedtext.head_hash
}