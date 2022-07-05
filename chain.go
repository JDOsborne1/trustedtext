package main

type Trustedtext_chain_i interface {
	Genesis(_author string, _tags []string) Trustedtext_chain_i
	Amend(_existing_ttc Trustedtext_chain_i, _author string, _body string) Trustedtext_chain_i
	Most_recent_hash(_existing_ttc Trustedtext_chain_i) string
	Head_hash(_existing_ttc Trustedtext_chain_i) string
	Move_head_hash(_existing_ttc Trustedtext_chain_i, _new_head_hash string) Trustedtext_chain_i
	Return_head_block(_existing_ttc Trustedtext_chain_i) trustedtext_s
}

type trustedtext_chain_s struct {
	original_author string
	tt_chain        []trustedtext_s
	head_hash string
}

// Genesis is the function called to create a new trusted text chain. 
// This is always initiated with the same first message. Partly because 
// the first element lacks a 'previous hash' which makes it harder to 
// consider all the text bodies 'validated'. By defering the first real 
// message to actually be the second message, we can ensure that all real
// text data is held in blocks with both sides of the chain. 
func Genesis(_author string, _tags []string) trustedtext_chain_s {
	first_element,_ := Instantiate(
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

// Amend is the function called to increment a chain with a new tt block. This is 'stateless' 
// such that it creates a new chain, which is a copy of the previous, but for the inclusion of 
// a new block at the end. 
func Amend(_existing_ttc trustedtext_chain_s, _author string, _body string) trustedtext_chain_s {
	new_element,_ := Instantiate(
		_author,
		_existing_ttc.tt_chain[0].tags,
		_body,
	)

	new_element.previous_hash = Most_recent_hash(_existing_ttc)

	_existing_ttc.tt_chain =  append(_existing_ttc.tt_chain, new_element)
	return _existing_ttc
}

// Most_recent_hash is the function called to find one of the core identifiers of a chain,
// its last hash. For trustedtext this is only one of 2 hashes which are key to operations. 
func Most_recent_hash(_existing_ttc trustedtext_chain_s) string {
	chain_length := len(_existing_ttc.tt_chain)
	last_element := _existing_ttc.tt_chain[chain_length-1]
	return last_element.hash
}

// Head_hash is a function called to find a core chain identifier. This is the hash of the header block. 
// This header block may be moved over time, and points to the block which contains the current definitive 
// record of the trusted text element.
func Head_hash(_existing_trustedtext trustedtext_chain_s) string {
	return _existing_trustedtext.head_hash
}