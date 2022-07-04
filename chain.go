package main

type Trustedtext_chain_i interface {
	Genesis(_author string, _tags []string) Trustedtext_chain_i
	Amend(_existing_ttc Trustedtext_chain_i, _author string, _body string) Trustedtext_chain_i
}

type trustedtext_chain_s struct {
	original_author string
	tt_chain        []trustedtext_s
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
	}
	return new_chain
}

func Amend(_existing_ttc trustedtext_chain_s, _author string, _body string) trustedtext_chain_s {
	new_element := Instantiate(
		_author,
		_existing_ttc.tt_chain[0].tags,
		_body,
	)
	_existing_ttc.tt_chain =  append(_existing_ttc.tt_chain, new_element)
	return _existing_ttc
}
